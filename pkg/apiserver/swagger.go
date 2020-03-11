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

package apiserver

import "github.com/go-openapi/spec"

// EnrichSwagger provides the swagger config
func EnrichSwagger(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Appvia Kore API",
			Description: "Kore API provides the frontend API for the Appvia Kore (kore.appvia.io)",
			Contact: &spec.ContactInfo{
				Name:  "Rohith Jayawardene",
				Email: "info@appvia.io",
				URL:   "https://appvia.io",
			},
			License: &spec.License{
				Name: "GPLV2",
				URL:  "http://mit.org",
			},
			Version: "0.0.1",
		},
	}
}
