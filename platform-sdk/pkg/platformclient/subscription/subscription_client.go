// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new subscription API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscription API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CancelSubscription(params *CancelSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*CancelSubscriptionOK, *CancelSubscriptionNotFound, *CancelSubscriptionConflict, error)

	CheckUserSubscriptionSubscribableByItemID(params *CheckUserSubscriptionSubscribableByItemIDParams, authInfo runtime.ClientAuthInfoWriter) (*CheckUserSubscriptionSubscribableByItemIDOK, error)

	DeleteUserSubscription(params *DeleteUserSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserSubscriptionNoContent, error)

	GetUserSubscription(params *GetUserSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserSubscriptionOK, *GetUserSubscriptionNotFound, error)

	GetUserSubscriptionActivities(params *GetUserSubscriptionActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserSubscriptionActivitiesOK, error)

	GetUserSubscriptionBillingHistories(params *GetUserSubscriptionBillingHistoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserSubscriptionBillingHistoriesOK, error)

	GrantDaysToSubscription(params *GrantDaysToSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*GrantDaysToSubscriptionOK, *GrantDaysToSubscriptionNotFound, error)

	PlatformSubscribeSubscription(params *PlatformSubscribeSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PlatformSubscribeSubscriptionOK, *PlatformSubscribeSubscriptionCreated, *PlatformSubscribeSubscriptionBadRequest, *PlatformSubscribeSubscriptionNotFound, *PlatformSubscribeSubscriptionUnprocessableEntity, error)

	ProcessUserSubscriptionNotification(params *ProcessUserSubscriptionNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*ProcessUserSubscriptionNotificationNoContent, *ProcessUserSubscriptionNotificationBadRequest, error)

	PublicCancelSubscription(params *PublicCancelSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PublicCancelSubscriptionOK, *PublicCancelSubscriptionNotFound, *PublicCancelSubscriptionConflict, error)

	PublicChangeSubscriptionBillingAccount(params *PublicChangeSubscriptionBillingAccountParams, authInfo runtime.ClientAuthInfoWriter) (*PublicChangeSubscriptionBillingAccountOK, *PublicChangeSubscriptionBillingAccountBadRequest, *PublicChangeSubscriptionBillingAccountNotFound, *PublicChangeSubscriptionBillingAccountConflict, error)

	PublicCheckUserSubscriptionSubscribableByItemID(params *PublicCheckUserSubscriptionSubscribableByItemIDParams, authInfo runtime.ClientAuthInfoWriter) (*PublicCheckUserSubscriptionSubscribableByItemIDOK, error)

	PublicGetUserSubscription(params *PublicGetUserSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetUserSubscriptionOK, *PublicGetUserSubscriptionNotFound, error)

	PublicGetUserSubscriptionBillingHistories(params *PublicGetUserSubscriptionBillingHistoriesParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetUserSubscriptionBillingHistoriesOK, error)

	PublicQueryUserSubscriptions(params *PublicQueryUserSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*PublicQueryUserSubscriptionsOK, error)

	PublicSubscribeSubscription(params *PublicSubscribeSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PublicSubscribeSubscriptionCreated, *PublicSubscribeSubscriptionBadRequest, *PublicSubscribeSubscriptionForbidden, *PublicSubscribeSubscriptionNotFound, *PublicSubscribeSubscriptionConflict, *PublicSubscribeSubscriptionUnprocessableEntity, error)

	QuerySubscriptions(params *QuerySubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*QuerySubscriptionsOK, error)

	QueryUserSubscriptions(params *QueryUserSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*QueryUserSubscriptionsOK, error)

	RecurringChargeSubscription(params *RecurringChargeSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*RecurringChargeSubscriptionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CancelSubscription cancels a subscription

  Cancel a subscription, only ACTIVE subscription can be cancelled. <b>Ensure successfully cancel, recommend at least 1 day before current period ends, otherwise it may be charging or charged.</b><br>Set immediate true, the subscription will be terminated immediately, otherwise till the end of current billing cycle.<br>Set force true, will ignore the error if subscription is during recurring charging.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=4 (UPDATE)</li><li><i>Returns</i>: cancelled subscription</li></ul>
*/
func (a *Client) CancelSubscription(params *CancelSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*CancelSubscriptionOK, *CancelSubscriptionNotFound, *CancelSubscriptionConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelSubscriptionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cancelSubscription",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CancelSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CancelSubscriptionOK:
		return v, nil, nil, nil
	case *CancelSubscriptionNotFound:
		return nil, v, nil, nil
	case *CancelSubscriptionConflict:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  CheckUserSubscriptionSubscribableByItemID checks user subscription subscribable

  Check user subscription subscribable by itemId, ACTIVE USER subscription can't do subscribe again.<p>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=2 (READ)</li><li><i>Returns</i>: subscribable info</li></ul>
*/
func (a *Client) CheckUserSubscriptionSubscribableByItemID(params *CheckUserSubscriptionSubscribableByItemIDParams, authInfo runtime.ClientAuthInfoWriter) (*CheckUserSubscriptionSubscribableByItemIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckUserSubscriptionSubscribableByItemIDParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkUserSubscriptionSubscribableByItemId",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/subscriptions/subscribable/byItemId",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckUserSubscriptionSubscribableByItemIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *CheckUserSubscriptionSubscribableByItemIDOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DeleteUserSubscription deletes user subscription

  <b>[TEST FACILITY ONLY]</b> Delete user subscription.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=8 (DELETE)</li></ul>
*/
func (a *Client) DeleteUserSubscription(params *DeleteUserSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserSubscriptionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserSubscriptionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUserSubscription",
		Method:             "DELETE",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteUserSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DeleteUserSubscriptionNoContent:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetUserSubscription gets user subscription

  Get user subscription.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=2 (READ)</li><li><i>Returns</i>: subscription</li></ul>
*/
func (a *Client) GetUserSubscription(params *GetUserSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserSubscriptionOK, *GetUserSubscriptionNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserSubscriptionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserSubscription",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetUserSubscriptionOK:
		return v, nil, nil
	case *GetUserSubscriptionNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetUserSubscriptionActivities gets user subscription activity

  Get user subscription activity.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=2 (READ)</li><li><i>Returns</i>: paginated subscription activity</li></ul>
*/
func (a *Client) GetUserSubscriptionActivities(params *GetUserSubscriptionActivitiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserSubscriptionActivitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserSubscriptionActivitiesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserSubscriptionActivities",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/subscriptions/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserSubscriptionActivitiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetUserSubscriptionActivitiesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetUserSubscriptionBillingHistories gets user subscription billing histories

  Get user subscription billing histories.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=2 (READ)</li><li><i>Returns</i>: paginated subscription billing history</li></ul>
*/
func (a *Client) GetUserSubscriptionBillingHistories(params *GetUserSubscriptionBillingHistoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserSubscriptionBillingHistoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserSubscriptionBillingHistoriesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserSubscriptionBillingHistories",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserSubscriptionBillingHistoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetUserSubscriptionBillingHistoriesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GrantDaysToSubscription grants days to a subscription

  Grant days to a subscription, if grantDays is positive, it will add free days and push the next billing date by the amount of day.<br>if the grantDays is negative or zero, it only apply to active/cancelled subscription, remove days will decrease current period end, and move the next billing date closer.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=4 (UPDATE)</li><li><i>Returns</i>: updated subscription</li></ul>
*/
func (a *Client) GrantDaysToSubscription(params *GrantDaysToSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*GrantDaysToSubscriptionOK, *GrantDaysToSubscriptionNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGrantDaysToSubscriptionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "grantDaysToSubscription",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/grant",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GrantDaysToSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GrantDaysToSubscriptionOK:
		return v, nil, nil
	case *GrantDaysToSubscriptionNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PlatformSubscribeSubscription frees subscribe by platform

  Free subscribe by platform, can used by other justice service to redeem/reward the subscription.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=1 (CREATE)</li><li><i>Returns</i>: result subscription</li></ul>
*/
func (a *Client) PlatformSubscribeSubscription(params *PlatformSubscribeSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PlatformSubscribeSubscriptionOK, *PlatformSubscribeSubscriptionCreated, *PlatformSubscribeSubscriptionBadRequest, *PlatformSubscribeSubscriptionNotFound, *PlatformSubscribeSubscriptionUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPlatformSubscribeSubscriptionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "platformSubscribeSubscription",
		Method:             "POST",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/subscriptions/platformSubscribe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PlatformSubscribeSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PlatformSubscribeSubscriptionOK:
		return v, nil, nil, nil, nil, nil
	case *PlatformSubscribeSubscriptionCreated:
		return nil, v, nil, nil, nil, nil
	case *PlatformSubscribeSubscriptionBadRequest:
		return nil, nil, v, nil, nil, nil
	case *PlatformSubscribeSubscriptionNotFound:
		return nil, nil, nil, v, nil, nil
	case *PlatformSubscribeSubscriptionUnprocessableEntity:
		return nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  ProcessUserSubscriptionNotification webs hook for payment notification

  <b>[SERVICE COMMUNICATION ONLY]</b> This API is used as a web hook for payment notification from justice payment service.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=4 (UPDATE)</li><li><i>Returns</i>: Process result</li></ul>
*/
func (a *Client) ProcessUserSubscriptionNotification(params *ProcessUserSubscriptionNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*ProcessUserSubscriptionNotificationNoContent, *ProcessUserSubscriptionNotificationBadRequest, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProcessUserSubscriptionNotificationParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "processUserSubscriptionNotification",
		Method:             "POST",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProcessUserSubscriptionNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *ProcessUserSubscriptionNotificationNoContent:
		return v, nil, nil
	case *ProcessUserSubscriptionNotificationBadRequest:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicCancelSubscription cancels a subscription

  Cancel a subscription, only ACTIVE subscription can be cancelled. <b>Ensure successfully cancel, recommend at least 1 day before current period ends, otherwise it may be charging or charged.</b><br>Set immediate true, the subscription will be terminated immediately, otherwise till the end of current billing cycle.<br>Other detail info: <ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=4 (UPDATE)</li><li><i>Returns</i>: cancelled subscription</li></ul>
*/
func (a *Client) PublicCancelSubscription(params *PublicCancelSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PublicCancelSubscriptionOK, *PublicCancelSubscriptionNotFound, *PublicCancelSubscriptionConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicCancelSubscriptionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicCancelSubscription",
		Method:             "PUT",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicCancelSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicCancelSubscriptionOK:
		return v, nil, nil, nil
	case *PublicCancelSubscriptionNotFound:
		return nil, v, nil, nil
	case *PublicCancelSubscriptionConflict:
		return nil, nil, v, nil
	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicChangeSubscriptionBillingAccount requests to change a subscription billing account

  Request to change a subscription billing account, this will guide user to payment station. The actual change will happen at the 0 payment notification successfully handled.<br>Only ACTIVE USER subscription with real currency billing account can be changed.<br>Other detail info: <ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=4 (UPDATE)</li><li><i>Returns</i>: updated subscription</li></ul>
*/
func (a *Client) PublicChangeSubscriptionBillingAccount(params *PublicChangeSubscriptionBillingAccountParams, authInfo runtime.ClientAuthInfoWriter) (*PublicChangeSubscriptionBillingAccountOK, *PublicChangeSubscriptionBillingAccountBadRequest, *PublicChangeSubscriptionBillingAccountNotFound, *PublicChangeSubscriptionBillingAccountConflict, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicChangeSubscriptionBillingAccountParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicChangeSubscriptionBillingAccount",
		Method:             "PUT",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/billingAccount",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicChangeSubscriptionBillingAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicChangeSubscriptionBillingAccountOK:
		return v, nil, nil, nil, nil
	case *PublicChangeSubscriptionBillingAccountBadRequest:
		return nil, v, nil, nil, nil
	case *PublicChangeSubscriptionBillingAccountNotFound:
		return nil, nil, v, nil, nil
	case *PublicChangeSubscriptionBillingAccountConflict:
		return nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicCheckUserSubscriptionSubscribableByItemID checks user subscription subscribable

  Check user subscription subscribable by itemId, ACTIVE USER subscription can't do subscribe again.<p>Other detail info: <ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=2 (READ)</li><li><i>Returns</i>: subscribable info</li></ul>
*/
func (a *Client) PublicCheckUserSubscriptionSubscribableByItemID(params *PublicCheckUserSubscriptionSubscribableByItemIDParams, authInfo runtime.ClientAuthInfoWriter) (*PublicCheckUserSubscriptionSubscribableByItemIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicCheckUserSubscriptionSubscribableByItemIDParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicCheckUserSubscriptionSubscribableByItemId",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/subscriptions/subscribable/byItemId",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicCheckUserSubscriptionSubscribableByItemIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicCheckUserSubscriptionSubscribableByItemIDOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetUserSubscription gets user subscription

  Get user subscription.<br>Other detail info: <ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=2 (READ)</li><li><i>Returns</i>: subscription</li></ul>
*/
func (a *Client) PublicGetUserSubscription(params *PublicGetUserSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetUserSubscriptionOK, *PublicGetUserSubscriptionNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetUserSubscriptionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetUserSubscription",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicGetUserSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *PublicGetUserSubscriptionOK:
		return v, nil, nil
	case *PublicGetUserSubscriptionNotFound:
		return nil, v, nil
	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicGetUserSubscriptionBillingHistories gets user subscription billing histories

  Get user subscription billing histories.<br>Other detail info: <ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=2 (READ)</li><li><i>Returns</i>: paginated subscription history</li></ul>
*/
func (a *Client) PublicGetUserSubscriptionBillingHistories(params *PublicGetUserSubscriptionBillingHistoriesParams, authInfo runtime.ClientAuthInfoWriter) (*PublicGetUserSubscriptionBillingHistoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicGetUserSubscriptionBillingHistoriesParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicGetUserSubscriptionBillingHistories",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicGetUserSubscriptionBillingHistoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicGetUserSubscriptionBillingHistoriesOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicQueryUserSubscriptions queries user subscriptions

  Query user subscriptions.<br>Other detail info: <ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=2 (READ)</li><li><i>Returns</i>: paginated subscription</li></ul>
*/
func (a *Client) PublicQueryUserSubscriptions(params *PublicQueryUserSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*PublicQueryUserSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicQueryUserSubscriptionsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicQueryUserSubscriptions",
		Method:             "GET",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicQueryUserSubscriptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *PublicQueryUserSubscriptionsOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  PublicSubscribeSubscription subscribes a subscription

  Subscribe a subscription. Support both real and virtual payment. Need go through payment flow using the paymentOrderNo if paymentFlowRequired true.<br><b>ACTIVE USER subscription can't do subscribe again.</b><br><b>The next billing date will be X(default 4) hours before the current period ends if correctly subscribed.</b><br>User with permission SANDBOX will create sandbox subscription that not real paid.<br>Other detail info: <ul><li><i>Required permission</i>: resource="NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=1 (CREATE)</li><li><i>Optional permission(user with this permission will create sandbox subscription)</i>: resource="SANDBOX", action=1 (CREATE)</li><li>It will be forbidden while the user is banned: ORDER_INITIATE or ORDER_AND_PAYMENT</li><li><i>Returns</i>: created subscription</li></ul>
*/
func (a *Client) PublicSubscribeSubscription(params *PublicSubscribeSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PublicSubscribeSubscriptionCreated, *PublicSubscribeSubscriptionBadRequest, *PublicSubscribeSubscriptionForbidden, *PublicSubscribeSubscriptionNotFound, *PublicSubscribeSubscriptionConflict, *PublicSubscribeSubscriptionUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicSubscribeSubscriptionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "publicSubscribeSubscription",
		Method:             "POST",
		PathPattern:        "/public/namespaces/{namespace}/users/{userId}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicSubscribeSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *PublicSubscribeSubscriptionCreated:
		return v, nil, nil, nil, nil, nil, nil
	case *PublicSubscribeSubscriptionBadRequest:
		return nil, v, nil, nil, nil, nil, nil
	case *PublicSubscribeSubscriptionForbidden:
		return nil, nil, v, nil, nil, nil, nil
	case *PublicSubscribeSubscriptionNotFound:
		return nil, nil, nil, v, nil, nil, nil
	case *PublicSubscribeSubscriptionConflict:
		return nil, nil, nil, nil, v, nil, nil
	case *PublicSubscribeSubscriptionUnprocessableEntity:
		return nil, nil, nil, nil, nil, v, nil
	default:
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  QuerySubscriptions queries subscriptions

  Query subscriptions.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:SUBSCRIPTION", action=2 (READ)</li><li><i>Returns</i>: paginated subscriptions</li></ul>
*/
func (a *Client) QuerySubscriptions(params *QuerySubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*QuerySubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySubscriptionsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "querySubscriptions",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuerySubscriptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *QuerySubscriptionsOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  QueryUserSubscriptions queries user subscriptions

  Query user subscriptions.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:USER:{userId}:SUBSCRIPTION", action=2 (READ)</li><li><i>Returns</i>: paginated subscription</li></ul>
*/
func (a *Client) QueryUserSubscriptions(params *QueryUserSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*QueryUserSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryUserSubscriptionsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryUserSubscriptions",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryUserSubscriptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *QueryUserSubscriptionsOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  RecurringChargeSubscription recurrings charge subscription

  <b>[TEST FACILITY ONLY]</b> Recurring charge subscription, it will trigger recurring charge if the USER subscription status is ACTIVE, nextBillingDate is before now and no fail recurring charge within X(default 12) hours.<br>Other detail info: <ul><li><i>Required permission</i>: resource="ADMIN:NAMESPACE:{namespace}:SUBSCRIPTION", action=4 (UPDATE)</li><li><i>Returns</i>: recurring charge result</li></ul>
*/
func (a *Client) RecurringChargeSubscription(params *RecurringChargeSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*RecurringChargeSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRecurringChargeSubscriptionParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "recurringChargeSubscription",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/subscriptions/{subscriptionId}/recurring",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RecurringChargeSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *RecurringChargeSubscriptionOK:
		return v, nil
	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
