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

	"github.com/appvia/kore/pkg/store"
	"github.com/appvia/kore/pkg/utils/validation"
	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
)

// ServiceDeployments is the interface to manage service deployments
type ServiceDeployments interface {
	// CheckDelete verifies whether the service deployment can be deleted
	CheckDelete(context.Context, *servicesv1.ServiceDeployment, ...DeleteOptionFunc) error
	// Delete is used to delete a service deployment in kore
	Delete(context.Context, string, ...DeleteOptionFunc) (*servicesv1.ServiceDeployment, error)
	// Get returns the service deployment
	Get(context.Context, string) (*servicesv1.ServiceDeployment, error)
	// List returns the existing service deployments
	// The optional filter functions can be used to include items only for which all functions return true
	List(context.Context, ...func(servicesv1.ServiceDeployment) bool) (*servicesv1.ServiceDeploymentList, error)
	// Has checks if a service deployment exists
	Has(context.Context, string) (bool, error)
	// Update is responsible for updating a service deployment
	Update(_ context.Context, _ *servicesv1.ServiceDeployment, allowSystemServices bool) error
}

type serviceDeploymentsImpl struct {
	Interface
	// team is the name
	team string
}

// Update is responsible for updating a service deployment
func (s serviceDeploymentsImpl) Update(ctx context.Context, deployment *servicesv1.ServiceDeployment, allowSystemServices bool) error {
	if err := IsValidResourceName("service deployment", deployment.Name); err != nil {
		return err
	}

	if deployment.Namespace != s.team {
		return validation.NewError("%q failed validation", deployment.Name).
			WithFieldErrorf("namespace", validation.InvalidValue, "must be %q", s.team)
	}

	if s.team != HubAdminTeam {
		if len(deployment.Spec.ClusterSelector.Teams) != 1 || deployment.Spec.ClusterSelector.Teams[0] != s.team {
			return validation.NewError("%q failed validation", deployment.Name).
				WithFieldErrorf("spec.clusterSelector.Teams", validation.InvalidValue, "must only reference %q team", s.team)
		}
	}

	plan, err := s.ServicePlans().Get(ctx, deployment.Spec.Plan)
	if err != nil {
		if err == ErrNotFound {
			return validation.NewError("%q failed validation", deployment.Name).
				WithFieldErrorf("spec.plan", validation.MustExist, "%q does not exist", deployment.Spec.Plan)
		}
		return err
	}

	if plan.Spec.Kind != deployment.Spec.Kind {
		return validation.NewError("%q failed validation", deployment.Name).
			WithFieldErrorf("spec.plan", validation.InvalidType, "deployment has kind %q, but plan has %q", deployment.Spec.Kind, plan.Spec.Kind)
	}

	if !allowSystemServices {
		if plan.Annotations[AnnotationSystem] == AnnotationValueTrue {
			return validation.NewError("%q failed validation", deployment.Name).
				WithFieldError("spec.plan", validation.InvalidType, "system plans can not be used to create new services")
		}
	}

	err = s.Store().Client().Update(ctx,
		store.UpdateOptions.To(deployment),
		store.UpdateOptions.WithCreate(true),
		store.UpdateOptions.WithForce(true),
	)
	if err != nil {
		log.WithError(err).Error("failed to update a service deployment")

		return err
	}

	return nil
}

// CheckDelete verifies whether the service deployment can be deleted
func (s serviceDeploymentsImpl) CheckDelete(ctx context.Context, serviceDeployment *servicesv1.ServiceDeployment, o ...DeleteOptionFunc) error {
	return nil
}

// Delete is used to delete a service deployment in kore
func (s serviceDeploymentsImpl) Delete(ctx context.Context, name string, o ...DeleteOptionFunc) (*servicesv1.ServiceDeployment, error) {
	opts := ResolveDeleteOptions(o)

	deployment := &servicesv1.ServiceDeployment{}
	err := s.Store().Client().Get(ctx,
		store.GetOptions.InNamespace(s.team),
		store.GetOptions.InTo(deployment),
		store.GetOptions.WithName(name),
	)
	if err != nil {
		if kerrors.IsNotFound(err) {
			return nil, ErrNotFound
		}
		log.WithError(err).Error("failed to retrieve the service deployment")

		return nil, err
	}

	if err := opts.Check(deployment, func(o ...DeleteOptionFunc) error { return s.CheckDelete(ctx, deployment, o...) }); err != nil {
		return nil, err
	}

	if err := s.Store().Client().Delete(ctx, append(opts.StoreOptions(), store.DeleteOptions.From(deployment))...); err != nil {
		log.WithError(err).Error("failed to delete the service deployment")

		return nil, err
	}

	return deployment, nil
}

// Get returns the service deployment
func (s serviceDeploymentsImpl) Get(ctx context.Context, name string) (*servicesv1.ServiceDeployment, error) {
	deployment := &servicesv1.ServiceDeployment{}

	if found, err := s.Has(ctx, name); err != nil {
		return nil, err
	} else if !found {
		return nil, ErrNotFound
	}

	return deployment, s.Store().Client().Get(ctx,
		store.GetOptions.InNamespace(s.team),
		store.GetOptions.WithName(name),
		store.GetOptions.InTo(deployment),
	)
}

// List returns the existing service deployments
func (s serviceDeploymentsImpl) List(ctx context.Context, filters ...func(servicesv1.ServiceDeployment) bool) (*servicesv1.ServiceDeploymentList, error) {
	list := &servicesv1.ServiceDeploymentList{}

	err := s.Store().Client().List(ctx,
		store.ListOptions.InNamespace(s.team),
		store.ListOptions.InTo(list),
	)
	if err != nil {
		return nil, err
	}

	if len(filters) == 0 {
		return list, nil
	}

	res := []servicesv1.ServiceDeployment{}
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

// Has checks if a service deployment exists
func (s serviceDeploymentsImpl) Has(ctx context.Context, name string) (bool, error) {
	return s.Store().Client().Has(ctx,
		store.HasOptions.InNamespace(s.team),
		store.HasOptions.From(&servicesv1.ServiceDeployment{}),
		store.HasOptions.WithName(name),
	)
}
