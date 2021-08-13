// Code generated by go-swagger; DO NOT EDIT.

package users_v4

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users v4 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users v4 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AdminAddUserRoleV4(params *AdminAddUserRoleV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminAddUserRoleV4OK, *AdminAddUserRoleV4BadRequest, *AdminAddUserRoleV4Forbidden, *AdminAddUserRoleV4NotFound, *AdminAddUserRoleV4UnprocessableEntity, *AdminAddUserRoleV4InternalServerError, error)

	AdminListUserRolesV4(params *AdminListUserRolesV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminListUserRolesV4OK, *AdminListUserRolesV4Forbidden, *AdminListUserRolesV4NotFound, *AdminListUserRolesV4InternalServerError, error)

	AdminRemoveUserRoleV4(params *AdminRemoveUserRoleV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminRemoveUserRoleV4NoContent, *AdminRemoveUserRoleV4BadRequest, *AdminRemoveUserRoleV4Forbidden, *AdminRemoveUserRoleV4NotFound, *AdminRemoveUserRoleV4UnprocessableEntity, *AdminRemoveUserRoleV4InternalServerError, error)

	AdminUpdateUserEmailAddressV4(params *AdminUpdateUserEmailAddressV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminUpdateUserEmailAddressV4NoContent, *AdminUpdateUserEmailAddressV4BadRequest, *AdminUpdateUserEmailAddressV4Unauthorized, *AdminUpdateUserEmailAddressV4NotFound, *AdminUpdateUserEmailAddressV4Conflict, *AdminUpdateUserEmailAddressV4InternalServerError, error)

	AdminUpdateUserRoleV4(params *AdminUpdateUserRoleV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminUpdateUserRoleV4OK, *AdminUpdateUserRoleV4BadRequest, *AdminUpdateUserRoleV4Forbidden, *AdminUpdateUserRoleV4NotFound, *AdminUpdateUserRoleV4UnprocessableEntity, *AdminUpdateUserRoleV4InternalServerError, error)

	AdminUpdateUserV4(params *AdminUpdateUserV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminUpdateUserV4OK, *AdminUpdateUserV4BadRequest, *AdminUpdateUserV4Unauthorized, *AdminUpdateUserV4Forbidden, *AdminUpdateUserV4NotFound, *AdminUpdateUserV4Conflict, *AdminUpdateUserV4InternalServerError, error)

	CreateUserFromInvitationV4(params *CreateUserFromInvitationV4Params, authInfo runtime.ClientAuthInfoWriter) (*CreateUserFromInvitationV4Created, *CreateUserFromInvitationV4BadRequest, *CreateUserFromInvitationV4NotFound, *CreateUserFromInvitationV4InternalServerError, error)

	PublicCreateUserV4(params *PublicCreateUserV4Params, authInfo runtime.ClientAuthInfoWriter) (*PublicCreateUserV4Created, *PublicCreateUserV4BadRequest, *PublicCreateUserV4NotFound, *PublicCreateUserV4Conflict, *PublicCreateUserV4InternalServerError, error)

	PublicUpdateUserEmailAddressV4(params *PublicUpdateUserEmailAddressV4Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdateUserEmailAddressV4NoContent, *PublicUpdateUserEmailAddressV4BadRequest, *PublicUpdateUserEmailAddressV4Unauthorized, *PublicUpdateUserEmailAddressV4NotFound, *PublicUpdateUserEmailAddressV4Conflict, *PublicUpdateUserEmailAddressV4InternalServerError, error)

	PublicUpdateUserV4(params *PublicUpdateUserV4Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdateUserV4OK, *PublicUpdateUserV4BadRequest, *PublicUpdateUserV4Unauthorized, *PublicUpdateUserV4Conflict, *PublicUpdateUserV4InternalServerError, error)

	PublicUpgradeHeadlessAccountV4(params *PublicUpgradeHeadlessAccountV4Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpgradeHeadlessAccountV4OK, *PublicUpgradeHeadlessAccountV4BadRequest, *PublicUpgradeHeadlessAccountV4Unauthorized, *PublicUpgradeHeadlessAccountV4NotFound, *PublicUpgradeHeadlessAccountV4Conflict, *PublicUpgradeHeadlessAccountV4InternalServerError, error)

	PublicUpgradeHeadlessAccountWithVerificationCodeV4(params *PublicUpgradeHeadlessAccountWithVerificationCodeV4Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpgradeHeadlessAccountWithVerificationCodeV4OK, *PublicUpgradeHeadlessAccountWithVerificationCodeV4BadRequest, *PublicUpgradeHeadlessAccountWithVerificationCodeV4Unauthorized, *PublicUpgradeHeadlessAccountWithVerificationCodeV4Forbidden, *PublicUpgradeHeadlessAccountWithVerificationCodeV4NotFound, *PublicUpgradeHeadlessAccountWithVerificationCodeV4Conflict, *PublicUpgradeHeadlessAccountWithVerificationCodeV4InternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AdminAddUserRoleV4 admins add user s role v4

  This endpoint requires ADMIN:NAMESPACE:{namespace}:ROLE:USER:* [UPDATE] permission.

New role will be appended to user's current roles. Request body need to specify allowed namespace for given role to support new role restriction.
Skipped the check whether the user performing the request is a role manager / assigner since there is a plan to discard the role manager / assigner.

*/
func (a *Client) AdminAddUserRoleV4(params *AdminAddUserRoleV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminAddUserRoleV4OK, *AdminAddUserRoleV4BadRequest, *AdminAddUserRoleV4Forbidden, *AdminAddUserRoleV4NotFound, *AdminAddUserRoleV4UnprocessableEntity, *AdminAddUserRoleV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminAddUserRoleV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AdminAddUserRoleV4",
		Method:             "POST",
		PathPattern:        "/iam/v4/admin/namespaces/{namespace}/users/{userId}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdminAddUserRoleV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminAddUserRoleV4OK:
		return v, nil, nil, nil, nil, nil, nil
	case *AdminAddUserRoleV4BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *AdminAddUserRoleV4Forbidden:
		return nil, nil, v, nil, nil, nil, nil
	case *AdminAddUserRoleV4NotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *AdminAddUserRoleV4UnprocessableEntity:
		return nil, nil, nil, nil, v, nil, nil
	case *AdminAddUserRoleV4InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AdminListUserRolesV4 admins list user s roles v4

  This endpoint requires ADMIN:NAMESPACE:{namespace}:ROLE:USER:* [READ] permission.

List roles assigned to a user

*/
func (a *Client) AdminListUserRolesV4(params *AdminListUserRolesV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminListUserRolesV4OK, *AdminListUserRolesV4Forbidden, *AdminListUserRolesV4NotFound, *AdminListUserRolesV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminListUserRolesV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AdminListUserRolesV4",
		Method:             "GET",
		PathPattern:        "/iam/v4/admin/namespaces/{namespace}/users/{userId}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdminListUserRolesV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminListUserRolesV4OK:
		return v, nil, nil, nil, nil
	case *AdminListUserRolesV4Forbidden:
		return nil, v, nil, nil, nil
	case *AdminListUserRolesV4NotFound:
		return nil, nil, v, nil, nil
	case *AdminListUserRolesV4InternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AdminRemoveUserRoleV4 admins remove user role v4

  This endpoint requires ADMIN:NAMESPACE:{namespace}:ROLE:USER:* [Delete] permission.

Remove a role from user's roles.

*/
func (a *Client) AdminRemoveUserRoleV4(params *AdminRemoveUserRoleV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminRemoveUserRoleV4NoContent, *AdminRemoveUserRoleV4BadRequest, *AdminRemoveUserRoleV4Forbidden, *AdminRemoveUserRoleV4NotFound, *AdminRemoveUserRoleV4UnprocessableEntity, *AdminRemoveUserRoleV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminRemoveUserRoleV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AdminRemoveUserRoleV4",
		Method:             "DELETE",
		PathPattern:        "/iam/v4/admin/namespaces/{namespace}/users/{userId}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdminRemoveUserRoleV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminRemoveUserRoleV4NoContent:
		return v, nil, nil, nil, nil, nil, nil
	case *AdminRemoveUserRoleV4BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *AdminRemoveUserRoleV4Forbidden:
		return nil, nil, v, nil, nil, nil, nil
	case *AdminRemoveUserRoleV4NotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *AdminRemoveUserRoleV4UnprocessableEntity:
		return nil, nil, nil, nil, v, nil, nil
	case *AdminRemoveUserRoleV4InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AdminUpdateUserEmailAddressV4 updates a user email address

  <p>Required permission <pre>'ADMIN:NAMESPACE:{namespace}:USER:{userId} [UPDATE]'</pre></p>

<br><p>This is the endpoint for an admin to update a user email address.
This endpoint need a valid user token from an admin to verify its identity (email) before updating a user.</p>

*/
func (a *Client) AdminUpdateUserEmailAddressV4(params *AdminUpdateUserEmailAddressV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminUpdateUserEmailAddressV4NoContent, *AdminUpdateUserEmailAddressV4BadRequest, *AdminUpdateUserEmailAddressV4Unauthorized, *AdminUpdateUserEmailAddressV4NotFound, *AdminUpdateUserEmailAddressV4Conflict, *AdminUpdateUserEmailAddressV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminUpdateUserEmailAddressV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AdminUpdateUserEmailAddressV4",
		Method:             "PUT",
		PathPattern:        "/iam/v4/admin/namespaces/{namespace}/users/{userId}/email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdminUpdateUserEmailAddressV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminUpdateUserEmailAddressV4NoContent:
		return v, nil, nil, nil, nil, nil, nil
	case *AdminUpdateUserEmailAddressV4BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *AdminUpdateUserEmailAddressV4Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *AdminUpdateUserEmailAddressV4NotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *AdminUpdateUserEmailAddressV4Conflict:
		return nil, nil, nil, nil, v, nil, nil
	case *AdminUpdateUserEmailAddressV4InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AdminUpdateUserRoleV4 admins update user s role v4

  This endpoint requires ADMIN:NAMESPACE:{namespace}:ROLE:USER:* [UPDATE] permission.

User's roles will be replaced with roles from request body.

*/
func (a *Client) AdminUpdateUserRoleV4(params *AdminUpdateUserRoleV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminUpdateUserRoleV4OK, *AdminUpdateUserRoleV4BadRequest, *AdminUpdateUserRoleV4Forbidden, *AdminUpdateUserRoleV4NotFound, *AdminUpdateUserRoleV4UnprocessableEntity, *AdminUpdateUserRoleV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminUpdateUserRoleV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AdminUpdateUserRoleV4",
		Method:             "PUT",
		PathPattern:        "/iam/v4/admin/namespaces/{namespace}/users/{userId}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdminUpdateUserRoleV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminUpdateUserRoleV4OK:
		return v, nil, nil, nil, nil, nil, nil
	case *AdminUpdateUserRoleV4BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *AdminUpdateUserRoleV4Forbidden:
		return nil, nil, v, nil, nil, nil, nil
	case *AdminUpdateUserRoleV4NotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *AdminUpdateUserRoleV4UnprocessableEntity:
		return nil, nil, nil, nil, v, nil, nil
	case *AdminUpdateUserRoleV4InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  AdminUpdateUserV4 updates user

  <p>This endpoint requires ADMIN:NAMESPACE:{namespace}:USER:{userId} [UPDATE] permission</p>
<br><p>This Endpoint support update user based on given data. <b>Single request can update single field or multi fields.</b></p>
<p>Supported field {country, displayName, languageTag, dateOfBirth}</p>
<p>Country use ISO3166-1 alpha-2 two letter, e.g. US.</p>
<p>Date of Birth format : YYYY-MM-DD, e.g. 2019-04-29.</p>
<br><b>Several case of updating email address</b>
<p>action code : 10103 </p>
*/
func (a *Client) AdminUpdateUserV4(params *AdminUpdateUserV4Params, authInfo runtime.ClientAuthInfoWriter) (*AdminUpdateUserV4OK, *AdminUpdateUserV4BadRequest, *AdminUpdateUserV4Unauthorized, *AdminUpdateUserV4Forbidden, *AdminUpdateUserV4NotFound, *AdminUpdateUserV4Conflict, *AdminUpdateUserV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminUpdateUserV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AdminUpdateUserV4",
		Method:             "PUT",
		PathPattern:        "/iam/v4/admin/namespaces/{namespace}/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdminUpdateUserV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminUpdateUserV4OK:
		return v, nil, nil, nil, nil, nil, nil, nil
	case *AdminUpdateUserV4BadRequest:
		return nil, v, nil, nil, nil, nil, nil, nil
	case *AdminUpdateUserV4Unauthorized:
		return nil, nil, v, nil, nil, nil, nil, nil
	case *AdminUpdateUserV4Forbidden:
		return nil, nil, nil, v, nil, nil, nil, nil
	case *AdminUpdateUserV4NotFound:
		return nil, nil, nil, nil, v, nil, nil, nil
	case *AdminUpdateUserV4Conflict:
		return nil, nil, nil, nil, nil, v, nil, nil
	case *AdminUpdateUserV4InternalServerError:
		return nil, nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CreateUserFromInvitationV4 creates user admin from invitation

  This endpoint create user from saved roles when creating invitation and submitted data.
User will be able to login after completing submitting the data through this endpoint.
Available Authentication Types:

	EMAILPASSWD: an authentication type used for new user registration through email.

Country use ISO3166-1 alpha-2 two letter, e.g. US.

Date of Birth format : YYYY-MM-DD, e.g. 2019-04-29.

Required attributes:
- authType: possible value is EMAILPASSWD (see above)
- country: ISO3166-1 alpha-2 two letter, e.g. US.
- dateOfBirth: YYYY-MM-DD, e.g. 1990-01-01. valid values are between 1905-01-01 until current date.
- displayName: case insensitive, alphanumeric with allowed symbols dash (-), comma (,), and dot (.)
- password: 8 to 32 characters, satisfy at least 3 out of 4 conditions(uppercase, lowercase letters, numbers and special characters) and should not have more than 2 equal characters in a row.
- username: case insensitive, alphanumeric with allowed symbols underscore (_) and dot (.)

*/
func (a *Client) CreateUserFromInvitationV4(params *CreateUserFromInvitationV4Params, authInfo runtime.ClientAuthInfoWriter) (*CreateUserFromInvitationV4Created, *CreateUserFromInvitationV4BadRequest, *CreateUserFromInvitationV4NotFound, *CreateUserFromInvitationV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserFromInvitationV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateUserFromInvitationV4",
		Method:             "POST",
		PathPattern:        "/iam/v4/public/namespaces/{namespace}/users/invite/{invitationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateUserFromInvitationV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateUserFromInvitationV4Created:
		return v, nil, nil, nil, nil
	case *CreateUserFromInvitationV4BadRequest:
		return nil, v, nil, nil, nil
	case *CreateUserFromInvitationV4NotFound:
		return nil, nil, v, nil, nil
	case *CreateUserFromInvitationV4InternalServerError:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicCreateUserV4 creates user

  Create a new user with unique email address and username.
		<p>
		<b>Required attributes:</b>
		- authType: possible value is EMAILPASSWD
		- emailAddress
		- username: case insensitive, alphanumeric with allowed symbols underscore (_) and dot (.)
		- password: 8 to 32 characters, satisfy at least 3 out of 4 conditions(uppercase, lowercase letters, numbers and special characters) and should not have more than 2 equal characters in a row.
		- country: ISO3166-1 alpha-2 two letter, e.g. US.
		- dateOfBirth: YYYY-MM-DD, e.g. 1990-01-01. valid values are between 1905-01-01 until current date.
		</p>
		<p>This endpoint support accepting agreements for the created user. Supply the accepted agreements in acceptedPolicies attribute.</p>

*/
func (a *Client) PublicCreateUserV4(params *PublicCreateUserV4Params, authInfo runtime.ClientAuthInfoWriter) (*PublicCreateUserV4Created, *PublicCreateUserV4BadRequest, *PublicCreateUserV4NotFound, *PublicCreateUserV4Conflict, *PublicCreateUserV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicCreateUserV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PublicCreateUserV4",
		Method:             "POST",
		PathPattern:        "/iam/v4/public/namespaces/{namespace}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicCreateUserV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicCreateUserV4Created:
		return v, nil, nil, nil, nil, nil
	case *PublicCreateUserV4BadRequest:
		return nil, v, nil, nil, nil, nil
	case *PublicCreateUserV4NotFound:
		return nil, nil, v, nil, nil, nil
	case *PublicCreateUserV4Conflict:
		return nil, nil, nil, v, nil, nil
	case *PublicCreateUserV4InternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicUpdateUserEmailAddressV4 updates my email address

  <p>The endpoint to update my email address. </p>
<p>It requires a verification code from <pre>/users/me/code/request</pre> with <b>UpdateEmailAddress</b> context.</p>

*/
func (a *Client) PublicUpdateUserEmailAddressV4(params *PublicUpdateUserEmailAddressV4Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdateUserEmailAddressV4NoContent, *PublicUpdateUserEmailAddressV4BadRequest, *PublicUpdateUserEmailAddressV4Unauthorized, *PublicUpdateUserEmailAddressV4NotFound, *PublicUpdateUserEmailAddressV4Conflict, *PublicUpdateUserEmailAddressV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicUpdateUserEmailAddressV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PublicUpdateUserEmailAddressV4",
		Method:             "PUT",
		PathPattern:        "/iam/v4/public/namespaces/{namespace}/users/me/email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicUpdateUserEmailAddressV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicUpdateUserEmailAddressV4NoContent:
		return v, nil, nil, nil, nil, nil, nil
	case *PublicUpdateUserEmailAddressV4BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *PublicUpdateUserEmailAddressV4Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *PublicUpdateUserEmailAddressV4NotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *PublicUpdateUserEmailAddressV4Conflict:
		return nil, nil, nil, nil, v, nil, nil
	case *PublicUpdateUserEmailAddressV4InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicUpdateUserV4 updates user

  <p>Requires valid user access token </p>
<br><p>This Endpoint support update user based on given data. <b>Single request can update single field or multi fields.</b></p>
<p>Supported field {country, displayName, languageTag, dateOfBirth}</p>
<p>Country use ISO3166-1 alpha-2 two letter, e.g. US.</p>
<p>Date of Birth format : YYYY-MM-DD, e.g. 2019-04-29.</p>
<br><b>Several case of updating email address</b>
<p>action code : 10103 </p>
*/
func (a *Client) PublicUpdateUserV4(params *PublicUpdateUserV4Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpdateUserV4OK, *PublicUpdateUserV4BadRequest, *PublicUpdateUserV4Unauthorized, *PublicUpdateUserV4Conflict, *PublicUpdateUserV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicUpdateUserV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PublicUpdateUserV4",
		Method:             "PATCH",
		PathPattern:        "/iam/v4/public/namespaces/{namespace}/users/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicUpdateUserV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicUpdateUserV4OK:
		return v, nil, nil, nil, nil, nil
	case *PublicUpdateUserV4BadRequest:
		return nil, v, nil, nil, nil, nil
	case *PublicUpdateUserV4Unauthorized:
		return nil, nil, v, nil, nil, nil
	case *PublicUpdateUserV4Conflict:
		return nil, nil, nil, v, nil, nil
	case *PublicUpdateUserV4InternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicUpgradeHeadlessAccountV4 upgrades user account to full account

  Require valid user authorization
			Upgrade headless account to full account without verifying email address. Client does not need to provide verification code which sent to email address.
			<br>action code : 10124 </p>
*/
func (a *Client) PublicUpgradeHeadlessAccountV4(params *PublicUpgradeHeadlessAccountV4Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpgradeHeadlessAccountV4OK, *PublicUpgradeHeadlessAccountV4BadRequest, *PublicUpgradeHeadlessAccountV4Unauthorized, *PublicUpgradeHeadlessAccountV4NotFound, *PublicUpgradeHeadlessAccountV4Conflict, *PublicUpgradeHeadlessAccountV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicUpgradeHeadlessAccountV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PublicUpgradeHeadlessAccountV4",
		Method:             "POST",
		PathPattern:        "/iam/v4/public/namespaces/{namespace}/users/me/headless/verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicUpgradeHeadlessAccountV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicUpgradeHeadlessAccountV4OK:
		return v, nil, nil, nil, nil, nil, nil
	case *PublicUpgradeHeadlessAccountV4BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *PublicUpgradeHeadlessAccountV4Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *PublicUpgradeHeadlessAccountV4NotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *PublicUpgradeHeadlessAccountV4Conflict:
		return nil, nil, nil, nil, v, nil, nil
	case *PublicUpgradeHeadlessAccountV4InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicUpgradeHeadlessAccountWithVerificationCodeV4 upgrades headless account and automatically verified the email address if it is succeeded

  Require valid user access token.
        	The endpoint upgrades a headless account by linking the headless account with the email address, username, and password.
			By upgrading the headless account into a full account, the user could use the email address, username, and password for using Justice IAM.
        	<br>
			The endpoint is a shortcut for upgrading a headless account and verifying the email address in one call.
			In order to get a verification code for the endpoint, please check the <a href="#operations-Users-PublicSendVerificationCodeV3">send verification code endpoint</a>.
        	<br>
			This endpoint also have an ability to update user data (if the user data field is specified) right after the upgrade account process is done.
			Supported user data fields:
				<ul>
					<li>displayName</li>
					<li>dateOfBirth : format YYYY-MM-DD, e.g. 2019-04-29</li>
					<li>country : format ISO3166-1 alpha-2 two letter, e.g. US</li>
				</ul>
        	action code : 10124
*/
func (a *Client) PublicUpgradeHeadlessAccountWithVerificationCodeV4(params *PublicUpgradeHeadlessAccountWithVerificationCodeV4Params, authInfo runtime.ClientAuthInfoWriter) (*PublicUpgradeHeadlessAccountWithVerificationCodeV4OK, *PublicUpgradeHeadlessAccountWithVerificationCodeV4BadRequest, *PublicUpgradeHeadlessAccountWithVerificationCodeV4Unauthorized, *PublicUpgradeHeadlessAccountWithVerificationCodeV4Forbidden, *PublicUpgradeHeadlessAccountWithVerificationCodeV4NotFound, *PublicUpgradeHeadlessAccountWithVerificationCodeV4Conflict, *PublicUpgradeHeadlessAccountWithVerificationCodeV4InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicUpgradeHeadlessAccountWithVerificationCodeV4Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PublicUpgradeHeadlessAccountWithVerificationCodeV4",
		Method:             "POST",
		PathPattern:        "/iam/v4/public/namespaces/{namespace}/users/me/headless/code/verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicUpgradeHeadlessAccountWithVerificationCodeV4Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicUpgradeHeadlessAccountWithVerificationCodeV4OK:
		return v, nil, nil, nil, nil, nil, nil, nil
	case *PublicUpgradeHeadlessAccountWithVerificationCodeV4BadRequest:
		return nil, v, nil, nil, nil, nil, nil, nil
	case *PublicUpgradeHeadlessAccountWithVerificationCodeV4Unauthorized:
		return nil, nil, v, nil, nil, nil, nil, nil
	case *PublicUpgradeHeadlessAccountWithVerificationCodeV4Forbidden:
		return nil, nil, nil, v, nil, nil, nil, nil
	case *PublicUpgradeHeadlessAccountWithVerificationCodeV4NotFound:
		return nil, nil, nil, nil, v, nil, nil, nil
	case *PublicUpgradeHeadlessAccountWithVerificationCodeV4Conflict:
		return nil, nil, nil, nil, nil, v, nil, nil
	case *PublicUpgradeHeadlessAccountWithVerificationCodeV4InternalServerError:
		return nil, nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
