// Code generated by go-swagger; DO NOT EDIT.

package iamclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelUserResponseV3 model user response v3
//
// swagger:model model.UserResponseV3
type ModelUserResponseV3 struct {

	// auth type
	// Required: true
	AuthType *string `json:"authType"`

	// bans
	// Required: true
	Bans []*ModelUserActiveBanResponseV3 `json:"bans"`

	// country
	// Required: true
	Country *string `json:"country"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// date of birth
	// Required: true
	// Format: date-time
	DateOfBirth *strfmt.DateTime `json:"dateOfBirth"`

	// deletion status
	// Required: true
	DeletionStatus *bool `json:"deletionStatus"`

	// display name
	// Required: true
	DisplayName *string `json:"displayName"`

	// email address
	// Required: true
	EmailAddress *string `json:"emailAddress"`

	// email verified
	// Required: true
	EmailVerified *bool `json:"emailVerified"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// last date of birth changed time
	// Required: true
	// Format: date-time
	LastDateOfBirthChangedTime *strfmt.DateTime `json:"lastDateOfBirthChangedTime"`

	// last enabled changed time
	// Required: true
	// Format: date-time
	LastEnabledChangedTime *strfmt.DateTime `json:"lastEnabledChangedTime"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// namespace roles
	// Required: true
	NamespaceRoles []*AccountcommonNamespaceRole `json:"namespaceRoles"`

	// new email address
	NewEmailAddress string `json:"newEmailAddress,omitempty"`

	// old email address
	// Required: true
	OldEmailAddress *string `json:"oldEmailAddress"`

	// permissions
	// Required: true
	Permissions []*ModelUserPermissionsResponseV3 `json:"permissions"`

	// phone number
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// phone verified
	// Required: true
	PhoneVerified *bool `json:"phoneVerified"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`

	// platform user Id
	PlatformUserID string `json:"platformUserId,omitempty"`

	// roles
	// Required: true
	Roles []string `json:"roles"`

	// user Id
	// Required: true
	UserID *string `json:"userId"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this model user response v3
func (m *ModelUserResponseV3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateOfBirth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailVerified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastDateOfBirthChangedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastEnabledChangedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOldEmailAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneVerified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
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

func (m *ModelUserResponseV3) validateAuthType(formats strfmt.Registry) error {

	if err := validate.Required("authType", "body", m.AuthType); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateBans(formats strfmt.Registry) error {

	if err := validate.Required("bans", "body", m.Bans); err != nil {
		return err
	}

	for i := 0; i < len(m.Bans); i++ {
		if swag.IsZero(m.Bans[i]) { // not required
			continue
		}

		if m.Bans[i] != nil {
			if err := m.Bans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelUserResponseV3) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateDateOfBirth(formats strfmt.Registry) error {

	if err := validate.Required("dateOfBirth", "body", m.DateOfBirth); err != nil {
		return err
	}

	if err := validate.FormatOf("dateOfBirth", "body", "date-time", m.DateOfBirth.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateDeletionStatus(formats strfmt.Registry) error {

	if err := validate.Required("deletionStatus", "body", m.DeletionStatus); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("emailAddress", "body", m.EmailAddress); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateEmailVerified(formats strfmt.Registry) error {

	if err := validate.Required("emailVerified", "body", m.EmailVerified); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateLastDateOfBirthChangedTime(formats strfmt.Registry) error {

	if err := validate.Required("lastDateOfBirthChangedTime", "body", m.LastDateOfBirthChangedTime); err != nil {
		return err
	}

	if err := validate.FormatOf("lastDateOfBirthChangedTime", "body", "date-time", m.LastDateOfBirthChangedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateLastEnabledChangedTime(formats strfmt.Registry) error {

	if err := validate.Required("lastEnabledChangedTime", "body", m.LastEnabledChangedTime); err != nil {
		return err
	}

	if err := validate.FormatOf("lastEnabledChangedTime", "body", "date-time", m.LastEnabledChangedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateNamespaceRoles(formats strfmt.Registry) error {

	if err := validate.Required("namespaceRoles", "body", m.NamespaceRoles); err != nil {
		return err
	}

	for i := 0; i < len(m.NamespaceRoles); i++ {
		if swag.IsZero(m.NamespaceRoles[i]) { // not required
			continue
		}

		if m.NamespaceRoles[i] != nil {
			if err := m.NamespaceRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namespaceRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelUserResponseV3) validateOldEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("oldEmailAddress", "body", m.OldEmailAddress); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelUserResponseV3) validatePhoneVerified(formats strfmt.Registry) error {

	if err := validate.Required("phoneVerified", "body", m.PhoneVerified); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

func (m *ModelUserResponseV3) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelUserResponseV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelUserResponseV3) UnmarshalBinary(b []byte) error {
	var res ModelUserResponseV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
