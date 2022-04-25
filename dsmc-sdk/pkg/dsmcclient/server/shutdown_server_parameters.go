// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewShutdownServerParams creates a new ShutdownServerParams object
// with the default values initialized.
func NewShutdownServerParams() *ShutdownServerParams {
	var ()
	return &ShutdownServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShutdownServerParamsWithTimeout creates a new ShutdownServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShutdownServerParamsWithTimeout(timeout time.Duration) *ShutdownServerParams {
	var ()
	return &ShutdownServerParams{

		timeout: timeout,
	}
}

// NewShutdownServerParamsWithContext creates a new ShutdownServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewShutdownServerParamsWithContext(ctx context.Context) *ShutdownServerParams {
	var ()
	return &ShutdownServerParams{

		Context: ctx,
	}
}

// NewShutdownServerParamsWithHTTPClient creates a new ShutdownServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShutdownServerParamsWithHTTPClient(client *http.Client) *ShutdownServerParams {
	var ()
	return &ShutdownServerParams{
		HTTPClient: client,
	}
}

/*ShutdownServerParams contains all the parameters to send to the API endpoint
for the shutdown server operation typically these are written to a http.Request
*/
type ShutdownServerParams struct {

	/*Body*/
	Body *dsmcclientmodels.ModelsShutdownServerRequest
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the shutdown server params
func (o *ShutdownServerParams) WithTimeout(timeout time.Duration) *ShutdownServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shutdown server params
func (o *ShutdownServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shutdown server params
func (o *ShutdownServerParams) WithContext(ctx context.Context) *ShutdownServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shutdown server params
func (o *ShutdownServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shutdown server params
func (o *ShutdownServerParams) WithHTTPClient(client *http.Client) *ShutdownServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shutdown server params
func (o *ShutdownServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the shutdown server params
func (o *ShutdownServerParams) WithBody(body *dsmcclientmodels.ModelsShutdownServerRequest) *ShutdownServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the shutdown server params
func (o *ShutdownServerParams) SetBody(body *dsmcclientmodels.ModelsShutdownServerRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the shutdown server params
func (o *ShutdownServerParams) WithNamespace(namespace string) *ShutdownServerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the shutdown server params
func (o *ShutdownServerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *ShutdownServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
