// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_0_extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new o auth2 0 extension API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for o auth2 0 extension API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetCountryLocationV3(params *GetCountryLocationV3Params, authInfo runtime.ClientAuthInfoWriter) (*GetCountryLocationV3OK, error)
	GetCountryLocationV3Short(params *GetCountryLocationV3Params, authInfo runtime.ClientAuthInfoWriter) (*GetCountryLocationV3OK, error)
	Logout(params *LogoutParams, authInfo runtime.ClientAuthInfoWriter) (*LogoutNoContent, error)
	LogoutShort(params *LogoutParams, authInfo runtime.ClientAuthInfoWriter) (*LogoutNoContent, error)
	PlatformAuthenticationV3(params *PlatformAuthenticationV3Params, authInfo runtime.ClientAuthInfoWriter) (*PlatformAuthenticationV3Found, error)
	PlatformAuthenticationV3Short(params *PlatformAuthenticationV3Params, authInfo runtime.ClientAuthInfoWriter) (*PlatformAuthenticationV3Found, error)
	UserAuthenticationV3(params *UserAuthenticationV3Params, authInfo runtime.ClientAuthInfoWriter) (*UserAuthenticationV3Found, error)
	UserAuthenticationV3Short(params *UserAuthenticationV3Params, authInfo runtime.ClientAuthInfoWriter) (*UserAuthenticationV3Found, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCountryLocationV3 gets country location

  &lt;p&gt;This endpoint get country location based on the request.&lt;/p&gt;
*/
func (a *Client) GetCountryLocationV3(params *GetCountryLocationV3Params, authInfo runtime.ClientAuthInfoWriter) (*GetCountryLocationV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCountryLocationV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCountryLocationV3",
		Method:             "GET",
		PathPattern:        "/iam/v3/location/country",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCountryLocationV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetCountryLocationV3OK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) GetCountryLocationV3Short(params *GetCountryLocationV3Params, authInfo runtime.ClientAuthInfoWriter) (*GetCountryLocationV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCountryLocationV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCountryLocationV3",
		Method:             "GET",
		PathPattern:        "/iam/v3/location/country",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCountryLocationV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetCountryLocationV3OK:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  Logout logouts

  &lt;p&gt;This endpoint is used to remove &lt;b&gt;access_token&lt;/b&gt; cookie and &lt;b&gt;refresh_token&lt;/b&gt; cookie.&lt;/p&gt;
*/
func (a *Client) Logout(params *LogoutParams, authInfo runtime.ClientAuthInfoWriter) (*LogoutNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogoutParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Logout",
		Method:             "POST",
		PathPattern:        "/iam/v3/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LogoutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *LogoutNoContent:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) LogoutShort(params *LogoutParams, authInfo runtime.ClientAuthInfoWriter) (*LogoutNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogoutParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Logout",
		Method:             "POST",
		PathPattern:        "/iam/v3/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LogoutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *LogoutNoContent:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PlatformAuthenticationV3 platforms authentication API

  This endpoint authenticates user platform. It validates user to its
          respective platforms. Deactivated or login-banned users are unable to login. &lt;br&gt;
          &lt;h2&gt;Supported platforms:&lt;/h2&gt;&lt;ul&gt;
          &lt;li&gt;&lt;strong&gt;steamopenid&lt;/strong&gt;&lt;/li&gt;Steam login page will redirects to this endpoint after login success
          as previously defined on openID request parameter &lt;code&gt;openid.return_to&lt;/code&gt; when request login to steam
          https://openid.net/specs/openid-authentication-2_0.html#anchor27
          &lt;li&gt;&lt;strong&gt;ps4web&lt;/strong&gt;&lt;/li&gt;PS4 login page will redirects to this endpoint after login success
          as previously defined on authorize request parameter &lt;code&gt;redirect_uri&lt;/code&gt;
		  https://ps4.siedev.net/resources/documents/WebAPI/1/Auth_WebAPI-Reference/0002.html#0GetAccessTokenUsingAuthorizationCode
          &lt;li&gt;&lt;strong&gt;xblweb&lt;/strong&gt;&lt;/li&gt;XBL login page will redirects to this endpoint after login success
          as previously defined on authorize request parameter &lt;code&gt;redirect_uri&lt;/code&gt;
          &lt;li&gt;&lt;strong&gt;epicgames&lt;/strong&gt;&lt;/li&gt;Epicgames login page will redirects to this endpoint after login success
          or an error occurred. If error, it redirects to the login page.
          &lt;li&gt;&lt;strong&gt;twitch&lt;/strong&gt;&lt;/li&gt;Twitch login page will redirects to this endpoint after login success
          as previously defined on authorize request parameter &lt;code&gt;redirect_uri&lt;/code&gt;
          &lt;/ul&gt; action code : 10709
*/
func (a *Client) PlatformAuthenticationV3(params *PlatformAuthenticationV3Params, authInfo runtime.ClientAuthInfoWriter) (*PlatformAuthenticationV3Found, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPlatformAuthenticationV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PlatformAuthenticationV3",
		Method:             "GET",
		PathPattern:        "/iam/v3/platforms/{platformId}/authenticate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PlatformAuthenticationV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PlatformAuthenticationV3Found:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) PlatformAuthenticationV3Short(params *PlatformAuthenticationV3Params, authInfo runtime.ClientAuthInfoWriter) (*PlatformAuthenticationV3Found, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPlatformAuthenticationV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PlatformAuthenticationV3",
		Method:             "GET",
		PathPattern:        "/iam/v3/platforms/{platformId}/authenticate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PlatformAuthenticationV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PlatformAuthenticationV3Found:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UserAuthenticationV3 authentications API

  &lt;p&gt;This endpoint is being used to authenticate a user account.
					It validates user&#39;s email / username and password. Deactivated or login-banned users are unable to login
					Redirect URI and Client ID must be specified as a pair and only used to redirect to the specified
					redirect URI in case the requestId is no longer valid.&lt;/p&gt;
					&lt;br&gt;action code: 10801
*/
func (a *Client) UserAuthenticationV3(params *UserAuthenticationV3Params, authInfo runtime.ClientAuthInfoWriter) (*UserAuthenticationV3Found, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserAuthenticationV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserAuthenticationV3",
		Method:             "POST",
		PathPattern:        "/iam/v3/authenticate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserAuthenticationV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *UserAuthenticationV3Found:
		return v, nil

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) UserAuthenticationV3Short(params *UserAuthenticationV3Params, authInfo runtime.ClientAuthInfoWriter) (*UserAuthenticationV3Found, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserAuthenticationV3Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserAuthenticationV3",
		Method:             "POST",
		PathPattern:        "/iam/v3/authenticate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserAuthenticationV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *UserAuthenticationV3Found:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
