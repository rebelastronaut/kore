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

package kubernetes

import (
	"context"
	"time"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	"github.com/appvia/kore/pkg/controllers"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils/kubernetes"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// EnsureCloudProvider is responsible for checking the cloud provider
func (a k8sCtrl) EnsureCloudProvider(object *clustersv1.Kubernetes) controllers.EnsureFunc {
	return func(ctx kore.Context) (reconcile.Result, error) {
		if !kore.IsProviderBacked(object) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, nil
	}
}

// EnsureKoreNamespaces is responsible for ensuring the namespace are present
func (a k8sCtrl) EnsureKoreNamespaces(ctx context.Context, client client.Client) error {
	for _, name := range []string{kore.HubNamespace, kore.HubSystem, kore.HubOperators} {
		ns := &v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
				Annotations: map[string]string{
					kore.Label("owned"): "true",
				},
			},
		}

		if err := kubernetes.EnsureNamespace(ctx, client, ns); err != nil {
			return err
		}
	}

	return nil

}

// EnsureDeleteStatus is responsible for ensure the status is set to deleting
func (a k8sCtrl) EnsureDeleteStatus(object *clustersv1.Kubernetes) controllers.EnsureFunc {
	return func(ctx kore.Context) (reconcile.Result, error) {
		if object.Status.Status != corev1.DeletingStatus {
			object.Status.Status = corev1.DeletingStatus

			return reconcile.Result{Requeue: true}, nil
		}

		return reconcile.Result{}, nil
	}
}

// EnsureServiceDeletion is responsible for cleanup the resources in the cluster
// @note: at present this is only done for EKS as GKE performs it's own cleanup
func (a k8sCtrl) EnsureServiceDeletion(object *clustersv1.Kubernetes) controllers.EnsureFunc {
	return func(ctx kore.Context) (reconcile.Result, error) {
		logger := log.WithFields(log.Fields{
			"name":      object.Name,
			"namespace": object.Namespace,
		})
		logger.Debug("attempting to delete the kubernetes resource")

		if !kore.IsProviderBacked(object) {
			return reconcile.Result{}, nil
		}
		if object.Spec.Provider.Kind != "EKS" {
			return reconcile.Result{}, nil
		}

		result, err := func() (reconcile.Result, error) {
			// @step: retrieve the provider credentials secret
			token, err := controllers.GetConfigSecret(ctx, ctx.Client(), object.Namespace, object.Name)
			if err != nil {
				if kerrors.IsNotFound(err) {
					return reconcile.Result{}, nil
				}

				return reconcile.Result{}, err
			}

			// @step: create a client for the remote cluster
			cc, err := kubernetes.NewRuntimeClientFromConfigSecret(token)
			if err != nil {
				return reconcile.Result{}, err
			}

			// @note: we need to look for any namespaces with loadbalancer types and
			// delete them to free up the ELB and security groups. Deleting namespace isn't easier
			// as you will probably end up with namespace in a forever loop due to you deleting
			// the controllers which is responsible for finalizing them
			list, err := kubernetes.ListServicesByTypes(ctx, cc, v1.NamespaceAll, string(v1.ServiceTypeLoadBalancer))
			if err != nil {
				return reconcile.Result{}, err
			}
			if len(list.Items) > 0 {
				logger.Debug("cluster still has resource left to cleanup")

				for _, x := range list.Items {
					if x.GetDeletionTimestamp() != nil {
						continue
					}
					if err := cc.Delete(ctx, &x); err != nil {
						return reconcile.Result{}, err
					}
				}

				return reconcile.Result{RequeueAfter: 30 * time.Second}, nil
			}

			return reconcile.Result{}, nil
		}()
		if err != nil {
			logger.WithError(err).Error("trying to ensure the cluster is cleaned out")

			object.Status.Components.SetCondition(corev1.Component{
				Name:    ComponentKubernetesCleanup,
				Message: "Failed trying to cleanup in cluster resources",
				Detail:  err.Error(),
				Status:  corev1.FailureStatus,
			})

			return reconcile.Result{}, err
		}

		return result, nil
	}
}

// EnsureSecretDeletion is responsible for deletion the admin token
func (a k8sCtrl) EnsureSecretDeletion(object *clustersv1.Kubernetes) controllers.EnsureFunc {
	return func(ctx kore.Context) (reconcile.Result, error) {
		// @step: we should delete the secert from api
		if err := kubernetes.DeleteIfExists(ctx, ctx.Client(), &configv1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      object.Name,
				Namespace: object.Namespace,
			},
		}); err != nil {
			log.WithError(err).Error("trying to delete the secret from api")

			return reconcile.Result{}, err
		}

		return reconcile.Result{}, nil
	}
}

// EnsureFinalizerRemoved removes the finalizer now
func (a k8sCtrl) EnsureFinalizerRemoved(object *clustersv1.Kubernetes) controllers.EnsureFunc {
	return func(ctx kore.Context) (reconcile.Result, error) {
		// @cool we can remove the finalizer now
		finalizer := kubernetes.NewFinalizer(ctx.Client(), finalizerName)
		if finalizer.IsDeletionCandidate(object) {
			if err := finalizer.Remove(object); err != nil {
				return reconcile.Result{}, err
			}
		}

		return reconcile.Result{}, nil
	}
}
