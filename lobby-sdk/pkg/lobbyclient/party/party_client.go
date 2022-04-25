// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package party

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new party API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for party API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AdminGetPartyDataV1(params *AdminGetPartyDataV1Params, authInfo runtime.ClientAuthInfoWriter) (*AdminGetPartyDataV1OK, *AdminGetPartyDataV1BadRequest, *AdminGetPartyDataV1Unauthorized, *AdminGetPartyDataV1Forbidden, *AdminGetPartyDataV1NotFound, *AdminGetPartyDataV1InternalServerError, error)
	AdminGetPartyDataV1Short(params *AdminGetPartyDataV1Params, authInfo runtime.ClientAuthInfoWriter) (*AdminGetPartyDataV1OK, error)
	AdminGetUserPartyV1(params *AdminGetUserPartyV1Params, authInfo runtime.ClientAuthInfoWriter) (*AdminGetUserPartyV1OK, *AdminGetUserPartyV1BadRequest, *AdminGetUserPartyV1Unauthorized, *AdminGetUserPartyV1Forbidden, *AdminGetUserPartyV1NotFound, *AdminGetUserPartyV1InternalServerError, error)
	AdminGetUserPartyV1Short(params *AdminGetUserPartyV1Params, authInfo runtime.ClientAuthInfoWriter) (*AdminGetUserPartyV1OK, error)
	PublicGetPartyDataV1(params *PublicGetPartyDataV1Params, authInfo runtime.ClientAuthInfoWriter) (*PublicGetPartyDataV1OK, *PublicGetPartyDataV1BadRequest, *PublicGetPartyDataV1Unauthorized, *PublicGetPartyDataV1Forbidden, *PublicGetPartyDataV1NotFound, *PublicGetPartyDataV1InternalServerError, error)
	PublicGetPartyDataV1Short(params *PublicGetPartyDataV1Params, authInfo runtime.ClientAuthInfoWriter) (*PublicGetPartyDataV1OK, error)
	PublicUpdatePartyAttributesV1(params *PublicUpdatePartyAttributesV1Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdatePartyAttributesV1OK, *PublicUpdatePartyAttributesV1BadRequest, *PublicUpdatePartyAttributesV1Unauthorized, *PublicUpdatePartyAttributesV1Forbidden, *PublicUpdatePartyAttributesV1NotFound, *PublicUpdatePartyAttributesV1PreconditionFailed, *PublicUpdatePartyAttributesV1InternalServerError, error)
	PublicUpdatePartyAttributesV1Short(params *PublicUpdatePartyAttributesV1Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdatePartyAttributesV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AdminGetPartyDataV1 admins get party data

  Required permission : &lt;code&gt;ADMIN:NAMESPACE:{namespace}:PARTY:STORAGE [READ]&lt;/code&gt; with scope &lt;code&gt;social&lt;/code&gt;
			&lt;br&gt;get party data in a namespace.
*/
func (a *Client) AdminGetPartyDataV1(params *AdminGetPartyDataV1Params, authInfo runtime.ClientAuthInfoWriter) (*AdminGetPartyDataV1OK, *AdminGetPartyDataV1BadRequest, *AdminGetPartyDataV1Unauthorized, *AdminGetPartyDataV1Forbidden, *AdminGetPartyDataV1NotFound, *AdminGetPartyDataV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminGetPartyDataV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adminGetPartyDataV1",
		Method:             "GET",
		PathPattern:        "/lobby/v1/admin/party/namespaces/{namespace}/parties/{partyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AdminGetPartyDataV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminGetPartyDataV1OK:
		return v, nil, nil, nil, nil, nil, nil

	case *AdminGetPartyDataV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *AdminGetPartyDataV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *AdminGetPartyDataV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *AdminGetPartyDataV1NotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *AdminGetPartyDataV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) AdminGetPartyDataV1Short(params *AdminGetPartyDataV1Params, authInfo runtime.ClientAuthInfoWriter) (*AdminGetPartyDataV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminGetPartyDataV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adminGetPartyDataV1",
		Method:             "GET",
		PathPattern:        "/lobby/v1/admin/party/namespaces/{namespace}/parties/{partyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AdminGetPartyDataV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AdminGetPartyDataV1OK:
		return v, nil
	case *AdminGetPartyDataV1BadRequest:
		return nil, v
	case *AdminGetPartyDataV1Unauthorized:
		return nil, v
	case *AdminGetPartyDataV1Forbidden:
		return nil, v
	case *AdminGetPartyDataV1NotFound:
		return nil, v
	case *AdminGetPartyDataV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AdminGetUserPartyV1 admins get user party data

  Required permission : &lt;code&gt;ADMIN:NAMESPACE:{namespace}:PARTY:STORAGE [READ]&lt;/code&gt; with scope &lt;code&gt;social&lt;/code&gt;
			&lt;br&gt;get party data in a namespace.
*/
func (a *Client) AdminGetUserPartyV1(params *AdminGetUserPartyV1Params, authInfo runtime.ClientAuthInfoWriter) (*AdminGetUserPartyV1OK, *AdminGetUserPartyV1BadRequest, *AdminGetUserPartyV1Unauthorized, *AdminGetUserPartyV1Forbidden, *AdminGetUserPartyV1NotFound, *AdminGetUserPartyV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminGetUserPartyV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adminGetUserPartyV1",
		Method:             "GET",
		PathPattern:        "/lobby/v1/admin/party/namespaces/{namespace}/users/{userId}/party",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AdminGetUserPartyV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminGetUserPartyV1OK:
		return v, nil, nil, nil, nil, nil, nil

	case *AdminGetUserPartyV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *AdminGetUserPartyV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *AdminGetUserPartyV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *AdminGetUserPartyV1NotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *AdminGetUserPartyV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) AdminGetUserPartyV1Short(params *AdminGetUserPartyV1Params, authInfo runtime.ClientAuthInfoWriter) (*AdminGetUserPartyV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminGetUserPartyV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adminGetUserPartyV1",
		Method:             "GET",
		PathPattern:        "/lobby/v1/admin/party/namespaces/{namespace}/users/{userId}/party",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AdminGetUserPartyV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AdminGetUserPartyV1OK:
		return v, nil
	case *AdminGetUserPartyV1BadRequest:
		return nil, v
	case *AdminGetUserPartyV1Unauthorized:
		return nil, v
	case *AdminGetUserPartyV1Forbidden:
		return nil, v
	case *AdminGetUserPartyV1NotFound:
		return nil, v
	case *AdminGetUserPartyV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetPartyDataV1 gets party data by party id

  Required valid user authorization &lt;br/&gt;
			&lt;br&gt;load personal party data in a namespace based on Party ID &lt;br/&gt;
			Action Code: 50101
*/
func (a *Client) PublicGetPartyDataV1(params *PublicGetPartyDataV1Params, authInfo runtime.ClientAuthInfoWriter) (*PublicGetPartyDataV1OK, *PublicGetPartyDataV1BadRequest, *PublicGetPartyDataV1Unauthorized, *PublicGetPartyDataV1Forbidden, *PublicGetPartyDataV1NotFound, *PublicGetPartyDataV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetPartyDataV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetPartyDataV1",
		Method:             "GET",
		PathPattern:        "/lobby/v1/public/party/namespaces/{namespace}/parties/{partyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetPartyDataV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicGetPartyDataV1OK:
		return v, nil, nil, nil, nil, nil, nil

	case *PublicGetPartyDataV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *PublicGetPartyDataV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *PublicGetPartyDataV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *PublicGetPartyDataV1NotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *PublicGetPartyDataV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) PublicGetPartyDataV1Short(params *PublicGetPartyDataV1Params, authInfo runtime.ClientAuthInfoWriter) (*PublicGetPartyDataV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetPartyDataV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetPartyDataV1",
		Method:             "GET",
		PathPattern:        "/lobby/v1/public/party/namespaces/{namespace}/parties/{partyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetPartyDataV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicGetPartyDataV1OK:
		return v, nil
	case *PublicGetPartyDataV1BadRequest:
		return nil, v
	case *PublicGetPartyDataV1Unauthorized:
		return nil, v
	case *PublicGetPartyDataV1Forbidden:
		return nil, v
	case *PublicGetPartyDataV1NotFound:
		return nil, v
	case *PublicGetPartyDataV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicUpdatePartyAttributesV1 updates party attributes

  Required valid user authorization &lt;br/&gt;
			&lt;br&gt;update party attributes in a namespace.
*/
func (a *Client) PublicUpdatePartyAttributesV1(params *PublicUpdatePartyAttributesV1Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdatePartyAttributesV1OK, *PublicUpdatePartyAttributesV1BadRequest, *PublicUpdatePartyAttributesV1Unauthorized, *PublicUpdatePartyAttributesV1Forbidden, *PublicUpdatePartyAttributesV1NotFound, *PublicUpdatePartyAttributesV1PreconditionFailed, *PublicUpdatePartyAttributesV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicUpdatePartyAttributesV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicUpdatePartyAttributesV1",
		Method:             "PUT",
		PathPattern:        "/lobby/v1/public/party/namespaces/{namespace}/parties/{partyId}/attributes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicUpdatePartyAttributesV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicUpdatePartyAttributesV1OK:
		return v, nil, nil, nil, nil, nil, nil, nil

	case *PublicUpdatePartyAttributesV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil, nil

	case *PublicUpdatePartyAttributesV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil, nil

	case *PublicUpdatePartyAttributesV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil, nil

	case *PublicUpdatePartyAttributesV1NotFound:
		return nil, nil, nil, nil, v, nil, nil, nil

	case *PublicUpdatePartyAttributesV1PreconditionFailed:
		return nil, nil, nil, nil, nil, v, nil, nil

	case *PublicUpdatePartyAttributesV1InternalServerError:
		return nil, nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) PublicUpdatePartyAttributesV1Short(params *PublicUpdatePartyAttributesV1Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdatePartyAttributesV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicUpdatePartyAttributesV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicUpdatePartyAttributesV1",
		Method:             "PUT",
		PathPattern:        "/lobby/v1/public/party/namespaces/{namespace}/parties/{partyId}/attributes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicUpdatePartyAttributesV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicUpdatePartyAttributesV1OK:
		return v, nil
	case *PublicUpdatePartyAttributesV1BadRequest:
		return nil, v
	case *PublicUpdatePartyAttributesV1Unauthorized:
		return nil, v
	case *PublicUpdatePartyAttributesV1Forbidden:
		return nil, v
	case *PublicUpdatePartyAttributesV1NotFound:
		return nil, v
	case *PublicUpdatePartyAttributesV1PreconditionFailed:
		return nil, v
	case *PublicUpdatePartyAttributesV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
