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

package eksnodegroup

import (
	"context"

	"github.com/appvia/kore/pkg/kore"

	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	eks "github.com/appvia/kore/pkg/apis/eks/v1alpha1"
	"github.com/appvia/kore/pkg/controllers"

	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Delete is responsible for deleting the aws eks nodegroup
func (n *ctrl) Delete(request reconcile.Request) (reconcile.Result, error) {
	ctx := context.Background()
	logger := log.WithFields(log.Fields{
		"name":      request.NamespacedName.Name,
		"namespace": request.NamespacedName.Namespace,
	})
	logger.Info("attempting to delete eks cluster nodegroup")

	// @step: retrieve the resource from the api
	resource := &eks.EKSNodeGroup{}
	if err := n.mgr.GetClient().Get(ctx, request.NamespacedName, resource); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, err
	}
	original := resource.DeepCopy()

	koreCtx := kore.NewContext(ctx, logger, n.mgr.GetClient(), n)
	result, err := func() (reconcile.Result, error) {
		return controllers.DefaultEnsureHandler.Run(koreCtx,
			[]controllers.EnsureFunc{
				n.EnsureDeletionStatus(resource),
				n.EnsureDeletion(resource),
				n.EnsureRoleDeletion(resource),
				n.EnsureRemoveFinalizer(resource),
			},
		)
	}()
	if err != nil {
		logger.WithError(err).Error("attempting to delete the eks cluster")
		resource.Status.Status = corev1.FailureStatus
	}

	// @step: we update always update the status before throwing any error
	if err := controllers.PatchStatus(ctx, n.mgr.GetClient(), resource, original); err != nil {
		logger.WithError(err).Error("trying to update the resource status")

		return reconcile.Result{}, err
	}

	return result, err
}
