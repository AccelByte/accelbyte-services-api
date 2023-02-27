// Code generated by go-swagger; DO NOT EDIT.

package sessionbrowserclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsGetSessionHistorySearchResponseItemV2 models get session history search response item v2
//
// swagger:model models.GetSessionHistorySearchResponseItemV2
type ModelsGetSessionHistorySearchResponseItemV2 struct {

	// id
	// Required: true
	ID string `json:"_id"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at"`

	// game mode
	// Required: true
	GameMode string `json:"game_mode"`

	// joinable
	// Required: true
	Joinable bool `json:"joinable"`

	// joining
	// Required: true
	Joining []*ModelsSessionPlayerJoining `json:"joining"`

	// match id
	// Required: true
	MatchID string `json:"match_id"`

	// namespace
	// Required: true
	Namespace string `json:"namespace"`

	// players
	// Required: true
	Players []*ModelsSessionPlayerHistory `json:"players"`

	// removed reason
	// Required: true
	RemovedReason *string `json:"removed_reason"`

	// session type
	// Required: true
	SessionType *string `json:"session_type"`

	// status
	// Required: true
	Status string `json:"status"`

	// sub game mode
	// Required: true
	SubGameMode []string `json:"sub_game_mode"`

	// user id
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this models get session history search response item v2
func (m *ModelsGetSessionHistorySearchResponseItemV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGameMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJoinable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJoining(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemovedReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubGameMode(formats); err != nil {
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

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("_id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateGameMode(formats strfmt.Registry) error {

	if err := validate.RequiredString("game_mode", "body", string(m.GameMode)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateJoinable(formats strfmt.Registry) error {

	if err := validate.Required("joinable", "body", bool(m.Joinable)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateJoining(formats strfmt.Registry) error {

	if err := validate.Required("joining", "body", m.Joining); err != nil {
		return err
	}

	for i := 0; i < len(m.Joining); i++ {
		if swag.IsZero(m.Joining[i]) { // not required
			continue
		}

		if m.Joining[i] != nil {
			if err := m.Joining[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("joining" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateMatchID(formats strfmt.Registry) error {

	if err := validate.RequiredString("match_id", "body", string(m.MatchID)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateNamespace(formats strfmt.Registry) error {

	if err := validate.RequiredString("namespace", "body", string(m.Namespace)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validatePlayers(formats strfmt.Registry) error {

	if err := validate.Required("players", "body", m.Players); err != nil {
		return err
	}

	for i := 0; i < len(m.Players); i++ {
		if swag.IsZero(m.Players[i]) { // not required
			continue
		}

		if m.Players[i] != nil {
			if err := m.Players[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("players" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateRemovedReason(formats strfmt.Registry) error {

	if err := validate.Required("removed_reason", "body", m.RemovedReason); err != nil {
		return err
	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateSessionType(formats strfmt.Registry) error {

	if err := validate.Required("session_type", "body", m.SessionType); err != nil {
		return err
	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateSubGameMode(formats strfmt.Registry) error {

	if err := validate.Required("sub_game_mode", "body", m.SubGameMode); err != nil {
		return err
	}

	return nil
}

func (m *ModelsGetSessionHistorySearchResponseItemV2) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsGetSessionHistorySearchResponseItemV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsGetSessionHistorySearchResponseItemV2) UnmarshalBinary(b []byte) error {
	var res ModelsGetSessionHistorySearchResponseItemV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
