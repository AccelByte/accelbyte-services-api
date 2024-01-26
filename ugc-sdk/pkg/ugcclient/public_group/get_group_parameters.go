// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package public_group

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

// NewGetGroupParams creates a new GetGroupParams object
// with the default values initialized.
func NewGetGroupParams() *GetGroupParams {
	var ()
	return &GetGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupParamsWithTimeout creates a new GetGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupParamsWithTimeout(timeout time.Duration) *GetGroupParams {
	var ()
	return &GetGroupParams{

		timeout: timeout,
	}
}

// NewGetGroupParamsWithContext creates a new GetGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupParamsWithContext(ctx context.Context) *GetGroupParams {
	var ()
	return &GetGroupParams{

		Context: ctx,
	}
}

// NewGetGroupParamsWithHTTPClient creates a new GetGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupParamsWithHTTPClient(client *http.Client) *GetGroupParams {
	var ()
	return &GetGroupParams{
		HTTPClient: client,
	}
}

/*GetGroupParams contains all the parameters to send to the API endpoint
for the get group operation typically these are written to a http.Request
*/
type GetGroupParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*GroupID
	  group ID

	*/
	GroupID string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*UserID
	  user ID

	*/
	UserID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the get group params
func (o *GetGroupParams) WithTimeout(timeout time.Duration) *GetGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get group params
func (o *GetGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get group params
func (o *GetGroupParams) WithContext(ctx context.Context) *GetGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get group params
func (o *GetGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the get group params
func (o *GetGroupParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the get group params
func (o *GetGroupParams) WithHTTPClient(client *http.Client) *GetGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get group params
func (o *GetGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the get group params
func (o *GetGroupParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *GetGroupParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithGroupID adds the groupID to the get group params
func (o *GetGroupParams) WithGroupID(groupID string) *GetGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get group params
func (o *GetGroupParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithNamespace adds the namespace to the get group params
func (o *GetGroupParams) WithNamespace(namespace string) *GetGroupParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get group params
func (o *GetGroupParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the get group params
func (o *GetGroupParams) WithUserID(userID string) *GetGroupParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get group params
func (o *GetGroupParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
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
