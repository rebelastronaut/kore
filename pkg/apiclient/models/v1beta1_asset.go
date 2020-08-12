// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1beta1Asset v1beta1 asset
//
// swagger:model v1beta1.Asset
type V1beta1Asset struct {

	// asset identifier
	AssetIdentifier string `json:"assetIdentifier,omitempty"`

	// kore identifier
	KoreIdentifier string `json:"koreIdentifier,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`

	// team identifier
	TeamIdentifier string `json:"teamIdentifier,omitempty"`
}

// Validate validates this v1beta1 asset
func (m *V1beta1Asset) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1Asset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1Asset) UnmarshalBinary(b []byte) error {
	var res V1beta1Asset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
