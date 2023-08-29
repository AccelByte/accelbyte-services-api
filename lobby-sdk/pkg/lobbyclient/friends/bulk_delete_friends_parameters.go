// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package friends

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// NewBulkDeleteFriendsParams creates a new BulkDeleteFriendsParams object
// with the default values initialized.
func NewBulkDeleteFriendsParams() *BulkDeleteFriendsParams {
	var ()
	return &BulkDeleteFriendsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBulkDeleteFriendsParamsWithTimeout creates a new BulkDeleteFriendsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBulkDeleteFriendsParamsWithTimeout(timeout time.Duration) *BulkDeleteFriendsParams {
	var ()
	return &BulkDeleteFriendsParams{

		timeout: timeout,
	}
}

// NewBulkDeleteFriendsParamsWithContext creates a new BulkDeleteFriendsParams object
// with the default values initialized, and the ability to set a context for a request
func NewBulkDeleteFriendsParamsWithContext(ctx context.Context) *BulkDeleteFriendsParams {
	var ()
	return &BulkDeleteFriendsParams{

		Context: ctx,
	}
}

// NewBulkDeleteFriendsParamsWithHTTPClient creates a new BulkDeleteFriendsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBulkDeleteFriendsParamsWithHTTPClient(client *http.Client) *BulkDeleteFriendsParams {
	var ()
	return &BulkDeleteFriendsParams{
		HTTPClient: client,
	}
}

/*BulkDeleteFriendsParams contains all the parameters to send to the API endpoint
for the bulk delete friends operation typically these are written to a http.Request
*/
type BulkDeleteFriendsParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body *lobbyclientmodels.ModelBulkFriendsRequest
	/*Namespace
	  namespace

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
}

// WithTimeout adds the timeout to the bulk delete friends params
func (o *BulkDeleteFriendsParams) WithTimeout(timeout time.Duration) *BulkDeleteFriendsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk delete friends params
func (o *BulkDeleteFriendsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk delete friends params
func (o *BulkDeleteFriendsParams) WithContext(ctx context.Context) *BulkDeleteFriendsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk delete friends params
func (o *BulkDeleteFriendsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the bulk delete friends params
func (o *BulkDeleteFriendsParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the bulk delete friends params
func (o *BulkDeleteFriendsParams) WithHTTPClient(client *http.Client) *BulkDeleteFriendsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk delete friends params
func (o *BulkDeleteFriendsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the bulk delete friends params
func (o *BulkDeleteFriendsParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithBody adds the body to the bulk delete friends params
func (o *BulkDeleteFriendsParams) WithBody(body *lobbyclientmodels.ModelBulkFriendsRequest) *BulkDeleteFriendsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bulk delete friends params
func (o *BulkDeleteFriendsParams) SetBody(body *lobbyclientmodels.ModelBulkFriendsRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the bulk delete friends params
func (o *BulkDeleteFriendsParams) WithNamespace(namespace string) *BulkDeleteFriendsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the bulk delete friends params
func (o *BulkDeleteFriendsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the bulk delete friends params
func (o *BulkDeleteFriendsParams) WithUserID(userID string) *BulkDeleteFriendsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the bulk delete friends params
func (o *BulkDeleteFriendsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *BulkDeleteFriendsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
