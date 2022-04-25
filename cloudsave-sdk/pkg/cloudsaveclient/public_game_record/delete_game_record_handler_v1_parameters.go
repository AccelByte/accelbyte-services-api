// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package public_game_record

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

// NewDeleteGameRecordHandlerV1Params creates a new DeleteGameRecordHandlerV1Params object
// with the default values initialized.
func NewDeleteGameRecordHandlerV1Params() *DeleteGameRecordHandlerV1Params {
	var ()
	return &DeleteGameRecordHandlerV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGameRecordHandlerV1ParamsWithTimeout creates a new DeleteGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteGameRecordHandlerV1ParamsWithTimeout(timeout time.Duration) *DeleteGameRecordHandlerV1Params {
	var ()
	return &DeleteGameRecordHandlerV1Params{

		timeout: timeout,
	}
}

// NewDeleteGameRecordHandlerV1ParamsWithContext creates a new DeleteGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteGameRecordHandlerV1ParamsWithContext(ctx context.Context) *DeleteGameRecordHandlerV1Params {
	var ()
	return &DeleteGameRecordHandlerV1Params{

		Context: ctx,
	}
}

// NewDeleteGameRecordHandlerV1ParamsWithHTTPClient creates a new DeleteGameRecordHandlerV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteGameRecordHandlerV1ParamsWithHTTPClient(client *http.Client) *DeleteGameRecordHandlerV1Params {
	var ()
	return &DeleteGameRecordHandlerV1Params{
		HTTPClient: client,
	}
}

/*DeleteGameRecordHandlerV1Params contains all the parameters to send to the API endpoint
for the delete game record handler v1 operation typically these are written to a http.Request
*/
type DeleteGameRecordHandlerV1Params struct {

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

// WithTimeout adds the timeout to the delete game record handler v1 params
func (o *DeleteGameRecordHandlerV1Params) WithTimeout(timeout time.Duration) *DeleteGameRecordHandlerV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete game record handler v1 params
func (o *DeleteGameRecordHandlerV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete game record handler v1 params
func (o *DeleteGameRecordHandlerV1Params) WithContext(ctx context.Context) *DeleteGameRecordHandlerV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete game record handler v1 params
func (o *DeleteGameRecordHandlerV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete game record handler v1 params
func (o *DeleteGameRecordHandlerV1Params) WithHTTPClient(client *http.Client) *DeleteGameRecordHandlerV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete game record handler v1 params
func (o *DeleteGameRecordHandlerV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the delete game record handler v1 params
func (o *DeleteGameRecordHandlerV1Params) WithKey(key string) *DeleteGameRecordHandlerV1Params {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete game record handler v1 params
func (o *DeleteGameRecordHandlerV1Params) SetKey(key string) {
	o.Key = key
}

// WithNamespace adds the namespace to the delete game record handler v1 params
func (o *DeleteGameRecordHandlerV1Params) WithNamespace(namespace string) *DeleteGameRecordHandlerV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete game record handler v1 params
func (o *DeleteGameRecordHandlerV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGameRecordHandlerV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
