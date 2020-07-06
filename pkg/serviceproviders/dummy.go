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
	"strings"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func init() {
	kore.RegisterServiceProviderFactory(DummyFactory{})
}

const dummySchema = `{
	"$id": "https://appvia.io/kore/schemas/services/dummy/dummy.json",
	"$schema": "http://json-schema.org/draft-07/schema#",
	"description": "Dummy service plan schema",
	"type": "object",
	"additionalProperties": false,
	"required": [
		"foo"
	],
	"properties": {
		"foo": {
			"type": "string",
			"minLength": 1
		}
	}
}`

const dummyCredentialSchema = `{
	"$id": "https://appvia.io/kore/schemas/services/dummy/dummy-credentials.json",
	"$schema": "http://json-schema.org/draft-07/schema#",
	"description": "Dummy service plan credentials schema",
	"type": "object",
	"additionalProperties": false,
	"required": [
		"bar"
	],
	"properties": {
		"bar": {
			"type": "string",
			"minLength": 1
		}
	}
}`

type DummyFactory struct{}

func (d DummyFactory) Type() string {
	return "dummy"
}

func (d DummyFactory) JSONSchema() string {
	return `{
		"$id": "https://appvia.io/kore/schemas/serviceprovider/dummy.json",
		"$schema": "http://json-schema.org/draft-07/schema#",
		"description": "Dummy service plan schema",
		"type": "object",
		"additionalProperties": false,
		"required": [
			"iAmDummy"
		],
		"properties": {
			"iAmDummy": {
				"type": "string",
				"minLength": 1
			}
		}
	}`
}

func (d DummyFactory) Create(ctx kore.Context, serviceProvider *servicesv1.ServiceProvider) (kore.ServiceProvider, error) {
	return Dummy{name: serviceProvider.Name}, nil
}

func (d DummyFactory) SetUp(_ kore.Context, _ *servicesv1.ServiceProvider) (complete bool, _ error) {
	return true, nil
}

func (d DummyFactory) TearDown(_ kore.Context, _ *servicesv1.ServiceProvider) (complete bool, _ error) {
	return true, nil
}

func (d DummyFactory) DefaultProviders() []servicesv1.ServiceProvider {
	return nil
}

var _ kore.ServiceProvider = Dummy{}

type Dummy struct {
	name string
}

func (d Dummy) Name() string {
	return d.name
}

func (d Dummy) Catalog(_ kore.Context, _ *servicesv1.ServiceProvider) (kore.ServiceProviderCatalog, error) {
	return kore.ServiceProviderCatalog{
		Plans: d.plans(),
		Kinds: d.kinds(),
	}, nil
}

func (d Dummy) kinds() []servicesv1.ServiceKind {
	var serviceKinds []servicesv1.ServiceKind

	for _, platform := range []string{"AWS", "GCP", "Azure", "Kubernetes"} {
		serviceKinds = append(serviceKinds, servicesv1.ServiceKind{
			TypeMeta: metav1.TypeMeta{
				Kind:       servicesv1.ServiceKindGVK.Kind,
				APIVersion: servicesv1.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "dummy-" + strings.ToLower(platform),
				Namespace: kore.HubNamespace,
				Labels: map[string]string{
					kore.Label("platform"): platform,
				},
			},
			Spec: servicesv1.ServiceKindSpec{
				DisplayName:          platform + " Dummy",
				Summary:              platform + " dummy service kind used for testing",
				Enabled:              true,
				ServiceAccessEnabled: true,
				Schema:               dummySchema,
				CredentialSchema:     dummyCredentialSchema,
			},
		})

		serviceKinds = append(serviceKinds, servicesv1.ServiceKind{
			TypeMeta: metav1.TypeMeta{
				Kind:       servicesv1.ServiceKindGVK.Kind,
				APIVersion: servicesv1.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "dummy-" + strings.ToLower(platform) + "-no-schema",
				Namespace: kore.HubNamespace,
				Labels: map[string]string{
					kore.Label("platform"): platform,
				},
			},
			Spec: servicesv1.ServiceKindSpec{
				DisplayName:          platform + " Dummy (no schema)",
				Summary:              platform + " dummy service kind used for testing",
				Enabled:              true,
				ServiceAccessEnabled: true,
			},
		})

		serviceKinds = append(serviceKinds, servicesv1.ServiceKind{
			TypeMeta: metav1.TypeMeta{
				Kind:       servicesv1.ServiceKindGVK.Kind,
				APIVersion: servicesv1.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "dummy-" + strings.ToLower(platform) + "-no-creds-schema",
				Namespace: kore.HubNamespace,
				Labels: map[string]string{
					kore.Label("platform"): platform,
				},
			},
			Spec: servicesv1.ServiceKindSpec{
				DisplayName:          platform + " Dummy (no creds schema)",
				Summary:              platform + " dummy service kind used for testing",
				Enabled:              true,
				ServiceAccessEnabled: true,
				Schema:               dummySchema,
			},
		})

		serviceKinds = append(serviceKinds, servicesv1.ServiceKind{
			TypeMeta: metav1.TypeMeta{
				Kind:       servicesv1.ServiceKindGVK.Kind,
				APIVersion: servicesv1.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "dummy-" + strings.ToLower(platform) + "-no-access",
				Namespace: kore.HubNamespace,
				Labels: map[string]string{
					kore.Label("platform"): platform,
				},
			},
			Spec: servicesv1.ServiceKindSpec{
				DisplayName:          platform + " Dummy (no access)",
				Summary:              platform + " dummy service kind used for testing",
				Enabled:              true,
				ServiceAccessEnabled: false,
				Schema:               dummySchema,
			},
		})
	}

	return serviceKinds
}

func (d Dummy) plans() []servicesv1.ServicePlan {
	var servicePlans []servicesv1.ServicePlan

	for _, serviceKind := range d.kinds() {
		servicePlan := servicesv1.ServicePlan{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ServicePlan",
				APIVersion: servicesv1.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      serviceKind.Name + "-default",
				Namespace: "kore",
			},
			Spec: servicesv1.ServicePlanSpec{
				Kind:        serviceKind.Name,
				DisplayName: "Default",
				Summary:     serviceKind.Labels[kore.Label("platform")] + " dummy service plan used for testing",
				Description: "Testing, testing, 1, 2, 3",
			},
		}
		if serviceKind.Spec.Schema != "" {
			servicePlan.Spec.Configuration = &v1beta1.JSON{Raw: []byte(`{"foo":"bar"}`)}
		}
		servicePlans = append(servicePlans, servicePlan)

		if serviceKind.Spec.ServiceAccessEnabled {
			servicePlanNoAccess := servicePlan.DeepCopy()
			servicePlanNoAccess.Name = serviceKind.Name + "-no-service-access"
			servicePlanNoAccess.Spec.DisplayName = "No service access"
			servicePlanNoAccess.Spec.ServiceAccessDisabled = true
			servicePlans = append(servicePlans, *servicePlanNoAccess)
		}
	}

	return servicePlans
}

func (d Dummy) AdminServices() []servicesv1.Service {
	return nil
}

func (d Dummy) Reconcile(
	ctx kore.Context,
	service *servicesv1.Service,
) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}

func (d Dummy) Delete(
	ctx kore.Context,
	service *servicesv1.Service,
) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}

func (d Dummy) ReconcileCredentials(
	ctx kore.Context,
	service *servicesv1.Service,
	creds *servicesv1.ServiceCredentials,
) (reconcile.Result, map[string]string, error) {
	res := map[string]string{
		"superSecret": creds.Name + "-secret",
	}
	return reconcile.Result{}, res, nil
}

func (d Dummy) DeleteCredentials(
	ctx kore.Context,
	service *servicesv1.Service,
	creds *servicesv1.ServiceCredentials,
) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}
