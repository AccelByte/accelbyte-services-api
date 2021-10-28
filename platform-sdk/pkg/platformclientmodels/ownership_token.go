// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OwnershipToken ownership token
//
// swagger:model OwnershipToken
type OwnershipToken struct {

	// ownership token
	OwnershipToken string `json:"ownershipToken,omitempty"`
}

// Validate validates this ownership token
func (m *OwnershipToken) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OwnershipToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OwnershipToken) UnmarshalBinary(b []byte) error {
	var res OwnershipToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
