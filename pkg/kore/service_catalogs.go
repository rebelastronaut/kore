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

package kore

import (
	"context"
	fmt "fmt"

	"github.com/appvia/kore/pkg/utils"

	"github.com/appvia/kore/pkg/store"
	"github.com/appvia/kore/pkg/utils/kubernetes"
	"github.com/appvia/kore/pkg/utils/validation"
	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
)

// ServiceCatalogs is the interface to manage service catalogs
type ServiceCatalogs interface {
	// CheckDelete verifies whether the service catalog can be deleted
	CheckDelete(context.Context, *servicesv1.ServiceCatalog, ...DeleteOptionFunc) error
	// Delete is used to delete a service catalog in kore
	Delete(context.Context, string, ...DeleteOptionFunc) (*servicesv1.ServiceCatalog, error)
	// Get returns the service catalog
	Get(context.Context, string) (*servicesv1.ServiceCatalog, error)
	// List returns the existing service catalogs
	// The optional filter functions can be used to include items only for which all functions return true
	List(context.Context, ...func(servicesv1.ServiceCatalog) bool) (*servicesv1.ServiceCatalogList, error)
	// Has checks if a service catalog exists
	Has(context.Context, string) (bool, error)
	// Update is responsible for updating a service catalog
	Update(context.Context, *servicesv1.ServiceCatalog) error
}

type serviceCatalogsImpl struct {
	Interface
}

// Update is responsible for updating a service catalog
func (p serviceCatalogsImpl) Update(ctx context.Context, catalog *servicesv1.ServiceCatalog) error {
	if err := IsValidResourceName("service catalog", catalog.Name); err != nil {
		return err
	}

	if catalog.Namespace != HubNamespace {
		return validation.NewError("%q failed validation", catalog.Name).
			WithFieldErrorf("namespace", validation.InvalidValue, "must be %q", HubNamespace)
	}

	err := p.Store().Client().Update(ctx,
		store.UpdateOptions.To(catalog),
		store.UpdateOptions.WithCreate(true),
		store.UpdateOptions.WithForce(true),
	)
	if err != nil {
		log.WithError(err).Error("failed to update a service catalog")

		return err
	}

	return nil
}

// CheckDelete verifies whether the service catalog can be deleted
func (p serviceCatalogsImpl) CheckDelete(ctx context.Context, serviceCatalog *servicesv1.ServiceCatalog, o ...DeleteOptionFunc) error {
	opts := ResolveDeleteOptions(o)

	if !opts.Cascade {
		var dependents []kubernetes.DependentReference

		servicePlans, err := p.ServicePlans().List(ctx, func(p servicesv1.ServicePlan) bool {
			return kubernetes.HasOwnerReference(&p, serviceCatalog)
		})
		if err != nil {
			return err
		}

		if len(servicePlans.Items) == 0 {
			return nil
		}

		var servicePlanNames []string
		for _, p := range servicePlans.Items {
			servicePlanNames = append(servicePlanNames, p.Name)
		}

		teamList, err := p.Teams().List(ctx)
		if err != nil {
			return fmt.Errorf("failed to list teams: %w", err)
		}

		for _, team := range teamList.Items {
			services, err := p.Teams().Team(team.Name).Services().List(ctx, func(s servicesv1.Service) bool {
				return utils.Contains(s.Spec.Plan, servicePlanNames)
			})
			if err != nil {
				return fmt.Errorf("failed to list services: %w", err)
			}
			for _, item := range services.Items {
				dependents = append(dependents, kubernetes.DependentReferenceFromObject(&item))
			}
		}

		if len(dependents) > 0 {
			return validation.ErrDependencyViolation{
				Message:    "the following objects need to be deleted first",
				Dependents: dependents,
			}
		}
	}

	return nil
}

// Delete is used to delete a service catalog in kore
func (p serviceCatalogsImpl) Delete(ctx context.Context, name string, o ...DeleteOptionFunc) (*servicesv1.ServiceCatalog, error) {
	opts := ResolveDeleteOptions(o)

	catalog := &servicesv1.ServiceCatalog{}
	err := p.Store().Client().Get(ctx,
		store.GetOptions.InNamespace(HubNamespace),
		store.GetOptions.InTo(catalog),
		store.GetOptions.WithName(name),
	)
	if err != nil {
		if kerrors.IsNotFound(err) {
			return nil, ErrNotFound
		}
		log.WithError(err).Error("failed to retrieve the service catalog")

		return nil, err
	}

	if err := opts.Check(catalog, func(o ...DeleteOptionFunc) error { return p.CheckDelete(ctx, catalog, o...) }); err != nil {
		return nil, err
	}

	if err := p.Store().Client().Delete(ctx, append(opts.StoreOptions(), store.DeleteOptions.From(catalog))...); err != nil {
		log.WithError(err).Error("failed to delete the service catalog")

		return nil, err
	}

	return catalog, nil
}

// Get returns the service catalog
func (p serviceCatalogsImpl) Get(ctx context.Context, name string) (*servicesv1.ServiceCatalog, error) {
	catalog := &servicesv1.ServiceCatalog{}

	if found, err := p.Has(ctx, name); err != nil {
		return nil, err
	} else if !found {
		return nil, ErrNotFound
	}

	return catalog, p.Store().Client().Get(ctx,
		store.GetOptions.InNamespace(HubNamespace),
		store.GetOptions.WithName(name),
		store.GetOptions.InTo(catalog),
	)
}

// List returns the existing service catalogs
func (p serviceCatalogsImpl) List(ctx context.Context, filters ...func(servicesv1.ServiceCatalog) bool) (*servicesv1.ServiceCatalogList, error) {
	list := &servicesv1.ServiceCatalogList{}

	err := p.Store().Client().List(ctx,
		store.ListOptions.InNamespace(HubNamespace),
		store.ListOptions.InTo(list),
	)
	if err != nil {
		return nil, err
	}

	if len(filters) == 0 {
		return list, nil
	}

	res := []servicesv1.ServiceCatalog{}
	for _, item := range list.Items {
		if func() bool {
			for _, filter := range filters {
				if !filter(item) {
					return false
				}
			}
			return true
		}() {
			res = append(res, item)
		}
	}
	list.Items = res

	return list, nil
}

// Has checks if a service catalog exists
func (p serviceCatalogsImpl) Has(ctx context.Context, name string) (bool, error) {
	return p.Store().Client().Has(ctx,
		store.HasOptions.InNamespace(HubNamespace),
		store.HasOptions.From(&servicesv1.ServiceCatalog{}),
		store.HasOptions.WithName(name),
	)
}
