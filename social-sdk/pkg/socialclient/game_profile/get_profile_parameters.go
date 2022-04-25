// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package game_profile

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

// NewGetProfileParams creates a new GetProfileParams object
// with the default values initialized.
func NewGetProfileParams() *GetProfileParams {
	var ()
	return &GetProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProfileParamsWithTimeout creates a new GetProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProfileParamsWithTimeout(timeout time.Duration) *GetProfileParams {
	var ()
	return &GetProfileParams{

		timeout: timeout,
	}
}

// NewGetProfileParamsWithContext creates a new GetProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProfileParamsWithContext(ctx context.Context) *GetProfileParams {
	var ()
	return &GetProfileParams{

		Context: ctx,
	}
}

// NewGetProfileParamsWithHTTPClient creates a new GetProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProfileParamsWithHTTPClient(client *http.Client) *GetProfileParams {
	var ()
	return &GetProfileParams{
		HTTPClient: client,
	}
}

/*GetProfileParams contains all the parameters to send to the API endpoint
for the get profile operation typically these are written to a http.Request
*/
type GetProfileParams struct {

	/*Namespace
	  Namespace ID

	*/
	Namespace string
	/*ProfileID
	  Game profile ID

	*/
	ProfileID string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get profile params
func (o *GetProfileParams) WithTimeout(timeout time.Duration) *GetProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get profile params
func (o *GetProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get profile params
func (o *GetProfileParams) WithContext(ctx context.Context) *GetProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get profile params
func (o *GetProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get profile params
func (o *GetProfileParams) WithHTTPClient(client *http.Client) *GetProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get profile params
func (o *GetProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get profile params
func (o *GetProfileParams) WithNamespace(namespace string) *GetProfileParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get profile params
func (o *GetProfileParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithProfileID adds the profileID to the get profile params
func (o *GetProfileParams) WithProfileID(profileID string) *GetProfileParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the get profile params
func (o *GetProfileParams) SetProfileID(profileID string) {
	o.ProfileID = profileID
}

// WithUserID adds the userID to the get profile params
func (o *GetProfileParams) WithUserID(userID string) *GetProfileParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get profile params
func (o *GetProfileParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param profileId
	if err := r.SetPathParam("profileId", o.ProfileID); err != nil {
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
