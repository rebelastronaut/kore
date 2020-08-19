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
	"fmt"

	"github.com/appvia/kore/pkg/utils/configuration"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"
	osb "github.com/kubernetes-sigs/go-open-service-broker-client/v2"
)

func init() {
	kore.RegisterServiceProviderFactory(ProviderFactory{})
}

type ProviderFactory struct{}

func (p ProviderFactory) Type() string {
	return "osb"
}

// JSONSchemas returns all JSON schema versions for the provider's configuration
func (d ProviderFactory) JSONSchemas() map[string]string {
	return map[string]string{
		"https://appvia.io/kore/schemas/serviceprovider/osb/v1.json": providerSchemaV1,
	}
}

func (p ProviderFactory) Create(ctx kore.Context, serviceProvider *servicesv1.ServiceProvider) (kore.ServiceProvider, error) {
	var config = ProviderConfiguration{}
	config.Name = serviceProvider.Name

	if _, err := configuration.ParseObjectConfiguration(ctx, ctx.Client(), serviceProvider, &config); err != nil {
		return nil, fmt.Errorf("failed to process service provider configuration: %w", err)
	}

	osbClient, err := osb.NewClient(&config.ClientConfiguration)
	if err != nil {
		return nil, err
	}

	provider := NewProvider(serviceProvider.Name, config, osbClient)

	return provider, nil
}

func (p ProviderFactory) SetUp(ctx kore.Context, serviceProvider *servicesv1.ServiceProvider) (complete bool, _ error) {
	return true, nil
}

func (p ProviderFactory) TearDown(ctx kore.Context, serviceProvider *servicesv1.ServiceProvider) (complete bool, _ error) {
	return true, nil
}

func (d ProviderFactory) DefaultProviders() []servicesv1.ServiceProvider {
	return nil
}
