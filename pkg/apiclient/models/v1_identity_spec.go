// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1IdentitySpec v1 identity spec
//
// swagger:model v1.IdentitySpec
type V1IdentitySpec struct {

	// account type
	// Required: true
	AccountType *string `json:"accountType"`

	// basic auth
	BasicAuth *V1BasicAuth `json:"basicAuth,omitempty"`

	// idp user
	IdpUser *V1IDPUser `json:"idpUser,omitempty"`

	// user
	// Required: true
	User *V1User `json:"user"`
}

// Validate validates this v1 identity spec
func (m *V1IdentitySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBasicAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdpUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IdentitySpec) validateAccountType(formats strfmt.Registry) error {

	if err := validate.Required("accountType", "body", m.AccountType); err != nil {
		return err
	}

	return nil
}

func (m *V1IdentitySpec) validateBasicAuth(formats strfmt.Registry) error {

	if swag.IsZero(m.BasicAuth) { // not required
		return nil
	}

	if m.BasicAuth != nil {
		if err := m.BasicAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basicAuth")
			}
			return err
		}
	}

	return nil
}

func (m *V1IdentitySpec) validateIdpUser(formats strfmt.Registry) error {

	if swag.IsZero(m.IdpUser) { // not required
		return nil
	}

	if m.IdpUser != nil {
		if err := m.IdpUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("idpUser")
			}
			return err
		}
	}

	return nil
}

func (m *V1IdentitySpec) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IdentitySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IdentitySpec) UnmarshalBinary(b []byte) error {
	var res V1IdentitySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
