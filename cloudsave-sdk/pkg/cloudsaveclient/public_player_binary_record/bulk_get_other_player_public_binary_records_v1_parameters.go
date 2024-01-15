// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package public_player_binary_record

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
)

// NewBulkGetOtherPlayerPublicBinaryRecordsV1Params creates a new BulkGetOtherPlayerPublicBinaryRecordsV1Params object
// with the default values initialized.
func NewBulkGetOtherPlayerPublicBinaryRecordsV1Params() *BulkGetOtherPlayerPublicBinaryRecordsV1Params {
	var ()
	return &BulkGetOtherPlayerPublicBinaryRecordsV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewBulkGetOtherPlayerPublicBinaryRecordsV1ParamsWithTimeout creates a new BulkGetOtherPlayerPublicBinaryRecordsV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewBulkGetOtherPlayerPublicBinaryRecordsV1ParamsWithTimeout(timeout time.Duration) *BulkGetOtherPlayerPublicBinaryRecordsV1Params {
	var ()
	return &BulkGetOtherPlayerPublicBinaryRecordsV1Params{

		timeout: timeout,
	}
}

// NewBulkGetOtherPlayerPublicBinaryRecordsV1ParamsWithContext creates a new BulkGetOtherPlayerPublicBinaryRecordsV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewBulkGetOtherPlayerPublicBinaryRecordsV1ParamsWithContext(ctx context.Context) *BulkGetOtherPlayerPublicBinaryRecordsV1Params {
	var ()
	return &BulkGetOtherPlayerPublicBinaryRecordsV1Params{

		Context: ctx,
	}
}

// NewBulkGetOtherPlayerPublicBinaryRecordsV1ParamsWithHTTPClient creates a new BulkGetOtherPlayerPublicBinaryRecordsV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBulkGetOtherPlayerPublicBinaryRecordsV1ParamsWithHTTPClient(client *http.Client) *BulkGetOtherPlayerPublicBinaryRecordsV1Params {
	var ()
	return &BulkGetOtherPlayerPublicBinaryRecordsV1Params{
		HTTPClient: client,
	}
}

/*BulkGetOtherPlayerPublicBinaryRecordsV1Params contains all the parameters to send to the API endpoint
for the bulk get other player public binary records v1 operation typically these are written to a http.Request
*/
type BulkGetOtherPlayerPublicBinaryRecordsV1Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body *cloudsaveclientmodels.ModelsBulkGetPlayerRecordsRequest
	/*Namespace
	  namespace of the game, only accept alphabet and numeric

	*/
	Namespace string
	/*UserID
	  user ID who own the record, should follow UUID version 4 without hyphen format

	*/
	UserID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) WithTimeout(timeout time.Duration) *BulkGetOtherPlayerPublicBinaryRecordsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) WithContext(ctx context.Context) *BulkGetOtherPlayerPublicBinaryRecordsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) WithHTTPClient(client *http.Client) *BulkGetOtherPlayerPublicBinaryRecordsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithBody adds the body to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) WithBody(body *cloudsaveclientmodels.ModelsBulkGetPlayerRecordsRequest) *BulkGetOtherPlayerPublicBinaryRecordsV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) SetBody(body *cloudsaveclientmodels.ModelsBulkGetPlayerRecordsRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) WithNamespace(namespace string) *BulkGetOtherPlayerPublicBinaryRecordsV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) WithUserID(userID string) *BulkGetOtherPlayerPublicBinaryRecordsV1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the bulk get other player public binary records v1 params
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *BulkGetOtherPlayerPublicBinaryRecordsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
