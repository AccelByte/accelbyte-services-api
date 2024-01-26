// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package clients

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteClientPermissionParams creates a new DeleteClientPermissionParams object
// with the default values initialized.
func NewDeleteClientPermissionParams() *DeleteClientPermissionParams {
	var ()
	return &DeleteClientPermissionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClientPermissionParamsWithTimeout creates a new DeleteClientPermissionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClientPermissionParamsWithTimeout(timeout time.Duration) *DeleteClientPermissionParams {
	var ()
	return &DeleteClientPermissionParams{

		timeout: timeout,
	}
}

// NewDeleteClientPermissionParamsWithContext creates a new DeleteClientPermissionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClientPermissionParamsWithContext(ctx context.Context) *DeleteClientPermissionParams {
	var ()
	return &DeleteClientPermissionParams{

		Context: ctx,
	}
}

// NewDeleteClientPermissionParamsWithHTTPClient creates a new DeleteClientPermissionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClientPermissionParamsWithHTTPClient(client *http.Client) *DeleteClientPermissionParams {
	var ()
	return &DeleteClientPermissionParams{
		HTTPClient: client,
	}
}

/*DeleteClientPermissionParams contains all the parameters to send to the API endpoint
for the delete client permission operation typically these are written to a http.Request
*/
type DeleteClientPermissionParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Action
	  Action

	*/
	Action int64
	/*ClientID
	  Client ID

	*/
	ClientID string
	/*Resource
	  Resource Name

	*/
	Resource string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the delete client permission params
func (o *DeleteClientPermissionParams) WithTimeout(timeout time.Duration) *DeleteClientPermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete client permission params
func (o *DeleteClientPermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete client permission params
func (o *DeleteClientPermissionParams) WithContext(ctx context.Context) *DeleteClientPermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete client permission params
func (o *DeleteClientPermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the delete client permission params
func (o *DeleteClientPermissionParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the delete client permission params
func (o *DeleteClientPermissionParams) WithHTTPClient(client *http.Client) *DeleteClientPermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete client permission params
func (o *DeleteClientPermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the delete client permission params
func (o *DeleteClientPermissionParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *DeleteClientPermissionParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithAction adds the action to the delete client permission params
func (o *DeleteClientPermissionParams) WithAction(action int64) *DeleteClientPermissionParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the delete client permission params
func (o *DeleteClientPermissionParams) SetAction(action int64) {
	o.Action = action
}

// WithClientID adds the clientID to the delete client permission params
func (o *DeleteClientPermissionParams) WithClientID(clientID string) *DeleteClientPermissionParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the delete client permission params
func (o *DeleteClientPermissionParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithResource adds the resource to the delete client permission params
func (o *DeleteClientPermissionParams) WithResource(resource string) *DeleteClientPermissionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the delete client permission params
func (o *DeleteClientPermissionParams) SetResource(resource string) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClientPermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action
	if err := r.SetPathParam("action", swag.FormatInt64(o.Action)); err != nil {
		return err
	}

	// path param clientId
	if err := r.SetPathParam("clientId", o.ClientID); err != nil {
		return err
	}

	// path param resource
	if err := r.SetPathParam("resource", o.Resource); err != nil {
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
