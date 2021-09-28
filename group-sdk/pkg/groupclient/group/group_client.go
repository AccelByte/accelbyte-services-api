// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateNewGroupPublicV1(params *CreateNewGroupPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateNewGroupPublicV1Created, *CreateNewGroupPublicV1BadRequest, *CreateNewGroupPublicV1Unauthorized, *CreateNewGroupPublicV1Forbidden, *CreateNewGroupPublicV1Conflict, *CreateNewGroupPublicV1InternalServerError, error)

	DeleteGroupAdminV1(params *DeleteGroupAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteGroupAdminV1NoContent, *DeleteGroupAdminV1BadRequest, *DeleteGroupAdminV1Unauthorized, *DeleteGroupAdminV1Forbidden, *DeleteGroupAdminV1NotFound, *DeleteGroupAdminV1InternalServerError, error)

	DeleteGroupPredefinedRulePublicV1(params *DeleteGroupPredefinedRulePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteGroupPredefinedRulePublicV1NoContent, *DeleteGroupPredefinedRulePublicV1BadRequest, *DeleteGroupPredefinedRulePublicV1Unauthorized, *DeleteGroupPredefinedRulePublicV1Forbidden, *DeleteGroupPredefinedRulePublicV1NotFound, *DeleteGroupPredefinedRulePublicV1InternalServerError, error)

	DeleteGroupPublicV1(params *DeleteGroupPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteGroupPublicV1NoContent, *DeleteGroupPublicV1BadRequest, *DeleteGroupPublicV1Unauthorized, *DeleteGroupPublicV1Forbidden, *DeleteGroupPublicV1NotFound, *DeleteGroupPublicV1InternalServerError, error)

	GetGroupListAdminV1(params *GetGroupListAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetGroupListAdminV1OK, *GetGroupListAdminV1BadRequest, *GetGroupListAdminV1Unauthorized, *GetGroupListAdminV1Forbidden, *GetGroupListAdminV1InternalServerError, error)

	GetGroupListPublicV1(params *GetGroupListPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetGroupListPublicV1OK, *GetGroupListPublicV1BadRequest, *GetGroupListPublicV1Unauthorized, *GetGroupListPublicV1Forbidden, *GetGroupListPublicV1InternalServerError, error)

	GetSingleGroupAdminV1(params *GetSingleGroupAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetSingleGroupAdminV1OK, *GetSingleGroupAdminV1BadRequest, *GetSingleGroupAdminV1Unauthorized, *GetSingleGroupAdminV1Forbidden, *GetSingleGroupAdminV1NotFound, *GetSingleGroupAdminV1InternalServerError, error)

	GetSingleGroupPublicV1(params *GetSingleGroupPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetSingleGroupPublicV1OK, *GetSingleGroupPublicV1BadRequest, *GetSingleGroupPublicV1Unauthorized, *GetSingleGroupPublicV1Forbidden, *GetSingleGroupPublicV1NotFound, *GetSingleGroupPublicV1InternalServerError, error)

	UpdateGroupCustomAttributesPublicV1(params *UpdateGroupCustomAttributesPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateGroupCustomAttributesPublicV1OK, *UpdateGroupCustomAttributesPublicV1BadRequest, *UpdateGroupCustomAttributesPublicV1Unauthorized, *UpdateGroupCustomAttributesPublicV1Forbidden, *UpdateGroupCustomAttributesPublicV1NotFound, *UpdateGroupCustomAttributesPublicV1InternalServerError, error)

	UpdateGroupCustomRulePublicV1(params *UpdateGroupCustomRulePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateGroupCustomRulePublicV1OK, *UpdateGroupCustomRulePublicV1BadRequest, *UpdateGroupCustomRulePublicV1Unauthorized, *UpdateGroupCustomRulePublicV1Forbidden, *UpdateGroupCustomRulePublicV1NotFound, *UpdateGroupCustomRulePublicV1InternalServerError, error)

	UpdateGroupPredefinedRulePublicV1(params *UpdateGroupPredefinedRulePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateGroupPredefinedRulePublicV1OK, *UpdateGroupPredefinedRulePublicV1BadRequest, *UpdateGroupPredefinedRulePublicV1Unauthorized, *UpdateGroupPredefinedRulePublicV1Forbidden, *UpdateGroupPredefinedRulePublicV1NotFound, *UpdateGroupPredefinedRulePublicV1InternalServerError, error)

	UpdateSingleGroupPublicV1(params *UpdateSingleGroupPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateSingleGroupPublicV1OK, *UpdateSingleGroupPublicV1BadRequest, *UpdateSingleGroupPublicV1Unauthorized, *UpdateSingleGroupPublicV1Forbidden, *UpdateSingleGroupPublicV1NotFound, *UpdateSingleGroupPublicV1InternalServerError, error)

	UpdateSingleGroupV1(params *UpdateSingleGroupV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateSingleGroupV1OK, *UpdateSingleGroupV1BadRequest, *UpdateSingleGroupV1Unauthorized, *UpdateSingleGroupV1Forbidden, *UpdateSingleGroupV1NotFound, *UpdateSingleGroupV1InternalServerError, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateNewGroupPublicV1 creates new group

  <p>Required valid user authentication </p>
			<p>This endpoint is used to create new group</p>
			<p>There are some fields that needs to be fulfilled</p>
			<ul>
				<li><b>groupDescription</b>: the description of the group (optional)</li>
				<li><b>groupIcon</b>: group icon URL link (optional)</li>
				<li><b>groupName</b>: name of the group</li>
				<li><b>groupRegion</b>: region of the group</li>
				<li><b>groupRules</b>: rules for specific group. It consists of groupCustomRule that can be used to save custom rule, and groupPredefinedRules that has similar usage with configuration, but this rule only works in specific group</li>
				<li><b>allowedAction</b>: available action in group service. It consist of joinGroup and inviteGroup</li>
				<li><b>ruleAttribute</b>: attribute of the player that needs to be checked</li>
				<li><b>ruleCriteria</b>: criteria of the value. The value will be in enum of EQUAL, MINIMUM, MAXIMUM</li>
				<li><b>ruleValue</b>: value that needs to be checked</li>
				<li><b>customAttributes</b>: additional custom group attributes (optional)</li>
			</ul>
			<p>Action Code: 73304</p>

*/
func (a *Client) CreateNewGroupPublicV1(params *CreateNewGroupPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*CreateNewGroupPublicV1Created, *CreateNewGroupPublicV1BadRequest, *CreateNewGroupPublicV1Unauthorized, *CreateNewGroupPublicV1Forbidden, *CreateNewGroupPublicV1Conflict, *CreateNewGroupPublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNewGroupPublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNewGroupPublicV1",
		Method:             "POST",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateNewGroupPublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CreateNewGroupPublicV1Created:
		return v, nil, nil, nil, nil, nil, nil
	case *CreateNewGroupPublicV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *CreateNewGroupPublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *CreateNewGroupPublicV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *CreateNewGroupPublicV1Conflict:
		return nil, nil, nil, nil, v, nil, nil
	case *CreateNewGroupPublicV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteGroupAdminV1 deletes existing group

  <p>Required Permission: "ADMIN:NAMESPACE:{namespace}:GROUP:{groupId} [DELETE]"</p>
			<p>Delete existing group. It will check whether the groupID is exist before doing the process to delete the group.</p>
			<p>Action Code: 73302</p>

*/
func (a *Client) DeleteGroupAdminV1(params *DeleteGroupAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteGroupAdminV1NoContent, *DeleteGroupAdminV1BadRequest, *DeleteGroupAdminV1Unauthorized, *DeleteGroupAdminV1Forbidden, *DeleteGroupAdminV1NotFound, *DeleteGroupAdminV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGroupAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGroupAdminV1",
		Method:             "DELETE",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteGroupAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteGroupAdminV1NoContent:
		return v, nil, nil, nil, nil, nil, nil
	case *DeleteGroupAdminV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *DeleteGroupAdminV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *DeleteGroupAdminV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *DeleteGroupAdminV1NotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *DeleteGroupAdminV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteGroupPredefinedRulePublicV1 deletes group predefined rule

  <p>Required valid user authentication </p>
			<p>Required Member Role Permission: "GROUP [UPDATE]"</p>
			<p>Delete group predefined rule based on the allowed action. This endpoint will check the group ID of the user based on the access token
			and compare it with the group ID in path parameter. It will also check the member role of the user based on
			the access token</p>
			<p>Action Code: 73309</p>

*/
func (a *Client) DeleteGroupPredefinedRulePublicV1(params *DeleteGroupPredefinedRulePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteGroupPredefinedRulePublicV1NoContent, *DeleteGroupPredefinedRulePublicV1BadRequest, *DeleteGroupPredefinedRulePublicV1Unauthorized, *DeleteGroupPredefinedRulePublicV1Forbidden, *DeleteGroupPredefinedRulePublicV1NotFound, *DeleteGroupPredefinedRulePublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGroupPredefinedRulePublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGroupPredefinedRulePublicV1",
		Method:             "DELETE",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/groups/{groupId}/rules/defined/{allowedAction}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteGroupPredefinedRulePublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteGroupPredefinedRulePublicV1NoContent:
		return v, nil, nil, nil, nil, nil, nil
	case *DeleteGroupPredefinedRulePublicV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *DeleteGroupPredefinedRulePublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *DeleteGroupPredefinedRulePublicV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *DeleteGroupPredefinedRulePublicV1NotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *DeleteGroupPredefinedRulePublicV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteGroupPublicV1 deletes existing group

  <p>Required valid user authentication </p>
			<p>Required Member Role Permission: "GROUP [DELETE]"</p>
			<p>Delete existing group. This endpoint will check the group ID of the user based on the access token
			and compare it with the group ID in path parameter. It will also check the member role of the user based on
			the access token</p>
			<p>Action Code: 73305</p>

*/
func (a *Client) DeleteGroupPublicV1(params *DeleteGroupPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*DeleteGroupPublicV1NoContent, *DeleteGroupPublicV1BadRequest, *DeleteGroupPublicV1Unauthorized, *DeleteGroupPublicV1Forbidden, *DeleteGroupPublicV1NotFound, *DeleteGroupPublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGroupPublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGroupPublicV1",
		Method:             "DELETE",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteGroupPublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DeleteGroupPublicV1NoContent:
		return v, nil, nil, nil, nil, nil, nil
	case *DeleteGroupPublicV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *DeleteGroupPublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *DeleteGroupPublicV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *DeleteGroupPublicV1NotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *DeleteGroupPublicV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetGroupListAdminV1 gets list of groups

  <p>Required Permission: "ADMIN:NAMESPACE:{namespace}:GROUP [READ]" </p>
			<p>Get list of groups. This endpoint will show any types of group</p>
			<p>Action Code: 73301</p>

*/
func (a *Client) GetGroupListAdminV1(params *GetGroupListAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetGroupListAdminV1OK, *GetGroupListAdminV1BadRequest, *GetGroupListAdminV1Unauthorized, *GetGroupListAdminV1Forbidden, *GetGroupListAdminV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupListAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGroupListAdminV1",
		Method:             "GET",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGroupListAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetGroupListAdminV1OK:
		return v, nil, nil, nil, nil, nil
	case *GetGroupListAdminV1BadRequest:
		return nil, v, nil, nil, nil, nil
	case *GetGroupListAdminV1Unauthorized:
		return nil, nil, v, nil, nil, nil
	case *GetGroupListAdminV1Forbidden:
		return nil, nil, nil, v, nil, nil
	case *GetGroupListAdminV1InternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetGroupListPublicV1 gets list of groups

  <p>Required valid user authentication </p>
			<p>Get list of groups. This endpoint will only show OPEN and PUBLIC group type. This endpoint can search based on the group name by filling the "groupName" query parameter</p>
			<p>Action Code: 73303</p>

*/
func (a *Client) GetGroupListPublicV1(params *GetGroupListPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetGroupListPublicV1OK, *GetGroupListPublicV1BadRequest, *GetGroupListPublicV1Unauthorized, *GetGroupListPublicV1Forbidden, *GetGroupListPublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupListPublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGroupListPublicV1",
		Method:             "GET",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGroupListPublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetGroupListPublicV1OK:
		return v, nil, nil, nil, nil, nil
	case *GetGroupListPublicV1BadRequest:
		return nil, v, nil, nil, nil, nil
	case *GetGroupListPublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil
	case *GetGroupListPublicV1Forbidden:
		return nil, nil, nil, v, nil, nil
	case *GetGroupListPublicV1InternalServerError:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetSingleGroupAdminV1 gets single group

  <p>Required Permission: "ADMIN:NAMESPACE:{namespace}:GROUP [READ]"</p>
			<p>Get single group information. This endpoint will show the group information by the groupId</p>
			<p>Action Code: 73306</p>

*/
func (a *Client) GetSingleGroupAdminV1(params *GetSingleGroupAdminV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetSingleGroupAdminV1OK, *GetSingleGroupAdminV1BadRequest, *GetSingleGroupAdminV1Unauthorized, *GetSingleGroupAdminV1Forbidden, *GetSingleGroupAdminV1NotFound, *GetSingleGroupAdminV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSingleGroupAdminV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSingleGroupAdminV1",
		Method:             "GET",
		PathPattern:        "/group/v1/admin/namespaces/{namespace}/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSingleGroupAdminV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetSingleGroupAdminV1OK:
		return v, nil, nil, nil, nil, nil, nil
	case *GetSingleGroupAdminV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *GetSingleGroupAdminV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *GetSingleGroupAdminV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *GetSingleGroupAdminV1NotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *GetSingleGroupAdminV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetSingleGroupPublicV1 gets single group

  <p>Required valid user authentication </p>
			<p>Get single group information. This endpoint will show the group information by the groupId</p>
			<p>Action Code: 73306</p>

*/
func (a *Client) GetSingleGroupPublicV1(params *GetSingleGroupPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*GetSingleGroupPublicV1OK, *GetSingleGroupPublicV1BadRequest, *GetSingleGroupPublicV1Unauthorized, *GetSingleGroupPublicV1Forbidden, *GetSingleGroupPublicV1NotFound, *GetSingleGroupPublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSingleGroupPublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSingleGroupPublicV1",
		Method:             "GET",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSingleGroupPublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *GetSingleGroupPublicV1OK:
		return v, nil, nil, nil, nil, nil, nil
	case *GetSingleGroupPublicV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *GetSingleGroupPublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *GetSingleGroupPublicV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *GetSingleGroupPublicV1NotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *GetSingleGroupPublicV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateGroupCustomAttributesPublicV1 updates group custom attributes

  <p>Requires valid user authentication </p>
			<p>Required Member Role Permission: "GROUP [UPDATE]"</p>
			<p>This endpoint replaces current group custom attributes entirely.
			This endpoint will check the group ID of the user based on the access token and compare it with the group ID in path parameter.
			It will also check the member role of the user based on the access token</p>
			<p>Action Code: 73311</p>

*/
func (a *Client) UpdateGroupCustomAttributesPublicV1(params *UpdateGroupCustomAttributesPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateGroupCustomAttributesPublicV1OK, *UpdateGroupCustomAttributesPublicV1BadRequest, *UpdateGroupCustomAttributesPublicV1Unauthorized, *UpdateGroupCustomAttributesPublicV1Forbidden, *UpdateGroupCustomAttributesPublicV1NotFound, *UpdateGroupCustomAttributesPublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGroupCustomAttributesPublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateGroupCustomAttributesPublicV1",
		Method:             "PUT",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/groups/{groupId}/attributes/custom",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGroupCustomAttributesPublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateGroupCustomAttributesPublicV1OK:
		return v, nil, nil, nil, nil, nil, nil
	case *UpdateGroupCustomAttributesPublicV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *UpdateGroupCustomAttributesPublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *UpdateGroupCustomAttributesPublicV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *UpdateGroupCustomAttributesPublicV1NotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *UpdateGroupCustomAttributesPublicV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateGroupCustomRulePublicV1 updates group custom rule

  <p>Required valid user authentication </p>
			<p>Update group custom rule. This endpoint will check the group ID of the user based on the access token
			and compare it with the group ID in path parameter. It will also check the member role of the user based
			on the access token</p>
			<p>Action Code: 73308</p>

*/
func (a *Client) UpdateGroupCustomRulePublicV1(params *UpdateGroupCustomRulePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateGroupCustomRulePublicV1OK, *UpdateGroupCustomRulePublicV1BadRequest, *UpdateGroupCustomRulePublicV1Unauthorized, *UpdateGroupCustomRulePublicV1Forbidden, *UpdateGroupCustomRulePublicV1NotFound, *UpdateGroupCustomRulePublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGroupCustomRulePublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateGroupCustomRulePublicV1",
		Method:             "PUT",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/groups/{groupId}/rules/custom",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGroupCustomRulePublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateGroupCustomRulePublicV1OK:
		return v, nil, nil, nil, nil, nil, nil
	case *UpdateGroupCustomRulePublicV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *UpdateGroupCustomRulePublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *UpdateGroupCustomRulePublicV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *UpdateGroupCustomRulePublicV1NotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *UpdateGroupCustomRulePublicV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateGroupPredefinedRulePublicV1 updates predefined group rule

  <p>Required valid user authentication </p>
			<p>Required Member Role Permission: "GROUP [UPDATE]"</p>
			<p>Update predefined group rule. This endpoint will check the group ID of the user based on the access token
			and compare it with the group ID in path parameter. It will also check the member role of the user based on
			the access token</p>
			<p>If the rule action is not defined in the group, it wil be added immediately to the predefined group rule</p>
			<p>Action Code: 73310</p>

*/
func (a *Client) UpdateGroupPredefinedRulePublicV1(params *UpdateGroupPredefinedRulePublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateGroupPredefinedRulePublicV1OK, *UpdateGroupPredefinedRulePublicV1BadRequest, *UpdateGroupPredefinedRulePublicV1Unauthorized, *UpdateGroupPredefinedRulePublicV1Forbidden, *UpdateGroupPredefinedRulePublicV1NotFound, *UpdateGroupPredefinedRulePublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGroupPredefinedRulePublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateGroupPredefinedRulePublicV1",
		Method:             "PUT",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/groups/{groupId}/rules/defined/{allowedAction}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGroupPredefinedRulePublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateGroupPredefinedRulePublicV1OK:
		return v, nil, nil, nil, nil, nil, nil
	case *UpdateGroupPredefinedRulePublicV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *UpdateGroupPredefinedRulePublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *UpdateGroupPredefinedRulePublicV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *UpdateGroupPredefinedRulePublicV1NotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *UpdateGroupPredefinedRulePublicV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateSingleGroupPublicV1 updates existing group

  <p>Required valid user authentication </p>
			<p>Required Member Role Permission: "GROUP [UPDATE]"</p>
			<p>Update existing group. This endpoint supports partial update. This endpoint will check the group ID of the user based on the access token and compare it with the group ID in path parameter.
			It will also check the member role of the user based on the access token</p>
			<p>Action Code: 73307</p>

*/
func (a *Client) UpdateSingleGroupPublicV1(params *UpdateSingleGroupPublicV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateSingleGroupPublicV1OK, *UpdateSingleGroupPublicV1BadRequest, *UpdateSingleGroupPublicV1Unauthorized, *UpdateSingleGroupPublicV1Forbidden, *UpdateSingleGroupPublicV1NotFound, *UpdateSingleGroupPublicV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSingleGroupPublicV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSingleGroupPublicV1",
		Method:             "PATCH",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSingleGroupPublicV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateSingleGroupPublicV1OK:
		return v, nil, nil, nil, nil, nil, nil
	case *UpdateSingleGroupPublicV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *UpdateSingleGroupPublicV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *UpdateSingleGroupPublicV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *UpdateSingleGroupPublicV1NotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *UpdateSingleGroupPublicV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  UpdateSingleGroupV1 updates existing group

  <p>Required valid user authentication </p>
			<p>Required Member Role Permission: "GROUP [UPDATE]"</p>
			<p>Update existing group. This endpoint supports partial update. This endpoint will check the group ID of the user based on the access token and compare it with the group ID in path parameter.
			It will also check the member role of the user based on the access token</p>
			<p>Action Code: 73307</p>

*/
func (a *Client) UpdateSingleGroupV1(params *UpdateSingleGroupV1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateSingleGroupV1OK, *UpdateSingleGroupV1BadRequest, *UpdateSingleGroupV1Unauthorized, *UpdateSingleGroupV1Forbidden, *UpdateSingleGroupV1NotFound, *UpdateSingleGroupV1InternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSingleGroupV1Params()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSingleGroupV1",
		Method:             "PUT",
		PathPattern:        "/group/v1/public/namespaces/{namespace}/groups/{groupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSingleGroupV1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *UpdateSingleGroupV1OK:
		return v, nil, nil, nil, nil, nil, nil
	case *UpdateSingleGroupV1BadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *UpdateSingleGroupV1Unauthorized:
		return nil, nil, v, nil, nil, nil, nil
	case *UpdateSingleGroupV1Forbidden:
		return nil, nil, nil, v, nil, nil, nil
	case *UpdateSingleGroupV1NotFound:
		return nil, nil, nil, nil, v, nil, nil
	case *UpdateSingleGroupV1InternalServerError:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
