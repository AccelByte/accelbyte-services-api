// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

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

// NewAdminPostGameRecordHandlerV1Params creates a new AdminPostGameRecordHandlerV1Params object
// with the default values initialized.
func NewAdminPostGameRecordHandlerV1Params() *AdminPostGameRecordHandlerV1Params {
	var ()
	return &AdminPostGameRecordHandlerV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminPostGameRecordHandlerV1ParamsWithTimeout creates a new AdminPostGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminPostGameRecordHandlerV1ParamsWithTimeout(timeout time.Duration) *AdminPostGameRecordHandlerV1Params {
	var ()
	return &AdminPostGameRecordHandlerV1Params{

		timeout: timeout,
	}
}

// NewAdminPostGameRecordHandlerV1ParamsWithContext creates a new AdminPostGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminPostGameRecordHandlerV1ParamsWithContext(ctx context.Context) *AdminPostGameRecordHandlerV1Params {
	var ()
	return &AdminPostGameRecordHandlerV1Params{

		Context: ctx,
	}
}

// NewAdminPostGameRecordHandlerV1ParamsWithHTTPClient creates a new AdminPostGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminPostGameRecordHandlerV1ParamsWithHTTPClient(client *http.Client) *AdminPostGameRecordHandlerV1Params {
	var ()
	return &AdminPostGameRecordHandlerV1Params{
		HTTPClient: client,
	}
}

/*AdminPostGameRecordHandlerV1Params contains all the parameters to send to the API endpoint
for the admin post game record handler v1 operation typically these are written to a http.Request
*/
type AdminPostGameRecordHandlerV1Params struct {

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

// WithTimeout adds the timeout to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) WithTimeout(timeout time.Duration) *AdminPostGameRecordHandlerV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) WithContext(ctx context.Context) *AdminPostGameRecordHandlerV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) WithHTTPClient(client *http.Client) *AdminPostGameRecordHandlerV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) WithBody(body cloudsaveclientmodels.ModelsGameRecordRequest) *AdminPostGameRecordHandlerV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) SetBody(body cloudsaveclientmodels.ModelsGameRecordRequest) {
	o.Body = body
}

// WithKey adds the key to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) WithKey(key string) *AdminPostGameRecordHandlerV1Params {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) SetKey(key string) {
	o.Key = key
}

// WithNamespace adds the namespace to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) WithNamespace(namespace string) *AdminPostGameRecordHandlerV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin post game record handler v1 params
func (o *AdminPostGameRecordHandlerV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminPostGameRecordHandlerV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
