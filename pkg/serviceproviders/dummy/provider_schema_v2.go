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

//go:generate go run github.com/appvia/kore/cmd/struct-gen ProviderSchemaV2
const providerSchemaV2 = `{
	"$id": "https://appvia.io/kore/schemas/serviceprovider/dummy/v2.json",
	"$schema": "http://json-schema.org/draft-07/schema#",
	"description": "Dummy service plan schema",
	"type": "object",
	"additionalProperties": false,
	"required": [
		"iAmDummy",
		"iAmVersionTwo"
	],
	"properties": {
		"iAmDummy": {
			"type": "string",
			"minLength": 1
		},
		"iAmVersionTwo": {
			"type": "string",
			"minLength": 1
		}
	}
}`
