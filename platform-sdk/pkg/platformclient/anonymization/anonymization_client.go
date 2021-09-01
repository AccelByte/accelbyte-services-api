// Code generated by go-swagger; DO NOT EDIT.

package anonymization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new anonymization API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for anonymization API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AnonymizeEntitlement(params *AnonymizeEntitlementParams, authInfo runtime.ClientAuthInfoWriter) (*AnonymizeEntitlementNoContent, error)

	AnonymizeFulfillment(params *AnonymizeFulfillmentParams, authInfo runtime.ClientAuthInfoWriter) (*AnonymizeFulfillmentNoContent, error)

	AnonymizeWallet(params *AnonymizeWalletParams, authInfo runtime.ClientAuthInfoWriter) (*AnonymizeWalletNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AnonymizeEntitlement anonymizes entitlement

  Anonymize entitlement. At current it will only anonymize entitlement, entitlement history and distribution receiver.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:ANONYMIZATION", action=8 (DELETE)</li></ul>
*/
func (a *Client) AnonymizeEntitlement(params *AnonymizeEntitlementParams, authInfo runtime.ClientAuthInfoWriter) (*AnonymizeEntitlementNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAnonymizeEntitlementParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "anonymizeEntitlement",
		Method:             "DELETE",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/anonymization/entitlements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AnonymizeEntitlementReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AnonymizeEntitlementNoContent:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AnonymizeFulfillment anonymizes fulfillment

  Anonymize fulfillment. At current it will only anonymize fulfillment history.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:ANONYMIZATION", action=8 (DELETE)</li></ul>
*/
func (a *Client) AnonymizeFulfillment(params *AnonymizeFulfillmentParams, authInfo runtime.ClientAuthInfoWriter) (*AnonymizeFulfillmentNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAnonymizeFulfillmentParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "anonymizeFulfillment",
		Method:             "DELETE",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/anonymization/fulfillment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AnonymizeFulfillmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AnonymizeFulfillmentNoContent:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AnonymizeWallet anonymizes wallet

  Anonymize wallet. At current it will only anonymize wallet, wallet transaction.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:ANONYMIZATION", action=8 (DELETE)</li></ul>
*/
func (a *Client) AnonymizeWallet(params *AnonymizeWalletParams, authInfo runtime.ClientAuthInfoWriter) (*AnonymizeWalletNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAnonymizeWalletParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "anonymizeWallet",
		Method:             "DELETE",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/anonymization/wallets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AnonymizeWalletReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AnonymizeWalletNoContent:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}