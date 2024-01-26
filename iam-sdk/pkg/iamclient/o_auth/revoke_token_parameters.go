// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package o_auth

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

// NewRevokeTokenParams creates a new RevokeTokenParams object
// with the default values initialized.
func NewRevokeTokenParams() *RevokeTokenParams {
	var ()
	return &RevokeTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeTokenParamsWithTimeout creates a new RevokeTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRevokeTokenParamsWithTimeout(timeout time.Duration) *RevokeTokenParams {
	var ()
	return &RevokeTokenParams{

		timeout: timeout,
	}
}

// NewRevokeTokenParamsWithContext creates a new RevokeTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewRevokeTokenParamsWithContext(ctx context.Context) *RevokeTokenParams {
	var ()
	return &RevokeTokenParams{

		Context: ctx,
	}
}

// NewRevokeTokenParamsWithHTTPClient creates a new RevokeTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRevokeTokenParamsWithHTTPClient(client *http.Client) *RevokeTokenParams {
	var ()
	return &RevokeTokenParams{
		HTTPClient: client,
	}
}

/*RevokeTokenParams contains all the parameters to send to the API endpoint
for the revoke token operation typically these are written to a http.Request
*/
type RevokeTokenParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Token
	  Token to be revoked

	*/
	Token string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the revoke token params
func (o *RevokeTokenParams) WithTimeout(timeout time.Duration) *RevokeTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke token params
func (o *RevokeTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke token params
func (o *RevokeTokenParams) WithContext(ctx context.Context) *RevokeTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke token params
func (o *RevokeTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the revoke token params
func (o *RevokeTokenParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the revoke token params
func (o *RevokeTokenParams) WithHTTPClient(client *http.Client) *RevokeTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke token params
func (o *RevokeTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the revoke token params
func (o *RevokeTokenParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *RevokeTokenParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithToken adds the token to the revoke token params
func (o *RevokeTokenParams) WithToken(token string) *RevokeTokenParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the revoke token params
func (o *RevokeTokenParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param token
	frToken := o.Token
	fToken := frToken
	if fToken != "" {
		if err := r.SetFormParam("token", fToken); err != nil {
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
