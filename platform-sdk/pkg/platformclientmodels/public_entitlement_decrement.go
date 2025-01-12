// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package platformclientmodels

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PublicEntitlementDecrement Public entitlement decrement
//
// swagger:model Public entitlement decrement.
type PublicEntitlementDecrement struct {

	// metadata, It is not stored in the database, only forwarded to AGS event
	Metadata *PublicEntitlementMetadata `json:"metadata,omitempty"`

	// options, it is only available when entitlement clazz is OPTIONBOX and value should be item id.
	Options []string `json:"options,omitempty"`

	// Request id(Optional), client should provide a unique request id to perform at most once execution, When a request id is resubmitted, it will return original successful response
	RequestID string `json:"requestId,omitempty"`

	// the count to decrease, 1 for default, option box entitlement only allow 1
	// Format: int32
	UseCount int32 `json:"useCount,omitempty"`
}

// Validate validates this Public entitlement decrement
func (m *PublicEntitlementDecrement) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PublicEntitlementDecrement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicEntitlementDecrement) UnmarshalBinary(b []byte) error {
	var res PublicEntitlementDecrement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
