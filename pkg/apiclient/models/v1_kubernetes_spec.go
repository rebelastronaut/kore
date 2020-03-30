// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1KubernetesSpec v1 kubernetes spec
//
// swagger:model v1.KubernetesSpec
type V1KubernetesSpec struct {

	// auth proxy allowed i ps
	AuthProxyAllowedIPs []string `json:"authProxyAllowedIPs"`

	// auth proxy image
	AuthProxyImage string `json:"authProxyImage,omitempty"`

	// cluster
	Cluster *V1Ownership `json:"cluster,omitempty"`

	// cluster users
	ClusterUsers []*V1ClusterUser `json:"clusterUsers"`

	// default team role
	DefaultTeamRole string `json:"defaultTeamRole,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// enable default traffic block
	EnableDefaultTrafficBlock bool `json:"enableDefaultTrafficBlock,omitempty"`

	// inherit team members
	InheritTeamMembers bool `json:"inheritTeamMembers,omitempty"`

	// provider
	Provider *V1Ownership `json:"provider,omitempty"`
}

// Validate validates this v1 kubernetes spec
func (m *V1KubernetesSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1KubernetesSpec) validateCluster(formats strfmt.Registry) error {

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

func (m *V1KubernetesSpec) validateClusterUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterUsers) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterUsers); i++ {
		if swag.IsZero(m.ClusterUsers[i]) { // not required
			continue
		}

		if m.ClusterUsers[i] != nil {
			if err := m.ClusterUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1KubernetesSpec) validateProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	if m.Provider != nil {
		if err := m.Provider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provider")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1KubernetesSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1KubernetesSpec) UnmarshalBinary(b []byte) error {
	var res V1KubernetesSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
