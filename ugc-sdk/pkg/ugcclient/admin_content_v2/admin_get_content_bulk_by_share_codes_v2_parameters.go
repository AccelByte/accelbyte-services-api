// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package admin_content_v2

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclientmodels"
)

// NewAdminGetContentBulkByShareCodesV2Params creates a new AdminGetContentBulkByShareCodesV2Params object
// with the default values initialized.
func NewAdminGetContentBulkByShareCodesV2Params() *AdminGetContentBulkByShareCodesV2Params {
	var ()
	return &AdminGetContentBulkByShareCodesV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetContentBulkByShareCodesV2ParamsWithTimeout creates a new AdminGetContentBulkByShareCodesV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetContentBulkByShareCodesV2ParamsWithTimeout(timeout time.Duration) *AdminGetContentBulkByShareCodesV2Params {
	var ()
	return &AdminGetContentBulkByShareCodesV2Params{

		timeout: timeout,
	}
}

// NewAdminGetContentBulkByShareCodesV2ParamsWithContext creates a new AdminGetContentBulkByShareCodesV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetContentBulkByShareCodesV2ParamsWithContext(ctx context.Context) *AdminGetContentBulkByShareCodesV2Params {
	var ()
	return &AdminGetContentBulkByShareCodesV2Params{

		Context: ctx,
	}
}

// NewAdminGetContentBulkByShareCodesV2ParamsWithHTTPClient creates a new AdminGetContentBulkByShareCodesV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetContentBulkByShareCodesV2ParamsWithHTTPClient(client *http.Client) *AdminGetContentBulkByShareCodesV2Params {
	var ()
	return &AdminGetContentBulkByShareCodesV2Params{
		HTTPClient: client,
	}
}

/*AdminGetContentBulkByShareCodesV2Params contains all the parameters to send to the API endpoint
for the admin get content bulk by share codes v2 operation typically these are written to a http.Request
*/
type AdminGetContentBulkByShareCodesV2Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body *ugcclientmodels.ModelsGetContentBulkByShareCodesRequest
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) WithTimeout(timeout time.Duration) *AdminGetContentBulkByShareCodesV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) WithContext(ctx context.Context) *AdminGetContentBulkByShareCodesV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) WithHTTPClient(client *http.Client) *AdminGetContentBulkByShareCodesV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// WithBody adds the body to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) WithBody(body *ugcclientmodels.ModelsGetContentBulkByShareCodesRequest) *AdminGetContentBulkByShareCodesV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) SetBody(body *ugcclientmodels.ModelsGetContentBulkByShareCodesRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) WithNamespace(namespace string) *AdminGetContentBulkByShareCodesV2Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin get content bulk by share codes v2 params
func (o *AdminGetContentBulkByShareCodesV2Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetContentBulkByShareCodesV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
