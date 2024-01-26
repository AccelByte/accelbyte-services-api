// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package admin_type

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

// NewAdminDeleteTypeParams creates a new AdminDeleteTypeParams object
// with the default values initialized.
func NewAdminDeleteTypeParams() *AdminDeleteTypeParams {
	var ()
	return &AdminDeleteTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminDeleteTypeParamsWithTimeout creates a new AdminDeleteTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminDeleteTypeParamsWithTimeout(timeout time.Duration) *AdminDeleteTypeParams {
	var ()
	return &AdminDeleteTypeParams{

		timeout: timeout,
	}
}

// NewAdminDeleteTypeParamsWithContext creates a new AdminDeleteTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminDeleteTypeParamsWithContext(ctx context.Context) *AdminDeleteTypeParams {
	var ()
	return &AdminDeleteTypeParams{

		Context: ctx,
	}
}

// NewAdminDeleteTypeParamsWithHTTPClient creates a new AdminDeleteTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminDeleteTypeParamsWithHTTPClient(client *http.Client) *AdminDeleteTypeParams {
	var ()
	return &AdminDeleteTypeParams{
		HTTPClient: client,
	}
}

/*AdminDeleteTypeParams contains all the parameters to send to the API endpoint
for the admin delete type operation typically these are written to a http.Request
*/
type AdminDeleteTypeParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*TypeID
	  type ID

	*/
	TypeID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the admin delete type params
func (o *AdminDeleteTypeParams) WithTimeout(timeout time.Duration) *AdminDeleteTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin delete type params
func (o *AdminDeleteTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin delete type params
func (o *AdminDeleteTypeParams) WithContext(ctx context.Context) *AdminDeleteTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin delete type params
func (o *AdminDeleteTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin delete type params
func (o *AdminDeleteTypeParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin delete type params
func (o *AdminDeleteTypeParams) WithHTTPClient(client *http.Client) *AdminDeleteTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin delete type params
func (o *AdminDeleteTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the admin delete type params
func (o *AdminDeleteTypeParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *AdminDeleteTypeParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithNamespace adds the namespace to the admin delete type params
func (o *AdminDeleteTypeParams) WithNamespace(namespace string) *AdminDeleteTypeParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin delete type params
func (o *AdminDeleteTypeParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithTypeID adds the typeID to the admin delete type params
func (o *AdminDeleteTypeParams) WithTypeID(typeID string) *AdminDeleteTypeParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the admin delete type params
func (o *AdminDeleteTypeParams) SetTypeID(typeID string) {
	o.TypeID = typeID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminDeleteTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param typeId
	if err := r.SetPathParam("typeId", o.TypeID); err != nil {
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
