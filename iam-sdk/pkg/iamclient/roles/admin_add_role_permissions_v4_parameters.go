// Code generated by go-swagger; DO NOT EDIT.

package roles

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

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// NewAdminAddRolePermissionsV4Params creates a new AdminAddRolePermissionsV4Params object
// with the default values initialized.
func NewAdminAddRolePermissionsV4Params() *AdminAddRolePermissionsV4Params {
	var ()
	return &AdminAddRolePermissionsV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminAddRolePermissionsV4ParamsWithTimeout creates a new AdminAddRolePermissionsV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminAddRolePermissionsV4ParamsWithTimeout(timeout time.Duration) *AdminAddRolePermissionsV4Params {
	var ()
	return &AdminAddRolePermissionsV4Params{

		timeout: timeout,
	}
}

// NewAdminAddRolePermissionsV4ParamsWithContext creates a new AdminAddRolePermissionsV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminAddRolePermissionsV4ParamsWithContext(ctx context.Context) *AdminAddRolePermissionsV4Params {
	var ()
	return &AdminAddRolePermissionsV4Params{

		Context: ctx,
	}
}

// NewAdminAddRolePermissionsV4ParamsWithHTTPClient creates a new AdminAddRolePermissionsV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminAddRolePermissionsV4ParamsWithHTTPClient(client *http.Client) *AdminAddRolePermissionsV4Params {
	var ()
	return &AdminAddRolePermissionsV4Params{
		HTTPClient: client,
	}
}

/*AdminAddRolePermissionsV4Params contains all the parameters to send to the API endpoint
for the admin add role permissions v4 operation typically these are written to a http.Request
*/
type AdminAddRolePermissionsV4Params struct {

	/*Body*/
	Body *iamclientmodels.AccountcommonPermissionsV3
	/*RoleID
	  Role ID, should follow UUID version 4 without hyphen

	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin add role permissions v4 params
func (o *AdminAddRolePermissionsV4Params) WithTimeout(timeout time.Duration) *AdminAddRolePermissionsV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin add role permissions v4 params
func (o *AdminAddRolePermissionsV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin add role permissions v4 params
func (o *AdminAddRolePermissionsV4Params) WithContext(ctx context.Context) *AdminAddRolePermissionsV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin add role permissions v4 params
func (o *AdminAddRolePermissionsV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin add role permissions v4 params
func (o *AdminAddRolePermissionsV4Params) WithHTTPClient(client *http.Client) *AdminAddRolePermissionsV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin add role permissions v4 params
func (o *AdminAddRolePermissionsV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin add role permissions v4 params
func (o *AdminAddRolePermissionsV4Params) WithBody(body *iamclientmodels.AccountcommonPermissionsV3) *AdminAddRolePermissionsV4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin add role permissions v4 params
func (o *AdminAddRolePermissionsV4Params) SetBody(body *iamclientmodels.AccountcommonPermissionsV3) {
	o.Body = body
}

// WithRoleID adds the roleID to the admin add role permissions v4 params
func (o *AdminAddRolePermissionsV4Params) WithRoleID(roleID string) *AdminAddRolePermissionsV4Params {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the admin add role permissions v4 params
func (o *AdminAddRolePermissionsV4Params) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminAddRolePermissionsV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}