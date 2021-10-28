// Code generated by go-swagger; DO NOT EDIT.

package payment_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new payment account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payment account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	PublicDeletePaymentAccount(params *PublicDeletePaymentAccountParams, authInfo runtime.ClientAuthInfoWriter) (*PublicDeletePaymentAccountNoContent, error)

	PublicGetPaymentAccounts(params *PublicGetPaymentAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetPaymentAccountsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PublicDeletePaymentAccount deletes payment account

  Delete payment account.<br>Other detail info: <ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:PAYMENT:ACCOUNT", action=8 (DELETE)</li><li><i>Returns</i>:</li></ul>
*/
func (a *Client) PublicDeletePaymentAccount(params *PublicDeletePaymentAccountParams, authInfo runtime.ClientAuthInfoWriter) (*PublicDeletePaymentAccountNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicDeletePaymentAccountParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicDeletePaymentAccount",
		Method:             "DELETE",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/payment/accounts/{type}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicDeletePaymentAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicDeletePaymentAccountNoContent:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetPaymentAccounts gets payment accounts

  Get payment accounts.<br>Other detail info: <ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:PAYMENT:ACCOUNT", action=2 (READ)</li><li><i>Returns</i>: Payment account list</li></ul>
*/
func (a *Client) PublicGetPaymentAccounts(params *PublicGetPaymentAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetPaymentAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetPaymentAccountsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetPaymentAccounts",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/payment/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicGetPaymentAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicGetPaymentAccountsOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
