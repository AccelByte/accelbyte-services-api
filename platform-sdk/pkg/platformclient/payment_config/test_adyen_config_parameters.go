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

// NewTestAdyenConfigParams creates a new TestAdyenConfigParams object
// with the default values initialized.
func NewTestAdyenConfigParams() *TestAdyenConfigParams {
	var (
		sandboxDefault = bool(true)
	)
	return &TestAdyenConfigParams{
		Sandbox: &sandboxDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTestAdyenConfigParamsWithTimeout creates a new TestAdyenConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestAdyenConfigParamsWithTimeout(timeout time.Duration) *TestAdyenConfigParams {
	var (
		sandboxDefault = bool(true)
	)
	return &TestAdyenConfigParams{
		Sandbox: &sandboxDefault,

		timeout: timeout,
	}
}

// NewTestAdyenConfigParamsWithContext creates a new TestAdyenConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestAdyenConfigParamsWithContext(ctx context.Context) *TestAdyenConfigParams {
	var (
		sandboxDefault = bool(true)
	)
	return &TestAdyenConfigParams{
		Sandbox: &sandboxDefault,

		Context: ctx,
	}
}

// NewTestAdyenConfigParamsWithHTTPClient creates a new TestAdyenConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestAdyenConfigParamsWithHTTPClient(client *http.Client) *TestAdyenConfigParams {
	var (
		sandboxDefault = bool(true)
	)
	return &TestAdyenConfigParams{
		Sandbox:    &sandboxDefault,
		HTTPClient: client,
	}
}

/*TestAdyenConfigParams contains all the parameters to send to the API endpoint
for the test adyen config operation typically these are written to a http.Request
*/
type TestAdyenConfigParams struct {

	/*Body*/
	Body *platformclientmodels.AdyenConfig
	/*Sandbox*/
	Sandbox *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test adyen config params
func (o *TestAdyenConfigParams) WithTimeout(timeout time.Duration) *TestAdyenConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test adyen config params
func (o *TestAdyenConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test adyen config params
func (o *TestAdyenConfigParams) WithContext(ctx context.Context) *TestAdyenConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test adyen config params
func (o *TestAdyenConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test adyen config params
func (o *TestAdyenConfigParams) WithHTTPClient(client *http.Client) *TestAdyenConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test adyen config params
func (o *TestAdyenConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test adyen config params
func (o *TestAdyenConfigParams) WithBody(body *platformclientmodels.AdyenConfig) *TestAdyenConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test adyen config params
func (o *TestAdyenConfigParams) SetBody(body *platformclientmodels.AdyenConfig) {
	o.Body = body
}

// WithSandbox adds the sandbox to the test adyen config params
func (o *TestAdyenConfigParams) WithSandbox(sandbox *bool) *TestAdyenConfigParams {
	o.SetSandbox(sandbox)
	return o
}

// SetSandbox adds the sandbox to the test adyen config params
func (o *TestAdyenConfigParams) SetSandbox(sandbox *bool) {
	o.Sandbox = sandbox
}

// WriteToRequest writes these params to a swagger request
func (o *TestAdyenConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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