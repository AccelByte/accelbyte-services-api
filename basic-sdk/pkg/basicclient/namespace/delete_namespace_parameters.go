// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package namespace

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteNamespaceParams creates a new DeleteNamespaceParams object
// with the default values initialized.
func NewDeleteNamespaceParams() *DeleteNamespaceParams {
	var ()
	return &DeleteNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNamespaceParamsWithTimeout creates a new DeleteNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNamespaceParamsWithTimeout(timeout time.Duration) *DeleteNamespaceParams {
	var ()
	return &DeleteNamespaceParams{

		timeout: timeout,
	}
}

// NewDeleteNamespaceParamsWithContext creates a new DeleteNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNamespaceParamsWithContext(ctx context.Context) *DeleteNamespaceParams {
	var ()
	return &DeleteNamespaceParams{

		Context: ctx,
	}
}

// NewDeleteNamespaceParamsWithHTTPClient creates a new DeleteNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNamespaceParamsWithHTTPClient(client *http.Client) *DeleteNamespaceParams {
	var ()
	return &DeleteNamespaceParams{
		HTTPClient: client,
	}
}

/*DeleteNamespaceParams contains all the parameters to send to the API endpoint
for the delete namespace operation typically these are written to a http.Request
*/
type DeleteNamespaceParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the delete namespace params
func (o *DeleteNamespaceParams) WithTimeout(timeout time.Duration) *DeleteNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete namespace params
func (o *DeleteNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete namespace params
func (o *DeleteNamespaceParams) WithContext(ctx context.Context) *DeleteNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete namespace params
func (o *DeleteNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the delete namespace params
func (o *DeleteNamespaceParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the delete namespace params
func (o *DeleteNamespaceParams) WithHTTPClient(client *http.Client) *DeleteNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete namespace params
func (o *DeleteNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the delete namespace params
func (o *DeleteNamespaceParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *DeleteNamespaceParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithNamespace adds the namespace to the delete namespace params
func (o *DeleteNamespaceParams) WithNamespace(namespace string) *DeleteNamespaceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete namespace params
func (o *DeleteNamespaceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
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
