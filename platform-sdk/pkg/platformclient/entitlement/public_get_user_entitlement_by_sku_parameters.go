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
)

// NewPublicGetUserEntitlementBySkuParams creates a new PublicGetUserEntitlementBySkuParams object
// with the default values initialized.
func NewPublicGetUserEntitlementBySkuParams() *PublicGetUserEntitlementBySkuParams {
	var ()
	return &PublicGetUserEntitlementBySkuParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetUserEntitlementBySkuParamsWithTimeout creates a new PublicGetUserEntitlementBySkuParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetUserEntitlementBySkuParamsWithTimeout(timeout time.Duration) *PublicGetUserEntitlementBySkuParams {
	var ()
	return &PublicGetUserEntitlementBySkuParams{

		timeout: timeout,
	}
}

// NewPublicGetUserEntitlementBySkuParamsWithContext creates a new PublicGetUserEntitlementBySkuParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetUserEntitlementBySkuParamsWithContext(ctx context.Context) *PublicGetUserEntitlementBySkuParams {
	var ()
	return &PublicGetUserEntitlementBySkuParams{

		Context: ctx,
	}
}

// NewPublicGetUserEntitlementBySkuParamsWithHTTPClient creates a new PublicGetUserEntitlementBySkuParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetUserEntitlementBySkuParamsWithHTTPClient(client *http.Client) *PublicGetUserEntitlementBySkuParams {
	var ()
	return &PublicGetUserEntitlementBySkuParams{
		HTTPClient: client,
	}
}

/*PublicGetUserEntitlementBySkuParams contains all the parameters to send to the API endpoint
for the public get user entitlement by sku operation typically these are written to a http.Request
*/
type PublicGetUserEntitlementBySkuParams struct {

	/*EntitlementClazz*/
	EntitlementClazz *string
	/*Namespace*/
	Namespace string
	/*Sku*/
	Sku string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) WithTimeout(timeout time.Duration) *PublicGetUserEntitlementBySkuParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) WithContext(ctx context.Context) *PublicGetUserEntitlementBySkuParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) WithHTTPClient(client *http.Client) *PublicGetUserEntitlementBySkuParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntitlementClazz adds the entitlementClazz to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) WithEntitlementClazz(entitlementClazz *string) *PublicGetUserEntitlementBySkuParams {
	o.SetEntitlementClazz(entitlementClazz)
	return o
}

// SetEntitlementClazz adds the entitlementClazz to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) SetEntitlementClazz(entitlementClazz *string) {
	o.EntitlementClazz = entitlementClazz
}

// WithNamespace adds the namespace to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) WithNamespace(namespace string) *PublicGetUserEntitlementBySkuParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSku adds the sku to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) WithSku(sku string) *PublicGetUserEntitlementBySkuParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) SetSku(sku string) {
	o.Sku = sku
}

// WithUserID adds the userID to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) WithUserID(userID string) *PublicGetUserEntitlementBySkuParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public get user entitlement by sku params
func (o *PublicGetUserEntitlementBySkuParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetUserEntitlementBySkuParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EntitlementClazz != nil {

		// query param entitlementClazz
		var qrEntitlementClazz string
		if o.EntitlementClazz != nil {
			qrEntitlementClazz = *o.EntitlementClazz
		}
		qEntitlementClazz := qrEntitlementClazz
		if qEntitlementClazz != "" {
			if err := r.SetQueryParam("entitlementClazz", qEntitlementClazz); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param sku
	qrSku := o.Sku
	qSku := qrSku
	if qSku != "" {
		if err := r.SetQueryParam("sku", qSku); err != nil {
			return err
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
