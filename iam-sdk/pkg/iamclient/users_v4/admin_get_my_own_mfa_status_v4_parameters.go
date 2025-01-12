// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package users_v4

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

// NewAdminGetMyOwnMFAStatusV4Params creates a new AdminGetMyOwnMFAStatusV4Params object
// with the default values initialized.
func NewAdminGetMyOwnMFAStatusV4Params() *AdminGetMyOwnMFAStatusV4Params {
	var ()
	return &AdminGetMyOwnMFAStatusV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetMyOwnMFAStatusV4ParamsWithTimeout creates a new AdminGetMyOwnMFAStatusV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetMyOwnMFAStatusV4ParamsWithTimeout(timeout time.Duration) *AdminGetMyOwnMFAStatusV4Params {
	var ()
	return &AdminGetMyOwnMFAStatusV4Params{

		timeout: timeout,
	}
}

// NewAdminGetMyOwnMFAStatusV4ParamsWithContext creates a new AdminGetMyOwnMFAStatusV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetMyOwnMFAStatusV4ParamsWithContext(ctx context.Context) *AdminGetMyOwnMFAStatusV4Params {
	var ()
	return &AdminGetMyOwnMFAStatusV4Params{

		Context: ctx,
	}
}

// NewAdminGetMyOwnMFAStatusV4ParamsWithHTTPClient creates a new AdminGetMyOwnMFAStatusV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetMyOwnMFAStatusV4ParamsWithHTTPClient(client *http.Client) *AdminGetMyOwnMFAStatusV4Params {
	var ()
	return &AdminGetMyOwnMFAStatusV4Params{
		HTTPClient: client,
	}
}

/*AdminGetMyOwnMFAStatusV4Params contains all the parameters to send to the API endpoint
for the admin get my own mfa status v4 operation typically these are written to a http.Request
*/
type AdminGetMyOwnMFAStatusV4Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the admin get my own mfa status v4 params
func (o *AdminGetMyOwnMFAStatusV4Params) WithTimeout(timeout time.Duration) *AdminGetMyOwnMFAStatusV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get my own mfa status v4 params
func (o *AdminGetMyOwnMFAStatusV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get my own mfa status v4 params
func (o *AdminGetMyOwnMFAStatusV4Params) WithContext(ctx context.Context) *AdminGetMyOwnMFAStatusV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get my own mfa status v4 params
func (o *AdminGetMyOwnMFAStatusV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin get my own mfa status v4 params
func (o *AdminGetMyOwnMFAStatusV4Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin get my own mfa status v4 params
func (o *AdminGetMyOwnMFAStatusV4Params) WithHTTPClient(client *http.Client) *AdminGetMyOwnMFAStatusV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get my own mfa status v4 params
func (o *AdminGetMyOwnMFAStatusV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the admin get my own mfa status v4 params
func (o *AdminGetMyOwnMFAStatusV4Params) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *AdminGetMyOwnMFAStatusV4Params) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetMyOwnMFAStatusV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
