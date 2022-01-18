// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SteamSyncRequest steam sync request
//
// swagger:model SteamSyncRequest
type SteamSyncRequest struct {

	// Steam app id
	// Required: true
	AppID *string `json:"appId"`

	// product price currency
	CurrencyCode string `json:"currencyCode,omitempty"`

	// language value from language tag, allowed format: en, en-US
	Language string `json:"language,omitempty"`

	// product price
	Price float64 `json:"price,omitempty"`

	// steamdefid
	ProductID string `json:"productId,omitempty"`

	// country value from ISO countries
	Region string `json:"region,omitempty"`

	// Steam ID of the user
	// Required: true
	SteamID *string `json:"steamId"`
}

// Validate validates this steam sync request
func (m *SteamSyncRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteamID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SteamSyncRequest) validateAppID(formats strfmt.Registry) error {

	if err := validate.Required("appId", "body", m.AppID); err != nil {
		return err
	}

	return nil
}

func (m *SteamSyncRequest) validateSteamID(formats strfmt.Registry) error {

	if err := validate.Required("steamId", "body", m.SteamID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SteamSyncRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SteamSyncRequest) UnmarshalBinary(b []byte) error {
	var res SteamSyncRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
