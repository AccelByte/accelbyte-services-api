// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

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

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewSyncStadiaEntitlementParams creates a new SyncStadiaEntitlementParams object
// with the default values initialized.
func NewSyncStadiaEntitlementParams() *SyncStadiaEntitlementParams {
	var ()
	return &SyncStadiaEntitlementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSyncStadiaEntitlementParamsWithTimeout creates a new SyncStadiaEntitlementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSyncStadiaEntitlementParamsWithTimeout(timeout time.Duration) *SyncStadiaEntitlementParams {
	var ()
	return &SyncStadiaEntitlementParams{

		timeout: timeout,
	}
}

// NewSyncStadiaEntitlementParamsWithContext creates a new SyncStadiaEntitlementParams object
// with the default values initialized, and the ability to set a context for a request
func NewSyncStadiaEntitlementParamsWithContext(ctx context.Context) *SyncStadiaEntitlementParams {
	var ()
	return &SyncStadiaEntitlementParams{

		Context: ctx,
	}
}

// NewSyncStadiaEntitlementParamsWithHTTPClient creates a new SyncStadiaEntitlementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSyncStadiaEntitlementParamsWithHTTPClient(client *http.Client) *SyncStadiaEntitlementParams {
	var ()
	return &SyncStadiaEntitlementParams{
		HTTPClient: client,
	}
}

/*SyncStadiaEntitlementParams contains all the parameters to send to the API endpoint
for the sync stadia entitlement operation typically these are written to a http.Request
*/
type SyncStadiaEntitlementParams struct {

	/*Body*/
	Body *platformclientmodels.StadiaSyncRequest
	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) WithTimeout(timeout time.Duration) *SyncStadiaEntitlementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) WithContext(ctx context.Context) *SyncStadiaEntitlementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) WithHTTPClient(client *http.Client) *SyncStadiaEntitlementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) WithBody(body *platformclientmodels.StadiaSyncRequest) *SyncStadiaEntitlementParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) SetBody(body *platformclientmodels.StadiaSyncRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) WithNamespace(namespace string) *SyncStadiaEntitlementParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) WithUserID(userID string) *SyncStadiaEntitlementParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the sync stadia entitlement params
func (o *SyncStadiaEntitlementParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *SyncStadiaEntitlementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
