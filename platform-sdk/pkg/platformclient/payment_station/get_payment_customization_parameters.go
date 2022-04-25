// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package payment_station

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

// NewGetPaymentCustomizationParams creates a new GetPaymentCustomizationParams object
// with the default values initialized.
func NewGetPaymentCustomizationParams() *GetPaymentCustomizationParams {
	var (
		sandboxDefault = bool(false)
	)
	return &GetPaymentCustomizationParams{
		Sandbox: &sandboxDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentCustomizationParamsWithTimeout creates a new GetPaymentCustomizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentCustomizationParamsWithTimeout(timeout time.Duration) *GetPaymentCustomizationParams {
	var (
		sandboxDefault = bool(false)
	)
	return &GetPaymentCustomizationParams{
		Sandbox: &sandboxDefault,

		timeout: timeout,
	}
}

// NewGetPaymentCustomizationParamsWithContext creates a new GetPaymentCustomizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentCustomizationParamsWithContext(ctx context.Context) *GetPaymentCustomizationParams {
	var (
		sandboxDefault = bool(false)
	)
	return &GetPaymentCustomizationParams{
		Sandbox: &sandboxDefault,

		Context: ctx,
	}
}

// NewGetPaymentCustomizationParamsWithHTTPClient creates a new GetPaymentCustomizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentCustomizationParamsWithHTTPClient(client *http.Client) *GetPaymentCustomizationParams {
	var (
		sandboxDefault = bool(false)
	)
	return &GetPaymentCustomizationParams{
		Sandbox:    &sandboxDefault,
		HTTPClient: client,
	}
}

/*GetPaymentCustomizationParams contains all the parameters to send to the API endpoint
for the get payment customization operation typically these are written to a http.Request
*/
type GetPaymentCustomizationParams struct {

	/*Namespace*/
	Namespace string
	/*PaymentProvider*/
	PaymentProvider string
	/*Region*/
	Region string
	/*Sandbox*/
	Sandbox *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payment customization params
func (o *GetPaymentCustomizationParams) WithTimeout(timeout time.Duration) *GetPaymentCustomizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment customization params
func (o *GetPaymentCustomizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment customization params
func (o *GetPaymentCustomizationParams) WithContext(ctx context.Context) *GetPaymentCustomizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment customization params
func (o *GetPaymentCustomizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment customization params
func (o *GetPaymentCustomizationParams) WithHTTPClient(client *http.Client) *GetPaymentCustomizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment customization params
func (o *GetPaymentCustomizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get payment customization params
func (o *GetPaymentCustomizationParams) WithNamespace(namespace string) *GetPaymentCustomizationParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get payment customization params
func (o *GetPaymentCustomizationParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPaymentProvider adds the paymentProvider to the get payment customization params
func (o *GetPaymentCustomizationParams) WithPaymentProvider(paymentProvider string) *GetPaymentCustomizationParams {
	o.SetPaymentProvider(paymentProvider)
	return o
}

// SetPaymentProvider adds the paymentProvider to the get payment customization params
func (o *GetPaymentCustomizationParams) SetPaymentProvider(paymentProvider string) {
	o.PaymentProvider = paymentProvider
}

// WithRegion adds the region to the get payment customization params
func (o *GetPaymentCustomizationParams) WithRegion(region string) *GetPaymentCustomizationParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get payment customization params
func (o *GetPaymentCustomizationParams) SetRegion(region string) {
	o.Region = region
}

// WithSandbox adds the sandbox to the get payment customization params
func (o *GetPaymentCustomizationParams) WithSandbox(sandbox *bool) *GetPaymentCustomizationParams {
	o.SetSandbox(sandbox)
	return o
}

// SetSandbox adds the sandbox to the get payment customization params
func (o *GetPaymentCustomizationParams) SetSandbox(sandbox *bool) {
	o.Sandbox = sandbox
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentCustomizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param paymentProvider
	qrPaymentProvider := o.PaymentProvider
	qPaymentProvider := qrPaymentProvider
	if qPaymentProvider != "" {
		if err := r.SetQueryParam("paymentProvider", qPaymentProvider); err != nil {
			return err
		}
	}

	// query param region
	qrRegion := o.Region
	qRegion := qrRegion
	if qRegion != "" {
		if err := r.SetQueryParam("region", qRegion); err != nil {
			return err
		}
	}

	if o.Sandbox != nil {

		// query param sandbox
		var qrSandbox bool
		if o.Sandbox != nil {
			qrSandbox = *o.Sandbox
		}
		qSandbox := swag.FormatBool(qrSandbox)
		if qSandbox != "" {
			if err := r.SetQueryParam("sandbox", qSandbox); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
