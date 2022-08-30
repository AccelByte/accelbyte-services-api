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

// BundledItemInfo bundled item info
//
// swagger:model BundledItemInfo
type BundledItemInfo struct {

	// App id, required when itemType is APP
	AppID string `json:"appId,omitempty"`

	// App type, required when itemType is APP
	// Enum: [DEMO DLC GAME SOFTWARE]
	AppType string `json:"appType,omitempty"`

	// Base app id
	BaseAppID string `json:"baseAppId,omitempty"`

	// booth name to get tickets while it's item type is CODE
	BoothName string `json:"boothName,omitempty"`

	// the items which this item being bounded to
	// Unique: true
	BoundItemIds []string `json:"boundItemIds"`

	// bundled qty
	BundledQty int32 `json:"bundledQty,omitempty"`

	// Item category path
	// Required: true
	CategoryPath *string `json:"categoryPath"`

	// customized item clazz
	Clazz string `json:"clazz,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// description info
	Description string `json:"description,omitempty"`

	// display order
	DisplayOrder int32 `json:"displayOrder,omitempty"`

	// Entitlement type
	// Required: true
	// Enum: [CONSUMABLE DURABLE]
	EntitlementType *string `json:"entitlementType"`

	// customized item properties
	Ext map[string]interface{} `json:"ext,omitempty"`

	// Features
	// Unique: true
	Features []string `json:"features"`

	// fresh
	Fresh bool `json:"fresh"`

	// images
	Images []*Image `json:"images"`

	// Item id
	// Required: true
	ItemID *string `json:"itemId"`

	// itemIds, should be empty if item type is not "BUNDLE"
	ItemIds []string `json:"itemIds"`

	// itemQty, the quantity of item in itemIds, default 1, should be empty if item type is not "BUNDLE"
	ItemQty map[string]int32 `json:"itemQty,omitempty"`

	// Item type
	// Required: true
	// Enum: [APP BUNDLE CODE COINS EXTENSION INGAMEITEM MEDIA OPTIONBOX SEASON SUBSCRIPTION]
	ItemType *string `json:"itemType"`

	// language
	// Required: true
	Language *string `json:"language"`

	// Whether can be visible in Store for public user
	Listable bool `json:"listable"`

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

	// option box config
	OptionBoxConfig *OptionBoxConfig `json:"optionBoxConfig,omitempty"`

	// Whether can be purchased
	Purchasable bool `json:"purchasable"`

	// purchase condition
	PurchaseCondition *PurchaseCondition `json:"purchaseCondition,omitempty"`

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
	Stackable bool `json:"stackable"`

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
	UpdatedAt strfmt.DateTime `json:"updatedAt"`

	// Item use count, required when entitlement type is consumable or itemType is COINS
	UseCount int32 `json:"useCount,omitempty"`
}

// Validate validates this bundled item info
func (m *BundledItemInfo) Validate(formats strfmt.Registry) error {
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

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionBoxConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseCondition(formats); err != nil {
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

var bundledItemInfoTypeAppTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEMO","DLC","GAME","SOFTWARE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bundledItemInfoTypeAppTypePropEnum = append(bundledItemInfoTypeAppTypePropEnum, v)
	}
}

const (

	// BundledItemInfoAppTypeDEMO captures enum value "DEMO"
	BundledItemInfoAppTypeDEMO string = "DEMO"

	// BundledItemInfoAppTypeDLC captures enum value "DLC"
	BundledItemInfoAppTypeDLC string = "DLC"

	// BundledItemInfoAppTypeGAME captures enum value "GAME"
	BundledItemInfoAppTypeGAME string = "GAME"

	// BundledItemInfoAppTypeSOFTWARE captures enum value "SOFTWARE"
	BundledItemInfoAppTypeSOFTWARE string = "SOFTWARE"
)

// prop value enum
func (m *BundledItemInfo) validateAppTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bundledItemInfoTypeAppTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BundledItemInfo) validateAppType(formats strfmt.Registry) error {

	if swag.IsZero(m.AppType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAppTypeEnum("appType", "body", m.AppType); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateBoundItemIds(formats strfmt.Registry) error {

	if swag.IsZero(m.BoundItemIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("boundItemIds", "body", m.BoundItemIds); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateCategoryPath(formats strfmt.Registry) error {

	if err := validate.Required("categoryPath", "body", m.CategoryPath); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var bundledItemInfoTypeEntitlementTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONSUMABLE","DURABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bundledItemInfoTypeEntitlementTypePropEnum = append(bundledItemInfoTypeEntitlementTypePropEnum, v)
	}
}

const (

	// BundledItemInfoEntitlementTypeCONSUMABLE captures enum value "CONSUMABLE"
	BundledItemInfoEntitlementTypeCONSUMABLE string = "CONSUMABLE"

	// BundledItemInfoEntitlementTypeDURABLE captures enum value "DURABLE"
	BundledItemInfoEntitlementTypeDURABLE string = "DURABLE"
)

// prop value enum
func (m *BundledItemInfo) validateEntitlementTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bundledItemInfoTypeEntitlementTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BundledItemInfo) validateEntitlementType(formats strfmt.Registry) error {

	if err := validate.Required("entitlementType", "body", m.EntitlementType); err != nil {
		return err
	}

	// value enum
	if err := m.validateEntitlementTypeEnum("entitlementType", "body", *m.EntitlementType); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if err := validate.UniqueItems("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateImages(formats strfmt.Registry) error {

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

func (m *BundledItemInfo) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("itemId", "body", m.ItemID); err != nil {
		return err
	}

	return nil
}

var bundledItemInfoTypeItemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APP","BUNDLE","CODE","COINS","EXTENSION","INGAMEITEM","MEDIA","OPTIONBOX","SEASON","SUBSCRIPTION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bundledItemInfoTypeItemTypePropEnum = append(bundledItemInfoTypeItemTypePropEnum, v)
	}
}

const (

	// BundledItemInfoItemTypeAPP captures enum value "APP"
	BundledItemInfoItemTypeAPP string = "APP"

	// BundledItemInfoItemTypeBUNDLE captures enum value "BUNDLE"
	BundledItemInfoItemTypeBUNDLE string = "BUNDLE"

	// BundledItemInfoItemTypeCODE captures enum value "CODE"
	BundledItemInfoItemTypeCODE string = "CODE"

	// BundledItemInfoItemTypeCOINS captures enum value "COINS"
	BundledItemInfoItemTypeCOINS string = "COINS"

	// BundledItemInfoItemTypeEXTENSION captures enum value "EXTENSION"
	BundledItemInfoItemTypeEXTENSION string = "EXTENSION"

	// BundledItemInfoItemTypeINGAMEITEM captures enum value "INGAMEITEM"
	BundledItemInfoItemTypeINGAMEITEM string = "INGAMEITEM"

	// BundledItemInfoItemTypeMEDIA captures enum value "MEDIA"
	BundledItemInfoItemTypeMEDIA string = "MEDIA"

	// BundledItemInfoItemTypeOPTIONBOX captures enum value "OPTIONBOX"
	BundledItemInfoItemTypeOPTIONBOX string = "OPTIONBOX"

	// BundledItemInfoItemTypeSEASON captures enum value "SEASON"
	BundledItemInfoItemTypeSEASON string = "SEASON"

	// BundledItemInfoItemTypeSUBSCRIPTION captures enum value "SUBSCRIPTION"
	BundledItemInfoItemTypeSUBSCRIPTION string = "SUBSCRIPTION"
)

// prop value enum
func (m *BundledItemInfo) validateItemTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bundledItemInfoTypeItemTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BundledItemInfo) validateItemType(formats strfmt.Registry) error {

	if err := validate.Required("itemType", "body", m.ItemType); err != nil {
		return err
	}

	// value enum
	if err := m.validateItemTypeEnum("itemType", "body", *m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateOptionBoxConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.OptionBoxConfig) { // not required
		return nil
	}

	if m.OptionBoxConfig != nil {
		if err := m.OptionBoxConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("optionBoxConfig")
			}
			return err
		}
	}

	return nil
}

func (m *BundledItemInfo) validatePurchaseCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.PurchaseCondition) { // not required
		return nil
	}

	if m.PurchaseCondition != nil {
		if err := m.PurchaseCondition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("purchaseCondition")
			}
			return err
		}
	}

	return nil
}

func (m *BundledItemInfo) validateRecurring(formats strfmt.Registry) error {

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

func (m *BundledItemInfo) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateRegionData(formats strfmt.Registry) error {

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

var bundledItemInfoTypeSeasonTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PASS","TIER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bundledItemInfoTypeSeasonTypePropEnum = append(bundledItemInfoTypeSeasonTypePropEnum, v)
	}
}

const (

	// BundledItemInfoSeasonTypePASS captures enum value "PASS"
	BundledItemInfoSeasonTypePASS string = "PASS"

	// BundledItemInfoSeasonTypeTIER captures enum value "TIER"
	BundledItemInfoSeasonTypeTIER string = "TIER"
)

// prop value enum
func (m *BundledItemInfo) validateSeasonTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bundledItemInfoTypeSeasonTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BundledItemInfo) validateSeasonType(formats strfmt.Registry) error {

	if swag.IsZero(m.SeasonType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeasonTypeEnum("seasonType", "body", m.SeasonType); err != nil {
		return err
	}

	return nil
}

var bundledItemInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bundledItemInfoTypeStatusPropEnum = append(bundledItemInfoTypeStatusPropEnum, v)
	}
}

const (

	// BundledItemInfoStatusACTIVE captures enum value "ACTIVE"
	BundledItemInfoStatusACTIVE string = "ACTIVE"

	// BundledItemInfoStatusINACTIVE captures enum value "INACTIVE"
	BundledItemInfoStatusINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *BundledItemInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bundledItemInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BundledItemInfo) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *BundledItemInfo) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BundledItemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BundledItemInfo) UnmarshalBinary(b []byte) error {
	var res BundledItemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
