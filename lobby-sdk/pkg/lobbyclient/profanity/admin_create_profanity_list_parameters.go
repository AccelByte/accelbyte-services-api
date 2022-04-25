// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

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

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// NewAdminCreateProfanityListParams creates a new AdminCreateProfanityListParams object
// with the default values initialized.
func NewAdminCreateProfanityListParams() *AdminCreateProfanityListParams {
	var ()
	return &AdminCreateProfanityListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminCreateProfanityListParamsWithTimeout creates a new AdminCreateProfanityListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminCreateProfanityListParamsWithTimeout(timeout time.Duration) *AdminCreateProfanityListParams {
	var ()
	return &AdminCreateProfanityListParams{

		timeout: timeout,
	}
}

// NewAdminCreateProfanityListParamsWithContext creates a new AdminCreateProfanityListParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminCreateProfanityListParamsWithContext(ctx context.Context) *AdminCreateProfanityListParams {
	var ()
	return &AdminCreateProfanityListParams{

		Context: ctx,
	}
}

// NewAdminCreateProfanityListParamsWithHTTPClient creates a new AdminCreateProfanityListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminCreateProfanityListParamsWithHTTPClient(client *http.Client) *AdminCreateProfanityListParams {
	var ()
	return &AdminCreateProfanityListParams{
		HTTPClient: client,
	}
}

/*AdminCreateProfanityListParams contains all the parameters to send to the API endpoint
for the admin create profanity list operation typically these are written to a http.Request
*/
type AdminCreateProfanityListParams struct {

	/*Body
	  request

	*/
	Body *lobbyclientmodels.ModelsAdminCreateProfanityListRequest
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin create profanity list params
func (o *AdminCreateProfanityListParams) WithTimeout(timeout time.Duration) *AdminCreateProfanityListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin create profanity list params
func (o *AdminCreateProfanityListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin create profanity list params
func (o *AdminCreateProfanityListParams) WithContext(ctx context.Context) *AdminCreateProfanityListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin create profanity list params
func (o *AdminCreateProfanityListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin create profanity list params
func (o *AdminCreateProfanityListParams) WithHTTPClient(client *http.Client) *AdminCreateProfanityListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin create profanity list params
func (o *AdminCreateProfanityListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin create profanity list params
func (o *AdminCreateProfanityListParams) WithBody(body *lobbyclientmodels.ModelsAdminCreateProfanityListRequest) *AdminCreateProfanityListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin create profanity list params
func (o *AdminCreateProfanityListParams) SetBody(body *lobbyclientmodels.ModelsAdminCreateProfanityListRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the admin create profanity list params
func (o *AdminCreateProfanityListParams) WithNamespace(namespace string) *AdminCreateProfanityListParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin create profanity list params
func (o *AdminCreateProfanityListParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminCreateProfanityListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
