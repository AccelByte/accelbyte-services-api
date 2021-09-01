// Code generated by go-swagger; DO NOT EDIT.

package profanity

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

// NewAdminGetProfanityListFiltersV1Params creates a new AdminGetProfanityListFiltersV1Params object
// with the default values initialized.
func NewAdminGetProfanityListFiltersV1Params() *AdminGetProfanityListFiltersV1Params {
	var ()
	return &AdminGetProfanityListFiltersV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetProfanityListFiltersV1ParamsWithTimeout creates a new AdminGetProfanityListFiltersV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetProfanityListFiltersV1ParamsWithTimeout(timeout time.Duration) *AdminGetProfanityListFiltersV1Params {
	var ()
	return &AdminGetProfanityListFiltersV1Params{

		timeout: timeout,
	}
}

// NewAdminGetProfanityListFiltersV1ParamsWithContext creates a new AdminGetProfanityListFiltersV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetProfanityListFiltersV1ParamsWithContext(ctx context.Context) *AdminGetProfanityListFiltersV1Params {
	var ()
	return &AdminGetProfanityListFiltersV1Params{

		Context: ctx,
	}
}

// NewAdminGetProfanityListFiltersV1ParamsWithHTTPClient creates a new AdminGetProfanityListFiltersV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetProfanityListFiltersV1ParamsWithHTTPClient(client *http.Client) *AdminGetProfanityListFiltersV1Params {
	var ()
	return &AdminGetProfanityListFiltersV1Params{
		HTTPClient: client,
	}
}

/*AdminGetProfanityListFiltersV1Params contains all the parameters to send to the API endpoint
for the admin get profanity list filters v1 operation typically these are written to a http.Request
*/
type AdminGetProfanityListFiltersV1Params struct {

	/*List
	  list

	*/
	List string
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin get profanity list filters v1 params
func (o *AdminGetProfanityListFiltersV1Params) WithTimeout(timeout time.Duration) *AdminGetProfanityListFiltersV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get profanity list filters v1 params
func (o *AdminGetProfanityListFiltersV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get profanity list filters v1 params
func (o *AdminGetProfanityListFiltersV1Params) WithContext(ctx context.Context) *AdminGetProfanityListFiltersV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get profanity list filters v1 params
func (o *AdminGetProfanityListFiltersV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get profanity list filters v1 params
func (o *AdminGetProfanityListFiltersV1Params) WithHTTPClient(client *http.Client) *AdminGetProfanityListFiltersV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get profanity list filters v1 params
func (o *AdminGetProfanityListFiltersV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithList adds the list to the admin get profanity list filters v1 params
func (o *AdminGetProfanityListFiltersV1Params) WithList(list string) *AdminGetProfanityListFiltersV1Params {
	o.SetList(list)
	return o
}

// SetList adds the list to the admin get profanity list filters v1 params
func (o *AdminGetProfanityListFiltersV1Params) SetList(list string) {
	o.List = list
}

// WithNamespace adds the namespace to the admin get profanity list filters v1 params
func (o *AdminGetProfanityListFiltersV1Params) WithNamespace(namespace string) *AdminGetProfanityListFiltersV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin get profanity list filters v1 params
func (o *AdminGetProfanityListFiltersV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetProfanityListFiltersV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param list
	if err := r.SetPathParam("list", o.List); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}