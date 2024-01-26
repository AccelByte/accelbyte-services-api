// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package game_session

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

// NewPublicRevokeGameSessionCodeParams creates a new PublicRevokeGameSessionCodeParams object
// with the default values initialized.
func NewPublicRevokeGameSessionCodeParams() *PublicRevokeGameSessionCodeParams {
	var ()
	return &PublicRevokeGameSessionCodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicRevokeGameSessionCodeParamsWithTimeout creates a new PublicRevokeGameSessionCodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicRevokeGameSessionCodeParamsWithTimeout(timeout time.Duration) *PublicRevokeGameSessionCodeParams {
	var ()
	return &PublicRevokeGameSessionCodeParams{

		timeout: timeout,
	}
}

// NewPublicRevokeGameSessionCodeParamsWithContext creates a new PublicRevokeGameSessionCodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicRevokeGameSessionCodeParamsWithContext(ctx context.Context) *PublicRevokeGameSessionCodeParams {
	var ()
	return &PublicRevokeGameSessionCodeParams{

		Context: ctx,
	}
}

// NewPublicRevokeGameSessionCodeParamsWithHTTPClient creates a new PublicRevokeGameSessionCodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicRevokeGameSessionCodeParamsWithHTTPClient(client *http.Client) *PublicRevokeGameSessionCodeParams {
	var ()
	return &PublicRevokeGameSessionCodeParams{
		HTTPClient: client,
	}
}

/*PublicRevokeGameSessionCodeParams contains all the parameters to send to the API endpoint
for the public revoke game session code operation typically these are written to a http.Request
*/
type PublicRevokeGameSessionCodeParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*SessionID
	  session ID

	*/
	SessionID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) WithTimeout(timeout time.Duration) *PublicRevokeGameSessionCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) WithContext(ctx context.Context) *PublicRevokeGameSessionCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) WithHTTPClient(client *http.Client) *PublicRevokeGameSessionCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *PublicRevokeGameSessionCodeParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithNamespace adds the namespace to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) WithNamespace(namespace string) *PublicRevokeGameSessionCodeParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSessionID adds the sessionID to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) WithSessionID(sessionID string) *PublicRevokeGameSessionCodeParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the public revoke game session code params
func (o *PublicRevokeGameSessionCodeParams) SetSessionID(sessionID string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicRevokeGameSessionCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param sessionId
	if err := r.SetPathParam("sessionId", o.SessionID); err != nil {
		return err
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
