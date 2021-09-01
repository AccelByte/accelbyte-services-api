// Code generated by go-swagger; DO NOT EDIT.

package notification

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
	"github.com/go-openapi/swag"

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// NewSendMultipleUsersFreeformNotificationV1AdminParams creates a new SendMultipleUsersFreeformNotificationV1AdminParams object
// with the default values initialized.
func NewSendMultipleUsersFreeformNotificationV1AdminParams() *SendMultipleUsersFreeformNotificationV1AdminParams {
	var ()
	return &SendMultipleUsersFreeformNotificationV1AdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendMultipleUsersFreeformNotificationV1AdminParamsWithTimeout creates a new SendMultipleUsersFreeformNotificationV1AdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendMultipleUsersFreeformNotificationV1AdminParamsWithTimeout(timeout time.Duration) *SendMultipleUsersFreeformNotificationV1AdminParams {
	var ()
	return &SendMultipleUsersFreeformNotificationV1AdminParams{

		timeout: timeout,
	}
}

// NewSendMultipleUsersFreeformNotificationV1AdminParamsWithContext creates a new SendMultipleUsersFreeformNotificationV1AdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendMultipleUsersFreeformNotificationV1AdminParamsWithContext(ctx context.Context) *SendMultipleUsersFreeformNotificationV1AdminParams {
	var ()
	return &SendMultipleUsersFreeformNotificationV1AdminParams{

		Context: ctx,
	}
}

// NewSendMultipleUsersFreeformNotificationV1AdminParamsWithHTTPClient creates a new SendMultipleUsersFreeformNotificationV1AdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendMultipleUsersFreeformNotificationV1AdminParamsWithHTTPClient(client *http.Client) *SendMultipleUsersFreeformNotificationV1AdminParams {
	var ()
	return &SendMultipleUsersFreeformNotificationV1AdminParams{
		HTTPClient: client,
	}
}

/*SendMultipleUsersFreeformNotificationV1AdminParams contains all the parameters to send to the API endpoint
for the send multiple users freeform notification v1 admin operation typically these are written to a http.Request
*/
type SendMultipleUsersFreeformNotificationV1AdminParams struct {

	/*Async
	  notification type

	*/
	Async *bool
	/*Body
	  notification content

	*/
	Body *lobbyclientmodels.ModelBulkUsersFreeFormNotificationRequestV1
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) WithTimeout(timeout time.Duration) *SendMultipleUsersFreeformNotificationV1AdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) WithContext(ctx context.Context) *SendMultipleUsersFreeformNotificationV1AdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) WithHTTPClient(client *http.Client) *SendMultipleUsersFreeformNotificationV1AdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsync adds the async to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) WithAsync(async *bool) *SendMultipleUsersFreeformNotificationV1AdminParams {
	o.SetAsync(async)
	return o
}

// SetAsync adds the async to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) SetAsync(async *bool) {
	o.Async = async
}

// WithBody adds the body to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) WithBody(body *lobbyclientmodels.ModelBulkUsersFreeFormNotificationRequestV1) *SendMultipleUsersFreeformNotificationV1AdminParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) SetBody(body *lobbyclientmodels.ModelBulkUsersFreeFormNotificationRequestV1) {
	o.Body = body
}

// WithNamespace adds the namespace to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) WithNamespace(namespace string) *SendMultipleUsersFreeformNotificationV1AdminParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the send multiple users freeform notification v1 admin params
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *SendMultipleUsersFreeformNotificationV1AdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Async != nil {

		// query param async
		var qrAsync bool
		if o.Async != nil {
			qrAsync = *o.Async
		}
		qAsync := swag.FormatBool(qrAsync)
		if qAsync != "" {
			if err := r.SetQueryParam("async", qAsync); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}