// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package public_content_legacy

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

// NewUpdateScreenshotsParams creates a new UpdateScreenshotsParams object
// with the default values initialized.
func NewUpdateScreenshotsParams() *UpdateScreenshotsParams {
	var ()
	return &UpdateScreenshotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateScreenshotsParamsWithTimeout creates a new UpdateScreenshotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateScreenshotsParamsWithTimeout(timeout time.Duration) *UpdateScreenshotsParams {
	var ()
	return &UpdateScreenshotsParams{

		timeout: timeout,
	}
}

// NewUpdateScreenshotsParamsWithContext creates a new UpdateScreenshotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateScreenshotsParamsWithContext(ctx context.Context) *UpdateScreenshotsParams {
	var ()
	return &UpdateScreenshotsParams{

		Context: ctx,
	}
}

// NewUpdateScreenshotsParamsWithHTTPClient creates a new UpdateScreenshotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateScreenshotsParamsWithHTTPClient(client *http.Client) *UpdateScreenshotsParams {
	var ()
	return &UpdateScreenshotsParams{
		HTTPClient: client,
	}
}

/*UpdateScreenshotsParams contains all the parameters to send to the API endpoint
for the update screenshots operation typically these are written to a http.Request
*/
type UpdateScreenshotsParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body *ugcclientmodels.ModelsUpdateScreenshotRequest
	/*ContentID
	  content ID

	*/
	ContentID string
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

// WithTimeout adds the timeout to the update screenshots params
func (o *UpdateScreenshotsParams) WithTimeout(timeout time.Duration) *UpdateScreenshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update screenshots params
func (o *UpdateScreenshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update screenshots params
func (o *UpdateScreenshotsParams) WithContext(ctx context.Context) *UpdateScreenshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update screenshots params
func (o *UpdateScreenshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the update screenshots params
func (o *UpdateScreenshotsParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the update screenshots params
func (o *UpdateScreenshotsParams) WithHTTPClient(client *http.Client) *UpdateScreenshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update screenshots params
func (o *UpdateScreenshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the update screenshots params
func (o *UpdateScreenshotsParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *UpdateScreenshotsParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithBody adds the body to the update screenshots params
func (o *UpdateScreenshotsParams) WithBody(body *ugcclientmodels.ModelsUpdateScreenshotRequest) *UpdateScreenshotsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update screenshots params
func (o *UpdateScreenshotsParams) SetBody(body *ugcclientmodels.ModelsUpdateScreenshotRequest) {
	o.Body = body
}

// WithContentID adds the contentID to the update screenshots params
func (o *UpdateScreenshotsParams) WithContentID(contentID string) *UpdateScreenshotsParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the update screenshots params
func (o *UpdateScreenshotsParams) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithNamespace adds the namespace to the update screenshots params
func (o *UpdateScreenshotsParams) WithNamespace(namespace string) *UpdateScreenshotsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update screenshots params
func (o *UpdateScreenshotsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the update screenshots params
func (o *UpdateScreenshotsParams) WithUserID(userID string) *UpdateScreenshotsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update screenshots params
func (o *UpdateScreenshotsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScreenshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param contentId
	if err := r.SetPathParam("contentId", o.ContentID); err != nil {
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
