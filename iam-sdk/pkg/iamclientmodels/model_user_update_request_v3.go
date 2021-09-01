// Code generated by go-swagger; DO NOT EDIT.

package iamclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelUserUpdateRequestV3 model user update request v3
//
// swagger:model model.UserUpdateRequestV3
type ModelUserUpdateRequestV3 struct {

	// country
	Country string `json:"country,omitempty"`

	// date of birth
	DateOfBirth string `json:"dateOfBirth,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// language tag
	LanguageTag string `json:"languageTag,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this model user update request v3
func (m *ModelUserUpdateRequestV3) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelUserUpdateRequestV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelUserUpdateRequestV3) UnmarshalBinary(b []byte) error {
	var res ModelUserUpdateRequestV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}