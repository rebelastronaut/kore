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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/serviceproviders/application"
	"github.com/appvia/kore/pkg/utils/httputils"
	"github.com/appvia/kore/pkg/utils/jsonschema"
	"github.com/appvia/kore/pkg/utils/kubernetes"
	"github.com/ghodss/yaml"
	"github.com/tidwall/sjson"
	"golang.org/x/mod/semver"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type catalogComponent struct {
	Catalog *servicesv1.ServiceCatalog
}

func newCatalogComponent(catalog *servicesv1.ServiceCatalog) *catalogComponent {
	return &catalogComponent{
		Catalog: catalog,
	}
}

func (c *catalogComponent) Reconcile(ctx kore.Context) (reconcile.Result, error) {
	httpClient := httputils.DefaultHTTPClient

	indexURL := fmt.Sprintf("%s/%s", strings.TrimRight(c.Catalog.Spec.URL, "/"), "index.yaml")

	resp, err := httpClient.Get(indexURL)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to fetch URL: %s: %w", indexURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return reconcile.Result{}, fmt.Errorf("URL %s returned non-200 response code: %d", indexURL, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to read repository index %s: %w", indexURL, err)
	}

	repoIndex := &RepoIndex{}
	if err := yaml.Unmarshal(body, repoIndex); err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to decode repository index %s: %w", indexURL, err)
	}

	helmAppKind, err := ctx.Kore().ServiceKinds().Get(ctx, application.ServiceTypeHelmApp)
	if err != nil {
		if err == kore.ErrNotFound {
			return reconcile.Result{RequeueAfter: 5 * time.Second}, nil
		}
		return reconcile.Result{}, fmt.Errorf("failed to get helm app service kind: %w", err)
	}

	schema := helmAppKind.Spec.Schema

	for chartName, chartVersions := range repoIndex.Entries {
		var latest *ChartVersion
		for _, chartVersion := range chartVersions {
			if latest == nil || semver.Compare(chartVersion.Version, latest.Version) == 1 {
				latest = &chartVersion
			}
		}

		if latest == nil {
			ctx.Logger().Info("Skipping chart %q as it doesn't have any valid versions", chartName)
			continue
		}

		chartRawData, _ := json.Marshal(latest)

		schema, err = c.updateSchema(schema, *latest)
		if err != nil {
			ctx.Logger().WithError(err).Errorf("%s/%s has an invalid schema", latest.Name, latest.Version)
		}

		annotations, labels := c.annotationsAndLabelsFromChart(*latest)

		serviceKind := &servicesv1.ServiceKind{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ServiceKind",
				APIVersion: servicesv1.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:        c.Catalog.Spec.ServiceKindPrefix + chartName,
				Namespace:   kore.HubNamespace,
				Annotations: annotations,
				Labels:      labels,
			},
			Spec: servicesv1.ServiceKindSpec{
				Type:             application.ServiceTypeHelmApp,
				Enabled:          true,
				DisplayName:      latest.Annotations[kore.Label("displayName")],
				Summary:          latest.Description,
				Description:      latest.Annotations[kore.Label("longDescription")],
				ImageURL:         latest.Icon,
				DocumentationURL: latest.Home,
				Schema:           schema,
				ProviderData:     &apiextv1.JSON{Raw: chartRawData},
			},
		}

		kubernetes.EnsureOwnerReference(serviceKind, c.Catalog, true)

		existing, err := ctx.Kore().ServiceKinds().Get(ctx, serviceKind.Name)
		if err != nil && err != kore.ErrNotFound {
			return reconcile.Result{}, fmt.Errorf("failed to get service kind %q: %w", serviceKind.Name, err)
		}

		if existing != nil && !existing.Spec.Enabled {
			serviceKind.Spec.Enabled = false
		}

		if err := ctx.Kore().ServiceKinds().Update(ctx, serviceKind); err != nil {
			return reconcile.Result{}, fmt.Errorf("failed to create or update service kind %q: %w", serviceKind.Name, err)
		}

		for _, chartVersion := range chartVersions {
			if strings.TrimSpace(chartVersion.Version) == "" ||
				len(chartVersion.URLs) == 0 ||
				strings.TrimSpace(chartVersion.URLs[0]) == "" {
				continue
			}

			chartRawData, _ := json.Marshal(chartVersion)

			helmAppConfig := &application.HelmAppV1{
				Source: application.HelmAppV1Source{
					Helm: &application.HelmAppV1Helm{
						URL:     c.Catalog.Spec.URL,
						Name:    chartName,
						Version: chartVersion.Version,
					},
				},
			}

			chartVersionSchema, err := c.updateSchema(schema, chartVersion)
			if err != nil {
				ctx.Logger().WithError(err).Errorf("%s/%s has an invalid schema", chartVersion.Name, chartVersion.Version)
			}

			annotations, labels := c.annotationsAndLabelsFromChart(chartVersion)

			servicePlan := &servicesv1.ServicePlan{
				TypeMeta: metav1.TypeMeta{
					Kind:       "ServicePlan",
					APIVersion: servicesv1.GroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:        serviceKind.Name + "-" + chartVersion.Version,
					Namespace:   kore.HubNamespace,
					Annotations: annotations,
					Labels:      labels,
				},
				Spec: servicesv1.ServicePlanSpec{
					Kind:         serviceKind.Name,
					DisplayName:  chartVersion.Version,
					Summary:      chartVersion.Description,
					Description:  chartVersion.Annotations[kore.Label("longDescription")],
					Schema:       chartVersionSchema,
					ProviderData: &apiextv1.JSON{Raw: chartRawData},
				},
			}

			if err := servicePlan.Spec.SetConfiguration(helmAppConfig); err != nil {
				return reconcile.Result{}, fmt.Errorf("failed to construct service plan %q: %w", servicePlan.Name, err)
			}

			kubernetes.EnsureOwnerReference(servicePlan, serviceKind, true)

			if err := ctx.Kore().ServicePlans().Update(ctx, servicePlan, true); err != nil {
				return reconcile.Result{}, fmt.Errorf("failed to create or update service plan %q: %w", servicePlan.Name, err)
			}
		}
	}

	return reconcile.Result{}, nil
}

func (c *catalogComponent) Delete(ctx kore.Context) (reconcile.Result, error) {
	serviceKinds, err := ctx.Kore().ServiceKinds().List(ctx, func(sk servicesv1.ServiceKind) bool {
		return kubernetes.HasOwnerReference(&sk, c.Catalog)
	})
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to list service kinds created by the service catalog: %w", err)
	}

	for _, sk := range serviceKinds.Items {
		if _, err := ctx.Kore().ServiceKinds().Delete(ctx, sk.Name); err != nil {
			return reconcile.Result{}, fmt.Errorf("failed to delete service kind %q: %w", sk.Name, err)
		}
	}

	return reconcile.Result{}, nil
}

func (c *catalogComponent) updateSchema(schema string, chart ChartVersion) (string, error) {
	chartSchema := chart.Annotations[kore.Label("schema")]
	if chartSchema == "" {
		return schema, nil
	}

	schemaObj := jsonschema.MetaSchemaDraft7Ext{}
	if err := json.Unmarshal([]byte(chartSchema), &schemaObj); err != nil {
		return schema, err
	}

	if schemaObj.Type != "object" {
		return schema, errors.New("schema must define an object type for the chart values")
	}

	updated, err := sjson.SetRaw(schema, "properties.values", chartSchema)
	if err != nil {
		return schema, err
	}

	err = jsonschema.Validate(jsonschema.MetaSchemaDraft07, "chart schema", updated)
	if err != nil {
		return schema, err
	}

	return updated, nil
}

func (c *catalogComponent) annotationsAndLabelsFromChart(chart ChartVersion) (map[string]string, map[string]string) {
	annotations := map[string]string{}
	labels := map[string]string{}
	for k, v := range chart.Annotations {
		if !strings.HasPrefix(k, kore.Label("")) {
			continue
		}
		switch k {
		case kore.Label("schema"), kore.Label("displayName"), kore.Label("longDescription"):
		case kore.AnnotationSystem, kore.AnnotationReadOnly:
			annotations[k] = v
		default:
			labels[k] = v
		}
	}

	annotations[kore.AnnotationReadOnly] = kore.AnnotationValueTrue

	return annotations, labels
}
