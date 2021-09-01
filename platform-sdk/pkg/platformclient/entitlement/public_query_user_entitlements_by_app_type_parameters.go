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

// NewPublicQueryUserEntitlementsByAppTypeParams creates a new PublicQueryUserEntitlementsByAppTypeParams object
// with the default values initialized.
func NewPublicQueryUserEntitlementsByAppTypeParams() *PublicQueryUserEntitlementsByAppTypeParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &PublicQueryUserEntitlementsByAppTypeParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicQueryUserEntitlementsByAppTypeParamsWithTimeout creates a new PublicQueryUserEntitlementsByAppTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicQueryUserEntitlementsByAppTypeParamsWithTimeout(timeout time.Duration) *PublicQueryUserEntitlementsByAppTypeParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &PublicQueryUserEntitlementsByAppTypeParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewPublicQueryUserEntitlementsByAppTypeParamsWithContext creates a new PublicQueryUserEntitlementsByAppTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicQueryUserEntitlementsByAppTypeParamsWithContext(ctx context.Context) *PublicQueryUserEntitlementsByAppTypeParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &PublicQueryUserEntitlementsByAppTypeParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewPublicQueryUserEntitlementsByAppTypeParamsWithHTTPClient creates a new PublicQueryUserEntitlementsByAppTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicQueryUserEntitlementsByAppTypeParamsWithHTTPClient(client *http.Client) *PublicQueryUserEntitlementsByAppTypeParams {
	var (
		limitDefault  = int32(20)
		offsetDefault = int32(0)
	)
	return &PublicQueryUserEntitlementsByAppTypeParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*PublicQueryUserEntitlementsByAppTypeParams contains all the parameters to send to the API endpoint
for the public query user entitlements by app type operation typically these are written to a http.Request
*/
type PublicQueryUserEntitlementsByAppTypeParams struct {

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

// WithTimeout adds the timeout to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) WithTimeout(timeout time.Duration) *PublicQueryUserEntitlementsByAppTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) WithContext(ctx context.Context) *PublicQueryUserEntitlementsByAppTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) WithHTTPClient(client *http.Client) *PublicQueryUserEntitlementsByAppTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppType adds the appType to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) WithAppType(appType string) *PublicQueryUserEntitlementsByAppTypeParams {
	o.SetAppType(appType)
	return o
}

// SetAppType adds the appType to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) SetAppType(appType string) {
	o.AppType = appType
}

// WithLimit adds the limit to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) WithLimit(limit *int32) *PublicQueryUserEntitlementsByAppTypeParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) WithNamespace(namespace string) *PublicQueryUserEntitlementsByAppTypeParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) WithOffset(offset *int32) *PublicQueryUserEntitlementsByAppTypeParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithUserID adds the userID to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) WithUserID(userID string) *PublicQueryUserEntitlementsByAppTypeParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public query user entitlements by app type params
func (o *PublicQueryUserEntitlementsByAppTypeParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicQueryUserEntitlementsByAppTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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