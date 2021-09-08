// Code generated by go-swagger; DO NOT EDIT.

package config

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

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// NewSaveConfigParams creates a new SaveConfigParams object
// with the default values initialized.
func NewSaveConfigParams() *SaveConfigParams {
	var ()
	return &SaveConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSaveConfigParamsWithTimeout creates a new SaveConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSaveConfigParamsWithTimeout(timeout time.Duration) *SaveConfigParams {
	var ()
	return &SaveConfigParams{

		timeout: timeout,
	}
}

// NewSaveConfigParamsWithContext creates a new SaveConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewSaveConfigParamsWithContext(ctx context.Context) *SaveConfigParams {
	var ()
	return &SaveConfigParams{

		Context: ctx,
	}
}

// NewSaveConfigParamsWithHTTPClient creates a new SaveConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSaveConfigParamsWithHTTPClient(client *http.Client) *SaveConfigParams {
	var ()
	return &SaveConfigParams{
		HTTPClient: client,
	}
}

/*SaveConfigParams contains all the parameters to send to the API endpoint
for the save config operation typically these are written to a http.Request
*/
type SaveConfigParams struct {

	/*Body*/
	Body *dsmcclientmodels.ModelsDSMConfigRecord

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the save config params
func (o *SaveConfigParams) WithTimeout(timeout time.Duration) *SaveConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save config params
func (o *SaveConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save config params
func (o *SaveConfigParams) WithContext(ctx context.Context) *SaveConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save config params
func (o *SaveConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save config params
func (o *SaveConfigParams) WithHTTPClient(client *http.Client) *SaveConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save config params
func (o *SaveConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the save config params
func (o *SaveConfigParams) WithBody(body *dsmcclientmodels.ModelsDSMConfigRecord) *SaveConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the save config params
func (o *SaveConfigParams) SetBody(body *dsmcclientmodels.ModelsDSMConfigRecord) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SaveConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
