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

// ItemSnapshot item snapshot
//
// swagger:model ItemSnapshot
type ItemSnapshot struct {

	// App id, required when itemType is APP
	AppID string `json:"appId,omitempty"`

	// App type, required when itemType is APP
	// Enum: [GAME SOFTWARE DLC DEMO]
	AppType string `json:"appType,omitempty"`

	// Base app Id
	BaseAppID string `json:"baseAppId,omitempty"`

	// booth name to get tickets while it's item type is CODE
	BoothName string `json:"boothName,omitempty"`

	// createdAt
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Entitlement type
	// Required: true
	// Enum: [DURABLE CONSUMABLE]
	EntitlementType *string `json:"entitlementType"`

	// supported features
	// Unique: true
	Features []string `json:"features"`

	// Item id
	// Required: true
	ItemID *string `json:"itemId"`

	// itemIds, should be empty if item type is not "BUNDLE"
	ItemIds []string `json:"itemIds"`

	// Item type
	// Required: true
	// Enum: [APP COINS INGAMEITEM BUNDLE CODE SUBSCRIPTION SEASON MEDIA]
	ItemType *string `json:"itemType"`

	// language
	// Required: true
	Language *string `json:"language"`

	// Whether can be visible in Store for public user
	Listable bool `json:"listable,omitempty"`

	// Max count, -1 means UNLIMITED, unset when itemType is CODE
	MaxCount int32 `json:"maxCount,omitempty"`

	// Max count per user, -1 means UNLIMITED
	MaxCountPerUser int32 `json:"maxCountPerUser,omitempty"`

	// Name
	// Required: true
	Name *string `json:"name"`

	// Item's namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// Whether can be purchased
	Purchasable bool `json:"purchasable,omitempty"`

	// recurring for subscription
	Recurring *Recurring `json:"recurring,omitempty"`

	// region
	// Required: true
	Region *string `json:"region"`

	// Region data
	RegionDataItem *RegionDataItem `json:"regionDataItem,omitempty"`

	// Season type, required while itemType is SEASON
	// Enum: [PASS TIER]
	SeasonType string `json:"seasonType,omitempty"`

	// Sku
	Sku string `json:"sku,omitempty"`

	// Whether stack the CONSUMABLE entitlement
	Stackable bool `json:"stackable,omitempty"`

	// Target currency code of coin item
	TargetCurrencyCode string `json:"targetCurrencyCode,omitempty"`

	// Target item id of mapping from game namespace to publisher namespace
	TargetItemID string `json:"targetItemId,omitempty"`

	// The target namespace of a cross namespace item
	TargetNamespace string `json:"targetNamespace,omitempty"`

	// thumbnail url
	ThumbnailURL string `json:"thumbnailUrl,omitempty"`

	// title info
	// Required: true
	Title *string `json:"title"`

	// updatedAt
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Item use count, required when entitlement type is consumable or itemType is COINS
	UseCount int32 `json:"useCount,omitempty"`
}

// Validate validates this item snapshot
func (m *ItemSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlementType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurring(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionDataItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeasonType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
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

var itemSnapshotTypeAppTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GAME","SOFTWARE","DLC","DEMO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemSnapshotTypeAppTypePropEnum = append(itemSnapshotTypeAppTypePropEnum, v)
	}
}

const (

	// ItemSnapshotAppTypeGAME captures enum value "GAME"
	ItemSnapshotAppTypeGAME string = "GAME"

	// ItemSnapshotAppTypeSOFTWARE captures enum value "SOFTWARE"
	ItemSnapshotAppTypeSOFTWARE string = "SOFTWARE"

	// ItemSnapshotAppTypeDLC captures enum value "DLC"
	ItemSnapshotAppTypeDLC string = "DLC"

	// ItemSnapshotAppTypeDEMO captures enum value "DEMO"
	ItemSnapshotAppTypeDEMO string = "DEMO"
)

// prop value enum
func (m *ItemSnapshot) validateAppTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemSnapshotTypeAppTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemSnapshot) validateAppType(formats strfmt.Registry) error {

	if swag.IsZero(m.AppType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppTypeEnum("appType", "body", m.AppType); err != nil {
		return err
	}

	return nil
}

func (m *ItemSnapshot) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var itemSnapshotTypeEntitlementTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DURABLE","CONSUMABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemSnapshotTypeEntitlementTypePropEnum = append(itemSnapshotTypeEntitlementTypePropEnum, v)
	}
}

const (

	// ItemSnapshotEntitlementTypeDURABLE captures enum value "DURABLE"
	ItemSnapshotEntitlementTypeDURABLE string = "DURABLE"

	// ItemSnapshotEntitlementTypeCONSUMABLE captures enum value "CONSUMABLE"
	ItemSnapshotEntitlementTypeCONSUMABLE string = "CONSUMABLE"
)

// prop value enum
func (m *ItemSnapshot) validateEntitlementTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemSnapshotTypeEntitlementTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemSnapshot) validateEntitlementType(formats strfmt.Registry) error {

	if err := validate.Required("entitlementType", "body", m.EntitlementType); err != nil {
		return err
	}

	// value enum
	if err := m.validateEntitlementTypeEnum("entitlementType", "body", *m.EntitlementType); err != nil {
		return err
	}

	return nil
}

func (m *ItemSnapshot) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if err := validate.UniqueItems("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *ItemSnapshot) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("itemId", "body", m.ItemID); err != nil {
		return err
	}

	return nil
}

var itemSnapshotTypeItemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APP","COINS","INGAMEITEM","BUNDLE","CODE","SUBSCRIPTION","SEASON","MEDIA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemSnapshotTypeItemTypePropEnum = append(itemSnapshotTypeItemTypePropEnum, v)
	}
}

const (

	// ItemSnapshotItemTypeAPP captures enum value "APP"
	ItemSnapshotItemTypeAPP string = "APP"

	// ItemSnapshotItemTypeCOINS captures enum value "COINS"
	ItemSnapshotItemTypeCOINS string = "COINS"

	// ItemSnapshotItemTypeINGAMEITEM captures enum value "INGAMEITEM"
	ItemSnapshotItemTypeINGAMEITEM string = "INGAMEITEM"

	// ItemSnapshotItemTypeBUNDLE captures enum value "BUNDLE"
	ItemSnapshotItemTypeBUNDLE string = "BUNDLE"

	// ItemSnapshotItemTypeCODE captures enum value "CODE"
	ItemSnapshotItemTypeCODE string = "CODE"

	// ItemSnapshotItemTypeSUBSCRIPTION captures enum value "SUBSCRIPTION"
	ItemSnapshotItemTypeSUBSCRIPTION string = "SUBSCRIPTION"

	// ItemSnapshotItemTypeSEASON captures enum value "SEASON"
	ItemSnapshotItemTypeSEASON string = "SEASON"

	// ItemSnapshotItemTypeMEDIA captures enum value "MEDIA"
	ItemSnapshotItemTypeMEDIA string = "MEDIA"
)

// prop value enum
func (m *ItemSnapshot) validateItemTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemSnapshotTypeItemTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemSnapshot) validateItemType(formats strfmt.Registry) error {

	if err := validate.Required("itemType", "body", m.ItemType); err != nil {
		return err
	}

	// value enum
	if err := m.validateItemTypeEnum("itemType", "body", *m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *ItemSnapshot) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *ItemSnapshot) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ItemSnapshot) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *ItemSnapshot) validateRecurring(formats strfmt.Registry) error {

	if swag.IsZero(m.Recurring) { // not required
		return nil
	}

	if m.Recurring != nil {
		if err := m.Recurring.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recurring")
			}
			return err
		}
	}

	return nil
}

func (m *ItemSnapshot) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *ItemSnapshot) validateRegionDataItem(formats strfmt.Registry) error {

	if swag.IsZero(m.RegionDataItem) { // not required
		return nil
	}

	if m.RegionDataItem != nil {
		if err := m.RegionDataItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regionDataItem")
			}
			return err
		}
	}

	return nil
}

var itemSnapshotTypeSeasonTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PASS","TIER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemSnapshotTypeSeasonTypePropEnum = append(itemSnapshotTypeSeasonTypePropEnum, v)
	}
}

const (

	// ItemSnapshotSeasonTypePASS captures enum value "PASS"
	ItemSnapshotSeasonTypePASS string = "PASS"

	// ItemSnapshotSeasonTypeTIER captures enum value "TIER"
	ItemSnapshotSeasonTypeTIER string = "TIER"
)

// prop value enum
func (m *ItemSnapshot) validateSeasonTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemSnapshotTypeSeasonTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemSnapshot) validateSeasonType(formats strfmt.Registry) error {

	if swag.IsZero(m.SeasonType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeasonTypeEnum("seasonType", "body", m.SeasonType); err != nil {
		return err
	}

	return nil
}

func (m *ItemSnapshot) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *ItemSnapshot) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemSnapshot) UnmarshalBinary(b []byte) error {
	var res ItemSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
