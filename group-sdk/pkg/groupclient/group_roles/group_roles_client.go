// Code generated by go-swagger; DO NOT EDIT.

package group_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group roles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group roles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMemberRoleAdminV1(params *CreateMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateMemberRoleAdminV1Created, *CreateMemberRoleAdminV1BadRequest, *CreateMemberRoleAdminV1Unauthorized, *CreateMemberRoleAdminV1Forbidden, *CreateMemberRoleAdminV1InternalServerError, error)
	CreateMemberRoleAdminV1Short(params *CreateMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateMemberRoleAdminV1Created, error)
	DeleteMemberRoleAdminV1(params *DeleteMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteMemberRoleAdminV1NoContent, *DeleteMemberRoleAdminV1BadRequest, *DeleteMemberRoleAdminV1Unauthorized, *DeleteMemberRoleAdminV1Forbidden, *DeleteMemberRoleAdminV1NotFound, *DeleteMemberRoleAdminV1InternalServerError, error)
	DeleteMemberRoleAdminV1Short(params *DeleteMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteMemberRoleAdminV1NoContent, error)
	DeleteMemberRolePublicV1(params *DeleteMemberRolePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteMemberRolePublicV1OK, *DeleteMemberRolePublicV1BadRequest, *DeleteMemberRolePublicV1Unauthorized, *DeleteMemberRolePublicV1Forbidden, *DeleteMemberRolePublicV1NotFound, *DeleteMemberRolePublicV1UnprocessableEntity, *DeleteMemberRolePublicV1InternalServerError, error)
	DeleteMemberRolePublicV1Short(params *DeleteMemberRolePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteMemberRolePublicV1OK, error)
	GetMemberRolesListAdminV1(params *GetMemberRolesListAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetMemberRolesListAdminV1OK, *GetMemberRolesListAdminV1BadRequest, *GetMemberRolesListAdminV1Unauthorized, *GetMemberRolesListAdminV1Forbidden, *GetMemberRolesListAdminV1InternalServerError, error)
	GetMemberRolesListAdminV1Short(params *GetMemberRolesListAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetMemberRolesListAdminV1OK, error)
	GetMemberRolesListPublicV1(params *GetMemberRolesListPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetMemberRolesListPublicV1OK, *GetMemberRolesListPublicV1BadRequest, *GetMemberRolesListPublicV1Unauthorized, *GetMemberRolesListPublicV1Forbidden, *GetMemberRolesListPublicV1InternalServerError, error)
	GetMemberRolesListPublicV1Short(params *GetMemberRolesListPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetMemberRolesListPublicV1OK, error)
	GetSingleMemberRoleAdminV1(params *GetSingleMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetSingleMemberRoleAdminV1OK, *GetSingleMemberRoleAdminV1BadRequest, *GetSingleMemberRoleAdminV1Unauthorized, *GetSingleMemberRoleAdminV1Forbidden, *GetSingleMemberRoleAdminV1NotFound, *GetSingleMemberRoleAdminV1InternalServerError, error)
	GetSingleMemberRoleAdminV1Short(params *GetSingleMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetSingleMemberRoleAdminV1OK, error)
	UpdateMemberRoleAdminV1(params *UpdateMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRoleAdminV1OK, *UpdateMemberRoleAdminV1BadRequest, *UpdateMemberRoleAdminV1Unauthorized, *UpdateMemberRoleAdminV1Forbidden, *UpdateMemberRoleAdminV1NotFound, *UpdateMemberRoleAdminV1InternalServerError, error)
	UpdateMemberRoleAdminV1Short(params *UpdateMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRoleAdminV1OK, error)
	UpdateMemberRolePermissionAdminV1(params *UpdateMemberRolePermissionAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRolePermissionAdminV1OK, *UpdateMemberRolePermissionAdminV1BadRequest, *UpdateMemberRolePermissionAdminV1Unauthorized, *UpdateMemberRolePermissionAdminV1Forbidden, *UpdateMemberRolePermissionAdminV1NotFound, *UpdateMemberRolePermissionAdminV1InternalServerError, error)
	UpdateMemberRolePermissionAdminV1Short(params *UpdateMemberRolePermissionAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRolePermissionAdminV1OK, error)
	UpdateMemberRolePublicV1(params *UpdateMemberRolePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRolePublicV1OK, *UpdateMemberRolePublicV1BadRequest, *UpdateMemberRolePublicV1Unauthorized, *UpdateMemberRolePublicV1Forbidden, *UpdateMemberRolePublicV1NotFound, *UpdateMemberRolePublicV1InternalServerError, error)
	UpdateMemberRolePublicV1Short(params *UpdateMemberRolePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRolePublicV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateMemberRoleAdminV1 creates new member role

  &lt;p&gt;Required permission ADMIN:NAMESPACE:{namespace}:GROUP:ROLE [CREATE]&lt;/p&gt;
			&lt;p&gt;This endpoint is used to create new member role&lt;/p&gt;
			&lt;p&gt;Action Code: 73202&lt;/p&gt;

*/
func (a *Client) CreateMemberRoleAdminV1(params *CreateMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateMemberRoleAdminV1Created, *CreateMemberRoleAdminV1BadRequest, *CreateMemberRoleAdminV1Unauthorized, *CreateMemberRoleAdminV1Forbidden, *CreateMemberRoleAdminV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMemberRoleAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMemberRoleAdminV1",
		Method:             "POST",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMemberRoleAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateMemberRoleAdminV1Created:
		return v, nil, nil, nil, nil, nil

	case *CreateMemberRoleAdminV1BadRequest:
		return nil, v, nil, nil, nil, nil

	case *CreateMemberRoleAdminV1Unauthorized:
		return nil, nil, v, nil, nil, nil

	case *CreateMemberRoleAdminV1Forbidden:
		return nil, nil, nil, v, nil, nil

	case *CreateMemberRoleAdminV1InternalServerError:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) CreateMemberRoleAdminV1Short(params *CreateMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateMemberRoleAdminV1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMemberRoleAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMemberRoleAdminV1",
		Method:             "POST",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMemberRoleAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *CreateMemberRoleAdminV1Created:
		return v, nil
	case *CreateMemberRoleAdminV1BadRequest:
		return nil, v
	case *CreateMemberRoleAdminV1Unauthorized:
		return nil, v
	case *CreateMemberRoleAdminV1Forbidden:
		return nil, v
	case *CreateMemberRoleAdminV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteMemberRoleAdminV1 deletes member role

  &lt;p&gt;Required permission ADMIN:NAMESPACE:{namespace}:GROUP:ROLE [DELETE]&lt;/p&gt;
			&lt;p&gt;This endpoint is used to delete member role. Any member role can&#39;t be deleted if the specific role is applied to the configuration (admin and member role)&lt;/p&gt;
			&lt;p&gt;Action Code: 73207&lt;/p&gt;

*/
func (a *Client) DeleteMemberRoleAdminV1(params *DeleteMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteMemberRoleAdminV1NoContent, *DeleteMemberRoleAdminV1BadRequest, *DeleteMemberRoleAdminV1Unauthorized, *DeleteMemberRoleAdminV1Forbidden, *DeleteMemberRoleAdminV1NotFound, *DeleteMemberRoleAdminV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMemberRoleAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMemberRoleAdminV1",
		Method:             "DELETE",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles/{memberRoleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMemberRoleAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteMemberRoleAdminV1NoContent:
		return v, nil, nil, nil, nil, nil, nil

	case *DeleteMemberRoleAdminV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *DeleteMemberRoleAdminV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *DeleteMemberRoleAdminV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *DeleteMemberRoleAdminV1NotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *DeleteMemberRoleAdminV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) DeleteMemberRoleAdminV1Short(params *DeleteMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteMemberRoleAdminV1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMemberRoleAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMemberRoleAdminV1",
		Method:             "DELETE",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles/{memberRoleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMemberRoleAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DeleteMemberRoleAdminV1NoContent:
		return v, nil
	case *DeleteMemberRoleAdminV1BadRequest:
		return nil, v
	case *DeleteMemberRoleAdminV1Unauthorized:
		return nil, v
	case *DeleteMemberRoleAdminV1Forbidden:
		return nil, v
	case *DeleteMemberRoleAdminV1NotFound:
		return nil, v
	case *DeleteMemberRoleAdminV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteMemberRolePublicV1 removes role from group member

  Required Member Role Permission: &#34;GROUP:ROLE [UPDATE]&#34;&lt;/p&gt;
			&lt;p&gt;This endpoint is used to remove role from group member&lt;/p&gt;
			&lt;p&gt;Action Code: 73204&lt;/p&gt;
*/
func (a *Client) DeleteMemberRolePublicV1(params *DeleteMemberRolePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteMemberRolePublicV1OK, *DeleteMemberRolePublicV1BadRequest, *DeleteMemberRolePublicV1Unauthorized, *DeleteMemberRolePublicV1Forbidden, *DeleteMemberRolePublicV1NotFound, *DeleteMemberRolePublicV1UnprocessableEntity, *DeleteMemberRolePublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMemberRolePublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMemberRolePublicV1",
		Method:             "DELETE",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/roles/{memberRoleId}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMemberRolePublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteMemberRolePublicV1OK:
		return v, nil, nil, nil, nil, nil, nil, nil

	case *DeleteMemberRolePublicV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil, nil

	case *DeleteMemberRolePublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil, nil

	case *DeleteMemberRolePublicV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil, nil

	case *DeleteMemberRolePublicV1NotFound:
		return nil, nil, nil, nil, v, nil, nil, nil

	case *DeleteMemberRolePublicV1UnprocessableEntity:
		return nil, nil, nil, nil, nil, v, nil, nil

	case *DeleteMemberRolePublicV1InternalServerError:
		return nil, nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) DeleteMemberRolePublicV1Short(params *DeleteMemberRolePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteMemberRolePublicV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMemberRolePublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMemberRolePublicV1",
		Method:             "DELETE",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/roles/{memberRoleId}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMemberRolePublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DeleteMemberRolePublicV1OK:
		return v, nil
	case *DeleteMemberRolePublicV1BadRequest:
		return nil, v
	case *DeleteMemberRolePublicV1Unauthorized:
		return nil, v
	case *DeleteMemberRolePublicV1Forbidden:
		return nil, v
	case *DeleteMemberRolePublicV1NotFound:
		return nil, v
	case *DeleteMemberRolePublicV1UnprocessableEntity:
		return nil, v
	case *DeleteMemberRolePublicV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetMemberRolesListAdminV1 gets all list of member roles

  &lt;p&gt;Required permission ADMIN:NAMESPACE:{namespace}:GROUP:ROLE [READ]&lt;/p&gt;
			&lt;p&gt;This endpoint is used to get list of member roles&lt;/p&gt;
			&lt;p&gt;Action Code: 73201&lt;/p&gt;

*/
func (a *Client) GetMemberRolesListAdminV1(params *GetMemberRolesListAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetMemberRolesListAdminV1OK, *GetMemberRolesListAdminV1BadRequest, *GetMemberRolesListAdminV1Unauthorized, *GetMemberRolesListAdminV1Forbidden, *GetMemberRolesListAdminV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMemberRolesListAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMemberRolesListAdminV1",
		Method:             "GET",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMemberRolesListAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetMemberRolesListAdminV1OK:
		return v, nil, nil, nil, nil, nil

	case *GetMemberRolesListAdminV1BadRequest:
		return nil, v, nil, nil, nil, nil

	case *GetMemberRolesListAdminV1Unauthorized:
		return nil, nil, v, nil, nil, nil

	case *GetMemberRolesListAdminV1Forbidden:
		return nil, nil, nil, v, nil, nil

	case *GetMemberRolesListAdminV1InternalServerError:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) GetMemberRolesListAdminV1Short(params *GetMemberRolesListAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetMemberRolesListAdminV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMemberRolesListAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMemberRolesListAdminV1",
		Method:             "GET",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMemberRolesListAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetMemberRolesListAdminV1OK:
		return v, nil
	case *GetMemberRolesListAdminV1BadRequest:
		return nil, v
	case *GetMemberRolesListAdminV1Unauthorized:
		return nil, v
	case *GetMemberRolesListAdminV1Forbidden:
		return nil, v
	case *GetMemberRolesListAdminV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetMemberRolesListPublicV1 gets all list of member roles

  &lt;p&gt;Required Member Role Permission: &#34;GROUP:ROLE [READ]&#34;&lt;/p&gt;
			&lt;p&gt;This endpoint is used to get list of member roles&lt;/p&gt;
			&lt;p&gt;Action Code: 73201&lt;/p&gt;

*/
func (a *Client) GetMemberRolesListPublicV1(params *GetMemberRolesListPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetMemberRolesListPublicV1OK, *GetMemberRolesListPublicV1BadRequest, *GetMemberRolesListPublicV1Unauthorized, *GetMemberRolesListPublicV1Forbidden, *GetMemberRolesListPublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMemberRolesListPublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMemberRolesListPublicV1",
		Method:             "GET",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMemberRolesListPublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetMemberRolesListPublicV1OK:
		return v, nil, nil, nil, nil, nil

	case *GetMemberRolesListPublicV1BadRequest:
		return nil, v, nil, nil, nil, nil

	case *GetMemberRolesListPublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil

	case *GetMemberRolesListPublicV1Forbidden:
		return nil, nil, nil, v, nil, nil

	case *GetMemberRolesListPublicV1InternalServerError:
		return nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) GetMemberRolesListPublicV1Short(params *GetMemberRolesListPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetMemberRolesListPublicV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMemberRolesListPublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMemberRolesListPublicV1",
		Method:             "GET",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMemberRolesListPublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetMemberRolesListPublicV1OK:
		return v, nil
	case *GetMemberRolesListPublicV1BadRequest:
		return nil, v
	case *GetMemberRolesListPublicV1Unauthorized:
		return nil, v
	case *GetMemberRolesListPublicV1Forbidden:
		return nil, v
	case *GetMemberRolesListPublicV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetSingleMemberRoleAdminV1 gets member role

  &lt;p&gt;Required permission ADMIN:NAMESPACE:{namespace}:GROUP:ROLE [READ]&lt;/p&gt;
			&lt;p&gt;This endpoint is used to get member role based on the role ID&lt;/p&gt;
			&lt;p&gt;Action Code: 73203&lt;/p&gt;

*/
func (a *Client) GetSingleMemberRoleAdminV1(params *GetSingleMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetSingleMemberRoleAdminV1OK, *GetSingleMemberRoleAdminV1BadRequest, *GetSingleMemberRoleAdminV1Unauthorized, *GetSingleMemberRoleAdminV1Forbidden, *GetSingleMemberRoleAdminV1NotFound, *GetSingleMemberRoleAdminV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSingleMemberRoleAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSingleMemberRoleAdminV1",
		Method:             "GET",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles/{memberRoleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSingleMemberRoleAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetSingleMemberRoleAdminV1OK:
		return v, nil, nil, nil, nil, nil, nil

	case *GetSingleMemberRoleAdminV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *GetSingleMemberRoleAdminV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *GetSingleMemberRoleAdminV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *GetSingleMemberRoleAdminV1NotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *GetSingleMemberRoleAdminV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) GetSingleMemberRoleAdminV1Short(params *GetSingleMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetSingleMemberRoleAdminV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSingleMemberRoleAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSingleMemberRoleAdminV1",
		Method:             "GET",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles/{memberRoleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSingleMemberRoleAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetSingleMemberRoleAdminV1OK:
		return v, nil
	case *GetSingleMemberRoleAdminV1BadRequest:
		return nil, v
	case *GetSingleMemberRoleAdminV1Unauthorized:
		return nil, v
	case *GetSingleMemberRoleAdminV1Forbidden:
		return nil, v
	case *GetSingleMemberRoleAdminV1NotFound:
		return nil, v
	case *GetSingleMemberRoleAdminV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateMemberRoleAdminV1 updates member role

  &lt;p&gt;Required permission ADMIN:NAMESPACE:{namespace}:GROUP:ROLE [UPDATE]&lt;/p&gt;
			&lt;p&gt;This endpoint is used to get member role&lt;/p&gt;
			&lt;p&gt;Action Code: 73204&lt;/p&gt;

*/
func (a *Client) UpdateMemberRoleAdminV1(params *UpdateMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRoleAdminV1OK, *UpdateMemberRoleAdminV1BadRequest, *UpdateMemberRoleAdminV1Unauthorized, *UpdateMemberRoleAdminV1Forbidden, *UpdateMemberRoleAdminV1NotFound, *UpdateMemberRoleAdminV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMemberRoleAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMemberRoleAdminV1",
		Method:             "PATCH",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles/{memberRoleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMemberRoleAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateMemberRoleAdminV1OK:
		return v, nil, nil, nil, nil, nil, nil

	case *UpdateMemberRoleAdminV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *UpdateMemberRoleAdminV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *UpdateMemberRoleAdminV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *UpdateMemberRoleAdminV1NotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *UpdateMemberRoleAdminV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) UpdateMemberRoleAdminV1Short(params *UpdateMemberRoleAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRoleAdminV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMemberRoleAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMemberRoleAdminV1",
		Method:             "PATCH",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles/{memberRoleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMemberRoleAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *UpdateMemberRoleAdminV1OK:
		return v, nil
	case *UpdateMemberRoleAdminV1BadRequest:
		return nil, v
	case *UpdateMemberRoleAdminV1Unauthorized:
		return nil, v
	case *UpdateMemberRoleAdminV1Forbidden:
		return nil, v
	case *UpdateMemberRoleAdminV1NotFound:
		return nil, v
	case *UpdateMemberRoleAdminV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateMemberRolePermissionAdminV1 updates member role permission

  &lt;p&gt;Required permission ADMIN:NAMESPACE:{namespace}:GROUP:ROLE [UPDATE]&lt;/p&gt;
			&lt;p&gt;This endpoint is used to update member role permission. It will replace the existing permission based on the request from this endpoint&lt;/p&gt;
			&lt;p&gt;Action Code: 73205&lt;/p&gt;

*/
func (a *Client) UpdateMemberRolePermissionAdminV1(params *UpdateMemberRolePermissionAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRolePermissionAdminV1OK, *UpdateMemberRolePermissionAdminV1BadRequest, *UpdateMemberRolePermissionAdminV1Unauthorized, *UpdateMemberRolePermissionAdminV1Forbidden, *UpdateMemberRolePermissionAdminV1NotFound, *UpdateMemberRolePermissionAdminV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMemberRolePermissionAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMemberRolePermissionAdminV1",
		Method:             "PUT",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles/{memberRoleId}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMemberRolePermissionAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateMemberRolePermissionAdminV1OK:
		return v, nil, nil, nil, nil, nil, nil

	case *UpdateMemberRolePermissionAdminV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *UpdateMemberRolePermissionAdminV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *UpdateMemberRolePermissionAdminV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *UpdateMemberRolePermissionAdminV1NotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *UpdateMemberRolePermissionAdminV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) UpdateMemberRolePermissionAdminV1Short(params *UpdateMemberRolePermissionAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRolePermissionAdminV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMemberRolePermissionAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMemberRolePermissionAdminV1",
		Method:             "PUT",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/roles/{memberRoleId}/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMemberRolePermissionAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *UpdateMemberRolePermissionAdminV1OK:
		return v, nil
	case *UpdateMemberRolePermissionAdminV1BadRequest:
		return nil, v
	case *UpdateMemberRolePermissionAdminV1Unauthorized:
		return nil, v
	case *UpdateMemberRolePermissionAdminV1Forbidden:
		return nil, v
	case *UpdateMemberRolePermissionAdminV1NotFound:
		return nil, v
	case *UpdateMemberRolePermissionAdminV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateMemberRolePublicV1 assigns role to group member

  Required Member Role Permission: &#34;GROUP:ROLE [UPDATE]&lt;/p&gt;
			&lt;p&gt;This endpoint is used to assign role to group member&lt;/p&gt;
			&lt;p&gt;Action Code: 73204&lt;/p&gt;
*/
func (a *Client) UpdateMemberRolePublicV1(params *UpdateMemberRolePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRolePublicV1OK, *UpdateMemberRolePublicV1BadRequest, *UpdateMemberRolePublicV1Unauthorized, *UpdateMemberRolePublicV1Forbidden, *UpdateMemberRolePublicV1NotFound, *UpdateMemberRolePublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMemberRolePublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMemberRolePublicV1",
		Method:             "POST",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/roles/{memberRoleId}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMemberRolePublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateMemberRolePublicV1OK:
		return v, nil, nil, nil, nil, nil, nil

	case *UpdateMemberRolePublicV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil

	case *UpdateMemberRolePublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil

	case *UpdateMemberRolePublicV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil

	case *UpdateMemberRolePublicV1NotFound:
		return nil, nil, nil, nil, v, nil, nil

	case *UpdateMemberRolePublicV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) UpdateMemberRolePublicV1Short(params *UpdateMemberRolePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateMemberRolePublicV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMemberRolePublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMemberRolePublicV1",
		Method:             "POST",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/roles/{memberRoleId}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMemberRolePublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *UpdateMemberRolePublicV1OK:
		return v, nil
	case *UpdateMemberRolePublicV1BadRequest:
		return nil, v
	case *UpdateMemberRolePublicV1Unauthorized:
		return nil, v
	case *UpdateMemberRolePublicV1Forbidden:
		return nil, v
	case *UpdateMemberRolePublicV1NotFound:
		return nil, v
	case *UpdateMemberRolePublicV1InternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
