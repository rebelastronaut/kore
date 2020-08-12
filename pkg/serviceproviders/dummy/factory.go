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
	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"
)

func init() {
	kore.RegisterServiceProviderFactory(DummyFactory{})
}

type DummyFactory struct{}

func (d DummyFactory) Type() string {
	return "dummy"
}

// JSONSchemas returns all JSON schema versions for the provider's configuration
func (d DummyFactory) JSONSchemas() map[string]string {
	return map[string]string{
		"https://appvia.io/kore/schemas/serviceprovider/dummy/v1.json": providerSchemaV1,
		"https://appvia.io/kore/schemas/serviceprovider/dummy/v2.json": providerSchemaV2,
	}
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
