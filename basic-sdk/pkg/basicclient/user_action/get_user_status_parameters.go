// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package user_action

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

// NewGetUserStatusParams creates a new GetUserStatusParams object
// with the default values initialized.
func NewGetUserStatusParams() *GetUserStatusParams {
	var ()
	return &GetUserStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserStatusParamsWithTimeout creates a new GetUserStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserStatusParamsWithTimeout(timeout time.Duration) *GetUserStatusParams {
	var ()
	return &GetUserStatusParams{

		timeout: timeout,
	}
}

// NewGetUserStatusParamsWithContext creates a new GetUserStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserStatusParamsWithContext(ctx context.Context) *GetUserStatusParams {
	var ()
	return &GetUserStatusParams{

		Context: ctx,
	}
}

// NewGetUserStatusParamsWithHTTPClient creates a new GetUserStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserStatusParamsWithHTTPClient(client *http.Client) *GetUserStatusParams {
	var ()
	return &GetUserStatusParams{
		HTTPClient: client,
	}
}

/*GetUserStatusParams contains all the parameters to send to the API endpoint
for the get user status operation typically these are written to a http.Request
*/
type GetUserStatusParams struct {

	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string
	/*UserID
	  user id, should follow UUID version 4 without hyphen

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user status params
func (o *GetUserStatusParams) WithTimeout(timeout time.Duration) *GetUserStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user status params
func (o *GetUserStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user status params
func (o *GetUserStatusParams) WithContext(ctx context.Context) *GetUserStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user status params
func (o *GetUserStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user status params
func (o *GetUserStatusParams) WithHTTPClient(client *http.Client) *GetUserStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user status params
func (o *GetUserStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get user status params
func (o *GetUserStatusParams) WithNamespace(namespace string) *GetUserStatusParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get user status params
func (o *GetUserStatusParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the get user status params
func (o *GetUserStatusParams) WithUserID(userID string) *GetUserStatusParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user status params
func (o *GetUserStatusParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param userId
	qrUserID := o.UserID
	qUserID := qrUserID
	if qUserID != "" {
		if err := r.SetQueryParam("userId", qUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
