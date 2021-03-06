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

package model

import (
	"time"
)

// TeamAssetCost represents an individual line item of cost data for a team asset
type TeamAssetCost struct {
	// CostIdentifier is the unique identifer for this line of cost data - cost providers must ensure that if a
	// cost line item is updated, it has the same identifier
	CostIdentifier string `sql:"type:varchar(255)" gorm:"PRIMARY_KEY"`
	// AssetIdentifier is the identity of the asset in question
	AssetIdentifier string `sql:"type:char(20)" gorm:"PRIMARY_KEY"`
	// Asset is the asset record this cost relates to
	Asset TeamAsset `gorm:"foreignkey:AssetIdentifier;association_foreignkey:AssetIdentifier"`
	// TeamIdentifier is the identity of the owning team
	TeamIdentifier string `sql:"type:char(20)"`
	// Team is the identity record for the team who owns this asset
	Team TeamIdentity `gorm:"foreignkey:TeamIdentifier;association_foreignkey:TeamIdentifier"`
	// UsageType is the provider-specific code or title for this type of usage (e.g. a SKU or similar)
	UsageType string
	// Description identifies the type of cost this line item refers to
	Description string
	// UsageStartTime identifies the start of the period for which this cost applies
	UsageStartTime time.Time
	// UsageEndTime identifies the end of the period for which this cost applies
	UsageEndTime time.Time
	// UsageAmount is the quantity of the resource used (e.g. amount of storage)
	UsageAmount float64
	// UsageUnit is the unit that UsageAmount is expressed in (e.g. seconds, gibibytes, etc)
	UsageUnit string
	// Cost is the cost (in microdollars) for the asset over the time period described. Negative amounts indicate a refund.
	Cost int64
	// Provider indicates which cloud provider this cost relates to
	Provider string
	// Account indicates which account / project / subscription this cost relates to
	Account string
	// Invoice identifies which invoice this cost is related to, in the format YYYYMM (e.g. 202008 for August 2020)
	Invoice string
	// RetrievedAt is the time at which this cost item was retrieved/refreshed from the provider
	RetrievedAt time.Time
}
