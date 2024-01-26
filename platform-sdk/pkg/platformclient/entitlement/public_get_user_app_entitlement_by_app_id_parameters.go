// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package entitlement

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPublicGetUserAppEntitlementByAppIDParams creates a new PublicGetUserAppEntitlementByAppIDParams object
// with the default values initialized.
func NewPublicGetUserAppEntitlementByAppIDParams() *PublicGetUserAppEntitlementByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementByAppIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetUserAppEntitlementByAppIDParamsWithTimeout creates a new PublicGetUserAppEntitlementByAppIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetUserAppEntitlementByAppIDParamsWithTimeout(timeout time.Duration) *PublicGetUserAppEntitlementByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementByAppIDParams{

		timeout: timeout,
	}
}

// NewPublicGetUserAppEntitlementByAppIDParamsWithContext creates a new PublicGetUserAppEntitlementByAppIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetUserAppEntitlementByAppIDParamsWithContext(ctx context.Context) *PublicGetUserAppEntitlementByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementByAppIDParams{

		Context: ctx,
	}
}

// NewPublicGetUserAppEntitlementByAppIDParamsWithHTTPClient creates a new PublicGetUserAppEntitlementByAppIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetUserAppEntitlementByAppIDParamsWithHTTPClient(client *http.Client) *PublicGetUserAppEntitlementByAppIDParams {
	var ()
	return &PublicGetUserAppEntitlementByAppIDParams{
		HTTPClient: client,
	}
}

/*PublicGetUserAppEntitlementByAppIDParams contains all the parameters to send to the API endpoint
for the public get user app entitlement by app id operation typically these are written to a http.Request
*/
type PublicGetUserAppEntitlementByAppIDParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string
	/*AppID*/
	AppID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithTimeout(timeout time.Duration) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithContext(ctx context.Context) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithHTTPClient(client *http.Client) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *PublicGetUserAppEntitlementByAppIDParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithNamespace adds the namespace to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithNamespace(namespace string) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithUserID(userID string) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithAppID adds the appID to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) WithAppID(appID string) *PublicGetUserAppEntitlementByAppIDParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the public get user app entitlement by app id params
func (o *PublicGetUserAppEntitlementByAppIDParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetUserAppEntitlementByAppIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	// query param appId
	qrAppID := o.AppID
	qAppID := qrAppID
	if qAppID != "" {
		if err := r.SetQueryParam("appId", qAppID); err != nil {
			return err
		}
	}

	// setting the default header value
	if err := r.SetHeaderParam("User-Agent", utils.UserAgentGen()); err != nil {
		return err
	}

	if err := r.SetHeaderParam("X-Amzn-Trace-Id", utils.AmazonTraceIDGen()); err != nil {
		return err
	}

	if o.XFlightId == nil {
		if err := r.SetHeaderParam("X-Flight-Id", utils.GetDefaultFlightID().Value); err != nil {
			return err
		}
	} else {
		if err := r.SetHeaderParam("X-Flight-Id", *o.XFlightId); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
