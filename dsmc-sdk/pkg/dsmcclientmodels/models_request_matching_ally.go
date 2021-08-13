// Code generated by go-swagger; DO NOT EDIT.

package dsmcclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsRequestMatchingAlly models request matching ally
//
// swagger:model models.RequestMatchingAlly
type ModelsRequestMatchingAlly struct {

	// matching parties
	// Required: true
	MatchingParties []*ModelsRequestMatchParty `json:"matching_parties"`
}

// Validate validates this models request matching ally
func (m *ModelsRequestMatchingAlly) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchingParties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsRequestMatchingAlly) validateMatchingParties(formats strfmt.Registry) error {

	if err := validate.Required("matching_parties", "body", m.MatchingParties); err != nil {
		return err
	}

	for i := 0; i < len(m.MatchingParties); i++ {
		if swag.IsZero(m.MatchingParties[i]) { // not required
			continue
		}

		if m.MatchingParties[i] != nil {
			if err := m.MatchingParties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matching_parties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsRequestMatchingAlly) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsRequestMatchingAlly) UnmarshalBinary(b []byte) error {
	var res ModelsRequestMatchingAlly
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
