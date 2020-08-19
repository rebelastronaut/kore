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

const providerSchemaV1 = `{
	"$id": "https://appvia.io/kore/schemas/serviceprovider/osb/v1.json",
	"$schema": "http://json-schema.org/draft-07/schema#",
	"description": "Open Service Broker Provider configuration schema",
	"type": "object",
	"additionalProperties": false,
	"required": [
		"url"
	],
	"properties": {
		"enable_alpha_features": {
			"type": "boolean"
		},
		"url": {
			"type": "string",
			"minLength": 1
		},
		"api_version": {
			"type": "string",
			"minLength": 1
		},
		"insecure": {
			"type": "boolean"
		},
		"ca_data": {
			"type": "string"
		},
		"auth_config": {
			"type": "object",
			"additionalProperties": false,
			"properties": {
				"basic_auth_config": {
					"type": "object",
					"additionalProperties": false,
					"required": [
						"username",
						"password"
					],
					"properties": {
						"username": {
							"type": "string",
							"minLength": 1
						},
						"password": {
							"type": "string"
						}
					}
				},
				"bearer_config": {
					"type": "object",
					"additionalProperties": false,
					"required": [
						"token"
					],
					"properties": {
						"token": {
							"type": "string",
							"minLength": 1
						}
					}
				}
			}
		},
		"allowEmptyCredentialSchema": {
			"type": "boolean",
			"default": false
		},
		"defaultPlans": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			}
		},
		"includeKinds": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			}
		},
		"excludeKinds": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			}
		},
		"includePlans": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			}
		},
		"excludePlans": {
			"type": "array",
			"items": {
				"type": "string",
				"minLength": 1
			}
		},
		"platformMapping": {
			"type": "object",
			"minProperties": 1,
			"additionalProperties": { "type": "string" }
		}
	}
}`
