// Code generated by go-swagger; DO NOT EDIT.

package socialclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatInfo stat info
//
// swagger:model StatInfo
type StatInfo struct {

	// created at
	// Required: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// default value
	// Required: true
	DefaultValue *float64 `json:"defaultValue"`

	// description
	Description string `json:"description,omitempty"`

	// increment only
	// Required: true
	IncrementOnly *bool `json:"incrementOnly"`

	// maximum
	Maximum float64 `json:"maximum,omitempty"`

	// minimum
	Minimum float64 `json:"minimum,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// set as global
	// Required: true
	SetAsGlobal *bool `json:"setAsGlobal"`

	// set by
	// Required: true
	// Enum: [CLIENT SERVER]
	SetBy *string `json:"setBy"`

	// stat code
	// Required: true
	StatCode *string `json:"statCode"`

	// status
	// Required: true
	// Enum: [INIT TIED]
	Status *string `json:"status"`

	// tags
	// Unique: true
	Tags []string `json:"tags"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt"`
}

// Validate validates this stat info
func (m *StatInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncrementOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSetAsGlobal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSetBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatInfo) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StatInfo) validateDefaultValue(formats strfmt.Registry) error {

	if err := validate.Required("defaultValue", "body", m.DefaultValue); err != nil {
		return err
	}

	return nil
}

func (m *StatInfo) validateIncrementOnly(formats strfmt.Registry) error {

	if err := validate.Required("incrementOnly", "body", m.IncrementOnly); err != nil {
		return err
	}

	return nil
}

func (m *StatInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StatInfo) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *StatInfo) validateSetAsGlobal(formats strfmt.Registry) error {

	if err := validate.Required("setAsGlobal", "body", m.SetAsGlobal); err != nil {
		return err
	}

	return nil
}

var statInfoTypeSetByPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLIENT","SERVER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statInfoTypeSetByPropEnum = append(statInfoTypeSetByPropEnum, v)
	}
}

const (

	// StatInfoSetByCLIENT captures enum value "CLIENT"
	StatInfoSetByCLIENT string = "CLIENT"

	// StatInfoSetBySERVER captures enum value "SERVER"
	StatInfoSetBySERVER string = "SERVER"
)

// prop value enum
func (m *StatInfo) validateSetByEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statInfoTypeSetByPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatInfo) validateSetBy(formats strfmt.Registry) error {

	if err := validate.Required("setBy", "body", m.SetBy); err != nil {
		return err
	}

	// value enum
	if err := m.validateSetByEnum("setBy", "body", *m.SetBy); err != nil {
		return err
	}

	return nil
}

func (m *StatInfo) validateStatCode(formats strfmt.Registry) error {

	if err := validate.Required("statCode", "body", m.StatCode); err != nil {
		return err
	}

	return nil
}

var statInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INIT","TIED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statInfoTypeStatusPropEnum = append(statInfoTypeStatusPropEnum, v)
	}
}

const (

	// StatInfoStatusINIT captures enum value "INIT"
	StatInfoStatusINIT string = "INIT"

	// StatInfoStatusTIED captures enum value "TIED"
	StatInfoStatusTIED string = "TIED"
)

// prop value enum
func (m *StatInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatInfo) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *StatInfo) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *StatInfo) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatInfo) UnmarshalBinary(b []byte) error {
	var res StatInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
