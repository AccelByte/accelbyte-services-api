// Code generated by go-swagger; DO NOT EDIT.

package platformclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentProviderConfigEdit A DTO object for creating/updating payment provider config
//
// swagger:model PaymentProviderConfigEdit
type PaymentProviderConfigEdit struct {

	// aggregate payment provider, allow empty value
	// Enum: [ADYEN XSOLLA]
	Aggregate string `json:"aggregate,omitempty"`

	// namespace, * indicates all namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// region, * indicates all regions
	// Required: true
	Region *string `json:"region"`

	// sandbox taxjar api token
	SandboxTaxJarAPIToken string `json:"sandboxTaxJarApiToken,omitempty"`

	// special payment providers = ['ALIPAY', 'WXPAY'], allow empty value
	Specials []string `json:"specials"`

	// taxjar api token, required when taxJarEnabled=true and useGlobalTaxJarApiToken=false
	TaxJarAPIToken string `json:"taxJarApiToken,omitempty"`

	// taxjar api integration enable
	TaxJarEnabled bool `json:"taxJarEnabled"`

	// only works when taxJarEnabled=true, and if useGlobalTaxJarApiToken is true, we will reset the taxJarApiToken as null
	UseGlobalTaxJarAPIToken bool `json:"useGlobalTaxJarApiToken"`
}

// Validate validates this payment provider config edit
func (m *PaymentProviderConfigEdit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var paymentProviderConfigEditTypeAggregatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ADYEN","XSOLLA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentProviderConfigEditTypeAggregatePropEnum = append(paymentProviderConfigEditTypeAggregatePropEnum, v)
	}
}

const (

	// PaymentProviderConfigEditAggregateADYEN captures enum value "ADYEN"
	PaymentProviderConfigEditAggregateADYEN string = "ADYEN"

	// PaymentProviderConfigEditAggregateXSOLLA captures enum value "XSOLLA"
	PaymentProviderConfigEditAggregateXSOLLA string = "XSOLLA"
)

// prop value enum
func (m *PaymentProviderConfigEdit) validateAggregateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paymentProviderConfigEditTypeAggregatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PaymentProviderConfigEdit) validateAggregate(formats strfmt.Registry) error {

	if swag.IsZero(m.Aggregate) { // not required
		return nil
	}

	// value enum
	if err := m.validateAggregateEnum("aggregate", "body", m.Aggregate); err != nil {
		return err
	}

	return nil
}

func (m *PaymentProviderConfigEdit) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *PaymentProviderConfigEdit) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

var paymentProviderConfigEditSpecialsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ADYEN","ALIPAY","CHECKOUT","PAYPAL","STRIPE","WALLET","WXPAY","XSOLLA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentProviderConfigEditSpecialsItemsEnum = append(paymentProviderConfigEditSpecialsItemsEnum, v)
	}
}

func (m *PaymentProviderConfigEdit) validateSpecialsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paymentProviderConfigEditSpecialsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PaymentProviderConfigEdit) validateSpecials(formats strfmt.Registry) error {

	if swag.IsZero(m.Specials) { // not required
		return nil
	}

	for i := 0; i < len(m.Specials); i++ {

		// value enum
		if err := m.validateSpecialsItemsEnum("specials"+"."+strconv.Itoa(i), "body", m.Specials[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentProviderConfigEdit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentProviderConfigEdit) UnmarshalBinary(b []byte) error {
	var res PaymentProviderConfigEdit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
