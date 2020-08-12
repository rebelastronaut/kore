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

package costs

import (
	"github.com/appvia/kore/pkg/metadata"
	"github.com/appvia/kore/pkg/persistence"
)

type Costs interface {
	// Estimates returns the interface to the cost estimation service
	Estimates() Estimates
	// Assets returns the interface to the assets service
	Assets() Assets
}

// New returns a new instance of the costs API
func New(metadata metadata.Metadata, persistence persistence.TeamAssets, getKoreIdentifier func() string) Costs {
	estimates := NewEstimates(metadata)
	assets := NewAssets(persistence, getKoreIdentifier)
	return &costsImpl{
		estimates,
		assets,
	}
}

type costsImpl struct {
	estimates Estimates
	assets    Assets
}

func (c *costsImpl) Estimates() Estimates {
	return c.estimates
}

func (c *costsImpl) Assets() Assets {
	return c.assets
}
