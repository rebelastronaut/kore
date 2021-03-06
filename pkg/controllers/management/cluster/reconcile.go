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

package cluster

import (
	"fmt"
	"reflect"
	"time"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	"github.com/appvia/kore/pkg/controllers"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils/kubernetes"

	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	finalizerName = "cluster.clusters.kore.appvia.io"
)

// Reconcile is the entrypoint for the reconciliation logic
func (c *Controller) Reconcile(ctx kore.Context, request reconcile.Request) (reconcile.Result, error) {
	ctx.Logger().Debug("attempting to reconcile the cluster")

	// @step: retrieve the object from the api
	cluster := &clustersv1.Cluster{}
	if err := ctx.Client().Get(ctx, request.NamespacedName, cluster); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		ctx.Logger().WithError(err).Error("trying to retrieve cluster from api")

		return reconcile.Result{}, err
	}
	original := cluster.DeepCopy()

	if cluster.Annotations[kore.AnnotationSystem] == kore.AnnotationValueTrue {
		cluster.Status.Status = corev1.SuccessStatus
		if err := ctx.Client().Status().Patch(ctx, cluster, client.MergeFrom(original)); err != nil {
			ctx.Logger().WithError(err).Error("failed to update the cluster status")
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	finalizer := kubernetes.NewFinalizer(ctx.Client(), finalizerName)
	if finalizer.IsDeletionCandidate(cluster) {
		return c.Delete(ctx, cluster)
	}

	// @logic:
	// - we retrieve the cloud provider (this is responsible for knowing which components to create)
	// - we generate on each iteration what we need
	// - we try and load the components if they exist
	// - we then use the provider to fill in any components from other i.e. a eks <- vpc
	// - we apply the components in order when theirs dependents are ready

	result, err := func() (reconcile.Result, error) {
		provider, exists := kore.GetClusterProvider(cluster.Spec.Kind)
		if !exists {
			return reconcile.Result{}, controllers.NewCriticalError(fmt.Errorf("%q cluster provider is invalid", cluster.Spec.Kind))
		}

		components := &kore.ClusterComponents{}

		return controllers.DefaultEnsureHandler.Run(ctx,
			[]controllers.EnsureFunc{
				c.AddFinalizer(cluster),
				c.SetPending(cluster),
				c.setComponents(cluster, components),
				c.setProviderComponents(provider, cluster, components),
				c.Load(cluster, components),
				func(ctx kore.Context) (reconcile.Result, error) {
					return reconcile.Result{}, provider.BeforeComponentsUpdate(ctx, cluster, components)
				},
				c.beforeComponentsUpdate(cluster, components),
				c.Apply(cluster, components),
				func(ctx kore.Context) (reconcile.Result, error) {
					return reconcile.Result{}, provider.SetProviderData(ctx, cluster, components)
				},
				c.Cleanup(cluster, components),
				c.SetClusterStatus(cluster, components),
			},
		)
	}()

	if err != nil {
		ctx.Logger().WithError(err).Error("trying to ensure the cluster")

		if controllers.IsCriticalError(err) {
			cluster.Status.Status = corev1.FailureStatus
			cluster.Status.Message = err.Error()
		}
	}

	if !reflect.DeepEqual(cluster, original) {
		if err := ctx.Client().Status().Patch(ctx, cluster, client.MergeFrom(original)); err != nil {
			ctx.Logger().WithError(err).Error("trying to patch the cluster status")

			return reconcile.Result{}, err
		}
	}

	return result, err
}

// AddFinalizer ensures the finalizer is on the resource
func (c *Controller) AddFinalizer(cluster *clustersv1.Cluster) controllers.EnsureFunc {
	return func(ctx kore.Context) (reconcile.Result, error) {
		finalizer := kubernetes.NewFinalizer(ctx.Client(), finalizerName)
		if finalizer.NeedToAdd(cluster) {
			if err := finalizer.Add(cluster); err != nil {
				ctx.Logger().WithError(err).Error("trying to add the finalizer")

				return reconcile.Result{}, err
			}

			return reconcile.Result{Requeue: true}, nil
		}

		return reconcile.Result{}, nil
	}
}

// SetPending ensures the state of the cluster is set to pending if not
func (c *Controller) SetPending(cluster *clustersv1.Cluster) controllers.EnsureFunc {
	return func(ctx kore.Context) (reconcile.Result, error) {
		switch cluster.Status.Status {
		case corev1.DeletingStatus:
			return reconcile.Result{RequeueAfter: 30 * time.Second}, nil
		}

		if cluster.Status.Status == "" {
			cluster.Status.Status = corev1.PendingStatus
			return reconcile.Result{Requeue: true}, nil
		}

		cluster.Status.Status = corev1.PendingStatus
		cluster.Status.Message = ""

		return reconcile.Result{}, nil
	}
}

// Apply is responsible for applying the component and updating the component status
func (c *Controller) Apply(cluster *clustersv1.Cluster, components *kore.ClusterComponents) controllers.EnsureFunc {
	return func(ctx kore.Context) (reconcile.Result, error) {
		if cluster.Status.Components == nil {
			cluster.Status.Components = corev1.Components{}
		}

		// We walk each of the components in order, we create them if required. If the resource
		// is not yet successful we wait and requeue. If the resource has failed, we throw
		// a critical failure and stop
		for _, comp := range *components {
			if comp.Absent {
				continue
			}

			result, err := c.applyComponent(ctx, cluster, comp)
			if err != nil || result.Requeue || result.RequeueAfter > 0 {
				return result, err
			}

			if comp.OnSuccess != nil {
				comp.OnSuccess(ctx, cluster, comp, components)
			}
		}

		return reconcile.Result{}, nil
	}
}

func (c *Controller) applyComponent(ctx kore.Context, cluster *clustersv1.Cluster, comp *kore.ClusterComponent) (reconcile.Result, error) {
	condition, found := cluster.Status.Components.GetComponent(comp.ComponentName())
	if !found {
		condition = &corev1.Component{
			Name:   comp.ComponentName(),
			Status: corev1.PendingStatus,
		}
	}
	defer func() {
		cluster.Status.Components.SetCondition(*condition)
	}()

	ownership := corev1.MustGetOwnershipFromObject(comp.Object)
	condition.Resource = &ownership
	logger := ctx.Logger().WithFields(log.Fields{
		"component": comp.ComponentName(),
		"condition": condition.Status,
		"existing":  comp.Exists(),
	})
	logger.Debug("attempting to reconciling the component")

	annotations := comp.Object.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}
	annotations[kore.AnnotationReadOnly] = kore.AnnotationValueTrue

	comp.Object.SetAnnotations(annotations)

	updated, err := comp.Update(ctx)
	if err != nil {
		return reconcile.Result{}, err
	}
	if updated {
		ctx.Logger().WithField("component", kubernetes.MustGetRuntimeSelfLink(comp.Object)).Debug("component has changed")
		return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
	}

	// @check if the resource is ready to reconcile
	status, err := GetObjectStatus(comp.Object)
	if err != nil {
		if err == kubernetes.ErrFieldNotFound {
			return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
		}

		logger.WithError(err).Error("trying to check the component status")
		return reconcile.Result{}, err
	}

	logger.WithField(
		"status", status,
	).Debug("current state of the resource")

	condition.Status = status
	condition.Message = ""
	condition.Detail = ""

	switch status {
	case corev1.SuccessStatus:
		return reconcile.Result{}, nil
	case corev1.FailureStatus:
		if cnd, err := GetObjectReasonForFailure(comp.Object); err == nil {
			condition.Message = cnd.Message
			condition.Detail = cnd.Detail
		} else {
			condition.Message = "Failed to provision the resource"
		}
	}

	return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
}

// Cleanup is responsible for deleting any components no longer required
func (c *Controller) Cleanup(cluster *clustersv1.Cluster, components *kore.ClusterComponents) controllers.EnsureFunc {
	return func(ctx kore.Context) (reconcile.Result, error) {
		// @logic:
		// - we remove any absent components first in reverse dependency order
		// - we find any components which are no longer referenced in status and we delete those too

		for i := len(*components) - 1; i >= 0; i-- {
			comp := (*components)[i]
			if !comp.Absent {
				continue
			}

			result, err := c.removeComponent(ctx, cluster, components, comp)
			if err != nil || result.Requeue || result.RequeueAfter > 0 {
				return result, err
			}

			cluster.Status.Components.RemoveComponent(comp.ComponentName())
		}

		for i := 0; i < len(cluster.Status.Components); i++ {
			statusComponent := cluster.Status.Components[i]
			if statusComponent.Resource == nil {
				continue
			}

			comp := components.Find(func(comp kore.ClusterComponent) bool {
				return kore.IsOwner(comp.Object, *statusComponent.Resource)
			})
			if comp != nil {
				continue
			}

			object, err := kubernetes.NewObject(statusComponent.Resource.GroupVersionKind())
			if err != nil {
				return reconcile.Result{}, err
			}
			object.SetName(statusComponent.Resource.Name)
			object.SetNamespace(statusComponent.Resource.Namespace)

			comp = &kore.ClusterComponent{Object: object}

			ctx.Logger().WithField("component", comp.ComponentName()).Debug("component is not defined anymore, deleting")

			res, err := c.removeComponent(ctx, cluster, components, comp)
			if err != nil || res.Requeue || res.RequeueAfter > 0 {
				return res, err
			}

			cluster.Status.Components.RemoveComponent(statusComponent.Name)
		}

		return reconcile.Result{}, nil
	}
}
