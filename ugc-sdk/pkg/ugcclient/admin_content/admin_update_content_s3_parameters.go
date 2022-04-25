// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package admin_content

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

	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclientmodels"
)

// NewAdminUpdateContentS3Params creates a new AdminUpdateContentS3Params object
// with the default values initialized.
func NewAdminUpdateContentS3Params() *AdminUpdateContentS3Params {
	var ()
	return &AdminUpdateContentS3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminUpdateContentS3ParamsWithTimeout creates a new AdminUpdateContentS3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminUpdateContentS3ParamsWithTimeout(timeout time.Duration) *AdminUpdateContentS3Params {
	var ()
	return &AdminUpdateContentS3Params{

		timeout: timeout,
	}
}

// NewAdminUpdateContentS3ParamsWithContext creates a new AdminUpdateContentS3Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminUpdateContentS3ParamsWithContext(ctx context.Context) *AdminUpdateContentS3Params {
	var ()
	return &AdminUpdateContentS3Params{

		Context: ctx,
	}
}

// NewAdminUpdateContentS3ParamsWithHTTPClient creates a new AdminUpdateContentS3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminUpdateContentS3ParamsWithHTTPClient(client *http.Client) *AdminUpdateContentS3Params {
	var ()
	return &AdminUpdateContentS3Params{
		HTTPClient: client,
	}
}

/*AdminUpdateContentS3Params contains all the parameters to send to the API endpoint
for the admin update content s3 operation typically these are written to a http.Request
*/
type AdminUpdateContentS3Params struct {

	/*Body*/
	Body *ugcclientmodels.ModelsCreateContentRequestS3
	/*ChannelID
	  channel ID

	*/
	ChannelID string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin update content s3 params
func (o *AdminUpdateContentS3Params) WithTimeout(timeout time.Duration) *AdminUpdateContentS3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin update content s3 params
func (o *AdminUpdateContentS3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin update content s3 params
func (o *AdminUpdateContentS3Params) WithContext(ctx context.Context) *AdminUpdateContentS3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin update content s3 params
func (o *AdminUpdateContentS3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin update content s3 params
func (o *AdminUpdateContentS3Params) WithHTTPClient(client *http.Client) *AdminUpdateContentS3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin update content s3 params
func (o *AdminUpdateContentS3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin update content s3 params
func (o *AdminUpdateContentS3Params) WithBody(body *ugcclientmodels.ModelsCreateContentRequestS3) *AdminUpdateContentS3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin update content s3 params
func (o *AdminUpdateContentS3Params) SetBody(body *ugcclientmodels.ModelsCreateContentRequestS3) {
	o.Body = body
}

// WithChannelID adds the channelID to the admin update content s3 params
func (o *AdminUpdateContentS3Params) WithChannelID(channelID string) *AdminUpdateContentS3Params {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the admin update content s3 params
func (o *AdminUpdateContentS3Params) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WithContentID adds the contentID to the admin update content s3 params
func (o *AdminUpdateContentS3Params) WithContentID(contentID string) *AdminUpdateContentS3Params {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the admin update content s3 params
func (o *AdminUpdateContentS3Params) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithNamespace adds the namespace to the admin update content s3 params
func (o *AdminUpdateContentS3Params) WithNamespace(namespace string) *AdminUpdateContentS3Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin update content s3 params
func (o *AdminUpdateContentS3Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the admin update content s3 params
func (o *AdminUpdateContentS3Params) WithUserID(userID string) *AdminUpdateContentS3Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin update content s3 params
func (o *AdminUpdateContentS3Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminUpdateContentS3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param channelId
	if err := r.SetPathParam("channelId", o.ChannelID); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
