// Code generated by go-swagger; DO NOT EDIT.

package lobbyclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelsConfigReq models config req
//
// swagger:model models.ConfigReq
type ModelsConfigReq struct {

	// auto kick on disconnect
	AutoKickOnDisconnect bool `json:"autoKickOnDisconnect,omitempty"`

	// chat rate limit burst
	ChatRateLimitBurst int32 `json:"chatRateLimitBurst,omitempty"`

	// chat rate limit duration
	ChatRateLimitDuration int64 `json:"chatRateLimitDuration,omitempty"`

	// concurrent users limit
	ConcurrentUsersLimit int32 `json:"concurrentUsersLimit,omitempty"`

	// enable chat
	EnableChat bool `json:"enableChat,omitempty"`

	// entitlement check
	EntitlementCheck bool `json:"entitlementCheck,omitempty"`

	// entitlement item ID
	EntitlementItemID string `json:"entitlementItemID,omitempty"`

	// general rate limit burst
	GeneralRateLimitBurst int32 `json:"generalRateLimitBurst,omitempty"`

	// general rate limit duration
	GeneralRateLimitDuration int64 `json:"generalRateLimitDuration,omitempty"`

	// max party member
	MaxPartyMember int32 `json:"maxPartyMember,omitempty"`

	// profanity filter
	ProfanityFilter bool `json:"profanityFilter,omitempty"`
}

// Validate validates this models config req
func (m *ModelsConfigReq) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsConfigReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsConfigReq) UnmarshalBinary(b []byte) error {
	var res ModelsConfigReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
