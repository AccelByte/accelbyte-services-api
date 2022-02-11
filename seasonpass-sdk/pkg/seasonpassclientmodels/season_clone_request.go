// Code generated by go-swagger; DO NOT EDIT.

package seasonpassclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SeasonCloneRequest A DTO object for clone season API call.
//
// swagger:model SeasonCloneRequest
type SeasonCloneRequest struct {

	// end date time
	// Required: true
	// Format: date-time
	End strfmt.DateTime `json:"end"`

	// name, max length is 127
	// Required: true
	Name *string `json:"name"`

	// start date time
	// Required: true
	// Format: date-time
	Start strfmt.DateTime `json:"start"`
}

// Validate validates this season clone request
func (m *SeasonCloneRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SeasonCloneRequest) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", strfmt.DateTime(m.End)); err != nil {
		return err
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SeasonCloneRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SeasonCloneRequest) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", strfmt.DateTime(m.Start)); err != nil {
		return err
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SeasonCloneRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeasonCloneRequest) UnmarshalBinary(b []byte) error {
	var res SeasonCloneRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
