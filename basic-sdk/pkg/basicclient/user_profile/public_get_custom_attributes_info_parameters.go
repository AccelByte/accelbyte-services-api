// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package user_profile

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

// NewPublicGetCustomAttributesInfoParams creates a new PublicGetCustomAttributesInfoParams object
// with the default values initialized.
func NewPublicGetCustomAttributesInfoParams() *PublicGetCustomAttributesInfoParams {
	var ()
	return &PublicGetCustomAttributesInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetCustomAttributesInfoParamsWithTimeout creates a new PublicGetCustomAttributesInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetCustomAttributesInfoParamsWithTimeout(timeout time.Duration) *PublicGetCustomAttributesInfoParams {
	var ()
	return &PublicGetCustomAttributesInfoParams{

		timeout: timeout,
	}
}

// NewPublicGetCustomAttributesInfoParamsWithContext creates a new PublicGetCustomAttributesInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetCustomAttributesInfoParamsWithContext(ctx context.Context) *PublicGetCustomAttributesInfoParams {
	var ()
	return &PublicGetCustomAttributesInfoParams{

		Context: ctx,
	}
}

// NewPublicGetCustomAttributesInfoParamsWithHTTPClient creates a new PublicGetCustomAttributesInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetCustomAttributesInfoParamsWithHTTPClient(client *http.Client) *PublicGetCustomAttributesInfoParams {
	var ()
	return &PublicGetCustomAttributesInfoParams{
		HTTPClient: client,
	}
}

/*PublicGetCustomAttributesInfoParams contains all the parameters to send to the API endpoint
for the public get custom attributes info operation typically these are written to a http.Request
*/
type PublicGetCustomAttributesInfoParams struct {

	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string
	/*UserID
	  user's id, should follow UUID version 4 without hyphen

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public get custom attributes info params
func (o *PublicGetCustomAttributesInfoParams) WithTimeout(timeout time.Duration) *PublicGetCustomAttributesInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get custom attributes info params
func (o *PublicGetCustomAttributesInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get custom attributes info params
func (o *PublicGetCustomAttributesInfoParams) WithContext(ctx context.Context) *PublicGetCustomAttributesInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get custom attributes info params
func (o *PublicGetCustomAttributesInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get custom attributes info params
func (o *PublicGetCustomAttributesInfoParams) WithHTTPClient(client *http.Client) *PublicGetCustomAttributesInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get custom attributes info params
func (o *PublicGetCustomAttributesInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the public get custom attributes info params
func (o *PublicGetCustomAttributesInfoParams) WithNamespace(namespace string) *PublicGetCustomAttributesInfoParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get custom attributes info params
func (o *PublicGetCustomAttributesInfoParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public get custom attributes info params
func (o *PublicGetCustomAttributesInfoParams) WithUserID(userID string) *PublicGetCustomAttributesInfoParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public get custom attributes info params
func (o *PublicGetCustomAttributesInfoParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetCustomAttributesInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
