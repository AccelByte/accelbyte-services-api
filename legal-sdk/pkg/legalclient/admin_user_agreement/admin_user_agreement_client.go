// Code generated by go-swagger; DO NOT EDIT.

package admin_user_agreement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new admin user agreement API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for admin user agreement API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	IndirectBulkAcceptVersionedPolicy(params *IndirectBulkAcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicyOK, error)
	IndirectBulkAcceptVersionedPolicyShort(params *IndirectBulkAcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  IndirectBulkAcceptVersionedPolicy admins bulk accept policy versions

  Accepts many legal policy versions all at once. Supply with localized version policy id and userId to accept an agreement. Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:USER:{userId}:LEGAL&#34;, action=1 (CREATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) IndirectBulkAcceptVersionedPolicy(params *IndirectBulkAcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndirectBulkAcceptVersionedPolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indirectBulkAcceptVersionedPolicy",
		Method:             "POST",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/agreements/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndirectBulkAcceptVersionedPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *IndirectBulkAcceptVersionedPolicyOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) IndirectBulkAcceptVersionedPolicyShort(params *IndirectBulkAcceptVersionedPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*IndirectBulkAcceptVersionedPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndirectBulkAcceptVersionedPolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "indirectBulkAcceptVersionedPolicy",
		Method:             "POST",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/agreements/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IndirectBulkAcceptVersionedPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *IndirectBulkAcceptVersionedPolicyOK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
