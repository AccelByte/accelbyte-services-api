// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package data_retrieval

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

// NewGetAdminEmailConfigurationParams creates a new GetAdminEmailConfigurationParams object
// with the default values initialized.
func NewGetAdminEmailConfigurationParams() *GetAdminEmailConfigurationParams {
	var ()
	return &GetAdminEmailConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdminEmailConfigurationParamsWithTimeout creates a new GetAdminEmailConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAdminEmailConfigurationParamsWithTimeout(timeout time.Duration) *GetAdminEmailConfigurationParams {
	var ()
	return &GetAdminEmailConfigurationParams{

		timeout: timeout,
	}
}

// NewGetAdminEmailConfigurationParamsWithContext creates a new GetAdminEmailConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAdminEmailConfigurationParamsWithContext(ctx context.Context) *GetAdminEmailConfigurationParams {
	var ()
	return &GetAdminEmailConfigurationParams{

		Context: ctx,
	}
}

// NewGetAdminEmailConfigurationParamsWithHTTPClient creates a new GetAdminEmailConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAdminEmailConfigurationParamsWithHTTPClient(client *http.Client) *GetAdminEmailConfigurationParams {
	var ()
	return &GetAdminEmailConfigurationParams{
		HTTPClient: client,
	}
}

/*GetAdminEmailConfigurationParams contains all the parameters to send to the API endpoint
for the get admin email configuration operation typically these are written to a http.Request
*/
type GetAdminEmailConfigurationParams struct {

	/*Namespace
	  namespace of the user

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get admin email configuration params
func (o *GetAdminEmailConfigurationParams) WithTimeout(timeout time.Duration) *GetAdminEmailConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get admin email configuration params
func (o *GetAdminEmailConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get admin email configuration params
func (o *GetAdminEmailConfigurationParams) WithContext(ctx context.Context) *GetAdminEmailConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get admin email configuration params
func (o *GetAdminEmailConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get admin email configuration params
func (o *GetAdminEmailConfigurationParams) WithHTTPClient(client *http.Client) *GetAdminEmailConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get admin email configuration params
func (o *GetAdminEmailConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get admin email configuration params
func (o *GetAdminEmailConfigurationParams) WithNamespace(namespace string) *GetAdminEmailConfigurationParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get admin email configuration params
func (o *GetAdminEmailConfigurationParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdminEmailConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
