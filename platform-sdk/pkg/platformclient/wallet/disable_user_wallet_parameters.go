// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package wallet

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

// NewDisableUserWalletParams creates a new DisableUserWalletParams object
// with the default values initialized.
func NewDisableUserWalletParams() *DisableUserWalletParams {
	var ()
	return &DisableUserWalletParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisableUserWalletParamsWithTimeout creates a new DisableUserWalletParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisableUserWalletParamsWithTimeout(timeout time.Duration) *DisableUserWalletParams {
	var ()
	return &DisableUserWalletParams{

		timeout: timeout,
	}
}

// NewDisableUserWalletParamsWithContext creates a new DisableUserWalletParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisableUserWalletParamsWithContext(ctx context.Context) *DisableUserWalletParams {
	var ()
	return &DisableUserWalletParams{

		Context: ctx,
	}
}

// NewDisableUserWalletParamsWithHTTPClient creates a new DisableUserWalletParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDisableUserWalletParamsWithHTTPClient(client *http.Client) *DisableUserWalletParams {
	var ()
	return &DisableUserWalletParams{
		HTTPClient: client,
	}
}

/*DisableUserWalletParams contains all the parameters to send to the API endpoint
for the disable user wallet operation typically these are written to a http.Request
*/
type DisableUserWalletParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*UserID
	  userId

	*/
	UserID string
	/*WalletID
	  walletId

	*/
	WalletID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the disable user wallet params
func (o *DisableUserWalletParams) WithTimeout(timeout time.Duration) *DisableUserWalletParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable user wallet params
func (o *DisableUserWalletParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable user wallet params
func (o *DisableUserWalletParams) WithContext(ctx context.Context) *DisableUserWalletParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable user wallet params
func (o *DisableUserWalletParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the disable user wallet params
func (o *DisableUserWalletParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the disable user wallet params
func (o *DisableUserWalletParams) WithHTTPClient(client *http.Client) *DisableUserWalletParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable user wallet params
func (o *DisableUserWalletParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the disable user wallet params
func (o *DisableUserWalletParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *DisableUserWalletParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithNamespace adds the namespace to the disable user wallet params
func (o *DisableUserWalletParams) WithNamespace(namespace string) *DisableUserWalletParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the disable user wallet params
func (o *DisableUserWalletParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the disable user wallet params
func (o *DisableUserWalletParams) WithUserID(userID string) *DisableUserWalletParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the disable user wallet params
func (o *DisableUserWalletParams) SetUserID(userID string) {
	o.UserID = userID
}

// WithWalletID adds the walletID to the disable user wallet params
func (o *DisableUserWalletParams) WithWalletID(walletID string) *DisableUserWalletParams {
	o.SetWalletID(walletID)
	return o
}

// SetWalletID adds the walletId to the disable user wallet params
func (o *DisableUserWalletParams) SetWalletID(walletID string) {
	o.WalletID = walletID
}

// WriteToRequest writes these params to a swagger request
func (o *DisableUserWalletParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param walletId
	if err := r.SetPathParam("walletId", o.WalletID); err != nil {
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
