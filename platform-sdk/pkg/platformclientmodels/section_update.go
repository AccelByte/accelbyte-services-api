// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SectionUpdate A DTO object for updating section API call.
//
// swagger:model SectionUpdate
type SectionUpdate struct {

	// active, default true
	Active bool `json:"active"`

	// display order
	DisplayOrder int32 `json:"displayOrder,omitempty"`

	// end date
	// Required: true
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate"`

	// section extension
	Ext map[string]interface{} `json:"ext,omitempty"`

	// fix period rotation config, don't allow null when rotation type is FIXED_PERIOD
	FixedPeriodRotationConfig *FixedPeriodRotationConfig `json:"fixedPeriodRotationConfig,omitempty"`

	// items
	Items []*SectionItem `json:"items"`

	// Localization, key language, value localization content
	// Required: true
	Localizations map[string]Localization `json:"localizations"`

	// section name
	// Required: true
	Name *string `json:"name"`

	// rotation type, default is NONE
	// Enum: [FIXED_PERIOD NONE]
	RotationType string `json:"rotationType,omitempty"`

	// start date
	// Required: true
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate"`

	// view id
	ViewID string `json:"viewId,omitempty"`
}

// Validate validates this section update
func (m *SectionUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedPeriodRotationConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRotationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SectionUpdate) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("endDate", "body", strfmt.DateTime(m.EndDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SectionUpdate) validateFixedPeriodRotationConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.FixedPeriodRotationConfig) { // not required
		return nil
	}

	if m.FixedPeriodRotationConfig != nil {
		if err := m.FixedPeriodRotationConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fixedPeriodRotationConfig")
			}
			return err
		}
	}

	return nil
}

func (m *SectionUpdate) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SectionUpdate) validateLocalizations(formats strfmt.Registry) error {

	for k := range m.Localizations {

		if err := validate.Required("localizations"+"."+k, "body", m.Localizations[k]); err != nil {
			return err
		}
		if val, ok := m.Localizations[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SectionUpdate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var sectionUpdateTypeRotationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FIXED_PERIOD","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sectionUpdateTypeRotationTypePropEnum = append(sectionUpdateTypeRotationTypePropEnum, v)
	}
}

const (

	// SectionUpdateRotationTypeFIXEDPERIOD captures enum value "FIXED_PERIOD"
	SectionUpdateRotationTypeFIXEDPERIOD string = "FIXED_PERIOD"

	// SectionUpdateRotationTypeNONE captures enum value "NONE"
	SectionUpdateRotationTypeNONE string = "NONE"
)

// prop value enum
func (m *SectionUpdate) validateRotationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sectionUpdateTypeRotationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SectionUpdate) validateRotationType(formats strfmt.Registry) error {

	if swag.IsZero(m.RotationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRotationTypeEnum("rotationType", "body", m.RotationType); err != nil {
		return err
	}

	return nil
}

func (m *SectionUpdate) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", strfmt.DateTime(m.StartDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SectionUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SectionUpdate) UnmarshalBinary(b []byte) error {
	var res SectionUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
