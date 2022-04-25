// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package configuration

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

// NewDeleteGroupConfigurationV1Params creates a new DeleteGroupConfigurationV1Params object
// with the default values initialized.
func NewDeleteGroupConfigurationV1Params() *DeleteGroupConfigurationV1Params {
	var ()
	return &DeleteGroupConfigurationV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGroupConfigurationV1ParamsWithTimeout creates a new DeleteGroupConfigurationV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteGroupConfigurationV1ParamsWithTimeout(timeout time.Duration) *DeleteGroupConfigurationV1Params {
	var ()
	return &DeleteGroupConfigurationV1Params{

		timeout: timeout,
	}
}

// NewDeleteGroupConfigurationV1ParamsWithContext creates a new DeleteGroupConfigurationV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteGroupConfigurationV1ParamsWithContext(ctx context.Context) *DeleteGroupConfigurationV1Params {
	var ()
	return &DeleteGroupConfigurationV1Params{

		Context: ctx,
	}
}

// NewDeleteGroupConfigurationV1ParamsWithHTTPClient creates a new DeleteGroupConfigurationV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteGroupConfigurationV1ParamsWithHTTPClient(client *http.Client) *DeleteGroupConfigurationV1Params {
	var ()
	return &DeleteGroupConfigurationV1Params{
		HTTPClient: client,
	}
}

/*DeleteGroupConfigurationV1Params contains all the parameters to send to the API endpoint
for the delete group configuration v1 operation typically these are written to a http.Request
*/
type DeleteGroupConfigurationV1Params struct {

	/*ConfigurationCode
	  Group Configuration Code

	*/
	ConfigurationCode string
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete group configuration v1 params
func (o *DeleteGroupConfigurationV1Params) WithTimeout(timeout time.Duration) *DeleteGroupConfigurationV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete group configuration v1 params
func (o *DeleteGroupConfigurationV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete group configuration v1 params
func (o *DeleteGroupConfigurationV1Params) WithContext(ctx context.Context) *DeleteGroupConfigurationV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete group configuration v1 params
func (o *DeleteGroupConfigurationV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete group configuration v1 params
func (o *DeleteGroupConfigurationV1Params) WithHTTPClient(client *http.Client) *DeleteGroupConfigurationV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete group configuration v1 params
func (o *DeleteGroupConfigurationV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigurationCode adds the configurationCode to the delete group configuration v1 params
func (o *DeleteGroupConfigurationV1Params) WithConfigurationCode(configurationCode string) *DeleteGroupConfigurationV1Params {
	o.SetConfigurationCode(configurationCode)
	return o
}

// SetConfigurationCode adds the configurationCode to the delete group configuration v1 params
func (o *DeleteGroupConfigurationV1Params) SetConfigurationCode(configurationCode string) {
	o.ConfigurationCode = configurationCode
}

// WithNamespace adds the namespace to the delete group configuration v1 params
func (o *DeleteGroupConfigurationV1Params) WithNamespace(namespace string) *DeleteGroupConfigurationV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete group configuration v1 params
func (o *DeleteGroupConfigurationV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGroupConfigurationV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param configurationCode
	if err := r.SetPathParam("configurationCode", o.ConfigurationCode); err != nil {
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
