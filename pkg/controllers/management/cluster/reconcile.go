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
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	eksv1alpha1 "github.com/appvia/kore/pkg/apis/eks/v1alpha1"
	"github.com/appvia/kore/pkg/controllers"
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
func (a *Controller) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	ctx := context.Background()

	logger := a.logger.WithFields(log.Fields{
		"name":      request.NamespacedName.Name,
		"namespace": request.NamespacedName.Namespace,
	})
	logger.Debug("attempting to reconcile the cluster")

	// @step: retrieve the object from the api
	cluster := &clustersv1.Cluster{}
	if err := a.mgr.GetClient().Get(ctx, request.NamespacedName, cluster); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		logger.WithError(err).Error("trying to retrieve cluster from api")

		return reconcile.Result{}, err
	}
	original := cluster.DeepCopy()

	finalizer := kubernetes.NewFinalizer(a.mgr.GetClient(), finalizerName)
	if finalizer.NeedToAdd(cluster) {
		err := finalizer.Add(cluster)
		if err != nil {
			logger.WithError(err).Error("failed to set the finalizer")
		}

		return reconcile.Result{Requeue: true}, err
	}

	if finalizer.IsDeletionCandidate(cluster) {
		return a.Delete(ctx, cluster)
	}

	err := func() error {
		cluster.Status.Status = corev1.PendingStatus
		if cluster.Status.Components == nil {
			cluster.Status.Components = corev1.Components{}
		}

		// @step: we generate a collection of resources based on the cluster type
		// GKE -> K8S, EKS -> VPC -> K8S etc.
		components, err := createClusterComponents(cluster)
		if err != nil {
			logger.WithError(err).Error("trying to create the cluster components")

			return controllers.NewCriticalError(err)
		}

		for componentName, c := range components {
			if err := a.loadComponent(ctx, cluster, c); err != nil {
				logger.WithError(err).Error("trying to load the cluster components")

				return fmt.Errorf("trying to load %s component: %w", componentName, err)
			}
		}

		for _, c := range components {
			switch r := c.(type) {
			case *eksv1alpha1.EKSVPC:
				applyEKSVPC(r, components)
			}
		}

		for componentName, c := range components {
			if readyForReconcile(c, components) {
				if err := a.createOrUpdateComponent(ctx, cluster, c); err != nil {
					logger.WithError(err).Error("trying create or update the cluster components")

					return fmt.Errorf("trying to create or update %s component: %w", componentName, err)
				}
			}
			switch r := c.(type) {
			case *clustersv1.Kubernetes:
				if r.Status.Status == corev1.SuccessStatus {
					cluster.Status.APIEndpoint = r.Status.APIEndpoint
					cluster.Status.AuthProxyEndpoint = r.Status.Endpoint
					cluster.Status.CaCertificate = r.Status.CaCertificate
				}
			}

			status, message := c.GetStatus()
			if status == "" {
				status = corev1.PendingStatus
			}
			if status.IsFailed() && message == "" {
				if err := c.GetComponents().Error(); err != nil {
					logger.WithError(err).Error("trying to retrieve the components")

					message = err.Error()
				}
			}
			component := corev1.Component{
				Name:    componentName,
				Status:  status,
				Message: message,
			}
			cluster.Status.Components.SetCondition(component)
		}

		// Check for deleted EKS node groups
		for _, existing := range cluster.Status.Components {
			if strings.HasPrefix(existing.Name, "EKSNodeGroup/") {
				if components[existing.Name] == nil {
					name := strings.TrimPrefix(existing.Name, "EKSNodeGroup/")
					deleted, err := deleteEKSNodeGroup(ctx, a.mgr.GetClient(), cluster, name)
					if err != nil {
						cluster.Status.Components.SetCondition(corev1.Component{
							Name:    existing.Name,
							Status:  corev1.DeleteFailedStatus,
							Message: err.Error(),
						})
						return err
					}

					cluster.Status.Components.SetCondition(corev1.Component{
						Name:   existing.Name,
						Status: corev1.DeletingStatus,
					})

					if deleted {
						logger.Infof("%s EKS Node Group was deleted", name)
						cluster.Status.Components.RemoveComponent(existing.Name)
					}
				}
			}
		}

		ready := cluster.Status.Components.HasStatusForAll(corev1.SuccessStatus)
		if ready {
			cluster.Status.Status = corev1.SuccessStatus
			cluster.Status.Message = "The cluster has been created successfully"

			return nil
		}

		if cluster.Status.Components.HasStatus(corev1.FailureStatus) {
			return controllers.NewCriticalError(errors.New("one or more components failed"))
		}

		return nil
	}()
	if err != nil {
		logger.WithError(err).Error("failed to reconcile the cluster")

		if controllers.IsCriticalError(err) {
			cluster.Status.Status = corev1.FailureStatus
			cluster.Status.Message = err.Error()
		}
	}

	if err := a.mgr.GetClient().Status().Patch(ctx, cluster, client.MergeFrom(original)); err != nil {
		logger.WithError(err).Error("failed to update the cluster status")

		return reconcile.Result{}, err
	}

	if cluster.Status.Status == corev1.SuccessStatus {
		return reconcile.Result{}, nil
	}
	if cluster.Status.Status == corev1.FailureStatus {
		return reconcile.Result{RequeueAfter: 2 * time.Minute}, nil
	}

	return reconcile.Result{RequeueAfter: 30 * time.Second}, err
}

func (a *Controller) loadComponent(ctx context.Context, cluster *clustersv1.Cluster, res clustersv1.ClusterComponent) error {
	_, err := kubernetes.GetIfExists(ctx, a.mgr.GetClient(), res)
	if err != nil {
		return err
	}

	if err := res.ApplyClusterConfiguration(cluster); err != nil {
		return controllers.NewCriticalError(err)
	}

	return nil
}

// createOrUpdateComponent is responsible for maintaining the components
func (a *Controller) createOrUpdateComponent(ctx context.Context, cluster *clustersv1.Cluster, resource clustersv1.ClusterComponent) error {
	found, err := kubernetes.CheckIfExists(ctx, a.mgr.GetClient(), resource.DeepCopyObject())
	if err != nil {
		return err
	}

	if !found {
		setClusterResourceVersion(resource, cluster.ResourceVersion)

		return kubernetes.CreateOrUpdateObject(ctx, a.mgr.GetClient(), resource)
	}

	// @step: get the revision of the component
	revision := getClusterResourceVersion(resource)
	if revision != cluster.ResourceVersion {
		setClusterResourceVersion(resource, cluster.ResourceVersion)

		return kubernetes.CreateOrUpdateObject(ctx, a.mgr.GetClient(), resource)
	}

	return nil
}
