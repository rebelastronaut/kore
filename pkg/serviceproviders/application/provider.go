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

package application

import (
	"fmt"

	"github.com/appvia/kore/pkg/utils/jsonutils"

	"github.com/appvia/kore/pkg/utils"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ kore.ServiceProvider = Provider{}

const (
	Type               = "application"
	ServiceTypeApp     = "app"
	ServiceTypeHelmApp = "helm-app"
)

type Provider struct {
	name  string
	plans []servicesv1.ServicePlan
}

func (p Provider) Name() string {
	return p.name
}

func (p Provider) Catalog(ctx kore.Context, provider *servicesv1.ServiceProvider) (kore.ServiceProviderCatalog, error) {
	return kore.ServiceProviderCatalog{
		Kinds: []servicesv1.ServiceKind{
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       servicesv1.ServiceKindGVK.Kind,
					APIVersion: servicesv1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      ServiceTypeApp,
					Namespace: kore.HubNamespace,
					Labels: map[string]string{
						kore.Label("platform"): "Kubernetes",
					},
					Annotations: map[string]string{
						kore.AnnotationReadOnly: kore.AnnotationValueTrue,
					},
				},
				Spec: servicesv1.ServiceKindSpec{
					Type:                 ServiceTypeApp,
					DisplayName:          "Kubernetes Application",
					Summary:              "Kubernetes Application",
					Enabled:              false,
					ServiceAccessEnabled: false,
					Schema:               string(jsonutils.MustCompact([]byte(appSchemaV1))),
				},
			},
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       servicesv1.ServiceKindGVK.Kind,
					APIVersion: servicesv1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      ServiceTypeHelmApp,
					Namespace: kore.HubNamespace,
					Labels: map[string]string{
						kore.Label("platform"): "Kubernetes",
					},
				},
				Spec: servicesv1.ServiceKindSpec{
					Type:                 ServiceTypeHelmApp,
					DisplayName:          "Kubernetes Helm Application",
					Summary:              "Kubernetes Helm Application",
					Enabled:              true,
					ServiceAccessEnabled: false,
					Schema:               string(jsonutils.MustCompact([]byte(helmAppSchemaV1))),
				},
			},
		},
		Plans: p.plans,
	}, nil
}

func (p Provider) AdminServices() []servicesv1.Service {
	cluster := corev1.Ownership{
		Group:     clustersv1.ClusterGVK.Group,
		Version:   clustersv1.ClusterGVK.Version,
		Kind:      clustersv1.ClusterGVK.Kind,
		Namespace: "kore-admin",
		Name:      "kore",
	}

	providerDeps := []string{"app-" + kore.AppAppManager, "app-" + kore.AppHelmOperator}

	var services []servicesv1.Service
	for _, servicePlan := range p.plans {
		if !utils.Contains(servicePlan.Name, providerDeps) {
			continue
		}

		services = append(services, CreateSystemServiceFromPlan(servicePlan, cluster, servicePlan.Name, kore.HubAdminTeam))
	}
	return services
}

func (p Provider) ReconcileCredentials(
	ctx kore.Context,
	service *servicesv1.Service,
	creds *servicesv1.ServiceCredentials,
) (reconcile.Result, map[string]string, error) {
	return reconcile.Result{}, nil, fmt.Errorf("can not create credentials for kubernetes services")
}

func (p Provider) DeleteCredentials(
	ctx kore.Context,
	service *servicesv1.Service,
	creds *servicesv1.ServiceCredentials,
) (reconcile.Result, error) {
	return reconcile.Result{}, fmt.Errorf("can not create credentials for kubernetes services")
}
