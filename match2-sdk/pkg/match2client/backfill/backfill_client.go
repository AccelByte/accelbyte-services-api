// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package backfill

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backfill API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backfill API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AcceptBackfill(params *AcceptBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptBackfillOK, *AcceptBackfillBadRequest, *AcceptBackfillUnauthorized, *AcceptBackfillForbidden, *AcceptBackfillNotFound, *AcceptBackfillInternalServerError, error)
	AcceptBackfillShort(params *AcceptBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptBackfillOK, error)
	CreateBackfill(params *CreateBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBackfillCreated, *CreateBackfillBadRequest, *CreateBackfillUnauthorized, *CreateBackfillForbidden, *CreateBackfillNotFound, *CreateBackfillConflict, *CreateBackfillInternalServerError, error)
	CreateBackfillShort(params *CreateBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBackfillCreated, error)
	RejectBackfill(params *RejectBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*RejectBackfillOK, *RejectBackfillBadRequest, *RejectBackfillUnauthorized, *RejectBackfillForbidden, *RejectBackfillNotFound, *RejectBackfillInternalServerError, error)
	RejectBackfillShort(params *RejectBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*RejectBackfillOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Deprecated: Use AcceptBackfillShort instead.

  AcceptBackfill accepts a backfill proposal

   Required Permission: NAMESPACE:{namespace}:MATCHMAKING:BACKFILL [UPDATE]

Required Scope: social

Accept backfill proposal

*/
func (a *Client) AcceptBackfill(params *AcceptBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptBackfillOK, *AcceptBackfillBadRequest, *AcceptBackfillUnauthorized, *AcceptBackfillForbidden, *AcceptBackfillNotFound, *AcceptBackfillInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcceptBackfillParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AcceptBackfill",
		Method:             "PUT",
		PathPattern:        "/match2/v1/namespaces/{namespace}/backfill/{backfillID}/proposal/accept",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptBackfillReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AcceptBackfillOK:
		return v, nil, nil, nil, nil, nil, nil

	case *AcceptBackfillBadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *AcceptBackfillUnauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *AcceptBackfillForbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *AcceptBackfillNotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *AcceptBackfillInternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AcceptBackfillShort accepts a backfill proposal

   Required Permission: NAMESPACE:{namespace}:MATCHMAKING:BACKFILL [UPDATE]

Required Scope: social

Accept backfill proposal

*/
func (a *Client) AcceptBackfillShort(params *AcceptBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptBackfillOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcceptBackfillParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AcceptBackfill",
		Method:             "PUT",
		PathPattern:        "/match2/v1/namespaces/{namespace}/backfill/{backfillID}/proposal/accept",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptBackfillReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AcceptBackfillOK:
		return v, nil
	case *AcceptBackfillBadRequest:
		return nil, v
	case *AcceptBackfillUnauthorized:
		return nil, v
	case *AcceptBackfillForbidden:
		return nil, v
	case *AcceptBackfillNotFound:
		return nil, v
	case *AcceptBackfillInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use CreateBackfillShort instead.

  CreateBackfill creates a backfill ticket

   Required Permission: NAMESPACE:{namespace}:MATCHMAKING:BACKFILL [CREATE]

Required Scope: social

Create backfill ticket

*/
func (a *Client) CreateBackfill(params *CreateBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBackfillCreated, *CreateBackfillBadRequest, *CreateBackfillUnauthorized, *CreateBackfillForbidden, *CreateBackfillNotFound, *CreateBackfillConflict, *CreateBackfillInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBackfillParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateBackfill",
		Method:             "POST",
		PathPattern:        "/match2/v1/namespaces/{namespace}/backfill",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBackfillReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateBackfillCreated:
		return v, nil, nil, nil, nil, nil, nil, nil

	case *CreateBackfillBadRequest:
		return nil, v, nil, nil, nil, nil, nil, nil

	case *CreateBackfillUnauthorized:
		return nil, nil, v, nil, nil, nil, nil, nil

	case *CreateBackfillForbidden:
		return nil, nil, nil, v, nil, nil, nil, nil

	case *CreateBackfillNotFound:
		return nil, nil, nil, nil, v, nil, nil, nil

	case *CreateBackfillConflict:
		return nil, nil, nil, nil, nil, v, nil, nil

	case *CreateBackfillInternalServerError:
		return nil, nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreateBackfillShort creates a backfill ticket

   Required Permission: NAMESPACE:{namespace}:MATCHMAKING:BACKFILL [CREATE]

Required Scope: social

Create backfill ticket

*/
func (a *Client) CreateBackfillShort(params *CreateBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBackfillCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBackfillParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateBackfill",
		Method:             "POST",
		PathPattern:        "/match2/v1/namespaces/{namespace}/backfill",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBackfillReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *CreateBackfillCreated:
		return v, nil
	case *CreateBackfillBadRequest:
		return nil, v
	case *CreateBackfillUnauthorized:
		return nil, v
	case *CreateBackfillForbidden:
		return nil, v
	case *CreateBackfillNotFound:
		return nil, v
	case *CreateBackfillConflict:
		return nil, v
	case *CreateBackfillInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: Use RejectBackfillShort instead.

  RejectBackfill rejects a backfill proposal

   Required Permission: NAMESPACE:{namespace}:MATCHMAKING:BACKFILL [UPDATE]

Required Scope: social

Reject backfill proposal

*/
func (a *Client) RejectBackfill(params *RejectBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*RejectBackfillOK, *RejectBackfillBadRequest, *RejectBackfillUnauthorized, *RejectBackfillForbidden, *RejectBackfillNotFound, *RejectBackfillInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRejectBackfillParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RejectBackfill",
		Method:             "PUT",
		PathPattern:        "/match2/v1/namespaces/{namespace}/backfill/{backfillID}/proposal/reject",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RejectBackfillReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *RejectBackfillOK:
		return v, nil, nil, nil, nil, nil, nil

	case *RejectBackfillBadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *RejectBackfillUnauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *RejectBackfillForbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *RejectBackfillNotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *RejectBackfillInternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RejectBackfillShort rejects a backfill proposal

   Required Permission: NAMESPACE:{namespace}:MATCHMAKING:BACKFILL [UPDATE]

Required Scope: social

Reject backfill proposal

*/
func (a *Client) RejectBackfillShort(params *RejectBackfillParams, authInfo runtime.ClientAuthInfoWriter) (*RejectBackfillOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRejectBackfillParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RejectBackfill",
		Method:             "PUT",
		PathPattern:        "/match2/v1/namespaces/{namespace}/backfill/{backfillID}/proposal/reject",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RejectBackfillReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RejectBackfillOK:
		return v, nil
	case *RejectBackfillBadRequest:
		return nil, v
	case *RejectBackfillUnauthorized:
		return nil, v
	case *RejectBackfillForbidden:
		return nil, v
	case *RejectBackfillNotFound:
		return nil, v
	case *RejectBackfillInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
