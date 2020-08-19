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

package dummy

import (
	"strings"

	"github.com/appvia/kore/pkg/utils/jsonutils"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

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
				Schema:               string(jsonutils.MustCompact([]byte(planSchemaV1))),
				CredentialSchema:     string(jsonutils.MustCompact([]byte(credentialSchemaV1))),
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
				Schema:               string(jsonutils.MustCompact([]byte(planSchemaV1))),
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
				Schema:               string(jsonutils.MustCompact([]byte(planSchemaV1))),
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
				Kind:             serviceKind.Name,
				DisplayName:      "Default",
				Summary:          serviceKind.Labels[kore.Label("platform")] + " dummy service plan used for testing",
				Description:      "Testing, testing, 1, 2, 3",
				Schema:           serviceKind.Spec.Schema,
				CredentialSchema: serviceKind.Spec.CredentialSchema,
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
