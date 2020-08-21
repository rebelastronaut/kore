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

package servicecatalog

import (
	"fmt"

	"github.com/appvia/kore/pkg/controllers"
	cc "github.com/appvia/kore/pkg/controllers/components"
	"github.com/appvia/kore/pkg/kore"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	finalizerName = "servicecatalog.kore.appvia.io"
)

// Reconcile is the entrypoint for the reconciliation logic
func (c *Controller) Reconcile(ctx kore.Context, request reconcile.Request) (reconcileResult reconcile.Result, reconcileError error) {
	ctx.Logger().Debug("attempting to reconcile the service catalog")

	serviceCatalog := &servicesv1.ServiceCatalog{}
	if err := ctx.Client().Get(ctx, request.NamespacedName, serviceCatalog); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, fmt.Errorf("failed to retrieve the service catalog: %w", err)
	}
	original := serviceCatalog.DeepCopyObject()

	components := controllers.Components{
		cc.NewFinalizer(finalizerName, serviceCatalog),
		newCatalogComponent(serviceCatalog),
	}

	res, err := components.Reconcile(ctx, serviceCatalog)
	if err != nil {
		ctx.Logger().WithError(err).Error("failed to reconcile the service catalog")
	}

	if err := ctx.Client().Status().Patch(ctx, serviceCatalog, client.MergeFrom(original)); err != nil {
		ctx.Logger().WithError(err).Error("failed to update the status of the service catalog")
		return reconcile.Result{}, fmt.Errorf("failed to update the status of the service catalog: %w", err)
	}

	return res, err
}
