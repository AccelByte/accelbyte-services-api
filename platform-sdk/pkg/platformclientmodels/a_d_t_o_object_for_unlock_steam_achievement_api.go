// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ADTOObjectForUnlockSteamAchievementAPI a d t o object for unlock steam achievement API
//
// swagger:model A DTO object for unlock steam achievement API
type ADTOObjectForUnlockSteamAchievementAPI struct {

	// achievements to be updated
	Achievements []*SteamAchievementRequest `json:"achievements"`

	// steam user id
	SteamUserID string `json:"steamUserId,omitempty"`
}

// Validate validates this a d t o object for unlock steam achievement API
func (m *ADTOObjectForUnlockSteamAchievementAPI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAchievements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ADTOObjectForUnlockSteamAchievementAPI) validateAchievements(formats strfmt.Registry) error {

	if swag.IsZero(m.Achievements) { // not required
		return nil
	}

	for i := 0; i < len(m.Achievements); i++ {
		if swag.IsZero(m.Achievements[i]) { // not required
			continue
		}

		if m.Achievements[i] != nil {
			if err := m.Achievements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("achievements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ADTOObjectForUnlockSteamAchievementAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ADTOObjectForUnlockSteamAchievementAPI) UnmarshalBinary(b []byte) error {
	var res ADTOObjectForUnlockSteamAchievementAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
