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

package openservicebroker_test

import (
	"encoding/json"

	"github.com/appvia/kore/pkg/serviceproviders/openservicebroker"
	"github.com/appvia/kore/pkg/utils/jsonschema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProviderFactory", func() {
	It("should have valid JSON schemas", func() {
		factory := openservicebroker.ProviderFactory{}
		for id, schema := range factory.JSONSchemas() {
			schemaObj := jsonschema.MetaSchemaDraft7Ext{}
			err := json.Unmarshal([]byte(schema), &schemaObj)
			var context string
			if err != nil {
				if jsonErr, ok := err.(*json.SyntaxError); ok {
					context = schema[jsonErr.Offset : jsonErr.Offset+100]
				}
			}
			Expect(err).ToNot(HaveOccurred(), "error at: %s", context)

			Expect(id).To(Equal(schemaObj.Id), "map key must match $id in schema")
		}
	})
})
