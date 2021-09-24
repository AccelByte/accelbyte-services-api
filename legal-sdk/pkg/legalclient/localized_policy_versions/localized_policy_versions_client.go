// Code generated by go-swagger; DO NOT EDIT.

package localized_policy_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new localized policy versions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for localized policy versions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateLocalizedPolicyVersion(params *CreateLocalizedPolicyVersionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLocalizedPolicyVersionCreated, *CreateLocalizedPolicyVersionBadRequest, *CreateLocalizedPolicyVersionConflict, error)

	RequestPresignedURL(params *RequestPresignedURLParams, authInfo runtime.ClientAuthInfoWriter) (*RequestPresignedURLCreated, *RequestPresignedURLBadRequest, error)

	RetrieveLocalizedPolicyVersions(params *RetrieveLocalizedPolicyVersionsParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveLocalizedPolicyVersionsOK, error)

	RetrieveSingleLocalizedPolicyVersion(params *RetrieveSingleLocalizedPolicyVersionParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveSingleLocalizedPolicyVersionOK, *RetrieveSingleLocalizedPolicyVersionBadRequest, error)

	RetrieveSingleLocalizedPolicyVersion1(params *RetrieveSingleLocalizedPolicyVersion1Params) (*RetrieveSingleLocalizedPolicyVersion1OK, *RetrieveSingleLocalizedPolicyVersion1NotFound, error)

	SetDefaultPolicy(params *SetDefaultPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*SetDefaultPolicyOK, error)

	UpdateLocalizedPolicyVersion(params *UpdateLocalizedPolicyVersionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLocalizedPolicyVersionOK, *UpdateLocalizedPolicyVersionBadRequest, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateLocalizedPolicyVersion creates a localized version from country specific policy

  Create a version of a particular country-specific policy.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:*:LEGAL", action=1 (CREATE)</li></ul>
*/
func (a *Client) CreateLocalizedPolicyVersion(params *CreateLocalizedPolicyVersionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLocalizedPolicyVersionCreated, *CreateLocalizedPolicyVersionBadRequest, *CreateLocalizedPolicyVersionConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLocalizedPolicyVersionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createLocalizedPolicyVersion",
		Method:             "POST",
		PathPattern:        "/admin/localized-policy-versions/versions/{policyVersionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateLocalizedPolicyVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateLocalizedPolicyVersionCreated:
		return v, nil, nil, nil
	case *CreateLocalizedPolicyVersionBadRequest:
		return nil, v, nil, nil
	case *CreateLocalizedPolicyVersionConflict:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RequestPresignedURL requests presigned URL for upload document

  Request presigned URL for  upload attachment for a particular localized version of base policy.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:*:LEGAL", action=1 (CREATE)</li></ul>
*/
func (a *Client) RequestPresignedURL(params *RequestPresignedURLParams, authInfo runtime.ClientAuthInfoWriter) (*RequestPresignedURLCreated, *RequestPresignedURLBadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestPresignedURLParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "requestPresignedURL",
		Method:             "POST",
		PathPattern:        "/admin/localized-policy-versions/{localizedPolicyVersionId}/attachments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RequestPresignedURLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *RequestPresignedURLCreated:
		return v, nil, nil
	case *RequestPresignedURLBadRequest:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveLocalizedPolicyVersions retrieves versions from country specific policy

  Retrieve versions of a particular country-specific policy.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:*:LEGAL", action=2 (READ)</li></ul>
*/
func (a *Client) RetrieveLocalizedPolicyVersions(params *RetrieveLocalizedPolicyVersionsParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveLocalizedPolicyVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveLocalizedPolicyVersionsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveLocalizedPolicyVersions",
		Method:             "GET",
		PathPattern:        "/admin/localized-policy-versions/versions/{policyVersionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RetrieveLocalizedPolicyVersionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RetrieveLocalizedPolicyVersionsOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveSingleLocalizedPolicyVersion retrieves a localized version from country specific policy

  Retrieve a version of a particular country-specific policy.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:*:LEGAL", action=2 (READ)</li></ul>
*/
func (a *Client) RetrieveSingleLocalizedPolicyVersion(params *RetrieveSingleLocalizedPolicyVersionParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveSingleLocalizedPolicyVersionOK, *RetrieveSingleLocalizedPolicyVersionBadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveSingleLocalizedPolicyVersionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveSingleLocalizedPolicyVersion",
		Method:             "GET",
		PathPattern:        "/admin/localized-policy-versions/{localizedPolicyVersionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RetrieveSingleLocalizedPolicyVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *RetrieveSingleLocalizedPolicyVersionOK:
		return v, nil, nil
	case *RetrieveSingleLocalizedPolicyVersionBadRequest:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RetrieveSingleLocalizedPolicyVersion1 retrieves a localized version

  Retrieve specific localized policy version including the policy version and base policy version where the localized policy version located.<br>Other detail info: <ul></ul>
*/
func (a *Client) RetrieveSingleLocalizedPolicyVersion1(params *RetrieveSingleLocalizedPolicyVersion1Params) (*RetrieveSingleLocalizedPolicyVersion1OK, *RetrieveSingleLocalizedPolicyVersion1NotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveSingleLocalizedPolicyVersion1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveSingleLocalizedPolicyVersion_1",
		Method:             "GET",
		PathPattern:        "/public/localized-policy-versions/{localizedPolicyVersionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RetrieveSingleLocalizedPolicyVersion1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *RetrieveSingleLocalizedPolicyVersion1OK:
		return v, nil, nil
	case *RetrieveSingleLocalizedPolicyVersion1NotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  SetDefaultPolicy sets default localized policy

  Update a localized version policy to be the default.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:*:LEGAL", action=4 (UPDATE)</li></ul>
*/
func (a *Client) SetDefaultPolicy(params *SetDefaultPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*SetDefaultPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDefaultPolicyParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setDefaultPolicy",
		Method:             "PATCH",
		PathPattern:        "/admin/localized-policy-versions/{localizedPolicyVersionId}/default",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetDefaultPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *SetDefaultPolicyOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateLocalizedPolicyVersion updates a localized version from country specific policy

  Update a version of a particular country-specific policy.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:*:LEGAL", action=4 (UPDATE)</li></ul>
*/
func (a *Client) UpdateLocalizedPolicyVersion(params *UpdateLocalizedPolicyVersionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLocalizedPolicyVersionOK, *UpdateLocalizedPolicyVersionBadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLocalizedPolicyVersionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateLocalizedPolicyVersion",
		Method:             "PUT",
		PathPattern:        "/admin/localized-policy-versions/{localizedPolicyVersionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateLocalizedPolicyVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateLocalizedPolicyVersionOK:
		return v, nil, nil
	case *UpdateLocalizedPolicyVersionBadRequest:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
