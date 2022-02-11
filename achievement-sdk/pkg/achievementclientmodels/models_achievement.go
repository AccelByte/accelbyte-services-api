// Code generated by go-swagger; DO NOT EDIT.

package achievementclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsAchievement models achievement
//
// swagger:model models.Achievement
type ModelsAchievement struct {

	// achievement code
	// Required: true
	AchievementCode *string `json:"AchievementCode"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"CreatedAt"`

	// default language
	// Required: true
	DefaultLanguage *string `json:"DefaultLanguage"`

	// description
	// Required: true
	Description map[string]string `json:"Description"`

	// goal value
	// Required: true
	GoalValue *float64 `json:"GoalValue"`

	// hidden
	// Required: true
	Hidden *bool `json:"Hidden"`

	// ID
	// Required: true
	ID *string `json:"ID"`

	// incremental
	// Required: true
	Incremental *bool `json:"Incremental"`

	// list order
	// Required: true
	ListOrder *int32 `json:"ListOrder"`

	// locked icons
	// Required: true
	LockedIcons []*ModelsIcon `json:"LockedIcons"`

	// name
	// Required: true
	Name map[string]string `json:"Name"`

	// namespace
	// Required: true
	Namespace *string `json:"Namespace"`

	// stat code
	// Required: true
	StatCode *string `json:"StatCode"`

	// tags
	// Required: true
	Tags []string `json:"Tags"`

	// unlocked icons
	// Required: true
	UnlockedIcons []*ModelsIcon `json:"UnlockedIcons"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"UpdatedAt"`
}

// Validate validates this models achievement
func (m *ModelsAchievement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAchievementCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGoalValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHidden(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncremental(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockedIcons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnlockedIcons(formats); err != nil {
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

func (m *ModelsAchievement) validateAchievementCode(formats strfmt.Registry) error {

	if err := validate.Required("AchievementCode", "body", m.AchievementCode); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("CreatedAt", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("CreatedAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateDefaultLanguage(formats strfmt.Registry) error {

	if err := validate.Required("DefaultLanguage", "body", m.DefaultLanguage); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateDescription(formats strfmt.Registry) error {

	return nil
}

func (m *ModelsAchievement) validateGoalValue(formats strfmt.Registry) error {

	if err := validate.Required("GoalValue", "body", m.GoalValue); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateHidden(formats strfmt.Registry) error {

	if err := validate.Required("Hidden", "body", m.Hidden); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateIncremental(formats strfmt.Registry) error {

	if err := validate.Required("Incremental", "body", m.Incremental); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateListOrder(formats strfmt.Registry) error {

	if err := validate.Required("ListOrder", "body", m.ListOrder); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateLockedIcons(formats strfmt.Registry) error {

	if err := validate.Required("LockedIcons", "body", m.LockedIcons); err != nil {
		return err
	}

	for i := 0; i < len(m.LockedIcons); i++ {
		if swag.IsZero(m.LockedIcons[i]) { // not required
			continue
		}

		if m.LockedIcons[i] != nil {
			if err := m.LockedIcons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LockedIcons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsAchievement) validateName(formats strfmt.Registry) error {

	return nil
}

func (m *ModelsAchievement) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("Namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateStatCode(formats strfmt.Registry) error {

	if err := validate.Required("StatCode", "body", m.StatCode); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("Tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *ModelsAchievement) validateUnlockedIcons(formats strfmt.Registry) error {

	if err := validate.Required("UnlockedIcons", "body", m.UnlockedIcons); err != nil {
		return err
	}

	for i := 0; i < len(m.UnlockedIcons); i++ {
		if swag.IsZero(m.UnlockedIcons[i]) { // not required
			continue
		}

		if m.UnlockedIcons[i] != nil {
			if err := m.UnlockedIcons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UnlockedIcons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsAchievement) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("UpdatedAt", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsAchievement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsAchievement) UnmarshalBinary(b []byte) error {
	var res ModelsAchievement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
