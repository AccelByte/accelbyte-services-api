// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package category

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

// NewGetRootCategoriesParams creates a new GetRootCategoriesParams object
// with the default values initialized.
func NewGetRootCategoriesParams() *GetRootCategoriesParams {
	var ()
	return &GetRootCategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRootCategoriesParamsWithTimeout creates a new GetRootCategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRootCategoriesParamsWithTimeout(timeout time.Duration) *GetRootCategoriesParams {
	var ()
	return &GetRootCategoriesParams{

		timeout: timeout,
	}
}

// NewGetRootCategoriesParamsWithContext creates a new GetRootCategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRootCategoriesParamsWithContext(ctx context.Context) *GetRootCategoriesParams {
	var ()
	return &GetRootCategoriesParams{

		Context: ctx,
	}
}

// NewGetRootCategoriesParamsWithHTTPClient creates a new GetRootCategoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRootCategoriesParamsWithHTTPClient(client *http.Client) *GetRootCategoriesParams {
	var ()
	return &GetRootCategoriesParams{
		HTTPClient: client,
	}
}

/*GetRootCategoriesParams contains all the parameters to send to the API endpoint
for the get root categories operation typically these are written to a http.Request
*/
type GetRootCategoriesParams struct {

	/*Namespace
	  Namespace

	*/
	Namespace string
	/*StoreID
	  default is published store id

	*/
	StoreID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get root categories params
func (o *GetRootCategoriesParams) WithTimeout(timeout time.Duration) *GetRootCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get root categories params
func (o *GetRootCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get root categories params
func (o *GetRootCategoriesParams) WithContext(ctx context.Context) *GetRootCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get root categories params
func (o *GetRootCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get root categories params
func (o *GetRootCategoriesParams) WithHTTPClient(client *http.Client) *GetRootCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get root categories params
func (o *GetRootCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get root categories params
func (o *GetRootCategoriesParams) WithNamespace(namespace string) *GetRootCategoriesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get root categories params
func (o *GetRootCategoriesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithStoreID adds the storeID to the get root categories params
func (o *GetRootCategoriesParams) WithStoreID(storeID *string) *GetRootCategoriesParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the get root categories params
func (o *GetRootCategoriesParams) SetStoreID(storeID *string) {
	o.StoreID = storeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRootCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.StoreID != nil {

		// query param storeId
		var qrStoreID string
		if o.StoreID != nil {
			qrStoreID = *o.StoreID
		}
		qStoreID := qrStoreID
		if qStoreID != "" {
			if err := r.SetQueryParam("storeId", qStoreID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
