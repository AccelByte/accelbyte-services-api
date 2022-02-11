// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

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

// NewGetTwitchIAPConfigParams creates a new GetTwitchIAPConfigParams object
// with the default values initialized.
func NewGetTwitchIAPConfigParams() *GetTwitchIAPConfigParams {
	var ()
	return &GetTwitchIAPConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTwitchIAPConfigParamsWithTimeout creates a new GetTwitchIAPConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTwitchIAPConfigParamsWithTimeout(timeout time.Duration) *GetTwitchIAPConfigParams {
	var ()
	return &GetTwitchIAPConfigParams{

		timeout: timeout,
	}
}

// NewGetTwitchIAPConfigParamsWithContext creates a new GetTwitchIAPConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTwitchIAPConfigParamsWithContext(ctx context.Context) *GetTwitchIAPConfigParams {
	var ()
	return &GetTwitchIAPConfigParams{

		Context: ctx,
	}
}

// NewGetTwitchIAPConfigParamsWithHTTPClient creates a new GetTwitchIAPConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTwitchIAPConfigParamsWithHTTPClient(client *http.Client) *GetTwitchIAPConfigParams {
	var ()
	return &GetTwitchIAPConfigParams{
		HTTPClient: client,
	}
}

/*GetTwitchIAPConfigParams contains all the parameters to send to the API endpoint
for the get twitch i a p config operation typically these are written to a http.Request
*/
type GetTwitchIAPConfigParams struct {

	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get twitch i a p config params
func (o *GetTwitchIAPConfigParams) WithTimeout(timeout time.Duration) *GetTwitchIAPConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get twitch i a p config params
func (o *GetTwitchIAPConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get twitch i a p config params
func (o *GetTwitchIAPConfigParams) WithContext(ctx context.Context) *GetTwitchIAPConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get twitch i a p config params
func (o *GetTwitchIAPConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get twitch i a p config params
func (o *GetTwitchIAPConfigParams) WithHTTPClient(client *http.Client) *GetTwitchIAPConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get twitch i a p config params
func (o *GetTwitchIAPConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get twitch i a p config params
func (o *GetTwitchIAPConfigParams) WithNamespace(namespace string) *GetTwitchIAPConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get twitch i a p config params
func (o *GetTwitchIAPConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetTwitchIAPConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
