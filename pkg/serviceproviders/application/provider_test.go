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

package application_test

import (
	"context"

	"github.com/appvia/kore/pkg/controllers/controllerstest"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/serviceproviders/application"
	"github.com/appvia/kore/pkg/utils/jsonschema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServicePlans", func() {
	It("All plans should be valid", func() {
		factory := application.Factory{}
		test := controllerstest.NewTest(context.Background())

		providerObj := factory.DefaultProviders()[0]

		provider, err := factory.Create(test.Context, &providerObj)
		Expect(err).ToNot(HaveOccurred())

		catalog, err := provider.Catalog(test.Context, &providerObj)
		Expect(err).ToNot(HaveOccurred())

		for _, plan := range catalog.Plans {
			var kind *servicesv1.ServiceKind
			for _, k := range catalog.Kinds {
				if k.Name == plan.Spec.Kind {
					kind = &k
					break
				}
			}

			Expect(kind).ToNot(Equal(""), "service plan %s doesn't have a valid service kind", plan.Name)

			err := jsonschema.Validate(plan.Spec.Schema, plan.Name, plan.Spec.Configuration)
			Expect(err).ToNot(HaveOccurred(), "%s plan is not valid: %s", plan.Name, err)
		}
	})

})
