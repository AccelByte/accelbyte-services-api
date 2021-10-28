// Code generated by go-swagger; DO NOT EDIT.

package public

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
)

// NewListServerParams creates a new ListServerParams object
// with the default values initialized.
func NewListServerParams() *ListServerParams {

	return &ListServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListServerParamsWithTimeout creates a new ListServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListServerParamsWithTimeout(timeout time.Duration) *ListServerParams {

	return &ListServerParams{

		timeout: timeout,
	}
}

// NewListServerParamsWithContext creates a new ListServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewListServerParamsWithContext(ctx context.Context) *ListServerParams {

	return &ListServerParams{

		Context: ctx,
	}
}

// NewListServerParamsWithHTTPClient creates a new ListServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListServerParamsWithHTTPClient(client *http.Client) *ListServerParams {

	return &ListServerParams{
		HTTPClient: client,
	}
}

/*ListServerParams contains all the parameters to send to the API endpoint
for the list server operation typically these are written to a http.Request
*/
type ListServerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list server params
func (o *ListServerParams) WithTimeout(timeout time.Duration) *ListServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list server params
func (o *ListServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list server params
func (o *ListServerParams) WithContext(ctx context.Context) *ListServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list server params
func (o *ListServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list server params
func (o *ListServerParams) WithHTTPClient(client *http.Client) *ListServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list server params
func (o *ListServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
