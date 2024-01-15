// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package public_player_record

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

// NewPublicDeletePlayerPublicRecordHandlerV1Params creates a new PublicDeletePlayerPublicRecordHandlerV1Params object
// with the default values initialized.
func NewPublicDeletePlayerPublicRecordHandlerV1Params() *PublicDeletePlayerPublicRecordHandlerV1Params {
	var ()
	return &PublicDeletePlayerPublicRecordHandlerV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicDeletePlayerPublicRecordHandlerV1ParamsWithTimeout creates a new PublicDeletePlayerPublicRecordHandlerV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicDeletePlayerPublicRecordHandlerV1ParamsWithTimeout(timeout time.Duration) *PublicDeletePlayerPublicRecordHandlerV1Params {
	var ()
	return &PublicDeletePlayerPublicRecordHandlerV1Params{

		timeout: timeout,
	}
}

// NewPublicDeletePlayerPublicRecordHandlerV1ParamsWithContext creates a new PublicDeletePlayerPublicRecordHandlerV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewPublicDeletePlayerPublicRecordHandlerV1ParamsWithContext(ctx context.Context) *PublicDeletePlayerPublicRecordHandlerV1Params {
	var ()
	return &PublicDeletePlayerPublicRecordHandlerV1Params{

		Context: ctx,
	}
}

// NewPublicDeletePlayerPublicRecordHandlerV1ParamsWithHTTPClient creates a new PublicDeletePlayerPublicRecordHandlerV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicDeletePlayerPublicRecordHandlerV1ParamsWithHTTPClient(client *http.Client) *PublicDeletePlayerPublicRecordHandlerV1Params {
	var ()
	return &PublicDeletePlayerPublicRecordHandlerV1Params{
		HTTPClient: client,
	}
}

/*PublicDeletePlayerPublicRecordHandlerV1Params contains all the parameters to send to the API endpoint
for the public delete player public record handler v1 operation typically these are written to a http.Request
*/
type PublicDeletePlayerPublicRecordHandlerV1Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Key
	  key of record

	*/
	Key string
	/*Namespace
	  namespace of the game, only accept alphabet and numeric

	*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) WithTimeout(timeout time.Duration) *PublicDeletePlayerPublicRecordHandlerV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) WithContext(ctx context.Context) *PublicDeletePlayerPublicRecordHandlerV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) WithHTTPClient(client *http.Client) *PublicDeletePlayerPublicRecordHandlerV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithKey adds the key to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) WithKey(key string) *PublicDeletePlayerPublicRecordHandlerV1Params {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) SetKey(key string) {
	o.Key = key
}

// WithNamespace adds the namespace to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) WithNamespace(namespace string) *PublicDeletePlayerPublicRecordHandlerV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public delete player public record handler v1 params
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *PublicDeletePlayerPublicRecordHandlerV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}
