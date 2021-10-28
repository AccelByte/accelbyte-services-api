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

// PopulatedItemInfo populated item info
//
// swagger:model PopulatedItemInfo
type PopulatedItemInfo struct {

	// App id, required when itemType is APP
	AppID string `json:"appId,omitempty"`

	// App type, required when itemType is APP
	// Enum: [GAME SOFTWARE DLC DEMO]
	AppType string `json:"appType,omitempty"`

	// Base app id
	BaseAppID string `json:"baseAppId,omitempty"`

	// booth name to get tickets while it's item type is CODE
	BoothName string `json:"boothName,omitempty"`

	// the items which this item being bounded to
	// Unique: true
	BoundItemIds []string `json:"boundItemIds"`

	// Item category path
	// Required: true
	CategoryPath *string `json:"categoryPath"`

	// customized item clazz
	Clazz string `json:"clazz,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// description info
	Description string `json:"description,omitempty"`

	// display order
	DisplayOrder int32 `json:"displayOrder,omitempty"`

	// Entitlement type
	// Required: true
	// Enum: [DURABLE CONSUMABLE]
	EntitlementType *string `json:"entitlementType"`

	// customized item properties
	Ext map[string]interface{} `json:"ext,omitempty"`

	// Features
	// Unique: true
	Features []string `json:"features"`

	// images
	Images []*Image `json:"images"`

	// Item id
	// Required: true
	ItemID *string `json:"itemId"`

	// itemIds, should be empty if item type is not "BUNDLE"
	ItemIds []string `json:"itemIds"`

	// Item type
	// Required: true
	// Enum: [APP COINS INGAMEITEM BUNDLE CODE SUBSCRIPTION SEASON MEDIA]
	ItemType *string `json:"itemType"`

	// bundle items, only has value when item is bundle and is populated
	Items []*ItemInfo `json:"items"`

	// language
	// Required: true
	Language *string `json:"language"`

	// Whether can be visible in Store for public user
	Listable bool `json:"listable,omitempty"`

	// local ext
	LocalExt map[string]interface{} `json:"localExt,omitempty"`

	// long description info
	LongDescription string `json:"longDescription,omitempty"`

	// Max count, -1 means UNLIMITED, unset when itemType is CODE
	MaxCount int32 `json:"maxCount,omitempty"`

	// Max count per user, -1 means UNLIMITED
	MaxCountPerUser int32 `json:"maxCountPerUser,omitempty"`

	// Name
	// Required: true
	Name *string `json:"name"`

	// Item namespace
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
	RegionData []*RegionDataItem `json:"regionData"`

	// Season type, required while itemType is SEASON
	// Enum: [PASS TIER]
	SeasonType string `json:"seasonType,omitempty"`

	// Sku
	Sku string `json:"sku,omitempty"`

	// Whether stack the CONSUMABLE entitlement
	Stackable bool `json:"stackable,omitempty"`

	// Item status
	// Required: true
	// Enum: [ACTIVE INACTIVE]
	Status *string `json:"status"`

	// Tags
	// Unique: true
	Tags []string `json:"tags"`

	// Target currency code of coin item
	TargetCurrencyCode string `json:"targetCurrencyCode,omitempty"`

	// Target item id if this item is mapping from game namespace
	TargetItemID string `json:"targetItemId,omitempty"`

	// The target namespace of a cross namespace item
	TargetNamespace string `json:"targetNamespace,omitempty"`

	// thumbnail url
	ThumbnailURL string `json:"thumbnailUrl,omitempty"`

	// title info
	// Required: true
	Title *string `json:"title"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`

	// Item use count, required when entitlement type is consumable or itemType is COINS
	UseCount int32 `json:"useCount,omitempty"`
}

// Validate validates this populated item info
func (m *PopulatedItemInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBoundItemIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategoryPath(formats); err != nil {
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

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
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

	if err := m.validateRegionData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeasonType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

var populatedItemInfoTypeAppTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GAME","SOFTWARE","DLC","DEMO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		populatedItemInfoTypeAppTypePropEnum = append(populatedItemInfoTypeAppTypePropEnum, v)
	}
}

const (

	// PopulatedItemInfoAppTypeGAME captures enum value "GAME"
	PopulatedItemInfoAppTypeGAME string = "GAME"

	// PopulatedItemInfoAppTypeSOFTWARE captures enum value "SOFTWARE"
	PopulatedItemInfoAppTypeSOFTWARE string = "SOFTWARE"

	// PopulatedItemInfoAppTypeDLC captures enum value "DLC"
	PopulatedItemInfoAppTypeDLC string = "DLC"

	// PopulatedItemInfoAppTypeDEMO captures enum value "DEMO"
	PopulatedItemInfoAppTypeDEMO string = "DEMO"
)

// prop value enum
func (m *PopulatedItemInfo) validateAppTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, populatedItemInfoTypeAppTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PopulatedItemInfo) validateAppType(formats strfmt.Registry) error {

	if swag.IsZero(m.AppType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppTypeEnum("appType", "body", m.AppType); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateBoundItemIds(formats strfmt.Registry) error {

	if swag.IsZero(m.BoundItemIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("boundItemIds", "body", m.BoundItemIds); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateCategoryPath(formats strfmt.Registry) error {

	if err := validate.Required("categoryPath", "body", m.CategoryPath); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var populatedItemInfoTypeEntitlementTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DURABLE","CONSUMABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		populatedItemInfoTypeEntitlementTypePropEnum = append(populatedItemInfoTypeEntitlementTypePropEnum, v)
	}
}

const (

	// PopulatedItemInfoEntitlementTypeDURABLE captures enum value "DURABLE"
	PopulatedItemInfoEntitlementTypeDURABLE string = "DURABLE"

	// PopulatedItemInfoEntitlementTypeCONSUMABLE captures enum value "CONSUMABLE"
	PopulatedItemInfoEntitlementTypeCONSUMABLE string = "CONSUMABLE"
)

// prop value enum
func (m *PopulatedItemInfo) validateEntitlementTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, populatedItemInfoTypeEntitlementTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PopulatedItemInfo) validateEntitlementType(formats strfmt.Registry) error {

	if err := validate.Required("entitlementType", "body", m.EntitlementType); err != nil {
		return err
	}

	// value enum
	if err := m.validateEntitlementTypeEnum("entitlementType", "body", *m.EntitlementType); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if err := validate.UniqueItems("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PopulatedItemInfo) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("itemId", "body", m.ItemID); err != nil {
		return err
	}

	return nil
}

var populatedItemInfoTypeItemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APP","COINS","INGAMEITEM","BUNDLE","CODE","SUBSCRIPTION","SEASON","MEDIA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		populatedItemInfoTypeItemTypePropEnum = append(populatedItemInfoTypeItemTypePropEnum, v)
	}
}

const (

	// PopulatedItemInfoItemTypeAPP captures enum value "APP"
	PopulatedItemInfoItemTypeAPP string = "APP"

	// PopulatedItemInfoItemTypeCOINS captures enum value "COINS"
	PopulatedItemInfoItemTypeCOINS string = "COINS"

	// PopulatedItemInfoItemTypeINGAMEITEM captures enum value "INGAMEITEM"
	PopulatedItemInfoItemTypeINGAMEITEM string = "INGAMEITEM"

	// PopulatedItemInfoItemTypeBUNDLE captures enum value "BUNDLE"
	PopulatedItemInfoItemTypeBUNDLE string = "BUNDLE"

	// PopulatedItemInfoItemTypeCODE captures enum value "CODE"
	PopulatedItemInfoItemTypeCODE string = "CODE"

	// PopulatedItemInfoItemTypeSUBSCRIPTION captures enum value "SUBSCRIPTION"
	PopulatedItemInfoItemTypeSUBSCRIPTION string = "SUBSCRIPTION"

	// PopulatedItemInfoItemTypeSEASON captures enum value "SEASON"
	PopulatedItemInfoItemTypeSEASON string = "SEASON"

	// PopulatedItemInfoItemTypeMEDIA captures enum value "MEDIA"
	PopulatedItemInfoItemTypeMEDIA string = "MEDIA"
)

// prop value enum
func (m *PopulatedItemInfo) validateItemTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, populatedItemInfoTypeItemTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PopulatedItemInfo) validateItemType(formats strfmt.Registry) error {

	if err := validate.Required("itemType", "body", m.ItemType); err != nil {
		return err
	}

	// value enum
	if err := m.validateItemTypeEnum("itemType", "body", *m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateItems(formats strfmt.Registry) error {

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

func (m *PopulatedItemInfo) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateRecurring(formats strfmt.Registry) error {

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

func (m *PopulatedItemInfo) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateRegionData(formats strfmt.Registry) error {

	if swag.IsZero(m.RegionData) { // not required
		return nil
	}

	for i := 0; i < len(m.RegionData); i++ {
		if swag.IsZero(m.RegionData[i]) { // not required
			continue
		}

		if m.RegionData[i] != nil {
			if err := m.RegionData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regionData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var populatedItemInfoTypeSeasonTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PASS","TIER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		populatedItemInfoTypeSeasonTypePropEnum = append(populatedItemInfoTypeSeasonTypePropEnum, v)
	}
}

const (

	// PopulatedItemInfoSeasonTypePASS captures enum value "PASS"
	PopulatedItemInfoSeasonTypePASS string = "PASS"

	// PopulatedItemInfoSeasonTypeTIER captures enum value "TIER"
	PopulatedItemInfoSeasonTypeTIER string = "TIER"
)

// prop value enum
func (m *PopulatedItemInfo) validateSeasonTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, populatedItemInfoTypeSeasonTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PopulatedItemInfo) validateSeasonType(formats strfmt.Registry) error {

	if swag.IsZero(m.SeasonType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeasonTypeEnum("seasonType", "body", m.SeasonType); err != nil {
		return err
	}

	return nil
}

var populatedItemInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		populatedItemInfoTypeStatusPropEnum = append(populatedItemInfoTypeStatusPropEnum, v)
	}
}

const (

	// PopulatedItemInfoStatusACTIVE captures enum value "ACTIVE"
	PopulatedItemInfoStatusACTIVE string = "ACTIVE"

	// PopulatedItemInfoStatusINACTIVE captures enum value "INACTIVE"
	PopulatedItemInfoStatusINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *PopulatedItemInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, populatedItemInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PopulatedItemInfo) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *PopulatedItemInfo) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PopulatedItemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PopulatedItemInfo) UnmarshalBinary(b []byte) error {
	var res PopulatedItemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
