// Code generated by go-swagger; DO NOT EDIT.

package global_statistic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new global statistic API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for global statistic API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetGlobalStatItems(params *GetGlobalStatItemsParams, authInfo runtime.ClientAuthInfoWriter) (*GetGlobalStatItemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetGlobalStatItems lists global stat items

  List global statItems by pagination.<br>Other detail info:<ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:STATITEM", action=2 (READ)</li><li><i>Returns</i>: stat items</li>ul
*/
func (a *Client) GetGlobalStatItems(params *GetGlobalStatItemsParams, authInfo runtime.ClientAuthInfoWriter) (*GetGlobalStatItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalStatItemsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGlobalStatItems",
		Method:             "GET",
		PathPattern:        "/v1/admin/namespaces/{namespace}/globalstatitems",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGlobalStatItemsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetGlobalStatItemsOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
