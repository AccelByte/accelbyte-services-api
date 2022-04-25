// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package payment_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewTestCheckoutConfigParams creates a new TestCheckoutConfigParams object
// with the default values initialized.
func NewTestCheckoutConfigParams() *TestCheckoutConfigParams {
	var (
		sandboxDefault = bool(true)
	)
	return &TestCheckoutConfigParams{
		Sandbox: &sandboxDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTestCheckoutConfigParamsWithTimeout creates a new TestCheckoutConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestCheckoutConfigParamsWithTimeout(timeout time.Duration) *TestCheckoutConfigParams {
	var (
		sandboxDefault = bool(true)
	)
	return &TestCheckoutConfigParams{
		Sandbox: &sandboxDefault,

		timeout: timeout,
	}
}

// NewTestCheckoutConfigParamsWithContext creates a new TestCheckoutConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestCheckoutConfigParamsWithContext(ctx context.Context) *TestCheckoutConfigParams {
	var (
		sandboxDefault = bool(true)
	)
	return &TestCheckoutConfigParams{
		Sandbox: &sandboxDefault,

		Context: ctx,
	}
}

// NewTestCheckoutConfigParamsWithHTTPClient creates a new TestCheckoutConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestCheckoutConfigParamsWithHTTPClient(client *http.Client) *TestCheckoutConfigParams {
	var (
		sandboxDefault = bool(true)
	)
	return &TestCheckoutConfigParams{
		Sandbox:    &sandboxDefault,
		HTTPClient: client,
	}
}

/*TestCheckoutConfigParams contains all the parameters to send to the API endpoint
for the test checkout config operation typically these are written to a http.Request
*/
type TestCheckoutConfigParams struct {

	/*Body*/
	Body *platformclientmodels.CheckoutConfig
	/*Sandbox*/
	Sandbox *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test checkout config params
func (o *TestCheckoutConfigParams) WithTimeout(timeout time.Duration) *TestCheckoutConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test checkout config params
func (o *TestCheckoutConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test checkout config params
func (o *TestCheckoutConfigParams) WithContext(ctx context.Context) *TestCheckoutConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test checkout config params
func (o *TestCheckoutConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test checkout config params
func (o *TestCheckoutConfigParams) WithHTTPClient(client *http.Client) *TestCheckoutConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test checkout config params
func (o *TestCheckoutConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test checkout config params
func (o *TestCheckoutConfigParams) WithBody(body *platformclientmodels.CheckoutConfig) *TestCheckoutConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test checkout config params
func (o *TestCheckoutConfigParams) SetBody(body *platformclientmodels.CheckoutConfig) {
	o.Body = body
}

// WithSandbox adds the sandbox to the test checkout config params
func (o *TestCheckoutConfigParams) WithSandbox(sandbox *bool) *TestCheckoutConfigParams {
	o.SetSandbox(sandbox)
	return o
}

// SetSandbox adds the sandbox to the test checkout config params
func (o *TestCheckoutConfigParams) SetSandbox(sandbox *bool) {
	o.Sandbox = sandbox
}

// WriteToRequest writes these params to a swagger request
func (o *TestCheckoutConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Sandbox != nil {

		// query param sandbox
		var qrSandbox bool
		if o.Sandbox != nil {
			qrSandbox = *o.Sandbox
		}
		qSandbox := swag.FormatBool(qrSandbox)
		if qSandbox != "" {
			if err := r.SetQueryParam("sandbox", qSandbox); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
