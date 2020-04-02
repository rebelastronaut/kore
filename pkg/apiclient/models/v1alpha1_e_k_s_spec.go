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

// V1alpha1EKSSpec v1alpha1 e k s spec
//
// swagger:model v1alpha1.EKSSpec
type V1alpha1EKSSpec struct {

	// cluster
	Cluster *V1Ownership `json:"cluster,omitempty"`

	// credentials
	// Required: true
	Credentials *V1Ownership `json:"credentials"`

	// name
	// Required: true
	Name *string `json:"name"`

	// region
	// Required: true
	Region *string `json:"region"`

	// role a r n
	// Required: true
	RoleARN *string `json:"roleARN"`

	// security group i ds
	SecurityGroupIDs []string `json:"securityGroupIDs"`

	// subnet i ds
	// Required: true
	SubnetIDs []string `json:"subnetIDs"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1alpha1 e k s spec
func (m *V1alpha1EKSSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleARN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetIDs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1EKSSpec) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
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

func (m *V1alpha1EKSSpec) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials", "body", m.Credentials); err != nil {
		return err
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1EKSSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1EKSSpec) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1EKSSpec) validateRoleARN(formats strfmt.Registry) error {

	if err := validate.Required("roleARN", "body", m.RoleARN); err != nil {
		return err
	}

	return nil
}

func (m *V1alpha1EKSSpec) validateSubnetIDs(formats strfmt.Registry) error {

	if err := validate.Required("subnetIDs", "body", m.SubnetIDs); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1EKSSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1EKSSpec) UnmarshalBinary(b []byte) error {
	var res V1alpha1EKSSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
