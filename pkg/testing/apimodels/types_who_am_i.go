// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesWhoAmI types who am i
//
// swagger:model types.WhoAmI
type TypesWhoAmI struct {

	// email
	Email string `json:"email,omitempty"`

	// teams
	Teams []string `json:"teams"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this types who am i
func (m *TypesWhoAmI) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TypesWhoAmI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesWhoAmI) UnmarshalBinary(b []byte) error {
	var res TypesWhoAmI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
