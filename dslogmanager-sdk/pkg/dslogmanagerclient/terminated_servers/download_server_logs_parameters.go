// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package terminated_servers

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

// NewDownloadServerLogsParams creates a new DownloadServerLogsParams object
// with the default values initialized.
func NewDownloadServerLogsParams() *DownloadServerLogsParams {
	var ()
	return &DownloadServerLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadServerLogsParamsWithTimeout creates a new DownloadServerLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadServerLogsParamsWithTimeout(timeout time.Duration) *DownloadServerLogsParams {
	var ()
	return &DownloadServerLogsParams{

		timeout: timeout,
	}
}

// NewDownloadServerLogsParamsWithContext creates a new DownloadServerLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadServerLogsParamsWithContext(ctx context.Context) *DownloadServerLogsParams {
	var ()
	return &DownloadServerLogsParams{

		Context: ctx,
	}
}

// NewDownloadServerLogsParamsWithHTTPClient creates a new DownloadServerLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadServerLogsParamsWithHTTPClient(client *http.Client) *DownloadServerLogsParams {
	var ()
	return &DownloadServerLogsParams{
		HTTPClient: client,
	}
}

/*DownloadServerLogsParams contains all the parameters to send to the API endpoint
for the download server logs operation typically these are written to a http.Request
*/
type DownloadServerLogsParams struct {

	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*PodName
	  name of the DS pod

	*/
	PodName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download server logs params
func (o *DownloadServerLogsParams) WithTimeout(timeout time.Duration) *DownloadServerLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download server logs params
func (o *DownloadServerLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download server logs params
func (o *DownloadServerLogsParams) WithContext(ctx context.Context) *DownloadServerLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download server logs params
func (o *DownloadServerLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download server logs params
func (o *DownloadServerLogsParams) WithHTTPClient(client *http.Client) *DownloadServerLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download server logs params
func (o *DownloadServerLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the download server logs params
func (o *DownloadServerLogsParams) WithNamespace(namespace string) *DownloadServerLogsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the download server logs params
func (o *DownloadServerLogsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPodName adds the podName to the download server logs params
func (o *DownloadServerLogsParams) WithPodName(podName string) *DownloadServerLogsParams {
	o.SetPodName(podName)
	return o
}

// SetPodName adds the podName to the download server logs params
func (o *DownloadServerLogsParams) SetPodName(podName string) {
	o.PodName = podName
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadServerLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param podName
	if err := r.SetPathParam("podName", o.PodName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
