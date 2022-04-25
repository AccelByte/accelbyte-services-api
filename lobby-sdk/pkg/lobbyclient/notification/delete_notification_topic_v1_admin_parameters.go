// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

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
)

// NewDeleteNotificationTopicV1AdminParams creates a new DeleteNotificationTopicV1AdminParams object
// with the default values initialized.
func NewDeleteNotificationTopicV1AdminParams() *DeleteNotificationTopicV1AdminParams {
	var ()
	return &DeleteNotificationTopicV1AdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNotificationTopicV1AdminParamsWithTimeout creates a new DeleteNotificationTopicV1AdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNotificationTopicV1AdminParamsWithTimeout(timeout time.Duration) *DeleteNotificationTopicV1AdminParams {
	var ()
	return &DeleteNotificationTopicV1AdminParams{

		timeout: timeout,
	}
}

// NewDeleteNotificationTopicV1AdminParamsWithContext creates a new DeleteNotificationTopicV1AdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNotificationTopicV1AdminParamsWithContext(ctx context.Context) *DeleteNotificationTopicV1AdminParams {
	var ()
	return &DeleteNotificationTopicV1AdminParams{

		Context: ctx,
	}
}

// NewDeleteNotificationTopicV1AdminParamsWithHTTPClient creates a new DeleteNotificationTopicV1AdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNotificationTopicV1AdminParamsWithHTTPClient(client *http.Client) *DeleteNotificationTopicV1AdminParams {
	var ()
	return &DeleteNotificationTopicV1AdminParams{
		HTTPClient: client,
	}
}

/*DeleteNotificationTopicV1AdminParams contains all the parameters to send to the API endpoint
for the delete notification topic v1 admin operation typically these are written to a http.Request
*/
type DeleteNotificationTopicV1AdminParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*TopicName
	  topic name

	*/
	TopicName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete notification topic v1 admin params
func (o *DeleteNotificationTopicV1AdminParams) WithTimeout(timeout time.Duration) *DeleteNotificationTopicV1AdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete notification topic v1 admin params
func (o *DeleteNotificationTopicV1AdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete notification topic v1 admin params
func (o *DeleteNotificationTopicV1AdminParams) WithContext(ctx context.Context) *DeleteNotificationTopicV1AdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete notification topic v1 admin params
func (o *DeleteNotificationTopicV1AdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete notification topic v1 admin params
func (o *DeleteNotificationTopicV1AdminParams) WithHTTPClient(client *http.Client) *DeleteNotificationTopicV1AdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete notification topic v1 admin params
func (o *DeleteNotificationTopicV1AdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete notification topic v1 admin params
func (o *DeleteNotificationTopicV1AdminParams) WithNamespace(namespace string) *DeleteNotificationTopicV1AdminParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete notification topic v1 admin params
func (o *DeleteNotificationTopicV1AdminParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithTopicName adds the topicName to the delete notification topic v1 admin params
func (o *DeleteNotificationTopicV1AdminParams) WithTopicName(topicName string) *DeleteNotificationTopicV1AdminParams {
	o.SetTopicName(topicName)
	return o
}

// SetTopicName adds the topicName to the delete notification topic v1 admin params
func (o *DeleteNotificationTopicV1AdminParams) SetTopicName(topicName string) {
	o.TopicName = topicName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNotificationTopicV1AdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param topicName
	if err := r.SetPathParam("topicName", o.TopicName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
