// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/karixtech/go-karix/models"
)

// NewPatchSubaccountParams creates a new PatchSubaccountParams object
// with the default values initialized.
func NewPatchSubaccountParams() *PatchSubaccountParams {
	var (
		aPIVersionDefault = string("1.0")
	)
	return &PatchSubaccountParams{
		APIVersion: &aPIVersionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSubaccountParamsWithTimeout creates a new PatchSubaccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSubaccountParamsWithTimeout(timeout time.Duration) *PatchSubaccountParams {
	var (
		aPIVersionDefault = string("1.0")
	)
	return &PatchSubaccountParams{
		APIVersion: &aPIVersionDefault,

		timeout: timeout,
	}
}

// NewPatchSubaccountParamsWithContext creates a new PatchSubaccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSubaccountParamsWithContext(ctx context.Context) *PatchSubaccountParams {
	var (
		apiVersionDefault = string("1.0")
	)
	return &PatchSubaccountParams{
		APIVersion: &apiVersionDefault,

		Context: ctx,
	}
}

// NewPatchSubaccountParamsWithHTTPClient creates a new PatchSubaccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSubaccountParamsWithHTTPClient(client *http.Client) *PatchSubaccountParams {
	var (
		apiVersionDefault = string("1.0")
	)
	return &PatchSubaccountParams{
		APIVersion: &apiVersionDefault,
		HTTPClient: client,
	}
}

/*PatchSubaccountParams contains all the parameters to send to the API endpoint
for the patch subaccount operation typically these are written to a http.Request
*/
type PatchSubaccountParams struct {

	/*APIVersion
	  API Version. If not specified your pinned verison is used.

	*/
	APIVersion *string
	/*Subaccount
	  Subaccount object

	*/
	Subaccount *models.EditAccount
	/*UID
	  Alphanumeric ID of the account/subaccount to edit.

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch subaccount params
func (o *PatchSubaccountParams) WithTimeout(timeout time.Duration) *PatchSubaccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch subaccount params
func (o *PatchSubaccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch subaccount params
func (o *PatchSubaccountParams) WithContext(ctx context.Context) *PatchSubaccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch subaccount params
func (o *PatchSubaccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch subaccount params
func (o *PatchSubaccountParams) WithHTTPClient(client *http.Client) *PatchSubaccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch subaccount params
func (o *PatchSubaccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the patch subaccount params
func (o *PatchSubaccountParams) WithAPIVersion(aPIVersion *string) *PatchSubaccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch subaccount params
func (o *PatchSubaccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithSubaccount adds the subaccount to the patch subaccount params
func (o *PatchSubaccountParams) WithSubaccount(subaccount *models.EditAccount) *PatchSubaccountParams {
	o.SetSubaccount(subaccount)
	return o
}

// SetSubaccount adds the subaccount to the patch subaccount params
func (o *PatchSubaccountParams) SetSubaccount(subaccount *models.EditAccount) {
	o.Subaccount = subaccount
}

// WithUID adds the uid to the patch subaccount params
func (o *PatchSubaccountParams) WithUID(uid string) *PatchSubaccountParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the patch subaccount params
func (o *PatchSubaccountParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSubaccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// header param api-version
		if err := r.SetHeaderParam("api-version", *o.APIVersion); err != nil {
			return err
		}

	}

	if o.Subaccount != nil {
		if err := r.SetBodyParam(o.Subaccount); err != nil {
			return err
		}
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
