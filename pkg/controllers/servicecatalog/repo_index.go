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
	"time"
)

type RepoIndex struct {
	APIVersion string                    `json:"apiVersion"`
	Generated  time.Time                 `json:"generated"`
	Entries    map[string][]ChartVersion `json:"entries"`
	PublicKeys []string                  `json:"publicKeys,omitempty"`
}

type ChartVersion struct {
	ChartMetadata `json:",inline"`
	URLs          []string  `json:"urls"`
	Created       time.Time `json:"created,omitempty"`
	Removed       bool      `json:"removed,omitempty"`
	Digest        string    `json:"digest,omitempty"`
}

type ChartMaintainer struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	URL   string `json:"url,omitempty"`
}

type ChartMetadata struct {
	Name         string             `json:"name,omitempty"`
	Home         string             `json:"home,omitempty"`
	Sources      []string           `json:"sources,omitempty"`
	Version      string             `json:"version,omitempty"`
	Description  string             `json:"description,omitempty"`
	Keywords     []string           `json:"keywords,omitempty"`
	Maintainers  []*ChartMaintainer `json:"maintainers,omitempty"`
	Icon         string             `json:"icon,omitempty"`
	APIVersion   string             `json:"apiVersion,omitempty"`
	Condition    string             `json:"condition,omitempty"`
	Tags         string             `json:"tags,omitempty"`
	AppVersion   string             `json:"appVersion,omitempty"`
	Deprecated   bool               `json:"deprecated,omitempty"`
	Annotations  map[string]string  `json:"annotations,omitempty"`
	KubeVersion  string             `json:"kubeVersion,omitempty"`
	Dependencies []*ChartDependency `json:"dependencies,omitempty"`
	Type         string             `json:"type,omitempty"`
}

type ChartDependency struct {
	Name         string        `json:"name"`
	Version      string        `json:"version,omitempty"`
	Repository   string        `json:"repository"`
	Condition    string        `json:"condition,omitempty"`
	Tags         []string      `json:"tags,omitempty"`
	Enabled      bool          `json:"enabled,omitempty"`
	ImportValues []interface{} `json:"import-values,omitempty"`
	Alias        string        `json:"alias,omitempty"`
}
