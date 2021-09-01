// Code generated by go-swagger; DO NOT EDIT.

package third_party

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new third party API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for third party API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AdminCreateThirdPartyConfig(params *AdminCreateThirdPartyConfigParams, authInfo runtime.ClientAuthInfoWriter) (*AdminCreateThirdPartyConfigCreated, *AdminCreateThirdPartyConfigBadRequest, *AdminCreateThirdPartyConfigUnauthorized, *AdminCreateThirdPartyConfigForbidden, *AdminCreateThirdPartyConfigInternalServerError, error)

	AdminDeleteThirdPartyConfig(params *AdminDeleteThirdPartyConfigParams, authInfo runtime.ClientAuthInfoWriter) (*AdminDeleteThirdPartyConfigNoContent, *AdminDeleteThirdPartyConfigBadRequest, *AdminDeleteThirdPartyConfigUnauthorized, *AdminDeleteThirdPartyConfigForbidden, *AdminDeleteThirdPartyConfigInternalServerError, error)

	AdminGetThirdPartyConfig(params *AdminGetThirdPartyConfigParams, authInfo runtime.ClientAuthInfoWriter) (*AdminGetThirdPartyConfigOK, *AdminGetThirdPartyConfigBadRequest, *AdminGetThirdPartyConfigUnauthorized, *AdminGetThirdPartyConfigForbidden, *AdminGetThirdPartyConfigInternalServerError, error)

	AdminUpdateThirdPartyConfig(params *AdminUpdateThirdPartyConfigParams, authInfo runtime.ClientAuthInfoWriter) (*AdminUpdateThirdPartyConfigOK, *AdminUpdateThirdPartyConfigBadRequest, *AdminUpdateThirdPartyConfigUnauthorized, *AdminUpdateThirdPartyConfigForbidden, *AdminUpdateThirdPartyConfigInternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AdminCreateThirdPartyConfig creates third party steam config

  Required permission : <code>ADMIN:NAMESPACE:{namespace}:THIRDPARTY:CONFIG [CREATE]</code> with scope <code>social</code>
			<br>create third party config in a namespace.
*/
func (a *Client) AdminCreateThirdPartyConfig(params *AdminCreateThirdPartyConfigParams, authInfo runtime.ClientAuthInfoWriter) (*AdminCreateThirdPartyConfigCreated, *AdminCreateThirdPartyConfigBadRequest, *AdminCreateThirdPartyConfigUnauthorized, *AdminCreateThirdPartyConfigForbidden, *AdminCreateThirdPartyConfigInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminCreateThirdPartyConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adminCreateThirdPartyConfig",
		Method:             "POST",
		PathPattern:        "/v1/admin/thirdparty/namespaces/{namespace}/config/steam",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdminCreateThirdPartyConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminCreateThirdPartyConfigCreated:
		return v, nil, nil, nil, nil, nil
	case *AdminCreateThirdPartyConfigBadRequest:
		return nil, v, nil, nil, nil, nil
	case *AdminCreateThirdPartyConfigUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *AdminCreateThirdPartyConfigForbidden:
		return nil, nil, nil, v, nil, nil
	case *AdminCreateThirdPartyConfigInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AdminDeleteThirdPartyConfig deletes third party steam config

  Required permission : <code>ADMIN:NAMESPACE:{namespace}:THIRDPARTY:CONFIG [DELETE]</code> with scope <code>social</code>
			<br>delete third party config in a namespace.
*/
func (a *Client) AdminDeleteThirdPartyConfig(params *AdminDeleteThirdPartyConfigParams, authInfo runtime.ClientAuthInfoWriter) (*AdminDeleteThirdPartyConfigNoContent, *AdminDeleteThirdPartyConfigBadRequest, *AdminDeleteThirdPartyConfigUnauthorized, *AdminDeleteThirdPartyConfigForbidden, *AdminDeleteThirdPartyConfigInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminDeleteThirdPartyConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adminDeleteThirdPartyConfig",
		Method:             "DELETE",
		PathPattern:        "/v1/admin/thirdparty/namespaces/{namespace}/config/steam",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdminDeleteThirdPartyConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminDeleteThirdPartyConfigNoContent:
		return v, nil, nil, nil, nil, nil
	case *AdminDeleteThirdPartyConfigBadRequest:
		return nil, v, nil, nil, nil, nil
	case *AdminDeleteThirdPartyConfigUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *AdminDeleteThirdPartyConfigForbidden:
		return nil, nil, nil, v, nil, nil
	case *AdminDeleteThirdPartyConfigInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AdminGetThirdPartyConfig gets third party steam config

  Required permission : <code>ADMIN:NAMESPACE:{namespace}:THIRDPARTY:CONFIG [READ]</code> with scope <code>social</code>
			<br>get third party config for specified namespace.
*/
func (a *Client) AdminGetThirdPartyConfig(params *AdminGetThirdPartyConfigParams, authInfo runtime.ClientAuthInfoWriter) (*AdminGetThirdPartyConfigOK, *AdminGetThirdPartyConfigBadRequest, *AdminGetThirdPartyConfigUnauthorized, *AdminGetThirdPartyConfigForbidden, *AdminGetThirdPartyConfigInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminGetThirdPartyConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adminGetThirdPartyConfig",
		Method:             "GET",
		PathPattern:        "/v1/admin/thirdparty/namespaces/{namespace}/config/steam",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdminGetThirdPartyConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminGetThirdPartyConfigOK:
		return v, nil, nil, nil, nil, nil
	case *AdminGetThirdPartyConfigBadRequest:
		return nil, v, nil, nil, nil, nil
	case *AdminGetThirdPartyConfigUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *AdminGetThirdPartyConfigForbidden:
		return nil, nil, nil, v, nil, nil
	case *AdminGetThirdPartyConfigInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AdminUpdateThirdPartyConfig updates third party steam config

  Required permission : <code>ADMIN:NAMESPACE:{namespace}:THIRDPARTY:CONFIG [UPDATE]</code> with scope <code>social</code>
			<br>Update third party config in a namespace.
*/
func (a *Client) AdminUpdateThirdPartyConfig(params *AdminUpdateThirdPartyConfigParams, authInfo runtime.ClientAuthInfoWriter) (*AdminUpdateThirdPartyConfigOK, *AdminUpdateThirdPartyConfigBadRequest, *AdminUpdateThirdPartyConfigUnauthorized, *AdminUpdateThirdPartyConfigForbidden, *AdminUpdateThirdPartyConfigInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminUpdateThirdPartyConfigParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adminUpdateThirdPartyConfig",
		Method:             "PUT",
		PathPattern:        "/v1/admin/thirdparty/namespaces/{namespace}/config/steam",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdminUpdateThirdPartyConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminUpdateThirdPartyConfigOK:
		return v, nil, nil, nil, nil, nil
	case *AdminUpdateThirdPartyConfigBadRequest:
		return nil, v, nil, nil, nil, nil
	case *AdminUpdateThirdPartyConfigUnauthorized:
		return nil, nil, v, nil, nil, nil
	case *AdminUpdateThirdPartyConfigForbidden:
		return nil, nil, nil, v, nil, nil
	case *AdminUpdateThirdPartyConfigInternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}