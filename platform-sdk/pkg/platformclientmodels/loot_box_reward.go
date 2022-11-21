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

// LootBoxReward loot box reward
//
// swagger:model LootBoxReward
type LootBoxReward struct {

	// reward items, if type is PROBABILITY_GROUP, will random get one of items in it, if type is REWARD_GROUP/REWARD, will grant all items to user when roll this reward
	LootBoxItems []*BoxItem `json:"lootBoxItems"`

	// reward name
	Name string `json:"name,omitempty"`

	// odds, automatic calculation based on weights.
	Odds float64 `json:"odds,omitempty"`

	// reward type
	// Enum: [PROBABILITY_GROUP REWARD REWARD_GROUP]
	Type string `json:"type,omitempty"`

	// reward weight
	Weight int32 `json:"weight,omitempty"`
}

// Validate validates this loot box reward
func (m *LootBoxReward) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLootBoxItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LootBoxReward) validateLootBoxItems(formats strfmt.Registry) error {

	if swag.IsZero(m.LootBoxItems) { // not required
		return nil
	}

	for i := 0; i < len(m.LootBoxItems); i++ {
		if swag.IsZero(m.LootBoxItems[i]) { // not required
			continue
		}

		if m.LootBoxItems[i] != nil {
			if err := m.LootBoxItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lootBoxItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var lootBoxRewardTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROBABILITY_GROUP","REWARD","REWARD_GROUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lootBoxRewardTypeTypePropEnum = append(lootBoxRewardTypeTypePropEnum, v)
	}
}

const (

	// LootBoxRewardTypePROBABILITYGROUP captures enum value "PROBABILITY_GROUP"
	LootBoxRewardTypePROBABILITYGROUP string = "PROBABILITY_GROUP"

	// LootBoxRewardTypeREWARD captures enum value "REWARD"
	LootBoxRewardTypeREWARD string = "REWARD"

	// LootBoxRewardTypeREWARDGROUP captures enum value "REWARD_GROUP"
	LootBoxRewardTypeREWARDGROUP string = "REWARD_GROUP"
)

// prop value enum
func (m *LootBoxReward) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, lootBoxRewardTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LootBoxReward) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LootBoxReward) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LootBoxReward) UnmarshalBinary(b []byte) error {
	var res LootBoxReward
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
