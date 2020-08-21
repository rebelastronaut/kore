/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package application

import (
	"errors"
	"fmt"
	"time"

	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/controllers"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils"
	"github.com/appvia/kore/pkg/utils/kubernetes"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	applicationv1beta "sigs.k8s.io/application/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func (p Provider) Reconcile(
	ctx kore.Context,
	service *servicesv1.Service,
) (reconcile.Result, error) {
	config, err := getAppConfiguration(ctx, service)
	if err != nil {
		return reconcile.Result{}, err
	}

	if service.Spec.Cluster.Name == "" || service.Spec.Cluster.Namespace == "" || service.Spec.ClusterNamespace == "" {
		return reconcile.Result{}, controllers.NewCriticalError(errors.New("a cluster and namespace must be defined on the service"))
	}

	clusterClient, err := controllers.CreateClient(ctx, ctx.Client(), service.Spec.Cluster)
	if err != nil {
		return reconcile.Result{}, err
	}
	if clusterClient == nil {
		return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
	}

	compiledResources, err := config.CompileResources(NewResourceParams(service, config))
	if err != nil {
		return reconcile.Result{}, err
	}

	if err := kubernetes.EnsureNamespace(ctx, clusterClient, &v1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: v1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: service.Spec.ClusterNamespace,
		},
	}); err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to create namespace: %q: %w", service.Spec.ClusterNamespace, err)
	}

	providerData := &ProviderData{}
	if err := service.Status.GetProviderData(providerData); err != nil {
		return reconcile.Result{}, err
	}

	// Check if we need to delete any resources
	for _, existingResource := range providerData.Resources {
		found := false
		for _, r := range compiledResources {
			if existingResource.Equals(corev1.MustGetOwnershipFromObject(r)) {
				found = true
				break
			}
		}
		if !found {
			u := existingResource.ToUnstructured()
			if err := kubernetes.DeleteIfExists(ctx, clusterClient, &u); err != nil {
				return reconcile.Result{}, fmt.Errorf("failed to delete %s: %w", utils.GetUnstructuredSelfLink(&u), err)
			}
		}
	}

	updatedProviderData := ProviderData{}

	for _, resource := range compiledResources {
		switch resource.(type) {
		case *v1.Namespace, *applicationv1beta.Application:
			continue
		}

		ctx.Logger().WithField("resource", kubernetes.MustGetRuntimeSelfLink(resource)).Trace("creating/updating resource")
		if err := ensureResource(ctx, clusterClient, resource.DeepCopyObject()); err != nil {
			return reconcile.Result{}, err
		}

		updatedProviderData.Resources = append(updatedProviderData.Resources, corev1.MustGetOwnershipFromObject(resource))
	}

	app := compiledResources.Application()

	if app != nil {
		ctx.Logger().WithField("application", kubernetes.MustGetRuntimeSelfLink(app)).Trace("creating/updating application object")
		if err := ensureResource(ctx, clusterClient, app.DeepCopyObject()); err != nil {
			return reconcile.Result{}, err
		}

		updatedProviderData.Resources = append(updatedProviderData.Resources, corev1.MustGetOwnershipFromObject(app))
	}

	if err := service.Status.SetProviderData(updatedProviderData); err != nil {
		return reconcile.Result{}, err
	}

	if app == nil {
		app = &applicationv1beta.Application{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Application",
				APIVersion: applicationv1beta.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      service.Name,
				Namespace: service.Spec.ClusterNamespace,
			},
		}
	}

	appExists, err := kubernetes.GetIfExists(ctx, clusterClient, app)
	if err != nil {
		if !utils.IsMissingKind(err) {
			return reconcile.Result{}, fmt.Errorf("failed to get application %q: %w", app.Name, err)
		}
	}

	if !appExists {
		ctx.Logger().Debug("application object does not exist, waiting")
		return reconcile.Result{RequeueAfter: 5 * time.Second}, nil
	}

	for _, condition := range app.Status.Conditions {
		switch condition.Type {
		case applicationv1beta.Error:
			if condition.Status == "True" {
				return reconcile.Result{}, errors.New(condition.Message)
			}
		case applicationv1beta.Ready:
			if condition.Status == "True" {
				// The Application status will be healthy even if there are no monitored resources, so we have to
				// explicitly check ComponentsReady for "0/0"
				if app.Status.ComponentsReady == "0/0" {
					return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
				}

				service.Status.Status = corev1.SuccessStatus
				service.Status.Message = condition.Message

				// We will actively monitor the application status and update the service
				return reconcile.Result{RequeueAfter: 1 * time.Minute}, nil
			} else {
				service.Status.Message = condition.Message
			}
		}
	}

	return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
}
