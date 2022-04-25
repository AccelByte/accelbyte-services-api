// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package slot

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

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// NewPublicUpdateUserNamespaceSlotMetadataParams creates a new PublicUpdateUserNamespaceSlotMetadataParams object
// with the default values initialized.
func NewPublicUpdateUserNamespaceSlotMetadataParams() *PublicUpdateUserNamespaceSlotMetadataParams {
	var ()
	return &PublicUpdateUserNamespaceSlotMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicUpdateUserNamespaceSlotMetadataParamsWithTimeout creates a new PublicUpdateUserNamespaceSlotMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicUpdateUserNamespaceSlotMetadataParamsWithTimeout(timeout time.Duration) *PublicUpdateUserNamespaceSlotMetadataParams {
	var ()
	return &PublicUpdateUserNamespaceSlotMetadataParams{

		timeout: timeout,
	}
}

// NewPublicUpdateUserNamespaceSlotMetadataParamsWithContext creates a new PublicUpdateUserNamespaceSlotMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicUpdateUserNamespaceSlotMetadataParamsWithContext(ctx context.Context) *PublicUpdateUserNamespaceSlotMetadataParams {
	var ()
	return &PublicUpdateUserNamespaceSlotMetadataParams{

		Context: ctx,
	}
}

// NewPublicUpdateUserNamespaceSlotMetadataParamsWithHTTPClient creates a new PublicUpdateUserNamespaceSlotMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicUpdateUserNamespaceSlotMetadataParamsWithHTTPClient(client *http.Client) *PublicUpdateUserNamespaceSlotMetadataParams {
	var ()
	return &PublicUpdateUserNamespaceSlotMetadataParams{
		HTTPClient: client,
	}
}

/*PublicUpdateUserNamespaceSlotMetadataParams contains all the parameters to send to the API endpoint
for the public update user namespace slot metadata operation typically these are written to a http.Request
*/
type PublicUpdateUserNamespaceSlotMetadataParams struct {

	/*Body*/
	Body *socialclientmodels.SlotMetadataUpdate
	/*Namespace
	  Namespace ID

	*/
	Namespace string
	/*SlotID
	  Slot ID

	*/
	SlotID string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) WithTimeout(timeout time.Duration) *PublicUpdateUserNamespaceSlotMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) WithContext(ctx context.Context) *PublicUpdateUserNamespaceSlotMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) WithHTTPClient(client *http.Client) *PublicUpdateUserNamespaceSlotMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) WithBody(body *socialclientmodels.SlotMetadataUpdate) *PublicUpdateUserNamespaceSlotMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) SetBody(body *socialclientmodels.SlotMetadataUpdate) {
	o.Body = body
}

// WithNamespace adds the namespace to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) WithNamespace(namespace string) *PublicUpdateUserNamespaceSlotMetadataParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSlotID adds the slotID to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) WithSlotID(slotID string) *PublicUpdateUserNamespaceSlotMetadataParams {
	o.SetSlotID(slotID)
	return o
}

// SetSlotID adds the slotId to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) SetSlotID(slotID string) {
	o.SlotID = slotID
}

// WithUserID adds the userID to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) WithUserID(userID string) *PublicUpdateUserNamespaceSlotMetadataParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public update user namespace slot metadata params
func (o *PublicUpdateUserNamespaceSlotMetadataParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicUpdateUserNamespaceSlotMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param slotId
	if err := r.SetPathParam("slotId", o.SlotID); err != nil {
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
