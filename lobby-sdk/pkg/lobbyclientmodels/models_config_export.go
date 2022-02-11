// Code generated by go-swagger; DO NOT EDIT.

package lobbyclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsConfigExport models config export
//
// swagger:model models.ConfigExport
type ModelsConfigExport struct {

	// auto kick on disconnect
	// Required: true
	AutoKickOnDisconnect *bool `json:"autoKickOnDisconnect"`

	// auto kick on disconnect delay
	// Required: true
	AutoKickOnDisconnectDelay *int64 `json:"autoKickOnDisconnectDelay"`

	// cancel ticket on disconnect
	// Required: true
	CancelTicketOnDisconnect *bool `json:"cancelTicketOnDisconnect"`

	// chat rate limit burst
	// Required: true
	ChatRateLimitBurst *int32 `json:"chatRateLimitBurst"`

	// chat rate limit duration
	// Required: true
	ChatRateLimitDuration *int64 `json:"chatRateLimitDuration"`

	// concurrent users limit
	// Required: true
	ConcurrentUsersLimit *int32 `json:"concurrentUsersLimit"`

	// disable party invitation token
	// Required: true
	DisablePartyInvitationToken *bool `json:"disablePartyInvitationToken"`

	// enable chat
	// Required: true
	EnableChat *bool `json:"enableChat"`

	// entitlement check
	// Required: true
	EntitlementCheck *bool `json:"entitlementCheck"`

	// entitlement item ID
	// Required: true
	EntitlementItemID *string `json:"entitlementItemID"`

	// general rate limit burst
	// Required: true
	GeneralRateLimitBurst *int32 `json:"generalRateLimitBurst"`

	// general rate limit duration
	// Required: true
	GeneralRateLimitDuration *int64 `json:"generalRateLimitDuration"`

	// max party member
	// Required: true
	MaxPartyMember *int32 `json:"maxPartyMember"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// profanity filter
	// Required: true
	ProfanityFilter *bool `json:"profanityFilter"`

	// ready consent timeout
	// Required: true
	ReadyConsentTimeout *int64 `json:"readyConsentTimeout"`
}

// Validate validates this models config export
func (m *ModelsConfigExport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoKickOnDisconnect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoKickOnDisconnectDelay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCancelTicketOnDisconnect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChatRateLimitBurst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChatRateLimitDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConcurrentUsersLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisablePartyInvitationToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnableChat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlementCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntitlementItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneralRateLimitBurst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneralRateLimitDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxPartyMember(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfanityFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadyConsentTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsConfigExport) validateAutoKickOnDisconnect(formats strfmt.Registry) error {

	if err := validate.Required("autoKickOnDisconnect", "body", m.AutoKickOnDisconnect); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateAutoKickOnDisconnectDelay(formats strfmt.Registry) error {

	if err := validate.Required("autoKickOnDisconnectDelay", "body", m.AutoKickOnDisconnectDelay); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateCancelTicketOnDisconnect(formats strfmt.Registry) error {

	if err := validate.Required("cancelTicketOnDisconnect", "body", m.CancelTicketOnDisconnect); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateChatRateLimitBurst(formats strfmt.Registry) error {

	if err := validate.Required("chatRateLimitBurst", "body", m.ChatRateLimitBurst); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateChatRateLimitDuration(formats strfmt.Registry) error {

	if err := validate.Required("chatRateLimitDuration", "body", m.ChatRateLimitDuration); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateConcurrentUsersLimit(formats strfmt.Registry) error {

	if err := validate.Required("concurrentUsersLimit", "body", m.ConcurrentUsersLimit); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateDisablePartyInvitationToken(formats strfmt.Registry) error {

	if err := validate.Required("disablePartyInvitationToken", "body", m.DisablePartyInvitationToken); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateEnableChat(formats strfmt.Registry) error {

	if err := validate.Required("enableChat", "body", m.EnableChat); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateEntitlementCheck(formats strfmt.Registry) error {

	if err := validate.Required("entitlementCheck", "body", m.EntitlementCheck); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateEntitlementItemID(formats strfmt.Registry) error {

	if err := validate.Required("entitlementItemID", "body", m.EntitlementItemID); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateGeneralRateLimitBurst(formats strfmt.Registry) error {

	if err := validate.Required("generalRateLimitBurst", "body", m.GeneralRateLimitBurst); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateGeneralRateLimitDuration(formats strfmt.Registry) error {

	if err := validate.Required("generalRateLimitDuration", "body", m.GeneralRateLimitDuration); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateMaxPartyMember(formats strfmt.Registry) error {

	if err := validate.Required("maxPartyMember", "body", m.MaxPartyMember); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateProfanityFilter(formats strfmt.Registry) error {

	if err := validate.Required("profanityFilter", "body", m.ProfanityFilter); err != nil {
		return err
	}

	return nil
}

func (m *ModelsConfigExport) validateReadyConsentTimeout(formats strfmt.Registry) error {

	if err := validate.Required("readyConsentTimeout", "body", m.ReadyConsentTimeout); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsConfigExport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsConfigExport) UnmarshalBinary(b []byte) error {
	var res ModelsConfigExport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
