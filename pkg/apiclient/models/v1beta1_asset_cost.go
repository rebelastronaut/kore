// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1beta1AssetCost v1beta1 asset cost
//
// swagger:model v1beta1.AssetCost
type V1beta1AssetCost struct {

	// account
	Account string `json:"account,omitempty"`

	// asset identifier
	AssetIdentifier string `json:"assetIdentifier,omitempty"`

	// cost
	Cost int64 `json:"cost,omitempty"`

	// cost identifier
	CostIdentifier string `json:"costIdentifier,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// invoice
	Invoice string `json:"invoice,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// retrieved at
	RetrievedAt string `json:"retrievedAt,omitempty"`

	// team identifier
	TeamIdentifier string `json:"teamIdentifier,omitempty"`

	// usage amount
	UsageAmount string `json:"usageAmount,omitempty"`

	// usage end time
	UsageEndTime string `json:"usageEndTime,omitempty"`

	// usage start time
	UsageStartTime string `json:"usageStartTime,omitempty"`

	// usage type
	UsageType string `json:"usageType,omitempty"`

	// usage unit
	UsageUnit string `json:"usageUnit,omitempty"`
}

// Validate validates this v1beta1 asset cost
func (m *V1beta1AssetCost) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1AssetCost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1AssetCost) UnmarshalBinary(b []byte) error {
	var res V1beta1AssetCost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
