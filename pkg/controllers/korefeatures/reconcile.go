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

package features

import (
	"context"
	"fmt"

	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/controllers"
	"github.com/appvia/kore/pkg/controllers/helpers"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils/kubernetes"

	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	finalizerName = "config.kore.appvia.io"
)

// Reconcile is responsible for handling the scanning of a kind
func (c *Controller) Reconcile(request reconcile.Request) (reconcileResult reconcile.Result, reconcileError error) {
	ctx := context.Background()

	logger := c.logger.WithFields(log.Fields{
		"name":      request.Name,
		"namespace": request.Namespace,
	})
	logger.Debug("attempting to reconcile feature")

	// @step: retrieve the object from the api
	feature := &configv1.KoreFeature{}
	if err := c.client.Get(ctx, request.NamespacedName, feature); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		logger.WithError(err).Error("failed to retrieve feature from api")

		return reconcile.Result{}, err
	}
	original := feature.DeepCopy()

	logger = logger.WithField("feature", feature.Name)

	defer func() {
		if err := c.client.Status().Patch(ctx, feature, client.MergeFrom(original)); err != nil {
			if !kerrors.IsNotFound(err) {
				logger.WithError(err).Error("failed to update the feature status")
				reconcileResult = reconcile.Result{}
				reconcileError = err
			}
		}
	}()

	finalizer := kubernetes.NewFinalizer(c.client, finalizerName)
	if finalizer.IsDeletionCandidate(feature) {
		return c.delete(ctx, logger, feature, finalizer)
	}

	result, err := func() (reconcile.Result, error) {
		ensure := []controllers.EnsureFunc{
			c.ensureFinalizer(feature, finalizer),
			c.ensurePending(feature),
			func(ctx kore.Context) (result reconcile.Result, err error) {
				var services []servicesv1.Service
				switch feature.Spec.FeatureType {
				case configv1.KoreFeatureCosts:
					services, err = c.getCostsServices(ctx, feature)
					if err != nil {
						return reconcile.Result{}, err
					}
				default:
					return reconcile.Result{}, fmt.Errorf("Unknown feature type %s", feature.Spec.FeatureType)
				}

				result, err = helpers.EnsureServices(
					kore.NewContext(ctx, logger, c.client, c.kore),
					services,
					feature,
					&feature.Status.Components,
				)
				if err != nil || result.Requeue || result.RequeueAfter > 0 {
					return result, err
				}

				return reconcile.Result{}, nil
			},
		}

		koreCtx := kore.NewContext(ctx, logger, c.client, c.kore)
		for _, handler := range ensure {
			result, err := handler(koreCtx)
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
		logger.WithError(err).Error("failed to reconcile the feature")

		feature.Status.Status = corev1.ErrorStatus
		feature.Status.Message = err.Error()

		if controllers.IsCriticalError(err) {
			feature.Status.Status = corev1.FailureStatus
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, err
	}

	if result.Requeue || result.RequeueAfter > 0 {
		return result, nil
	}

	feature.Status.Status = corev1.SuccessStatus
	feature.Status.Message = ""

	return result, nil
}
