// Code generated by go-swagger; DO NOT EDIT.

package number_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSearchNumberParams creates a new SearchNumberParams object
// with the default values initialized.
func NewSearchNumberParams() *SearchNumberParams {
	var (
		aPIVersionDefault = string("1.0")
		countryDefault    = string("US")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &SearchNumberParams{
		APIVersion: aPIVersionDefault,
		Country:    &countryDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchNumberParamsWithTimeout creates a new SearchNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchNumberParamsWithTimeout(timeout time.Duration) *SearchNumberParams {
	var (
		aPIVersionDefault = string("1.0")
		countryDefault    = string("US")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &SearchNumberParams{
		APIVersion: aPIVersionDefault,
		Country:    &countryDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		timeout: timeout,
	}
}

// NewSearchNumberParamsWithContext creates a new SearchNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchNumberParamsWithContext(ctx context.Context) *SearchNumberParams {
	var (
		apiVersionDefault = string("1.0")
		countryDefault    = string("US")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &SearchNumberParams{
		APIVersion: apiVersionDefault,
		Country:    &countryDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		Context: ctx,
	}
}

// NewSearchNumberParamsWithHTTPClient creates a new SearchNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchNumberParamsWithHTTPClient(client *http.Client) *SearchNumberParams {
	var (
		apiVersionDefault = string("1.0")
		countryDefault    = string("US")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &SearchNumberParams{
		APIVersion: apiVersionDefault,
		Country:    &countryDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*SearchNumberParams contains all the parameters to send to the API endpoint
for the search number operation typically these are written to a http.Request
*/
type SearchNumberParams struct {

	/*APIVersion
	  API Version. If not specified your pinned verison is used.

	*/
	APIVersion string
	/*Contains
	  Filter by numbers which contain this value

	*/
	Contains *string
	/*Country
	  Filter by country ISO. Only one country can be filtered at a time.
	If no country filter is provided then results for United States are returned by default.


	*/
	Country *string
	/*Limit
	  The numbers of items to return.

	*/
	Limit *int64
	/*NumberType
	  Filter by number type: fixed, mobile, tollfree

	*/
	NumberType []string
	/*Offset
	  The number of items to skip before starting to collect the result set.

	*/
	Offset *int64
	/*Prefix
	  Filter by numbers with this prefix after country code

	*/
	Prefix *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search number params
func (o *SearchNumberParams) WithTimeout(timeout time.Duration) *SearchNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search number params
func (o *SearchNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search number params
func (o *SearchNumberParams) WithContext(ctx context.Context) *SearchNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search number params
func (o *SearchNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search number params
func (o *SearchNumberParams) WithHTTPClient(client *http.Client) *SearchNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search number params
func (o *SearchNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the search number params
func (o *SearchNumberParams) WithAPIVersion(aPIVersion string) *SearchNumberParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the search number params
func (o *SearchNumberParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithContains adds the contains to the search number params
func (o *SearchNumberParams) WithContains(contains *string) *SearchNumberParams {
	o.SetContains(contains)
	return o
}

// SetContains adds the contains to the search number params
func (o *SearchNumberParams) SetContains(contains *string) {
	o.Contains = contains
}

// WithCountry adds the country to the search number params
func (o *SearchNumberParams) WithCountry(country *string) *SearchNumberParams {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the search number params
func (o *SearchNumberParams) SetCountry(country *string) {
	o.Country = country
}

// WithLimit adds the limit to the search number params
func (o *SearchNumberParams) WithLimit(limit *int64) *SearchNumberParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search number params
func (o *SearchNumberParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNumberType adds the numberType to the search number params
func (o *SearchNumberParams) WithNumberType(numberType []string) *SearchNumberParams {
	o.SetNumberType(numberType)
	return o
}

// SetNumberType adds the numberType to the search number params
func (o *SearchNumberParams) SetNumberType(numberType []string) {
	o.NumberType = numberType
}

// WithOffset adds the offset to the search number params
func (o *SearchNumberParams) WithOffset(offset *int64) *SearchNumberParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search number params
func (o *SearchNumberParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPrefix adds the prefix to the search number params
func (o *SearchNumberParams) WithPrefix(prefix *string) *SearchNumberParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the search number params
func (o *SearchNumberParams) SetPrefix(prefix *string) {
	o.Prefix = prefix
}

// WriteToRequest writes these params to a swagger request
func (o *SearchNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param api-version
	if err := r.SetHeaderParam("api-version", o.APIVersion); err != nil {
		return err
	}

	if o.Contains != nil {

		// query param contains
		var qrContains string
		if o.Contains != nil {
			qrContains = *o.Contains
		}
		qContains := qrContains
		if qContains != "" {
			if err := r.SetQueryParam("contains", qContains); err != nil {
				return err
			}
		}

	}

	if o.Country != nil {

		// query param country
		var qrCountry string
		if o.Country != nil {
			qrCountry = *o.Country
		}
		qCountry := qrCountry
		if qCountry != "" {
			if err := r.SetQueryParam("country", qCountry); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesNumberType := o.NumberType

	joinedNumberType := swag.JoinByFormat(valuesNumberType, "csv")
	// query array param number_type
	if err := r.SetQueryParam("number_type", joinedNumberType...); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Prefix != nil {

		// query param prefix
		var qrPrefix string
		if o.Prefix != nil {
			qrPrefix = *o.Prefix
		}
		qPrefix := qrPrefix
		if qPrefix != "" {
			if err := r.SetQueryParam("prefix", qPrefix); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}