// Code generated by go-swagger; DO NOT EDIT.

package eligibilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new eligibilities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for eligibilities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	RetrieveEligibilitiesPublic(params *RetrieveEligibilitiesPublicParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveEligibilitiesPublicOK, error)

	RetrieveEligibilitiesPublicIndirect(params *RetrieveEligibilitiesPublicIndirectParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveEligibilitiesPublicIndirectOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  RetrieveEligibilitiesPublic checks user legal eligibility

  Retrieve the active policies and its conformance status by user.<br>This process supports cross-namespace checking, that means if the active policy already accepted by the same user in other namespace, then it will be considered as eligible.<br/><br/>Other detail info: <ul><li><i>Required permission</i>: login user</li></ul>
*/
func (a *Client) RetrieveEligibilitiesPublic(params *RetrieveEligibilitiesPublicParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveEligibilitiesPublicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveEligibilitiesPublicParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveEligibilitiesPublic",
		Method:             "GET",
		PathPattern:        "/public/eligibilities/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RetrieveEligibilitiesPublicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveEligibilitiesPublicOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveEligibilitiesPublicIndirect checks user legal eligibility

  Retrieve the active policies and its conformance status by user. This endpoint used by Authentication Service during user login.<br>This process supports cross-namespace checking, that means if the active policy already accepted by the same user in other namespace, then it will be considered as eligible.<br/><br/>Other detail info: <ul><li><i>Required permission</i>: login user</li></ul>
*/
func (a *Client) RetrieveEligibilitiesPublicIndirect(params *RetrieveEligibilitiesPublicIndirectParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveEligibilitiesPublicIndirectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveEligibilitiesPublicIndirectParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveEligibilitiesPublicIndirect",
		Method:             "GET",
		PathPattern:        "/public/eligibilities/namespaces/{namespace}/countries/{countryCode}/clients/{clientId}/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RetrieveEligibilitiesPublicIndirectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveEligibilitiesPublicIndirectOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
