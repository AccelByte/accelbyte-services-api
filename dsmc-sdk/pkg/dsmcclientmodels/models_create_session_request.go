// Code generated by go-swagger; DO NOT EDIT.

package dsmcclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsCreateSessionRequest models create session request
//
// swagger:model models.CreateSessionRequest
type ModelsCreateSessionRequest struct {

	// match result
	// Required: true
	MatchResult *ModelsMatchResult `json:"MatchResult"`

	// client version
	// Required: true
	ClientVersion *string `json:"client_version"`

	// configuration
	// Required: true
	Configuration *string `json:"configuration"`

	// deployment
	// Required: true
	Deployment *string `json:"deployment"`

	// pod name
	// Required: true
	PodName *string `json:"pod_name"`

	// region
	// Required: true
	Region *string `json:"region"`
}

// Validate validates this models create session request
func (m *ModelsCreateSessionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsCreateSessionRequest) validateMatchResult(formats strfmt.Registry) error {

	if err := validate.Required("MatchResult", "body", m.MatchResult); err != nil {
		return err
	}

	if m.MatchResult != nil {
		if err := m.MatchResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MatchResult")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsCreateSessionRequest) validateClientVersion(formats strfmt.Registry) error {

	if err := validate.Required("client_version", "body", m.ClientVersion); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateSessionRequest) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateSessionRequest) validateDeployment(formats strfmt.Registry) error {

	if err := validate.Required("deployment", "body", m.Deployment); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateSessionRequest) validatePodName(formats strfmt.Registry) error {

	if err := validate.Required("pod_name", "body", m.PodName); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateSessionRequest) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsCreateSessionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsCreateSessionRequest) UnmarshalBinary(b []byte) error {
	var res ModelsCreateSessionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}