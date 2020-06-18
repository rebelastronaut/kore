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
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/template"

	"github.com/appvia/kore/pkg/utils"
	"github.com/appvia/kore/pkg/utils/jsonutils"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/store"
	"github.com/appvia/kore/pkg/utils/jsonschema"
	"github.com/appvia/kore/pkg/utils/validation"

	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
)

// ServicePlans is the interface to manage service plans
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ServicePlans
type ServicePlans interface {
	// Delete is used to delete a service plan in the kore
	Delete(ctx context.Context, name string, ignoreReadonly bool) (*servicesv1.ServicePlan, error)
	// Get returns the service plan
	Get(context.Context, string) (*servicesv1.ServicePlan, error)
	// GetSchema returns the service plan schema
	GetSchema(context.Context, string) (string, error)
	// GetSchemaForCluster returns the service plan schema generated for the given cluster
	GetSchemaForCluster(ctx context.Context, servicePlanName, team, clusterName string) (string, error)
	// GetCredentialSchema returns the service credential schema for the given plan
	GetCredentialSchema(context.Context, string) (string, error)
	// List returns the existing service plans
	List(context.Context) (*servicesv1.ServicePlanList, error)
	// ListFiltered returns a list of service plans using the given filter.
	ListFiltered(context.Context, func(servicesv1.ServicePlan) bool) (*servicesv1.ServicePlanList, error)
	// Has checks if a service plan exists
	Has(context.Context, string) (bool, error)
	// Update is responsible for updating a service plan
	Update(ctx context.Context, plan *servicesv1.ServicePlan, ignoreReadonly bool) error
	// GetEditablePlanParams returns with the editable service plan parameters for a specific team and service kind
	GetEditablePlanParams(ctx context.Context, team string, clusterKind string) (map[string]bool, error)
}

type servicePlansImpl struct {
	Interface
}

// Update is responsible for updating a service plan
func (p servicePlansImpl) Update(ctx context.Context, plan *servicesv1.ServicePlan, ignoreReadonly bool) error {
	if err := IsValidResourceName("plan", plan.Name); err != nil {
		return err
	}

	if !strings.HasPrefix(plan.Name, plan.Spec.Kind+"-") {
		return validation.NewError("%q failed validation", plan.Name).
			WithFieldErrorf("name", validation.InvalidValue, "must start with %s-", plan.Spec.Kind)
	}

	if plan.Namespace != HubNamespace {
		return validation.NewError("%q failed validation", plan.Name).
			WithFieldErrorf("namespace", validation.InvalidValue, "must be %q", HubNamespace)
	}

	existing, err := p.Get(ctx, plan.Name)
	if err != nil && err != ErrNotFound {
		return fmt.Errorf("failed to get plan %q: %w", plan.Name, err)
	}

	if !ignoreReadonly {
		if existing != nil && existing.Annotations[AnnotationReadOnly] == AnnotationValueTrue {
			return validation.NewError("the plan can not be updated").WithFieldError(validation.FieldRoot, validation.ReadOnly, "plan is read-only")
		}
		if plan.Annotations[AnnotationReadOnly] == AnnotationValueTrue {
			return validation.NewError("the plan can not be updated").WithFieldError(validation.FieldRoot, validation.ReadOnly, "read-only flag can not be set")
		}
	}

	if existing != nil {
		verr := validation.NewError("%q failed validation", plan.Name)
		if existing.Spec.Kind != plan.Spec.Kind {
			verr.AddFieldErrorf("kind", validation.ReadOnly, "can not be changed after the service plan was created")
		}
		if verr.HasErrors() {
			return verr
		}
	}

	schema := plan.Spec.Schema
	if schema == "" {
		kind, err := p.ServiceKinds().Get(ctx, plan.Spec.Kind)
		if err != nil {
			return fmt.Errorf("failed to get service kind %q: %w", plan.Spec.Kind, err)
		}
		schema = kind.Spec.Schema
	}

	if schema == "" && !utils.ApiExtJSONEmpty(plan.Spec.Configuration) {
		if existing == nil || !utils.ApiExtJSONEquals(plan.Spec.Configuration, existing.Spec.Configuration) {
			return validation.NewError("%q failed validation", plan.Name).
				WithFieldErrorf(
					"configuration",
					validation.ReadOnly,
					"the service provider doesn't have a JSON schema to validate the configuration",
				)
		}
	}

	if schema != "" {
		if err := jsonschema.Validate(schema, "plan", plan.Spec.Configuration); err != nil {
			return err
		}
	}

	err = p.Store().Client().Update(ctx,
		store.UpdateOptions.To(plan),
		store.UpdateOptions.WithCreate(true),
		store.UpdateOptions.WithForce(true),
	)
	if err != nil {
		log.WithError(err).Error("failed to update a service plan in the kore")

		return err
	}

	return nil
}

// Delete is used to delete a service plan in the kore
func (p servicePlansImpl) Delete(ctx context.Context, name string, ignoreReadonly bool) (*servicesv1.ServicePlan, error) {
	plan := &servicesv1.ServicePlan{}
	err := p.Store().Client().Get(ctx,
		store.GetOptions.InNamespace(HubNamespace),
		store.GetOptions.InTo(plan),
		store.GetOptions.WithName(name),
	)
	if err != nil {
		if kerrors.IsNotFound(err) {
			return nil, ErrNotFound
		}
		log.WithError(err).Error("failed to retrieve the service plan")

		return nil, err
	}

	if !ignoreReadonly {
		if plan.Annotations[AnnotationReadOnly] == AnnotationValueTrue {
			return nil, validation.NewError("the service plan can not be deleted").
				WithFieldError(validation.FieldRoot, validation.ReadOnly, "service plan is read-only")
		}
	}

	servicesWithPlan, err := p.getServicesWithPlan(ctx, name)
	if err != nil {
		return nil, err
	}
	if len(servicesWithPlan) > 0 {
		if len(servicesWithPlan) <= 5 {
			return nil, fmt.Errorf(
				"the service plan can not be deleted as there are %d services using it: %s",
				len(servicesWithPlan),
				strings.Join(servicesWithPlan, ", "),
			)
		}
		return nil, fmt.Errorf(
			"the service plan can not be deleted as there are %d services using it",
			len(servicesWithPlan),
		)
	}

	if err := p.Store().Client().Delete(ctx, store.DeleteOptions.From(plan)); err != nil {
		log.WithError(err).Error("failed to delete the service plan")

		return nil, err
	}

	return plan, nil
}

// Get returns the service plan
func (p servicePlansImpl) Get(ctx context.Context, name string) (*servicesv1.ServicePlan, error) {
	plan := &servicesv1.ServicePlan{}

	if found, err := p.Has(ctx, name); err != nil {
		return nil, err
	} else if !found {
		return nil, ErrNotFound
	}

	return plan, p.Store().Client().Get(ctx,
		store.GetOptions.InNamespace(HubNamespace),
		store.GetOptions.WithName(name),
		store.GetOptions.InTo(plan),
	)
}

// GetSchema returns the service plan schema
func (p servicePlansImpl) GetSchema(ctx context.Context, name string) (string, error) {
	plan, err := p.Get(ctx, name)
	if err != nil {
		return "", err
	}
	if plan.Spec.Schema != "" {
		return plan.Spec.Schema, nil
	}

	kind, err := p.ServiceKinds().Get(ctx, plan.Spec.Kind)
	if err != nil {
		return "", err
	}
	return kind.Spec.Schema, nil
}

// GetSchemaForCluster returns the service plan schema generated for the given cluster
func (p servicePlansImpl) GetSchemaForCluster(ctx context.Context, servicePlanName, team, clusterName string) (string, error) {
	cluster, err := p.Teams().Team(team).Clusters().Get(ctx, clusterName)
	if err != nil {
		return "", err
	}

	schema, err := p.GetSchema(ctx, servicePlanName)
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("schema").Parse(schema)
	if err != nil {
		return "", fmt.Errorf("failed to parse plan schema: %w", err)
	}

	clusterObj, err := jsonutils.ToMap(cluster)
	if err != nil {
		return "", fmt.Errorf("failed to encode cluster: %w", err)
	}

	tmplBuf := bytes.NewBuffer(make([]byte, 0, 16384))
	params := map[string]interface{}{
		"cluster": clusterObj,
	}
	if err := tmpl.Execute(tmplBuf, params); err != nil {
		return "", fmt.Errorf("failed to compile plan schema: %w", err)
	}

	return tmplBuf.String(), nil
}

// GetCredentialSchema returns the service credential schema for the given plan
func (p servicePlansImpl) GetCredentialSchema(ctx context.Context, name string) (string, error) {
	plan, err := p.Get(ctx, name)
	if err != nil {
		return "", err
	}
	if plan.Spec.CredentialSchema != "" {
		return plan.Spec.CredentialSchema, nil
	}

	kind, err := p.ServiceKinds().Get(ctx, plan.Spec.Kind)
	if err != nil {
		return "", err
	}
	return kind.Spec.CredentialSchema, nil
}

// List returns the existing service plans
func (p servicePlansImpl) List(ctx context.Context) (*servicesv1.ServicePlanList, error) {
	plans := &servicesv1.ServicePlanList{}

	return plans, p.Store().Client().List(ctx,
		store.ListOptions.InNamespace(HubNamespace),
		store.ListOptions.InTo(plans),
	)
}

// ListFiltered returns a list of service plans using the given filter.
// A service plan is included if the filter function returns true
func (p servicePlansImpl) ListFiltered(ctx context.Context, filter func(plan servicesv1.ServicePlan) bool) (*servicesv1.ServicePlanList, error) {
	res := []servicesv1.ServicePlan{}

	servicePlansList, err := p.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, servicePlan := range servicePlansList.Items {
		if filter(servicePlan) {
			res = append(res, servicePlan)
		}
	}

	servicePlansList.Items = res

	return servicePlansList, nil
}

// Has checks if a service plan exists
func (p servicePlansImpl) Has(ctx context.Context, name string) (bool, error) {
	return p.Store().Client().Has(ctx,
		store.HasOptions.InNamespace(HubNamespace),
		store.HasOptions.From(&servicesv1.ServicePlan{}),
		store.HasOptions.WithName(name),
	)
}

// GetEditablePlanParams returns with the editable service plan parameters for a specific team and service kind
func (p servicePlansImpl) GetEditablePlanParams(ctx context.Context, team string, clusterKind string) (map[string]bool, error) {
	// TODO: read this from the allocated service plan policies
	return map[string]bool{"*": true}, nil
}

func (p servicePlansImpl) getServicesWithPlan(ctx context.Context, clusterName string) ([]string, error) {
	var res []string

	teamList, err := p.Teams().List(ctx)
	if err != nil {
		return nil, err
	}

	for _, team := range teamList.Items {
		servicesList, err := p.Teams().Team(team.Name).Services().List(ctx)
		if err != nil {
			return nil, err
		}
		for _, service := range servicesList.Items {
			if service.Spec.Plan == clusterName {
				res = append(res, fmt.Sprintf("%s/%s", team.Name, service.Name))
			}
		}
	}

	return res, nil
}
