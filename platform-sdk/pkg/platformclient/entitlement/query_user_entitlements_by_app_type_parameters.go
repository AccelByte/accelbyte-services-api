// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package entitlement

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
	"github.com/go-openapi/swag"
)

// NewQueryUserEntitlementsByAppTypeParams creates a new QueryUserEntitlementsByAppTypeParams object
// with the default values initialized.
func NewQueryUserEntitlementsByAppTypeParams() *QueryUserEntitlementsByAppTypeParams {
	var (
		activeOnlyDefault = bool(true)
		limitDefault      = int32(20)
		offsetDefault     = int32(0)
	)
	return &QueryUserEntitlementsByAppTypeParams{
		ActiveOnly: &activeOnlyDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryUserEntitlementsByAppTypeParamsWithTimeout creates a new QueryUserEntitlementsByAppTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryUserEntitlementsByAppTypeParamsWithTimeout(timeout time.Duration) *QueryUserEntitlementsByAppTypeParams {
	var (
		activeOnlyDefault = bool(true)
		limitDefault      = int32(20)
		offsetDefault     = int32(0)
	)
	return &QueryUserEntitlementsByAppTypeParams{
		ActiveOnly: &activeOnlyDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		timeout: timeout,
	}
}

// NewQueryUserEntitlementsByAppTypeParamsWithContext creates a new QueryUserEntitlementsByAppTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryUserEntitlementsByAppTypeParamsWithContext(ctx context.Context) *QueryUserEntitlementsByAppTypeParams {
	var (
		activeOnlyDefault = bool(true)
		limitDefault      = int32(20)
		offsetDefault     = int32(0)
	)
	return &QueryUserEntitlementsByAppTypeParams{
		ActiveOnly: &activeOnlyDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		Context: ctx,
	}
}

// NewQueryUserEntitlementsByAppTypeParamsWithHTTPClient creates a new QueryUserEntitlementsByAppTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryUserEntitlementsByAppTypeParamsWithHTTPClient(client *http.Client) *QueryUserEntitlementsByAppTypeParams {
	var (
		activeOnlyDefault = bool(true)
		limitDefault      = int32(20)
		offsetDefault     = int32(0)
	)
	return &QueryUserEntitlementsByAppTypeParams{
		ActiveOnly: &activeOnlyDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*QueryUserEntitlementsByAppTypeParams contains all the parameters to send to the API endpoint
for the query user entitlements by app type operation typically these are written to a http.Request
*/
type QueryUserEntitlementsByAppTypeParams struct {

	/*ActiveOnly*/
	ActiveOnly *bool
	/*AppType*/
	AppType string
	/*Limit*/
	Limit *int32
	/*Namespace*/
	Namespace string
	/*Offset*/
	Offset *int32
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) WithTimeout(timeout time.Duration) *QueryUserEntitlementsByAppTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) WithContext(ctx context.Context) *QueryUserEntitlementsByAppTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) WithHTTPClient(client *http.Client) *QueryUserEntitlementsByAppTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActiveOnly adds the activeOnly to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) WithActiveOnly(activeOnly *bool) *QueryUserEntitlementsByAppTypeParams {
	o.SetActiveOnly(activeOnly)
	return o
}

// SetActiveOnly adds the activeOnly to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) SetActiveOnly(activeOnly *bool) {
	o.ActiveOnly = activeOnly
}

// WithAppType adds the appType to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) WithAppType(appType string) *QueryUserEntitlementsByAppTypeParams {
	o.SetAppType(appType)
	return o
}

// SetAppType adds the appType to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) SetAppType(appType string) {
	o.AppType = appType
}

// WithLimit adds the limit to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) WithLimit(limit *int32) *QueryUserEntitlementsByAppTypeParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) WithNamespace(namespace string) *QueryUserEntitlementsByAppTypeParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) WithOffset(offset *int32) *QueryUserEntitlementsByAppTypeParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithUserID adds the userID to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) WithUserID(userID string) *QueryUserEntitlementsByAppTypeParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the query user entitlements by app type params
func (o *QueryUserEntitlementsByAppTypeParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *QueryUserEntitlementsByAppTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActiveOnly != nil {

		// query param activeOnly
		var qrActiveOnly bool
		if o.ActiveOnly != nil {
			qrActiveOnly = *o.ActiveOnly
		}
		qActiveOnly := swag.FormatBool(qrActiveOnly)
		if qActiveOnly != "" {
			if err := r.SetQueryParam("activeOnly", qActiveOnly); err != nil {
				return err
			}
		}

	}

	// query param appType
	qrAppType := o.AppType
	qAppType := qrAppType
	if qAppType != "" {
		if err := r.SetQueryParam("appType", qAppType); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

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
