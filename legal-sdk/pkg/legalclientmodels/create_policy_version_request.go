// Code generated by go-swagger; DO NOT EDIT.

package legalclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreatePolicyVersionRequest create policy version request
//
// swagger:model CreatePolicyVersionRequest
type CreatePolicyVersionRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// display version
	DisplayVersion string `json:"displayVersion,omitempty"`

	// is committed
	IsCommitted bool `json:"isCommitted,omitempty"`
}

// Validate validates this create policy version request
func (m *CreatePolicyVersionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreatePolicyVersionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePolicyVersionRequest) UnmarshalBinary(b []byte) error {
	var res CreatePolicyVersionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
