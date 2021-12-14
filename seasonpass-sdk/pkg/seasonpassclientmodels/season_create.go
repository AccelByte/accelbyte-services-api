// Code generated by go-swagger; DO NOT EDIT.

package seasonpassclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SeasonCreate A DTO object for creating season API call.
//
// swagger:model SeasonCreate
type SeasonCreate struct {

	// whether auto claim rewards, default true
	AutoClaim bool `json:"autoClaim,omitempty"`

	// default language, BCP 47 language tag, default is en
	DefaultLanguage string `json:"defaultLanguage,omitempty"`

	// default exp required for a tier
	// Required: true
	DefaultRequiredExp *int32 `json:"defaultRequiredExp"`

	// draft store id
	// Required: true
	DraftStoreID *string `json:"draftStoreId"`

	// end date time
	// Required: true
	// Format: date-time
	End *strfmt.DateTime `json:"end"`

	// strategy while exceed final tier exp, default NONE
	ExcessStrategy *ExcessStrategy `json:"excessStrategy,omitempty"`

	// images
	Images []*Image `json:"images"`

	// localization, {language: localization} map
	// Required: true
	Localizations map[string]Localization `json:"localizations"`

	// name, max length is 127
	// Required: true
	Name *string `json:"name"`

	// start date time
	// Required: true
	// Format: date-time
	Start *strfmt.DateTime `json:"start"`

	// tier item id
	// Required: true
	TierItemID *string `json:"tierItemId"`
}

// Validate validates this season create
func (m *SeasonCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultRequiredExp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDraftStoreID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcessStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTierItemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SeasonCreate) validateDefaultRequiredExp(formats strfmt.Registry) error {

	if err := validate.Required("defaultRequiredExp", "body", m.DefaultRequiredExp); err != nil {
		return err
	}

	return nil
}

func (m *SeasonCreate) validateDraftStoreID(formats strfmt.Registry) error {

	if err := validate.Required("draftStoreId", "body", m.DraftStoreID); err != nil {
		return err
	}

	return nil
}

func (m *SeasonCreate) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SeasonCreate) validateExcessStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.ExcessStrategy) { // not required
		return nil
	}

	if m.ExcessStrategy != nil {
		if err := m.ExcessStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("excessStrategy")
			}
			return err
		}
	}

	return nil
}

func (m *SeasonCreate) validateImages(formats strfmt.Registry) error {

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

func (m *SeasonCreate) validateLocalizations(formats strfmt.Registry) error {

	for k := range m.Localizations {

		if err := validate.Required("localizations"+"."+k, "body", m.Localizations[k]); err != nil {
			return err
		}
		if val, ok := m.Localizations[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SeasonCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SeasonCreate) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SeasonCreate) validateTierItemID(formats strfmt.Registry) error {

	if err := validate.Required("tierItemId", "body", m.TierItemID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SeasonCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeasonCreate) UnmarshalBinary(b []byte) error {
	var res SeasonCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
