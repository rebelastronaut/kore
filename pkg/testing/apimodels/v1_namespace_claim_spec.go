// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1NamespaceClaimSpec v1 namespace claim spec
//
// swagger:model v1.NamespaceClaimSpec
type V1NamespaceClaimSpec struct {

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// cluster
	// Required: true
	Cluster *V1Ownership `json:"cluster"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this v1 namespace claim spec
func (m *V1NamespaceClaimSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NamespaceClaimSpec) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *V1NamespaceClaimSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NamespaceClaimSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NamespaceClaimSpec) UnmarshalBinary(b []byte) error {
	var res V1NamespaceClaimSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
