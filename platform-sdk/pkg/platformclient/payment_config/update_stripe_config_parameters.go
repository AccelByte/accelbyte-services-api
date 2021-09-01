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

// NewUpdateStripeConfigParams creates a new UpdateStripeConfigParams object
// with the default values initialized.
func NewUpdateStripeConfigParams() *UpdateStripeConfigParams {
	var (
		sandboxDefault  = bool(true)
		validateDefault = bool(false)
	)
	return &UpdateStripeConfigParams{
		Sandbox:  &sandboxDefault,
		Validate: &validateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateStripeConfigParamsWithTimeout creates a new UpdateStripeConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateStripeConfigParamsWithTimeout(timeout time.Duration) *UpdateStripeConfigParams {
	var (
		sandboxDefault  = bool(true)
		validateDefault = bool(false)
	)
	return &UpdateStripeConfigParams{
		Sandbox:  &sandboxDefault,
		Validate: &validateDefault,

		timeout: timeout,
	}
}

// NewUpdateStripeConfigParamsWithContext creates a new UpdateStripeConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateStripeConfigParamsWithContext(ctx context.Context) *UpdateStripeConfigParams {
	var (
		sandboxDefault  = bool(true)
		validateDefault = bool(false)
	)
	return &UpdateStripeConfigParams{
		Sandbox:  &sandboxDefault,
		Validate: &validateDefault,

		Context: ctx,
	}
}

// NewUpdateStripeConfigParamsWithHTTPClient creates a new UpdateStripeConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateStripeConfigParamsWithHTTPClient(client *http.Client) *UpdateStripeConfigParams {
	var (
		sandboxDefault  = bool(true)
		validateDefault = bool(false)
	)
	return &UpdateStripeConfigParams{
		Sandbox:    &sandboxDefault,
		Validate:   &validateDefault,
		HTTPClient: client,
	}
}

/*UpdateStripeConfigParams contains all the parameters to send to the API endpoint
for the update stripe config operation typically these are written to a http.Request
*/
type UpdateStripeConfigParams struct {

	/*Body*/
	Body *platformclientmodels.StripeConfig
	/*ID*/
	ID string
	/*Sandbox*/
	Sandbox *bool
	/*Validate*/
	Validate *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update stripe config params
func (o *UpdateStripeConfigParams) WithTimeout(timeout time.Duration) *UpdateStripeConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update stripe config params
func (o *UpdateStripeConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update stripe config params
func (o *UpdateStripeConfigParams) WithContext(ctx context.Context) *UpdateStripeConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update stripe config params
func (o *UpdateStripeConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update stripe config params
func (o *UpdateStripeConfigParams) WithHTTPClient(client *http.Client) *UpdateStripeConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update stripe config params
func (o *UpdateStripeConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update stripe config params
func (o *UpdateStripeConfigParams) WithBody(body *platformclientmodels.StripeConfig) *UpdateStripeConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update stripe config params
func (o *UpdateStripeConfigParams) SetBody(body *platformclientmodels.StripeConfig) {
	o.Body = body
}

// WithID adds the id to the update stripe config params
func (o *UpdateStripeConfigParams) WithID(id string) *UpdateStripeConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update stripe config params
func (o *UpdateStripeConfigParams) SetID(id string) {
	o.ID = id
}

// WithSandbox adds the sandbox to the update stripe config params
func (o *UpdateStripeConfigParams) WithSandbox(sandbox *bool) *UpdateStripeConfigParams {
	o.SetSandbox(sandbox)
	return o
}

// SetSandbox adds the sandbox to the update stripe config params
func (o *UpdateStripeConfigParams) SetSandbox(sandbox *bool) {
	o.Sandbox = sandbox
}

// WithValidate adds the validate to the update stripe config params
func (o *UpdateStripeConfigParams) WithValidate(validate *bool) *UpdateStripeConfigParams {
	o.SetValidate(validate)
	return o
}

// SetValidate adds the validate to the update stripe config params
func (o *UpdateStripeConfigParams) SetValidate(validate *bool) {
	o.Validate = validate
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateStripeConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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

	if o.Validate != nil {

		// query param validate
		var qrValidate bool
		if o.Validate != nil {
			qrValidate = *o.Validate
		}
		qValidate := swag.FormatBool(qrValidate)
		if qValidate != "" {
			if err := r.SetQueryParam("validate", qValidate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}