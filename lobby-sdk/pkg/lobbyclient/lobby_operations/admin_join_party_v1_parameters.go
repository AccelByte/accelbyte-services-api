// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package lobby_operations

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

// NewAdminJoinPartyV1Params creates a new AdminJoinPartyV1Params object
// with the default values initialized.
func NewAdminJoinPartyV1Params() *AdminJoinPartyV1Params {
	var ()
	return &AdminJoinPartyV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminJoinPartyV1ParamsWithTimeout creates a new AdminJoinPartyV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminJoinPartyV1ParamsWithTimeout(timeout time.Duration) *AdminJoinPartyV1Params {
	var ()
	return &AdminJoinPartyV1Params{

		timeout: timeout,
	}
}

// NewAdminJoinPartyV1ParamsWithContext creates a new AdminJoinPartyV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminJoinPartyV1ParamsWithContext(ctx context.Context) *AdminJoinPartyV1Params {
	var ()
	return &AdminJoinPartyV1Params{

		Context: ctx,
	}
}

// NewAdminJoinPartyV1ParamsWithHTTPClient creates a new AdminJoinPartyV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminJoinPartyV1ParamsWithHTTPClient(client *http.Client) *AdminJoinPartyV1Params {
	var ()
	return &AdminJoinPartyV1Params{
		HTTPClient: client,
	}
}

/*AdminJoinPartyV1Params contains all the parameters to send to the API endpoint
for the admin join party v1 operation typically these are written to a http.Request
*/
type AdminJoinPartyV1Params struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*PartyID
	  party ID

	*/
	PartyID string
	/*UserID
	  user ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin join party v1 params
func (o *AdminJoinPartyV1Params) WithTimeout(timeout time.Duration) *AdminJoinPartyV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin join party v1 params
func (o *AdminJoinPartyV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin join party v1 params
func (o *AdminJoinPartyV1Params) WithContext(ctx context.Context) *AdminJoinPartyV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin join party v1 params
func (o *AdminJoinPartyV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin join party v1 params
func (o *AdminJoinPartyV1Params) WithHTTPClient(client *http.Client) *AdminJoinPartyV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin join party v1 params
func (o *AdminJoinPartyV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the admin join party v1 params
func (o *AdminJoinPartyV1Params) WithNamespace(namespace string) *AdminJoinPartyV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin join party v1 params
func (o *AdminJoinPartyV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPartyID adds the partyID to the admin join party v1 params
func (o *AdminJoinPartyV1Params) WithPartyID(partyID string) *AdminJoinPartyV1Params {
	o.SetPartyID(partyID)
	return o
}

// SetPartyID adds the partyId to the admin join party v1 params
func (o *AdminJoinPartyV1Params) SetPartyID(partyID string) {
	o.PartyID = partyID
}

// WithUserID adds the userID to the admin join party v1 params
func (o *AdminJoinPartyV1Params) WithUserID(userID string) *AdminJoinPartyV1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin join party v1 params
func (o *AdminJoinPartyV1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminJoinPartyV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param partyId
	if err := r.SetPathParam("partyId", o.PartyID); err != nil {
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
