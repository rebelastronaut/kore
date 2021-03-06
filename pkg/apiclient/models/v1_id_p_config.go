// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1IDPConfig v1 ID p config
//
// swagger:model v1.IDPConfig
type V1IDPConfig struct {

	// github
	Github *V1GithubIDP `json:"github,omitempty"`

	// google
	Google *V1GoogleIDP `json:"google,omitempty"`

	// oidc
	Oidc *V1OIDCIDP `json:"oidc,omitempty"`

	// oidcdirect
	Oidcdirect *V1StaticOIDCIDP `json:"oidcdirect,omitempty"`

	// saml
	Saml *V1SAMLIDP `json:"saml,omitempty"`
}

// Validate validates this v1 ID p config
func (m *V1IDPConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGithub(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGoogle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidcdirect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSaml(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IDPConfig) validateGithub(formats strfmt.Registry) error {

	if swag.IsZero(m.Github) { // not required
		return nil
	}

	if m.Github != nil {
		if err := m.Github.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("github")
			}
			return err
		}
	}

	return nil
}

func (m *V1IDPConfig) validateGoogle(formats strfmt.Registry) error {

	if swag.IsZero(m.Google) { // not required
		return nil
	}

	if m.Google != nil {
		if err := m.Google.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("google")
			}
			return err
		}
	}

	return nil
}

func (m *V1IDPConfig) validateOidc(formats strfmt.Registry) error {

	if swag.IsZero(m.Oidc) { // not required
		return nil
	}

	if m.Oidc != nil {
		if err := m.Oidc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc")
			}
			return err
		}
	}

	return nil
}

func (m *V1IDPConfig) validateOidcdirect(formats strfmt.Registry) error {

	if swag.IsZero(m.Oidcdirect) { // not required
		return nil
	}

	if m.Oidcdirect != nil {
		if err := m.Oidcdirect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidcdirect")
			}
			return err
		}
	}

	return nil
}

func (m *V1IDPConfig) validateSaml(formats strfmt.Registry) error {

	if swag.IsZero(m.Saml) { // not required
		return nil
	}

	if m.Saml != nil {
		if err := m.Saml.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saml")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IDPConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IDPConfig) UnmarshalBinary(b []byte) error {
	var res V1IDPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
