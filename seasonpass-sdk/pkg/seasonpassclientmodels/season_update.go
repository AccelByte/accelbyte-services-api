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

// SeasonUpdate A DTO object for updating season API call.
//
// swagger:model SeasonUpdate
type SeasonUpdate struct {

	// whether auto claim rewards
	AutoClaim bool `json:"autoClaim,omitempty"`

	// default language, BCP 47 language tag
	DefaultLanguage string `json:"defaultLanguage,omitempty"`

	// default exp required for a tier
	DefaultRequiredExp int32 `json:"defaultRequiredExp,omitempty"`

	// draft store id
	DraftStoreID string `json:"draftStoreId,omitempty"`

	// end date time
	// Format: date-time
	End *strfmt.DateTime `json:"end,omitempty"`

	// strategy while exceed final tier exp
	ExcessStrategy *ExcessStrategy `json:"excessStrategy,omitempty"`

	// images
	Images []*Image `json:"images"`

	// localization, {language: localization} map
	Localizations map[string]Localization `json:"localizations,omitempty"`

	// name, max length is 127
	Name string `json:"name,omitempty"`

	// start date time
	// Format: date-time
	Start *strfmt.DateTime `json:"start,omitempty"`

	// tier item id
	TierItemID string `json:"tierItemId,omitempty"`
}

// Validate validates this season update
func (m *SeasonUpdate) Validate(formats strfmt.Registry) error {
	var res []error

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

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SeasonUpdate) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SeasonUpdate) validateExcessStrategy(formats strfmt.Registry) error {

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

func (m *SeasonUpdate) validateImages(formats strfmt.Registry) error {

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

func (m *SeasonUpdate) validateLocalizations(formats strfmt.Registry) error {

	if swag.IsZero(m.Localizations) { // not required
		return nil
	}

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

func (m *SeasonUpdate) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SeasonUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeasonUpdate) UnmarshalBinary(b []byte) error {
	var res SeasonUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
