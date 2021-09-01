// Code generated by go-swagger; DO NOT EDIT.

package admin_game_record

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

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
)

// NewPostAdminGameRecordHandlerV1Params creates a new PostAdminGameRecordHandlerV1Params object
// with the default values initialized.
func NewPostAdminGameRecordHandlerV1Params() *PostAdminGameRecordHandlerV1Params {
	var ()
	return &PostAdminGameRecordHandlerV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAdminGameRecordHandlerV1ParamsWithTimeout creates a new PostAdminGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAdminGameRecordHandlerV1ParamsWithTimeout(timeout time.Duration) *PostAdminGameRecordHandlerV1Params {
	var ()
	return &PostAdminGameRecordHandlerV1Params{

		timeout: timeout,
	}
}

// NewPostAdminGameRecordHandlerV1ParamsWithContext creates a new PostAdminGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewPostAdminGameRecordHandlerV1ParamsWithContext(ctx context.Context) *PostAdminGameRecordHandlerV1Params {
	var ()
	return &PostAdminGameRecordHandlerV1Params{

		Context: ctx,
	}
}

// NewPostAdminGameRecordHandlerV1ParamsWithHTTPClient creates a new PostAdminGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAdminGameRecordHandlerV1ParamsWithHTTPClient(client *http.Client) *PostAdminGameRecordHandlerV1Params {
	var ()
	return &PostAdminGameRecordHandlerV1Params{
		HTTPClient: client,
	}
}

/*PostAdminGameRecordHandlerV1Params contains all the parameters to send to the API endpoint
for the post admin game record handler v1 operation typically these are written to a http.Request
*/
type PostAdminGameRecordHandlerV1Params struct {

	/*Body*/
	Body cloudsaveclientmodels.ModelsGameRecordRequest
	/*Key
	  key of record

	*/
	Key string
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) WithTimeout(timeout time.Duration) *PostAdminGameRecordHandlerV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) WithContext(ctx context.Context) *PostAdminGameRecordHandlerV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) WithHTTPClient(client *http.Client) *PostAdminGameRecordHandlerV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) WithBody(body cloudsaveclientmodels.ModelsGameRecordRequest) *PostAdminGameRecordHandlerV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) SetBody(body cloudsaveclientmodels.ModelsGameRecordRequest) {
	o.Body = body
}

// WithKey adds the key to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) WithKey(key string) *PostAdminGameRecordHandlerV1Params {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) SetKey(key string) {
	o.Key = key
}

// WithNamespace adds the namespace to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) WithNamespace(namespace string) *PostAdminGameRecordHandlerV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the post admin game record handler v1 params
func (o *PostAdminGameRecordHandlerV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *PostAdminGameRecordHandlerV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
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