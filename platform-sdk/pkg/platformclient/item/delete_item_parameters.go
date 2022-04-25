// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package item

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

// NewDeleteItemParams creates a new DeleteItemParams object
// with the default values initialized.
func NewDeleteItemParams() *DeleteItemParams {
	var ()
	return &DeleteItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteItemParamsWithTimeout creates a new DeleteItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteItemParamsWithTimeout(timeout time.Duration) *DeleteItemParams {
	var ()
	return &DeleteItemParams{

		timeout: timeout,
	}
}

// NewDeleteItemParamsWithContext creates a new DeleteItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteItemParamsWithContext(ctx context.Context) *DeleteItemParams {
	var ()
	return &DeleteItemParams{

		Context: ctx,
	}
}

// NewDeleteItemParamsWithHTTPClient creates a new DeleteItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteItemParamsWithHTTPClient(client *http.Client) *DeleteItemParams {
	var ()
	return &DeleteItemParams{
		HTTPClient: client,
	}
}

/*DeleteItemParams contains all the parameters to send to the API endpoint
for the delete item operation typically these are written to a http.Request
*/
type DeleteItemParams struct {

	/*ItemID*/
	ItemID string
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

// WithTimeout adds the timeout to the delete item params
func (o *DeleteItemParams) WithTimeout(timeout time.Duration) *DeleteItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete item params
func (o *DeleteItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete item params
func (o *DeleteItemParams) WithContext(ctx context.Context) *DeleteItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete item params
func (o *DeleteItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete item params
func (o *DeleteItemParams) WithHTTPClient(client *http.Client) *DeleteItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete item params
func (o *DeleteItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItemID adds the itemID to the delete item params
func (o *DeleteItemParams) WithItemID(itemID string) *DeleteItemParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the delete item params
func (o *DeleteItemParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithNamespace adds the namespace to the delete item params
func (o *DeleteItemParams) WithNamespace(namespace string) *DeleteItemParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete item params
func (o *DeleteItemParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithStoreID adds the storeID to the delete item params
func (o *DeleteItemParams) WithStoreID(storeID *string) *DeleteItemParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the delete item params
func (o *DeleteItemParams) SetStoreID(storeID *string) {
	o.StoreID = storeID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param itemId
	if err := r.SetPathParam("itemId", o.ItemID); err != nil {
		return err
	}

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
