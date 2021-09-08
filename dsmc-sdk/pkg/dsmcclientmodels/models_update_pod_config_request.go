// Code generated by go-swagger; DO NOT EDIT.

package dsmcclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsUpdatePodConfigRequest models update pod config request
//
// swagger:model models.UpdatePodConfigRequest
type ModelsUpdatePodConfigRequest struct {

	// cpu limit
	// Required: true
	CPULimit *int32 `json:"cpu_limit"`

	// mem limit
	// Required: true
	MemLimit *int32 `json:"mem_limit"`

	// name
	// Required: true
	Name *string `json:"name"`

	// params
	// Required: true
	Params *string `json:"params"`
}

// Validate validates this models update pod config request
func (m *ModelsUpdatePodConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPULimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsUpdatePodConfigRequest) validateCPULimit(formats strfmt.Registry) error {

	if err := validate.Required("cpu_limit", "body", m.CPULimit); err != nil {
		return err
	}

	return nil
}

func (m *ModelsUpdatePodConfigRequest) validateMemLimit(formats strfmt.Registry) error {

	if err := validate.Required("mem_limit", "body", m.MemLimit); err != nil {
		return err
	}

	return nil
}

func (m *ModelsUpdatePodConfigRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ModelsUpdatePodConfigRequest) validateParams(formats strfmt.Registry) error {

	if err := validate.Required("params", "body", m.Params); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsUpdatePodConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsUpdatePodConfigRequest) UnmarshalBinary(b []byte) error {
	var res ModelsUpdatePodConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
