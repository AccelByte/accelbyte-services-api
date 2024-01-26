// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package roles

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// NewAdminUpdateRolePermissionsV4Params creates a new AdminUpdateRolePermissionsV4Params object
// with the default values initialized.
func NewAdminUpdateRolePermissionsV4Params() *AdminUpdateRolePermissionsV4Params {
	var ()
	return &AdminUpdateRolePermissionsV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminUpdateRolePermissionsV4ParamsWithTimeout creates a new AdminUpdateRolePermissionsV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminUpdateRolePermissionsV4ParamsWithTimeout(timeout time.Duration) *AdminUpdateRolePermissionsV4Params {
	var ()
	return &AdminUpdateRolePermissionsV4Params{

		timeout: timeout,
	}
}

// NewAdminUpdateRolePermissionsV4ParamsWithContext creates a new AdminUpdateRolePermissionsV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminUpdateRolePermissionsV4ParamsWithContext(ctx context.Context) *AdminUpdateRolePermissionsV4Params {
	var ()
	return &AdminUpdateRolePermissionsV4Params{

		Context: ctx,
	}
}

// NewAdminUpdateRolePermissionsV4ParamsWithHTTPClient creates a new AdminUpdateRolePermissionsV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminUpdateRolePermissionsV4ParamsWithHTTPClient(client *http.Client) *AdminUpdateRolePermissionsV4Params {
	var ()
	return &AdminUpdateRolePermissionsV4Params{
		HTTPClient: client,
	}
}

/*AdminUpdateRolePermissionsV4Params contains all the parameters to send to the API endpoint
for the admin update role permissions v4 operation typically these are written to a http.Request
*/
type AdminUpdateRolePermissionsV4Params struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*Body*/
	Body *iamclientmodels.AccountcommonPermissionsV3
	/*RoleID
	  Role ID, should follow UUID version 4 without hyphen

	*/
	RoleID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) WithTimeout(timeout time.Duration) *AdminUpdateRolePermissionsV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) WithContext(ctx context.Context) *AdminUpdateRolePermissionsV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) WithHTTPClient(client *http.Client) *AdminUpdateRolePermissionsV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *AdminUpdateRolePermissionsV4Params) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithBody adds the body to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) WithBody(body *iamclientmodels.AccountcommonPermissionsV3) *AdminUpdateRolePermissionsV4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) SetBody(body *iamclientmodels.AccountcommonPermissionsV3) {
	o.Body = body
}

// WithRoleID adds the roleID to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) WithRoleID(roleID string) *AdminUpdateRolePermissionsV4Params {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the admin update role permissions v4 params
func (o *AdminUpdateRolePermissionsV4Params) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminUpdateRolePermissionsV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
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
