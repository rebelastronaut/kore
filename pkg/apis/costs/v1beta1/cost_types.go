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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Asset represents a resource known to Kore which a cost provider should provide costs data for
// +k8s:openapi-gen=false
type Asset struct {
	// Tags are a set of tags which can be used to identify this asset
	Tags map[string]string `json:"tags,omitempty"`
	// TeamIdentifier is the unique identifier for the team that owns this asset
	TeamIdentifier string `json:"teamIdentifier,omitempty"`
	// AssetIdentifier is the unique identifier for this asset
	AssetIdentifier string `json:"assetIdentifier,omitempty"`
	// Name is the name of the resource in kore, for reference
	Name string `json:"name,omitempty"`
	// Provider is the cloud provider who provides this resource
	Provider string `json:"provider,omitempty"`
}

// AssetList is a list of cost assets known to kore for which costs can be provided by a cost
// provider
// +k8s:openapi-gen=false
type AssetList struct {
	// KoreIdentifier is the unique identifier for this instance of kore
	KoreIdentifier string `json:"koreIdentifier,omitempty"`
	// Items are the assets which should be monitored for costs
	Items []Asset `json:"items"`
}

// AssetCostList represents a collection of costs about one or more assets
// +k8s:openapi-gen=false
type AssetCostList struct {
	Items []AssetCost `json:"items"`
}

// CostSummary represents a total cost over a period of time
// +k8s:openapi-gen=false
type CostSummary struct {
	// Cost is the actual incurred cost total cost for the specified time period in
	// microdollars
	Cost int64 `json:"cost"`
	// StartTime indicates the start of the period this summary includes costs for
	StartTime metav1.Time `json:"usageStartTime,omitempty"`
	// EndTime indicates the end of the period this summary includes costs for
	EndTime metav1.Time `json:"usageEndTime,omitempty"`
}

// OverallCostSummary represents the total costs known to kore over a period of time,
// and acts as a container for TeamCostSummaries
// +k8s:openapi-gen=false
type OverallCostSummary struct {
	CostSummary
	TeamCosts []*TeamCostSummary `json:"teamCosts"`
}

// TeamCostSummary represents the total cost known to kore for a team (over a period of time)
// +k8s:openapi-gen=false
type TeamCostSummary struct {
	// TeamIdentifier is the unique identifier for the team these costs belongs to.
	TeamIdentifier string `json:"teamIdentifier,omitempty"`
	// TeamName is the name of the team that these costs belong to
	TeamName string `json:"teamName,omitempty"`
	// AssetCosts gives the detail of the assets which make up this team cost
	AssetCosts []*AssetCostSummary `json:"assetCosts"`
	CostSummary
}

// AssetCostSummary represents the total cost known to kore for an asset (over a period of time)
// +k8s:openapi-gen=false
type AssetCostSummary struct {
	// AssetIdentifier is the unique identifier assigned to the resource this cost applies to, e.g. the
	// unique cluster ID, etc.
	AssetIdentifier string `json:"assetIdentifier,omitempty"`
	// TeamIdentifier is the unique identifier for the team this resource belongs to.
	TeamIdentifier string `json:"teamIdentifier,omitempty"`
	// AssetName is the name of the asset these costs relate to
	AssetName string `json:"assetName,omitempty"`
	// AssetType is the type of the asset these costs relate to
	AssetType string `json:"assetType,omitempty"`
	// Provider is the cloud provider who provides this assset
	Provider string `json:"provider,omitempty"`
	// Details provides the individual cost line items that make up this summary
	Details []*AssetCost `json:"details,omitempty"`
	CostSummary
}

// AssetCost defines the details about a cost related to a piece of infrastructure deployed by Kore for
// a team. It is expected that any asset may have multiple AssetCosts covering a specific time period
// to represent the different charges levied by the provider for that piece of infrastructure.
// +k8s:openapi-gen=false
type AssetCost struct {
	// CostIdentifier is the unique identifer for this line of cost data - cost providers must ensure that if a
	// cost line item is updated, it has the same identifier, and that different line items have unique cost
	// identifiers for a given AssetIdentifier. If a cost provider provides immutable cost entries, i.e. they will
	// never be updated, then this can be left blank and Kore will assign a unique identifier.
	CostIdentifier string `json:"costIdentifier,omitempty"`
	// AssetIdentifier is the unique identifier assigned to the resource this cost applies to, e.g. the
	// unique cluster ID, etc.
	AssetIdentifier string `json:"assetIdentifier,omitempty"`
	// TeamIdentifier is the unique identifier for the team this resource belongs to.
	TeamIdentifier string `json:"teamIdentifier,omitempty"`
	// Cost is the actual incurred cost total cost for this piece of infrastructure for the
	// specified time period in microdollars
	Cost int64 `json:"cost,omitempty"`
	// UsageStartTime indicates the start of the period this cost is applicable for
	UsageStartTime metav1.Time `json:"usageStartTime,omitempty"`
	// UsageEndTime indicates the end of the period this cost is applicable for
	UsageEndTime metav1.Time `json:"usageEndTime,omitempty"`
	// UsageType is the provider-specific code or title for this type of usage (e.g. a SKU or similar)
	UsageType string `json:"usageType,omitempty"`
	// Description identifies the type of cost this line item refers to
	Description string `json:"description,omitempty"`
	// UsageAmount is the quantity of the resource used (e.g. amount of storage)
	UsageAmount string `json:"usageAmount,omitempty"`
	// UsageUnit is the unit that UsageAmount is expressed in (e.g. seconds, gibibytes, etc)
	UsageUnit string `json:"usageUnit,omitempty"`
	// Provider indicates which cloud provider this cost relates to
	Provider string `json:"provider,omitempty"`
	// Account indicates which account / project / subscription this cost relates to
	Account string `json:"account,omitempty"`
	// Invoice is the invoice on which this cost was billed (in the format YYYYMM, e.g. 202008 for August 2020)
	Invoice string `json:"invoice,omitempty"`
	// RetrievedAt is the time at which this cost item was retrieved/refreshed from the provider
	RetrievedAt metav1.Time `json:"retrievedAt,omitempty"`
}
