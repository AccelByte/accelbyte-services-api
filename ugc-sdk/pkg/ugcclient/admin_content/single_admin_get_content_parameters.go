// Code generated by go-swagger; DO NOT EDIT.

package admin_content

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
)

// NewSingleAdminGetContentParams creates a new SingleAdminGetContentParams object
// with the default values initialized.
func NewSingleAdminGetContentParams() *SingleAdminGetContentParams {
	var (
		limitDefault  = int64(1000)
		offsetDefault = int64(0)
	)
	return &SingleAdminGetContentParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSingleAdminGetContentParamsWithTimeout creates a new SingleAdminGetContentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSingleAdminGetContentParamsWithTimeout(timeout time.Duration) *SingleAdminGetContentParams {
	var (
		limitDefault  = int64(1000)
		offsetDefault = int64(0)
	)
	return &SingleAdminGetContentParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewSingleAdminGetContentParamsWithContext creates a new SingleAdminGetContentParams object
// with the default values initialized, and the ability to set a context for a request
func NewSingleAdminGetContentParamsWithContext(ctx context.Context) *SingleAdminGetContentParams {
	var (
		limitDefault  = int64(1000)
		offsetDefault = int64(0)
	)
	return &SingleAdminGetContentParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewSingleAdminGetContentParamsWithHTTPClient creates a new SingleAdminGetContentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSingleAdminGetContentParamsWithHTTPClient(client *http.Client) *SingleAdminGetContentParams {
	var (
		limitDefault  = int64(1000)
		offsetDefault = int64(0)
	)
	return &SingleAdminGetContentParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*SingleAdminGetContentParams contains all the parameters to send to the API endpoint
for the single admin get content operation typically these are written to a http.Request
*/
type SingleAdminGetContentParams struct {

	/*Limit
	  number of content per page

	*/
	Limit *int64
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Offset
	  the offset number to retrieve

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the single admin get content params
func (o *SingleAdminGetContentParams) WithTimeout(timeout time.Duration) *SingleAdminGetContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the single admin get content params
func (o *SingleAdminGetContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the single admin get content params
func (o *SingleAdminGetContentParams) WithContext(ctx context.Context) *SingleAdminGetContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the single admin get content params
func (o *SingleAdminGetContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the single admin get content params
func (o *SingleAdminGetContentParams) WithHTTPClient(client *http.Client) *SingleAdminGetContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the single admin get content params
func (o *SingleAdminGetContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the single admin get content params
func (o *SingleAdminGetContentParams) WithLimit(limit *int64) *SingleAdminGetContentParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the single admin get content params
func (o *SingleAdminGetContentParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the single admin get content params
func (o *SingleAdminGetContentParams) WithNamespace(namespace string) *SingleAdminGetContentParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the single admin get content params
func (o *SingleAdminGetContentParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the single admin get content params
func (o *SingleAdminGetContentParams) WithOffset(offset *int64) *SingleAdminGetContentParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the single admin get content params
func (o *SingleAdminGetContentParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *SingleAdminGetContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
