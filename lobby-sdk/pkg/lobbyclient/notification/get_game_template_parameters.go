// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package notification

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

// NewGetGameTemplateParams creates a new GetGameTemplateParams object
// with the default values initialized.
func NewGetGameTemplateParams() *GetGameTemplateParams {
	var ()
	return &GetGameTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGameTemplateParamsWithTimeout creates a new GetGameTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGameTemplateParamsWithTimeout(timeout time.Duration) *GetGameTemplateParams {
	var ()
	return &GetGameTemplateParams{

		timeout: timeout,
	}
}

// NewGetGameTemplateParamsWithContext creates a new GetGameTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGameTemplateParamsWithContext(ctx context.Context) *GetGameTemplateParams {
	var ()
	return &GetGameTemplateParams{

		Context: ctx,
	}
}

// NewGetGameTemplateParamsWithHTTPClient creates a new GetGameTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGameTemplateParamsWithHTTPClient(client *http.Client) *GetGameTemplateParams {
	var ()
	return &GetGameTemplateParams{
		HTTPClient: client,
	}
}

/*GetGameTemplateParams contains all the parameters to send to the API endpoint
for the get game template operation typically these are written to a http.Request
*/
type GetGameTemplateParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get game template params
func (o *GetGameTemplateParams) WithTimeout(timeout time.Duration) *GetGameTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get game template params
func (o *GetGameTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get game template params
func (o *GetGameTemplateParams) WithContext(ctx context.Context) *GetGameTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get game template params
func (o *GetGameTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get game template params
func (o *GetGameTemplateParams) WithHTTPClient(client *http.Client) *GetGameTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get game template params
func (o *GetGameTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get game template params
func (o *GetGameTemplateParams) WithNamespace(namespace string) *GetGameTemplateParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get game template params
func (o *GetGameTemplateParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetGameTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
