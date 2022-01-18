// Code generated by go-swagger; DO NOT EDIT.

package iamclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelValidationDetail model validation detail
//
// swagger:model model.ValidationDetail
type ModelValidationDetail struct {

	// allow digit
	// Required: true
	AllowDigit *bool `json:"allowDigit"`

	// allow letter
	// Required: true
	AllowLetter *bool `json:"allowLetter"`

	// allow space
	// Required: true
	AllowSpace *bool `json:"allowSpace"`

	// allow unicode
	// Required: true
	AllowUnicode *bool `json:"allowUnicode"`

	// description
	// Required: true
	Description []*AccountcommonInputValidationDescription `json:"description"`

	// is custom regex
	// Required: true
	IsCustomRegex *bool `json:"isCustomRegex"`

	// letter case
	// Required: true
	LetterCase *string `json:"letterCase"`

	// max length
	// Required: true
	MaxLength *int32 `json:"maxLength"`

	// max repeating alpha num
	// Required: true
	MaxRepeatingAlphaNum *int32 `json:"maxRepeatingAlphaNum"`

	// max repeating special character
	// Required: true
	MaxRepeatingSpecialCharacter *int32 `json:"maxRepeatingSpecialCharacter"`

	// min char type
	// Required: true
	MinCharType *int32 `json:"minCharType"`

	// min length
	// Required: true
	MinLength *int32 `json:"minLength"`

	// regex
	// Required: true
	Regex *string `json:"regex"`

	// special character location
	// Required: true
	SpecialCharacterLocation *string `json:"specialCharacterLocation"`

	// special characters
	// Required: true
	SpecialCharacters []string `json:"specialCharacters"`
}

// Validate validates this model validation detail
func (m *ModelValidationDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowDigit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowLetter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllowUnicode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsCustomRegex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLetterCase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxRepeatingAlphaNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxRepeatingSpecialCharacter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinCharType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecialCharacterLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecialCharacters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelValidationDetail) validateAllowDigit(formats strfmt.Registry) error {

	if err := validate.Required("allowDigit", "body", m.AllowDigit); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateAllowLetter(formats strfmt.Registry) error {

	if err := validate.Required("allowLetter", "body", m.AllowLetter); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateAllowSpace(formats strfmt.Registry) error {

	if err := validate.Required("allowSpace", "body", m.AllowSpace); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateAllowUnicode(formats strfmt.Registry) error {

	if err := validate.Required("allowUnicode", "body", m.AllowUnicode); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	for i := 0; i < len(m.Description); i++ {
		if swag.IsZero(m.Description[i]) { // not required
			continue
		}

		if m.Description[i] != nil {
			if err := m.Description[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("description" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelValidationDetail) validateIsCustomRegex(formats strfmt.Registry) error {

	if err := validate.Required("isCustomRegex", "body", m.IsCustomRegex); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateLetterCase(formats strfmt.Registry) error {

	if err := validate.Required("letterCase", "body", m.LetterCase); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("maxLength", "body", m.MaxLength); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateMaxRepeatingAlphaNum(formats strfmt.Registry) error {

	if err := validate.Required("maxRepeatingAlphaNum", "body", m.MaxRepeatingAlphaNum); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateMaxRepeatingSpecialCharacter(formats strfmt.Registry) error {

	if err := validate.Required("maxRepeatingSpecialCharacter", "body", m.MaxRepeatingSpecialCharacter); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateMinCharType(formats strfmt.Registry) error {

	if err := validate.Required("minCharType", "body", m.MinCharType); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateMinLength(formats strfmt.Registry) error {

	if err := validate.Required("minLength", "body", m.MinLength); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateRegex(formats strfmt.Registry) error {

	if err := validate.Required("regex", "body", m.Regex); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateSpecialCharacterLocation(formats strfmt.Registry) error {

	if err := validate.Required("specialCharacterLocation", "body", m.SpecialCharacterLocation); err != nil {
		return err
	}

	return nil
}

func (m *ModelValidationDetail) validateSpecialCharacters(formats strfmt.Registry) error {

	if err := validate.Required("specialCharacters", "body", m.SpecialCharacters); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelValidationDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelValidationDetail) UnmarshalBinary(b []byte) error {
	var res ModelValidationDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
