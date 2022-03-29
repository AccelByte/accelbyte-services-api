// Code generated by go-swagger; DO NOT EDIT.

package users_v4

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

// NewAdminEnableMyAuthenticatorV4Params creates a new AdminEnableMyAuthenticatorV4Params object
// with the default values initialized.
func NewAdminEnableMyAuthenticatorV4Params() *AdminEnableMyAuthenticatorV4Params {
	var ()
	return &AdminEnableMyAuthenticatorV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminEnableMyAuthenticatorV4ParamsWithTimeout creates a new AdminEnableMyAuthenticatorV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminEnableMyAuthenticatorV4ParamsWithTimeout(timeout time.Duration) *AdminEnableMyAuthenticatorV4Params {
	var ()
	return &AdminEnableMyAuthenticatorV4Params{

		timeout: timeout,
	}
}

// NewAdminEnableMyAuthenticatorV4ParamsWithContext creates a new AdminEnableMyAuthenticatorV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminEnableMyAuthenticatorV4ParamsWithContext(ctx context.Context) *AdminEnableMyAuthenticatorV4Params {
	var ()
	return &AdminEnableMyAuthenticatorV4Params{

		Context: ctx,
	}
}

// NewAdminEnableMyAuthenticatorV4ParamsWithHTTPClient creates a new AdminEnableMyAuthenticatorV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminEnableMyAuthenticatorV4ParamsWithHTTPClient(client *http.Client) *AdminEnableMyAuthenticatorV4Params {
	var ()
	return &AdminEnableMyAuthenticatorV4Params{
		HTTPClient: client,
	}
}

/*AdminEnableMyAuthenticatorV4Params contains all the parameters to send to the API endpoint
for the admin enable my authenticator v4 operation typically these are written to a http.Request
*/
type AdminEnableMyAuthenticatorV4Params struct {

	/*Code
	  code

	*/
	Code *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin enable my authenticator v4 params
func (o *AdminEnableMyAuthenticatorV4Params) WithTimeout(timeout time.Duration) *AdminEnableMyAuthenticatorV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin enable my authenticator v4 params
func (o *AdminEnableMyAuthenticatorV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin enable my authenticator v4 params
func (o *AdminEnableMyAuthenticatorV4Params) WithContext(ctx context.Context) *AdminEnableMyAuthenticatorV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin enable my authenticator v4 params
func (o *AdminEnableMyAuthenticatorV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin enable my authenticator v4 params
func (o *AdminEnableMyAuthenticatorV4Params) WithHTTPClient(client *http.Client) *AdminEnableMyAuthenticatorV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin enable my authenticator v4 params
func (o *AdminEnableMyAuthenticatorV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCode adds the code to the admin enable my authenticator v4 params
func (o *AdminEnableMyAuthenticatorV4Params) WithCode(code *string) *AdminEnableMyAuthenticatorV4Params {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the admin enable my authenticator v4 params
func (o *AdminEnableMyAuthenticatorV4Params) SetCode(code *string) {
	o.Code = code
}

// WriteToRequest writes these params to a swagger request
func (o *AdminEnableMyAuthenticatorV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Code != nil {

		// form param code
		var frCode string
		if o.Code != nil {
			frCode = *o.Code
		}
		fCode := frCode
		if fCode != "" {
			if err := r.SetFormParam("code", fCode); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
