// Code generated by go-swagger; DO NOT EDIT.

package game_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new game profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for game profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetProfile(params *GetProfileParams, authInfo runtime.ClientAuthInfoWriter) (*GetProfileOK, *GetProfileNotFound, error)

	GetUserProfiles(params *GetUserProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserProfilesOK, error)

	PublicCreateProfile(params *PublicCreateProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PublicCreateProfileCreated, *PublicCreateProfileUnprocessableEntity, error)

	PublicDeleteProfile(params *PublicDeleteProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PublicDeleteProfileNoContent, *PublicDeleteProfileNotFound, error)

	PublicGetProfile(params *PublicGetProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetProfileOK, *PublicGetProfileNotFound, error)

	PublicGetProfileAttribute(params *PublicGetProfileAttributeParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetProfileAttributeOK, *PublicGetProfileAttributeNotFound, error)

	PublicGetUserGameProfiles(params *PublicGetUserGameProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetUserGameProfilesOK, *PublicGetUserGameProfilesBadRequest, error)

	PublicGetUserProfiles(params *PublicGetUserProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetUserProfilesOK, error)

	PublicUpdateAttribute(params *PublicUpdateAttributeParams, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdateAttributeOK, *PublicUpdateAttributeBadRequest, *PublicUpdateAttributeNotFound, error)

	PublicUpdateProfile(params *PublicUpdateProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdateProfileOK, *PublicUpdateProfileNotFound, *PublicUpdateProfileUnprocessableEntity, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetProfile returns profile for a user

  Returns profile for a user.<br>Other detail info:<ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:GAMEPROFILE", action=2 (READ)</li><li><i>Returns</i>: game profile info</li></ul>
*/
func (a *Client) GetProfile(params *GetProfileParams, authInfo runtime.ClientAuthInfoWriter) (*GetProfileOK, *GetProfileNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProfileParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProfile",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/profiles/{profileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetProfileOK:
		return v, nil, nil
	case *GetProfileNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetUserProfiles returns all profiles header for a user

  Returns all profiles' header for a user.<br>Other detail info:<ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:GAMEPROFILE", action=2 (READ)</li><li><i>Returns</i>: list of profiles</li></ul>
*/
func (a *Client) GetUserProfiles(params *GetUserProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserProfilesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserProfiles",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserProfilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetUserProfilesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicCreateProfile creates a new profile for user

  Create new profile for user.<br>Other detail info:<ul><li><i>Required permission</li>: resource="NAMESPACE:{namespace}:USER:{userId}:GAMEPROFILE", action=1 (CREATE)</li><li><i>Returns</li>: created game profile</li></ul>
*/
func (a *Client) PublicCreateProfile(params *PublicCreateProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PublicCreateProfileCreated, *PublicCreateProfileUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicCreateProfileParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicCreateProfile",
		Method:             "POST",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicCreateProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PublicCreateProfileCreated:
		return v, nil, nil
	case *PublicCreateProfileUnprocessableEntity:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicDeleteProfile deletes game profile

  Deletes game profile.<br>Other detail info:<ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:GAMEPROFILE", action=8 (DELETE)</li></ul>
*/
func (a *Client) PublicDeleteProfile(params *PublicDeleteProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PublicDeleteProfileNoContent, *PublicDeleteProfileNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicDeleteProfileParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicDeleteProfile",
		Method:             "DELETE",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicDeleteProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PublicDeleteProfileNoContent:
		return v, nil, nil
	case *PublicDeleteProfileNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetProfile returns profile for a user

  Returns profile for a user.<br>Other detail info:<ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:GAMEPROFILE", action=2 (READ)</li><li><i>Returns</i>: game profile info</li></ul>
*/
func (a *Client) PublicGetProfile(params *PublicGetProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetProfileOK, *PublicGetProfileNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetProfileParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetProfile",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PublicGetProfileOK:
		return v, nil, nil
	case *PublicGetProfileNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetProfileAttribute returns game profile attribute

  Returns game profile attribute.<br>Other detail info:<ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:GAMEPROFILE", action=2 (READ)</li><li><i>Returns</i>: attribute info</li></ul>
*/
func (a *Client) PublicGetProfileAttribute(params *PublicGetProfileAttributeParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetProfileAttributeOK, *PublicGetProfileAttributeNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetProfileAttributeParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetProfileAttribute",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}/attributes/{attributeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetProfileAttributeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PublicGetProfileAttributeOK:
		return v, nil, nil
	case *PublicGetProfileAttributeNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetUserGameProfiles returns all profiles for specified users

  Returns all profiles for specified users.<br>Other detail info:<ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:GAMEPROFILE", action=2 (READ)
<li><i>Returns</i>: list of profiles</ul>
*/
func (a *Client) PublicGetUserGameProfiles(params *PublicGetUserGameProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetUserGameProfilesOK, *PublicGetUserGameProfilesBadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetUserGameProfilesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetUserGameProfiles",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetUserGameProfilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PublicGetUserGameProfilesOK:
		return v, nil, nil
	case *PublicGetUserGameProfilesBadRequest:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetUserProfiles returns all profiles header for a user

  Returns all profiles' header for a user.<br>Other detail info:<ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:GAMEPROFILE", action=2 (READ)</li><li><i>Returns</i>: list of profiles</li></ul>
*/
func (a *Client) PublicGetUserProfiles(params *PublicGetUserProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetUserProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetUserProfilesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetUserProfiles",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicGetUserProfilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicGetUserProfilesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicUpdateAttribute updates game profile attribute

  Updates game profile attribute, returns updated profile.<br>Other detail info:<ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:GAMEPROFILE", action=4 (UPDATE)</li><li><i>Returns</i>: updated attribute</li></ul>
*/
func (a *Client) PublicUpdateAttribute(params *PublicUpdateAttributeParams, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdateAttributeOK, *PublicUpdateAttributeBadRequest, *PublicUpdateAttributeNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicUpdateAttributeParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicUpdateAttribute",
		Method:             "PUT",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}/attributes/{attributeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicUpdateAttributeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicUpdateAttributeOK:
		return v, nil, nil, nil
	case *PublicUpdateAttributeBadRequest:
		return nil, v, nil, nil
	case *PublicUpdateAttributeNotFound:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicUpdateProfile updates user game profile

  Updates user game profile, returns updated profile.<br>Other detail info:<ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:GAMEPROFILE", action=4 (UPDATE)</li><li><i>Returns</i>: updated game profile</li></ul>
*/
func (a *Client) PublicUpdateProfile(params *PublicUpdateProfileParams, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdateProfileOK, *PublicUpdateProfileNotFound, *PublicUpdateProfileUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicUpdateProfileParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicUpdateProfile",
		Method:             "PUT",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicUpdateProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicUpdateProfileOK:
		return v, nil, nil, nil
	case *PublicUpdateProfileNotFound:
		return nil, v, nil, nil
	case *PublicUpdateProfileUnprocessableEntity:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}