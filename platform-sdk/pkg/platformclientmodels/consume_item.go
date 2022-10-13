// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsumeItem consume item
//
// swagger:model ConsumeItem
type ConsumeItem struct {

	// ext item Id
	ExtItemID string `json:"extItemId,omitempty"`

	// item identity
	ItemIdentity string `json:"itemIdentity,omitempty"`

	// item identity type
	// Enum: [ITEM_ID ITEM_SKU]
	ItemIdentityType string `json:"itemIdentityType,omitempty"`
}

// Validate validates this consume item
func (m *ConsumeItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemIdentityType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consumeItemTypeItemIdentityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ITEM_ID","ITEM_SKU"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consumeItemTypeItemIdentityTypePropEnum = append(consumeItemTypeItemIdentityTypePropEnum, v)
	}
}

const (

	// ConsumeItemItemIdentityTypeITEMID captures enum value "ITEM_ID"
	ConsumeItemItemIdentityTypeITEMID string = "ITEM_ID"

	// ConsumeItemItemIdentityTypeITEMSKU captures enum value "ITEM_SKU"
	ConsumeItemItemIdentityTypeITEMSKU string = "ITEM_SKU"
)

// prop value enum
func (m *ConsumeItem) validateItemIdentityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consumeItemTypeItemIdentityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsumeItem) validateItemIdentityType(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemIdentityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateItemIdentityTypeEnum("itemIdentityType", "body", m.ItemIdentityType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsumeItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsumeItem) UnmarshalBinary(b []byte) error {
	var res ConsumeItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
