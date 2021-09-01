// Code generated by go-swagger; DO NOT EDIT.

package user_profile

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

// NewGetUserZipCodeParams creates a new GetUserZipCodeParams object
// with the default values initialized.
func NewGetUserZipCodeParams() *GetUserZipCodeParams {
	var ()
	return &GetUserZipCodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserZipCodeParamsWithTimeout creates a new GetUserZipCodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserZipCodeParamsWithTimeout(timeout time.Duration) *GetUserZipCodeParams {
	var ()
	return &GetUserZipCodeParams{

		timeout: timeout,
	}
}

// NewGetUserZipCodeParamsWithContext creates a new GetUserZipCodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserZipCodeParamsWithContext(ctx context.Context) *GetUserZipCodeParams {
	var ()
	return &GetUserZipCodeParams{

		Context: ctx,
	}
}

// NewGetUserZipCodeParamsWithHTTPClient creates a new GetUserZipCodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserZipCodeParamsWithHTTPClient(client *http.Client) *GetUserZipCodeParams {
	var ()
	return &GetUserZipCodeParams{
		HTTPClient: client,
	}
}

/*GetUserZipCodeParams contains all the parameters to send to the API endpoint
for the get user zip code operation typically these are written to a http.Request
*/
type GetUserZipCodeParams struct {

	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user zip code params
func (o *GetUserZipCodeParams) WithTimeout(timeout time.Duration) *GetUserZipCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user zip code params
func (o *GetUserZipCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user zip code params
func (o *GetUserZipCodeParams) WithContext(ctx context.Context) *GetUserZipCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user zip code params
func (o *GetUserZipCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user zip code params
func (o *GetUserZipCodeParams) WithHTTPClient(client *http.Client) *GetUserZipCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user zip code params
func (o *GetUserZipCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get user zip code params
func (o *GetUserZipCodeParams) WithNamespace(namespace string) *GetUserZipCodeParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get user zip code params
func (o *GetUserZipCodeParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserZipCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}