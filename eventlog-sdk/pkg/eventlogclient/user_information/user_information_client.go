// Code generated by go-swagger; DO NOT EDIT.

package user_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user information API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user information API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteUserActivitiesHandler(params *DeleteUserActivitiesHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserActivitiesHandlerNoContent, *DeleteUserActivitiesHandlerUnauthorized, *DeleteUserActivitiesHandlerForbidden, *DeleteUserActivitiesHandlerNotFound, *DeleteUserActivitiesHandlerInternalServerError, error)
	DeleteUserActivitiesHandlerShort(params *DeleteUserActivitiesHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserActivitiesHandlerNoContent, error)
	GetUserActivitiesHandler(params *GetUserActivitiesHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserActivitiesHandlerOK, *GetUserActivitiesHandlerBadRequest, *GetUserActivitiesHandlerUnauthorized, *GetUserActivitiesHandlerForbidden, *GetUserActivitiesHandlerNotFound, *GetUserActivitiesHandlerInternalServerError, error)
	GetUserActivitiesHandlerShort(params *GetUserActivitiesHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserActivitiesHandlerOK, error)
	LastUserActivityTimeHandler(params *LastUserActivityTimeHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*LastUserActivityTimeHandlerOK, *LastUserActivityTimeHandlerUnauthorized, *LastUserActivityTimeHandlerForbidden, *LastUserActivityTimeHandlerNotFound, *LastUserActivityTimeHandlerInternalServerError, error)
	LastUserActivityTimeHandlerShort(params *LastUserActivityTimeHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*LastUserActivityTimeHandlerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteUserActivitiesHandler deletes all user s activities

  Required permission &lt;code&gt;NAMESPACE:{namespace}:EVENT [UPDATE]&lt;/code&gt;and scope &lt;code&gt;analytics&lt;/code&gt;
*/
func (a *Client) DeleteUserActivitiesHandler(params *DeleteUserActivitiesHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserActivitiesHandlerNoContent, *DeleteUserActivitiesHandlerUnauthorized, *DeleteUserActivitiesHandlerForbidden, *DeleteUserActivitiesHandlerNotFound, *DeleteUserActivitiesHandlerInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserActivitiesHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUserActivitiesHandler",
		Method:             "DELETE",
		PathPattern:        "/event/namespaces/{namespace}/users/{userId}/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserActivitiesHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteUserActivitiesHandlerNoContent:
		return v, nil, nil, nil, nil, nil

	case *DeleteUserActivitiesHandlerUnauthorized:
		return nil, v, nil, nil, nil, nil

	case *DeleteUserActivitiesHandlerForbidden:
		return nil, nil, v, nil, nil, nil

	case *DeleteUserActivitiesHandlerNotFound:
		return nil, nil, nil, v, nil, nil

	case *DeleteUserActivitiesHandlerInternalServerError:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) DeleteUserActivitiesHandlerShort(params *DeleteUserActivitiesHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserActivitiesHandlerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserActivitiesHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUserActivitiesHandler",
		Method:             "DELETE",
		PathPattern:        "/event/namespaces/{namespace}/users/{userId}/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserActivitiesHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DeleteUserActivitiesHandlerNoContent:
		return v, nil
	case *DeleteUserActivitiesHandlerUnauthorized:
		return nil, v
	case *DeleteUserActivitiesHandlerForbidden:
		return nil, v
	case *DeleteUserActivitiesHandlerNotFound:
		return nil, v
	case *DeleteUserActivitiesHandlerInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetUserActivitiesHandler gets all user s activities

  Required permission &lt;code&gt;NAMESPACE:{namespace}:EVENT [UPDATE]&lt;/code&gt;and scope &lt;code&gt;analytics&lt;/code&gt;
*/
func (a *Client) GetUserActivitiesHandler(params *GetUserActivitiesHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserActivitiesHandlerOK, *GetUserActivitiesHandlerBadRequest, *GetUserActivitiesHandlerUnauthorized, *GetUserActivitiesHandlerForbidden, *GetUserActivitiesHandlerNotFound, *GetUserActivitiesHandlerInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserActivitiesHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUserActivitiesHandler",
		Method:             "GET",
		PathPattern:        "/event/namespaces/{namespace}/users/{userId}/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserActivitiesHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetUserActivitiesHandlerOK:
		return v, nil, nil, nil, nil, nil, nil

	case *GetUserActivitiesHandlerBadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *GetUserActivitiesHandlerUnauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *GetUserActivitiesHandlerForbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *GetUserActivitiesHandlerNotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *GetUserActivitiesHandlerInternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) GetUserActivitiesHandlerShort(params *GetUserActivitiesHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserActivitiesHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserActivitiesHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUserActivitiesHandler",
		Method:             "GET",
		PathPattern:        "/event/namespaces/{namespace}/users/{userId}/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserActivitiesHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetUserActivitiesHandlerOK:
		return v, nil
	case *GetUserActivitiesHandlerBadRequest:
		return nil, v
	case *GetUserActivitiesHandlerUnauthorized:
		return nil, v
	case *GetUserActivitiesHandlerForbidden:
		return nil, v
	case *GetUserActivitiesHandlerNotFound:
		return nil, v
	case *GetUserActivitiesHandlerInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  LastUserActivityTimeHandler gets last activity time of a user in a namespace

  Required permission &lt;code&gt;NAMESPACE:{namespace}:EVENT [UPDATE]&lt;/code&gt;and scope &lt;code&gt;analytics&lt;/code&gt;
*/
func (a *Client) LastUserActivityTimeHandler(params *LastUserActivityTimeHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*LastUserActivityTimeHandlerOK, *LastUserActivityTimeHandlerUnauthorized, *LastUserActivityTimeHandlerForbidden, *LastUserActivityTimeHandlerNotFound, *LastUserActivityTimeHandlerInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLastUserActivityTimeHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LastUserActivityTimeHandler",
		Method:             "GET",
		PathPattern:        "/event/namespaces/{namespace}/users/{userId}/lastActivityTime",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LastUserActivityTimeHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *LastUserActivityTimeHandlerOK:
		return v, nil, nil, nil, nil, nil

	case *LastUserActivityTimeHandlerUnauthorized:
		return nil, v, nil, nil, nil, nil

	case *LastUserActivityTimeHandlerForbidden:
		return nil, nil, v, nil, nil, nil

	case *LastUserActivityTimeHandlerNotFound:
		return nil, nil, nil, v, nil, nil

	case *LastUserActivityTimeHandlerInternalServerError:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) LastUserActivityTimeHandlerShort(params *LastUserActivityTimeHandlerParams, authInfo runtime.ClientAuthInfoWriter) (*LastUserActivityTimeHandlerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLastUserActivityTimeHandlerParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LastUserActivityTimeHandler",
		Method:             "GET",
		PathPattern:        "/event/namespaces/{namespace}/users/{userId}/lastActivityTime",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LastUserActivityTimeHandlerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *LastUserActivityTimeHandlerOK:
		return v, nil
	case *LastUserActivityTimeHandlerUnauthorized:
		return nil, v
	case *LastUserActivityTimeHandlerForbidden:
		return nil, v
	case *LastUserActivityTimeHandlerNotFound:
		return nil, v
	case *LastUserActivityTimeHandlerInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
