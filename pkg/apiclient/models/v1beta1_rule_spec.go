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

// V1beta1RuleSpec v1beta1 rule spec
//
// swagger:model v1beta1.RuleSpec
type V1beta1RuleSpec struct {

	// raw rule
	// Required: true
	RawRule *string `json:"rawRule"`

	// resource
	// Required: true
	Resource *V1Ownership `json:"resource"`

	// rule ID
	RuleID string `json:"ruleID,omitempty"`

	// severity
	// Required: true
	Severity *string `json:"severity"`

	// source
	// Required: true
	Source *string `json:"source"`

	// summary
	// Required: true
	Summary *string `json:"summary"`
}

// Validate validates this v1beta1 rule spec
func (m *V1beta1RuleSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRawRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1RuleSpec) validateRawRule(formats strfmt.Registry) error {

	if err := validate.Required("rawRule", "body", m.RawRule); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1RuleSpec) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *V1beta1RuleSpec) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1RuleSpec) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *V1beta1RuleSpec) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("summary", "body", m.Summary); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1RuleSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1RuleSpec) UnmarshalBinary(b []byte) error {
	var res V1beta1RuleSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
