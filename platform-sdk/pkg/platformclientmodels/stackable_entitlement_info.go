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

// StackableEntitlementInfo stackable entitlement info
//
// swagger:model StackableEntitlementInfo
type StackableEntitlementInfo struct {

	// appId if entitlement is an app
	AppID string `json:"appId,omitempty"`

	// appType if entitlement is an app
	// Enum: [GAME SOFTWARE DLC DEMO]
	AppType string `json:"appType,omitempty"`

	// entitlement class
	// Required: true
	// Enum: [APP ENTITLEMENT DISTRIBUTION CODE SUBSCRIPTION]
	Clazz *string `json:"clazz"`

	// entitlement created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// distributed quantity for distribution, required if DISTRIBUTION
	DistributedQuantity int32 `json:"distributedQuantity,omitempty"`

	// entitlement end date
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// entitlement features
	// Unique: true
	Features []string `json:"features"`

	// entitlement granted at
	// Required: true
	// Format: date-time
	GrantedAt *strfmt.DateTime `json:"grantedAt"`

	// grantedCode, the granted code
	GrantedCode string `json:"grantedCode,omitempty"`

	// entitlement id
	// Required: true
	ID *string `json:"id"`

	// itemId of the entitlement
	// Required: true
	ItemID *string `json:"itemId"`

	// itemNamespace for the purchased item
	// Required: true
	ItemNamespace *string `json:"itemNamespace"`

	// itemSnapshot for distribution
	ItemSnapshot *ItemSnapshot `json:"itemSnapshot,omitempty"`

	// entitlement name
	// Required: true
	Name *string `json:"name"`

	// entitlement namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// purchased quantity for distribution, required if DISTRIBUTION
	Quantity int32 `json:"quantity,omitempty"`

	// sku for purchased item
	Sku string `json:"sku,omitempty"`

	// entitlement source
	// Required: true
	// Enum: [PURCHASE IAP PROMOTION ACHIEVEMENT REFERRAL_BONUS REDEEM_CODE REWARD GIFT OTHER]
	Source *string `json:"source"`

	// whether the CONSUMABLE entitlement is stackable
	Stackable bool `json:"stackable,omitempty"`

	// DISTRIBUTION entitlement stacked quantity
	StackedQuantity int32 `json:"stackedQuantity,omitempty"`

	// CONSUMABLE entitlement stacked use count
	StackedUseCount int32 `json:"stackedUseCount,omitempty"`

	// entitlement start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// entitlement status
	// Required: true
	// Enum: [ACTIVE INACTIVE CONSUMED DISTRIBUTED REVOKED]
	Status *string `json:"status"`

	// storeId of the item, published store if omitted
	StoreID string `json:"storeId,omitempty"`

	// targetNamespace for distribution
	TargetNamespace string `json:"targetNamespace,omitempty"`

	// entitlement type
	// Required: true
	// Enum: [DURABLE CONSUMABLE]
	Type *string `json:"type"`

	// entitlement updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`

	// useCount for entitlement
	UseCount int32 `json:"useCount,omitempty"`

	// userId for this entitlement
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this stackable entitlement info
func (m *StackableEntitlementInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClazz(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

var stackableEntitlementInfoTypeAppTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GAME","SOFTWARE","DLC","DEMO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackableEntitlementInfoTypeAppTypePropEnum = append(stackableEntitlementInfoTypeAppTypePropEnum, v)
	}
}

const (

	// StackableEntitlementInfoAppTypeGAME captures enum value "GAME"
	StackableEntitlementInfoAppTypeGAME string = "GAME"

	// StackableEntitlementInfoAppTypeSOFTWARE captures enum value "SOFTWARE"
	StackableEntitlementInfoAppTypeSOFTWARE string = "SOFTWARE"

	// StackableEntitlementInfoAppTypeDLC captures enum value "DLC"
	StackableEntitlementInfoAppTypeDLC string = "DLC"

	// StackableEntitlementInfoAppTypeDEMO captures enum value "DEMO"
	StackableEntitlementInfoAppTypeDEMO string = "DEMO"
)

// prop value enum
func (m *StackableEntitlementInfo) validateAppTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stackableEntitlementInfoTypeAppTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StackableEntitlementInfo) validateAppType(formats strfmt.Registry) error {

	if swag.IsZero(m.AppType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppTypeEnum("appType", "body", m.AppType); err != nil {
		return err
	}

	return nil
}

var stackableEntitlementInfoTypeClazzPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APP","ENTITLEMENT","DISTRIBUTION","CODE","SUBSCRIPTION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackableEntitlementInfoTypeClazzPropEnum = append(stackableEntitlementInfoTypeClazzPropEnum, v)
	}
}

const (

	// StackableEntitlementInfoClazzAPP captures enum value "APP"
	StackableEntitlementInfoClazzAPP string = "APP"

	// StackableEntitlementInfoClazzENTITLEMENT captures enum value "ENTITLEMENT"
	StackableEntitlementInfoClazzENTITLEMENT string = "ENTITLEMENT"

	// StackableEntitlementInfoClazzDISTRIBUTION captures enum value "DISTRIBUTION"
	StackableEntitlementInfoClazzDISTRIBUTION string = "DISTRIBUTION"

	// StackableEntitlementInfoClazzCODE captures enum value "CODE"
	StackableEntitlementInfoClazzCODE string = "CODE"

	// StackableEntitlementInfoClazzSUBSCRIPTION captures enum value "SUBSCRIPTION"
	StackableEntitlementInfoClazzSUBSCRIPTION string = "SUBSCRIPTION"
)

// prop value enum
func (m *StackableEntitlementInfo) validateClazzEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stackableEntitlementInfoTypeClazzPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StackableEntitlementInfo) validateClazz(formats strfmt.Registry) error {

	if err := validate.Required("clazz", "body", m.Clazz); err != nil {
		return err
	}

	// value enum
	if err := m.validateClazzEnum("clazz", "body", *m.Clazz); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if err := validate.UniqueItems("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateGrantedAt(formats strfmt.Registry) error {

	if err := validate.Required("grantedAt", "body", m.GrantedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("grantedAt", "body", "date-time", m.GrantedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("itemId", "body", m.ItemID); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateItemNamespace(formats strfmt.Registry) error {

	if err := validate.Required("itemNamespace", "body", m.ItemNamespace); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateItemSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemSnapshot) { // not required
		return nil
	}

	if m.ItemSnapshot != nil {
		if err := m.ItemSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemSnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *StackableEntitlementInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

var stackableEntitlementInfoTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PURCHASE","IAP","PROMOTION","ACHIEVEMENT","REFERRAL_BONUS","REDEEM_CODE","REWARD","GIFT","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackableEntitlementInfoTypeSourcePropEnum = append(stackableEntitlementInfoTypeSourcePropEnum, v)
	}
}

const (

	// StackableEntitlementInfoSourcePURCHASE captures enum value "PURCHASE"
	StackableEntitlementInfoSourcePURCHASE string = "PURCHASE"

	// StackableEntitlementInfoSourceIAP captures enum value "IAP"
	StackableEntitlementInfoSourceIAP string = "IAP"

	// StackableEntitlementInfoSourcePROMOTION captures enum value "PROMOTION"
	StackableEntitlementInfoSourcePROMOTION string = "PROMOTION"

	// StackableEntitlementInfoSourceACHIEVEMENT captures enum value "ACHIEVEMENT"
	StackableEntitlementInfoSourceACHIEVEMENT string = "ACHIEVEMENT"

	// StackableEntitlementInfoSourceREFERRALBONUS captures enum value "REFERRAL_BONUS"
	StackableEntitlementInfoSourceREFERRALBONUS string = "REFERRAL_BONUS"

	// StackableEntitlementInfoSourceREDEEMCODE captures enum value "REDEEM_CODE"
	StackableEntitlementInfoSourceREDEEMCODE string = "REDEEM_CODE"

	// StackableEntitlementInfoSourceREWARD captures enum value "REWARD"
	StackableEntitlementInfoSourceREWARD string = "REWARD"

	// StackableEntitlementInfoSourceGIFT captures enum value "GIFT"
	StackableEntitlementInfoSourceGIFT string = "GIFT"

	// StackableEntitlementInfoSourceOTHER captures enum value "OTHER"
	StackableEntitlementInfoSourceOTHER string = "OTHER"
)

// prop value enum
func (m *StackableEntitlementInfo) validateSourceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stackableEntitlementInfoTypeSourcePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StackableEntitlementInfo) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", *m.Source); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var stackableEntitlementInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE","CONSUMED","DISTRIBUTED","REVOKED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackableEntitlementInfoTypeStatusPropEnum = append(stackableEntitlementInfoTypeStatusPropEnum, v)
	}
}

const (

	// StackableEntitlementInfoStatusACTIVE captures enum value "ACTIVE"
	StackableEntitlementInfoStatusACTIVE string = "ACTIVE"

	// StackableEntitlementInfoStatusINACTIVE captures enum value "INACTIVE"
	StackableEntitlementInfoStatusINACTIVE string = "INACTIVE"

	// StackableEntitlementInfoStatusCONSUMED captures enum value "CONSUMED"
	StackableEntitlementInfoStatusCONSUMED string = "CONSUMED"

	// StackableEntitlementInfoStatusDISTRIBUTED captures enum value "DISTRIBUTED"
	StackableEntitlementInfoStatusDISTRIBUTED string = "DISTRIBUTED"

	// StackableEntitlementInfoStatusREVOKED captures enum value "REVOKED"
	StackableEntitlementInfoStatusREVOKED string = "REVOKED"
)

// prop value enum
func (m *StackableEntitlementInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stackableEntitlementInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StackableEntitlementInfo) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var stackableEntitlementInfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DURABLE","CONSUMABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackableEntitlementInfoTypeTypePropEnum = append(stackableEntitlementInfoTypeTypePropEnum, v)
	}
}

const (

	// StackableEntitlementInfoTypeDURABLE captures enum value "DURABLE"
	StackableEntitlementInfoTypeDURABLE string = "DURABLE"

	// StackableEntitlementInfoTypeCONSUMABLE captures enum value "CONSUMABLE"
	StackableEntitlementInfoTypeCONSUMABLE string = "CONSUMABLE"
)

// prop value enum
func (m *StackableEntitlementInfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stackableEntitlementInfoTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StackableEntitlementInfo) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StackableEntitlementInfo) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackableEntitlementInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackableEntitlementInfo) UnmarshalBinary(b []byte) error {
	var res StackableEntitlementInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
