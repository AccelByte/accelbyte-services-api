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

// EntitlementSummary entitlement summary
//
// swagger:model EntitlementSummary
type EntitlementSummary struct {

	// entitlement class
	// Required: true
	// Enum: [APP ENTITLEMENT DISTRIBUTION CODE SUBSCRIPTION]
	Clazz *string `json:"clazz"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// end date time
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// granted code
	GrantedCode string `json:"grantedCode,omitempty"`

	// entitlement id
	// Required: true
	ID *string `json:"id"`

	// item id
	ItemID string `json:"itemId,omitempty"`

	// entitlement namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// Whether the CONSUMABLE entitlement is stackable
	Stackable bool `json:"stackable,omitempty"`

	// DISTRIBUTION entitlement stacked quantity
	StackedQuantity int32 `json:"stackedQuantity,omitempty"`

	// CONSUMABLE entitlement stacked use count
	StackedUseCount int32 `json:"stackedUseCount,omitempty"`

	// start date time
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// item store id, null if published store
	StoreID string `json:"storeId,omitempty"`

	// entitlement type
	// Required: true
	// Enum: [DURABLE CONSUMABLE]
	Type *string `json:"type"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`

	// userId for this entitlement
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this entitlement summary
func (m *EntitlementSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClazz(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var entitlementSummaryTypeClazzPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APP","ENTITLEMENT","DISTRIBUTION","CODE","SUBSCRIPTION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		entitlementSummaryTypeClazzPropEnum = append(entitlementSummaryTypeClazzPropEnum, v)
	}
}

const (

	// EntitlementSummaryClazzAPP captures enum value "APP"
	EntitlementSummaryClazzAPP string = "APP"

	// EntitlementSummaryClazzENTITLEMENT captures enum value "ENTITLEMENT"
	EntitlementSummaryClazzENTITLEMENT string = "ENTITLEMENT"

	// EntitlementSummaryClazzDISTRIBUTION captures enum value "DISTRIBUTION"
	EntitlementSummaryClazzDISTRIBUTION string = "DISTRIBUTION"

	// EntitlementSummaryClazzCODE captures enum value "CODE"
	EntitlementSummaryClazzCODE string = "CODE"

	// EntitlementSummaryClazzSUBSCRIPTION captures enum value "SUBSCRIPTION"
	EntitlementSummaryClazzSUBSCRIPTION string = "SUBSCRIPTION"
)

// prop value enum
func (m *EntitlementSummary) validateClazzEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, entitlementSummaryTypeClazzPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EntitlementSummary) validateClazz(formats strfmt.Registry) error {

	if err := validate.Required("clazz", "body", m.Clazz); err != nil {
		return err
	}

	// value enum
	if err := m.validateClazzEnum("clazz", "body", *m.Clazz); err != nil {
		return err
	}

	return nil
}

func (m *EntitlementSummary) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EntitlementSummary) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EntitlementSummary) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *EntitlementSummary) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *EntitlementSummary) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var entitlementSummaryTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DURABLE","CONSUMABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		entitlementSummaryTypeTypePropEnum = append(entitlementSummaryTypeTypePropEnum, v)
	}
}

const (

	// EntitlementSummaryTypeDURABLE captures enum value "DURABLE"
	EntitlementSummaryTypeDURABLE string = "DURABLE"

	// EntitlementSummaryTypeCONSUMABLE captures enum value "CONSUMABLE"
	EntitlementSummaryTypeCONSUMABLE string = "CONSUMABLE"
)

// prop value enum
func (m *EntitlementSummary) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, entitlementSummaryTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EntitlementSummary) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *EntitlementSummary) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EntitlementSummary) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntitlementSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntitlementSummary) UnmarshalBinary(b []byte) error {
	var res EntitlementSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
