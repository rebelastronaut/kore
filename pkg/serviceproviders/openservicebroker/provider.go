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

package openservicebroker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/tidwall/sjson"

	"github.com/appvia/kore/pkg/utils"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils/jsonschema"

	osb "github.com/kubernetes-sigs/go-open-service-broker-client/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	MetadataKeyConfiguration    = "kore.appvia.io/configuration"
	MetadataKeyDisplayName      = "displayName"
	MetadataKeyImageURL         = "imageUrl"
	MetadataKeyDescription      = "longDescription"
	MetadataKeyDocumentationURL = "documentationUrl"
	ComponentProvision          = "Provision"
	ComponentUpdate             = "Update"
	ComponentDeprovision        = "Deprovision"
	ComponentBind               = "Bind"
	ComponentUnbind             = "Unbind"
)

var _ kore.ServiceProvider = &Provider{}

type Provider struct {
	name   string
	config ProviderConfiguration
	client osb.Client
}

// NewProvider creates a new service provider which is backed by an Open Service Broker API compatible HTTP service
func NewProvider(name string, config ProviderConfiguration, client osb.Client) *Provider {
	return &Provider{
		name:   name,
		config: config,
		client: client,
	}
}

func (p *Provider) Name() string {
	return p.name
}

func (p *Provider) Type() string {
	return "openservicebroker"
}

func (p *Provider) Catalog(ctx kore.Context, serviceProvider *servicesv1.ServiceProvider) (kore.ServiceProviderCatalog, error) {
	catalog := kore.ServiceProviderCatalog{}

	osbCatalog, err := p.client.GetCatalog()
	if err != nil {
		return kore.ServiceProviderCatalog{}, fmt.Errorf("failed to fetch catalog from service broker: %w", err)
	}

	if len(osbCatalog.Services) == 0 {
		return kore.ServiceProviderCatalog{}, fmt.Errorf("service broker returned an empty catalog")
	}

	for _, catalogService := range osbCatalog.Services {
		if len(p.config.IncludeKinds) > 0 && !utils.Contains(catalogService.Name, p.config.IncludeKinds) {
			continue
		}
		if utils.Contains(catalogService.Name, p.config.ExcludeKinds) {
			continue
		}

		if !kore.ResourceNameFilter.MatchString(catalogService.Name) {
			return kore.ServiceProviderCatalog{}, fmt.Errorf("%q service name is invalid, must match %s", catalogService.Name, kore.ResourceNameFilter.String())
		}

		summary := catalogService.Description
		if summary == "" {
			summary = strings.Title(catalogService.Name)
		}

		platform := p.config.PlatformMapping[catalogService.Name]
		if platform == "" {
			platform = p.config.PlatformMapping["*"]
		}

		serviceKind := servicesv1.ServiceKind{
			TypeMeta: metav1.TypeMeta{
				Kind:       servicesv1.ServiceKindGVK.Kind,
				APIVersion: servicesv1.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      catalogService.Name,
				Namespace: kore.HubNamespace,
				Labels: map[string]string{
					kore.Label("platform"): platform,
				},
			},
			Spec: servicesv1.ServiceKindSpec{
				Type:                 catalogService.Name,
				ServiceAccessEnabled: catalogService.Bindable,
				DisplayName:          getMetadataStringVal(catalogService.Metadata, MetadataKeyDisplayName, ""),
				Summary:              summary,
				Description:          getMetadataStringVal(catalogService.Metadata, MetadataKeyDescription, ""),
				ImageURL:             getMetadataStringVal(catalogService.Metadata, MetadataKeyImageURL, ""),
				DocumentationURL:     getMetadataStringVal(catalogService.Metadata, MetadataKeyDocumentationURL, ""),
			},
		}
		providerData := ServiceKindProviderData{ServiceID: catalogService.ID}

		for _, catalogPlan := range catalogService.Plans {
			if !kore.ResourceNameFilter.MatchString(catalogPlan.Name) {
				return kore.ServiceProviderCatalog{}, fmt.Errorf("%q plan name is invalid, must match %s", catalogPlan.Name, kore.ResourceNameFilter.String())
			}

			servicePlan, err := p.parseCatalogPlan(catalogService, catalogPlan)
			if err != nil {
				return kore.ServiceProviderCatalog{}, err
			}

			if utils.Contains(servicePlan.Name, p.config.DefaultPlans) {
				if servicePlan.Spec.Schema == "" {
					return kore.ServiceProviderCatalog{}, fmt.Errorf("%s plan does not have a schema for provisioning", servicePlan.Name)
				}
				if !p.config.AllowEmptyCredentialSchema {
					if catalogService.Bindable && (catalogPlan.Bindable == nil || *catalogPlan.Bindable) && servicePlan.Spec.CredentialSchema == "" {
						return kore.ServiceProviderCatalog{}, fmt.Errorf("%s plan does not have a schema for bind", servicePlan.Name)
					}
				}

				if providerData.DefaultPlanID != "" {
					return kore.ServiceProviderCatalog{}, fmt.Errorf("there are multiple default plans for the same service: %s", serviceKind.Name)
				}

				providerData.DefaultPlanID = catalogPlan.ID
				serviceKind.Spec.Schema = servicePlan.Spec.Schema
				serviceKind.Spec.CredentialSchema = servicePlan.Spec.CredentialSchema
			}

			if len(p.config.IncludePlans) > 0 && !utils.Contains(servicePlan.Name, p.config.IncludePlans) {
				continue
			}
			if utils.Contains(servicePlan.Name, p.config.ExcludePlans) {
				continue
			}

			catalog.Plans = append(catalog.Plans, *servicePlan)
		}

		if err := serviceKind.Spec.SetProviderData(providerData); err != nil {
			return kore.ServiceProviderCatalog{}, err
		}

		catalog.Kinds = append(catalog.Kinds, serviceKind)
	}

	return catalog, nil
}

func (p Provider) AdminServices() []servicesv1.Service {
	return nil
}

func (p Provider) parseCatalogPlan(service osb.Service, catalogPlan osb.Plan) (*servicesv1.ServicePlan, error) {
	schemaStr, err := getPlanSchema(catalogPlan)
	if err != nil {
		return nil, err
	}

	credentialsSchemaStr, err := getCredentialsSchema(catalogPlan)
	if err != nil {
		return nil, err
	}

	plan := &servicesv1.ServicePlan{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServicePlan",
			APIVersion: servicesv1.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      service.Name + "-" + catalogPlan.Name,
			Namespace: kore.HubNamespace,
		},
		Spec: servicesv1.ServicePlanSpec{
			Kind:                  service.Name,
			DisplayName:           getMetadataStringVal(catalogPlan.Metadata, MetadataKeyDisplayName, ""),
			Summary:               catalogPlan.Description,
			Description:           getMetadataStringVal(catalogPlan.Metadata, MetadataKeyDescription, ""),
			Schema:                schemaStr,
			CredentialSchema:      credentialsSchemaStr,
			ServiceAccessDisabled: catalogPlan.Bindable != nil && !*catalogPlan.Bindable,
		},
	}

	configuration := map[string]interface{}{}

	if rawConfiguration, hasConfig := catalogPlan.Metadata[MetadataKeyConfiguration]; hasConfig {
		var ok bool
		configuration, ok = rawConfiguration.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("%s-%s plan has an invalid configuration, it must be an object", service.Name, catalogPlan.Name)
		}
	}

	if plan.Spec.Schema != "" {
		var required utils.StringSet
		required, _ = utils.StringSliceFrom(gjson.Get(plan.Spec.Schema, "required").Value())

		for _, override := range p.config.PlanConfigurationOverrides["*"] {
			plan.Spec.Schema = overrideSchemaProperty(plan.Spec.Schema, &required, override)
		}

		for _, override := range p.config.PlanConfigurationOverrides[plan.Name] {
			plan.Spec.Schema = overrideSchemaProperty(plan.Spec.Schema, &required, override)
		}

		plan.Spec.Schema, _ = sjson.Set(plan.Spec.Schema, "required", required)
	}

	if plan.Spec.Schema != "" {
		schema := &jsonschema.Schema{}
		if err := json.Unmarshal([]byte(plan.Spec.Schema), schema); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON schema: %w", err)
		}

		for name, prop := range schema.Properties {
			if _, isSet := configuration[name]; !isSet {
				constValue, err := prop.ParseConst()
				if err != nil {
					return nil, fmt.Errorf("invalid const value %v in JSON schema: %w", prop.Const, err)
				}
				if constValue != nil {
					configuration[name] = constValue
					continue
				}

				defaultValue, err := prop.ParseDefault()
				if err != nil {
					return nil, fmt.Errorf("invalid default value %v in JSON schema: %w", prop.Default, err)
				}
				if defaultValue != nil {
					configuration[name] = defaultValue
				}
			}
		}
	}

	if err := plan.Spec.SetConfiguration(configuration); err != nil {
		return nil, err
	}

	if err := plan.Spec.SetProviderData(ServicePlanProviderData{
		ServiceID: service.ID,
		PlanID:    catalogPlan.ID,
	}); err != nil {
		return nil, err
	}

	return plan, nil
}

func overrideSchemaProperty(schema string, required *utils.StringSet, override PlanConfigurationOverride) string {
	if gjson.Get(schema, "properties."+override.Name).Exists() {
		schema, _ = sjson.Set(schema, "properties."+override.Name+".default", override.Value)
		if override.Const {
			schema, _ = sjson.Set(schema, "properties."+override.Name+".const", override.Value)
		}
		if override.Required != nil {
			required.MemberIf(override.Name, *override.Required)
		}
	}
	return schema
}

func getPlanSchema(plan osb.Plan) (string, error) {
	if plan.Schemas == nil || plan.Schemas.ServiceInstance == nil || plan.Schemas.ServiceInstance.Create == nil {
		return "", nil
	}
	return parseSchema(plan.Name+" plan", plan.Schemas.ServiceInstance.Create.Parameters)
}

func getCredentialsSchema(plan osb.Plan) (string, error) {
	if plan.Schemas == nil || plan.Schemas.ServiceBinding == nil || plan.Schemas.ServiceBinding.Create == nil {
		return "", nil
	}
	return parseSchema(plan.Name+" plan", plan.Schemas.ServiceBinding.Create.Parameters)
}

func parseSchema(subject string, val interface{}) (string, error) {
	if val == nil {
		return "", nil
	}

	var schema string
	switch reflect.TypeOf(val).Kind() {
	case reflect.Struct, reflect.Map:
		schemaBytes, err := json.Marshal(val)
		if err != nil {
			return "", fmt.Errorf("%s has an invalid provisioning schema", subject)
		}
		schema = string(schemaBytes)
	case reflect.String:
		schema = val.(string)
	default:
		return "", fmt.Errorf("%s has an invalid schema", subject)
	}

	if err := jsonschema.Validate(jsonschema.MetaSchemaDraft07, fmt.Sprintf("%s schema", subject), schema); err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(make([]byte, 0, len(schema)))
	if err := json.Compact(buf, []byte(schema)); err != nil {
		return "", fmt.Errorf("%s has an invalid schema: %w", subject, err)
	}

	return buf.String(), nil
}

func getMetadataStringVal(metadata map[string]interface{}, key, def string) string {
	val, ok := metadata[key]
	if ok {
		if strVal, ok := val.(string); ok && strVal != "" {
			return strVal
		}
	}

	return def
}
