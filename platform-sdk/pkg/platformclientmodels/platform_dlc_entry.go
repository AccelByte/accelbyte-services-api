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

// PlatformDlcEntry platform dlc entry
//
// swagger:model PlatformDlcEntry
type PlatformDlcEntry struct {

	// platform: steam, psn, xbl
	// Enum: [PSN STEAM XBOX]
	Platform string `json:"platform,omitempty"`

	// key is platform product id, value is dlc id
	PlatformDlcIDMap map[string]string `json:"platformDlcIdMap,omitempty"`
}

// Validate validates this platform dlc entry
func (m *PlatformDlcEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var platformDlcEntryTypePlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PSN","STEAM","XBOX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		platformDlcEntryTypePlatformPropEnum = append(platformDlcEntryTypePlatformPropEnum, v)
	}
}

const (

	// PlatformDlcEntryPlatformPSN captures enum value "PSN"
	PlatformDlcEntryPlatformPSN string = "PSN"

	// PlatformDlcEntryPlatformSTEAM captures enum value "STEAM"
	PlatformDlcEntryPlatformSTEAM string = "STEAM"

	// PlatformDlcEntryPlatformXBOX captures enum value "XBOX"
	PlatformDlcEntryPlatformXBOX string = "XBOX"
)

// prop value enum
func (m *PlatformDlcEntry) validatePlatformEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, platformDlcEntryTypePlatformPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlatformDlcEntry) validatePlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformEnum("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlatformDlcEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformDlcEntry) UnmarshalBinary(b []byte) error {
	var res PlatformDlcEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
