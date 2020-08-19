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

package serviceproviders

import (
	"fmt"
	"sort"
	"time"

	"github.com/appvia/kore/pkg/kore"

	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/controllers"
	"github.com/appvia/kore/pkg/controllers/helpers"
	"github.com/appvia/kore/pkg/utils/kubernetes"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	finalizerName = "serviceprovider.kore.appvia.io"
)

// Reconcile is the entrypoint for the reconciliation logic
func (c *Controller) Reconcile(ctx kore.Context, request reconcile.Request) (reconcileResult reconcile.Result, reconcileError error) {
	ctx.Logger().Debug("attempting to reconcile the service provider")

	// @step: retrieve the object from the api
	serviceProvider := &servicesv1.ServiceProvider{}
	if err := ctx.Client().Get(ctx, request.NamespacedName, serviceProvider); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		ctx.Logger().WithError(err).Error("trying to retrieve service provider from api")

		return reconcile.Result{}, err
	}
	original := serviceProvider.DeepCopy()

	defer func() {
		if err := ctx.Client().Status().Patch(ctx, serviceProvider, client.MergeFrom(original)); err != nil {
			ctx.Logger().WithError(err).Error("failed to update the service provider status")

			reconcileResult = reconcile.Result{}
			reconcileError = err
		}
	}()

	finalizer := kubernetes.NewFinalizer(ctx.Client(), finalizerName)
	if finalizer.IsDeletionCandidate(serviceProvider) {
		return c.delete(ctx, serviceProvider, finalizer)
	}

	result, err := func() (reconcile.Result, error) {
		ensure := []controllers.EnsureFunc{
			c.ensureFinalizer(serviceProvider, finalizer),
			c.ensurePending(serviceProvider),
			func(ctx kore.Context) (result reconcile.Result, err error) {
				complete, err := ctx.Kore().ServiceProviders().SetUp(ctx, serviceProvider)
				if err != nil {
					return reconcile.Result{}, fmt.Errorf("failed to set up service provider: %w", err)
				}
				if !complete {
					return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
				}

				provider, err := ctx.Kore().ServiceProviders().Register(ctx, serviceProvider)
				if err != nil {
					return reconcile.Result{}, fmt.Errorf("failed to register service provider: %w", err)
				}

				catalog, err := ctx.Kore().ServiceProviders().Catalog(ctx, serviceProvider)
				if err != nil {
					return reconcile.Result{}, fmt.Errorf("failed to load service catalog: %w", err)
				}

				var supportedKinds []string
				for _, kind := range catalog.Kinds {
					supportedKinds = append(supportedKinds, kind.Name)
				}
				sort.Strings(supportedKinds)

				serviceProvider.Status.SupportedKinds = supportedKinds

				kinds := map[string]*servicesv1.ServiceKind{}
				for _, kind := range catalog.Kinds {
					kind.Namespace = kore.HubNamespace

					existing := &servicesv1.ServiceKind{}
					existing.SetGroupVersionKind(kind.GroupVersionKind())
					existing.Name = kind.Name
					existing.Namespace = kind.Namespace
					exists, err := kubernetes.GetIfExists(ctx, ctx.Client(), existing)
					if err != nil {
						return reconcile.Result{}, fmt.Errorf("failed to retrieve service kind %q: %w", kind.Name, err)
					}
					if exists {
						kind.Spec.Enabled = existing.Spec.Enabled
					}

					kubernetes.EnsureOwnerReference(&kind, serviceProvider, true)

					if _, err := kubernetes.CreateOrUpdate(ctx, ctx.Client(), &kind); err != nil {
						return reconcile.Result{}, fmt.Errorf("failed to create or update service kind %q: %w", kind.Name, err)
					}

					kind.SetGroupVersionKind(servicesv1.ServiceKindGVK)
					kinds[kind.Name] = kind.DeepCopy()
				}

				for _, plan := range catalog.Plans {
					kubernetes.EnsureOwnerReference(&plan, kinds[plan.Spec.Kind], true)

					plan.Namespace = kore.HubNamespace
					if plan.Annotations == nil {
						plan.Annotations = map[string]string{}
					}
					plan.Annotations[kore.AnnotationReadOnly] = kore.AnnotationValueTrue

					if _, err := kubernetes.CreateOrUpdate(ctx, ctx.Client(), &plan); err != nil {
						return reconcile.Result{}, fmt.Errorf("failed to create or update service plan %q: %w", plan.Name, err)
					}
				}

				var adminServices []servicesv1.Service
				for _, service := range provider.AdminServices() {
					service.Namespace = kore.HubAdminTeam
					if service.Annotations == nil {
						service.Annotations = map[string]string{}
					}
					service.Annotations[kore.AnnotationSystem] = kore.AnnotationValueTrue
					service.Annotations[kore.AnnotationReadOnly] = kore.AnnotationValueTrue
					adminServices = append(adminServices, service)
				}

				result, err = helpers.EnsureServices(
					ctx,
					adminServices,
					serviceProvider,
					&serviceProvider.Status.Components,
				)
				if err != nil || result.Requeue || result.RequeueAfter > 0 {
					return result, err
				}

				return reconcile.Result{}, nil
			},
		}

		for _, handler := range ensure {
			result, err := handler(ctx)
			if err != nil {
				return reconcile.Result{}, err
			}
			if result.Requeue || result.RequeueAfter > 0 {
				return result, nil
			}
		}
		return reconcile.Result{}, nil
	}()

	if err != nil {
		ctx.Logger().WithError(err).Error("failed to reconcile the service provider")

		serviceProvider.Status.Status = corev1.ErrorStatus
		serviceProvider.Status.Message = err.Error()

		if controllers.IsCriticalError(err) {
			serviceProvider.Status.Status = corev1.FailureStatus
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, err
	}

	if result.Requeue || result.RequeueAfter > 0 {
		return result, nil
	}

	serviceProvider.Status.Status = corev1.SuccessStatus
	serviceProvider.Status.Message = ""

	return result, nil
}
