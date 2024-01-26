// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package item

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewCreateItemParams creates a new CreateItemParams object
// with the default values initialized.
func NewCreateItemParams() *CreateItemParams {
	var ()
	return &CreateItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateItemParamsWithTimeout creates a new CreateItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateItemParamsWithTimeout(timeout time.Duration) *CreateItemParams {
	var ()
	return &CreateItemParams{

		timeout: timeout,
	}
}

// NewCreateItemParamsWithContext creates a new CreateItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateItemParamsWithContext(ctx context.Context) *CreateItemParams {
	var ()
	return &CreateItemParams{

		Context: ctx,
	}
}

// NewCreateItemParamsWithHTTPClient creates a new CreateItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateItemParamsWithHTTPClient(client *http.Client) *CreateItemParams {
	var ()
	return &CreateItemParams{
		HTTPClient: client,
	}
}

/*CreateItemParams contains all the parameters to send to the API endpoint
for the create item operation typically these are written to a http.Request
*/
type CreateItemParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body *platformclientmodels.ItemCreate
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*StoreID*/
	StoreID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the create item params
func (o *CreateItemParams) WithTimeout(timeout time.Duration) *CreateItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create item params
func (o *CreateItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create item params
func (o *CreateItemParams) WithContext(ctx context.Context) *CreateItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create item params
func (o *CreateItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the create item params
func (o *CreateItemParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the create item params
func (o *CreateItemParams) WithHTTPClient(client *http.Client) *CreateItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create item params
func (o *CreateItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the create item params
func (o *CreateItemParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *CreateItemParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithBody adds the body to the create item params
func (o *CreateItemParams) WithBody(body *platformclientmodels.ItemCreate) *CreateItemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create item params
func (o *CreateItemParams) SetBody(body *platformclientmodels.ItemCreate) {
	o.Body = body
}

// WithNamespace adds the namespace to the create item params
func (o *CreateItemParams) WithNamespace(namespace string) *CreateItemParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create item params
func (o *CreateItemParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithStoreID adds the storeID to the create item params
func (o *CreateItemParams) WithStoreID(storeID string) *CreateItemParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the create item params
func (o *CreateItemParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param storeId
	qrStoreID := o.StoreID
	qStoreID := qrStoreID
	if qStoreID != "" {
		if err := r.SetQueryParam("storeId", qStoreID); err != nil {
			return err
		}
	}

	// setting the default header value
	if err := r.SetHeaderParam("User-Agent", utils.UserAgentGen()); err != nil {
		return err
	}

	if err := r.SetHeaderParam("X-Amzn-Trace-Id", utils.AmazonTraceIDGen()); err != nil {
		return err
	}

	if o.XFlightId == nil {
		if err := r.SetHeaderParam("X-Flight-Id", utils.GetDefaultFlightID().Value); err != nil {
			return err
		}
	} else {
		if err := r.SetHeaderParam("X-Flight-Id", *o.XFlightId); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
