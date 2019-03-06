// Code generated by go-swagger; DO NOT EDIT.

package number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/karixtech/go-karix/models"
)

// GetNumberReader is a Reader for the GetNumber structure.
type GetNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetNumberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetNumberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetNumberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNumberOK creates a GetNumberOK with default headers values
func NewGetNumberOK() *GetNumberOK {
	return &GetNumberOK{}
}

/*GetNumberOK handles this case with default header values.

A list of Account Number objects
*/
type GetNumberOK struct {
	Payload *GetNumberOKBody
}

func (o *GetNumberOK) Error() string {
	return fmt.Sprintf("[GET /number/][%d] getNumberOK  %+v", 200, o.Payload)
}

func (o *GetNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNumberOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNumberUnauthorized creates a GetNumberUnauthorized with default headers values
func NewGetNumberUnauthorized() *GetNumberUnauthorized {
	return &GetNumberUnauthorized{}
}

/*GetNumberUnauthorized handles this case with default header values.

Authentication information is missing or invalid
*/
type GetNumberUnauthorized struct {
	WWWAuthenticate string
}

func (o *GetNumberUnauthorized) Error() string {
	return fmt.Sprintf("[GET /number/][%d] getNumberUnauthorized ", 401)
}

func (o *GetNumberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header WWW_Authenticate
	o.WWWAuthenticate = response.GetHeader("WWW_Authenticate")

	return nil
}

// NewGetNumberForbidden creates a GetNumberForbidden with default headers values
func NewGetNumberForbidden() *GetNumberForbidden {
	return &GetNumberForbidden{}
}

/*GetNumberForbidden handles this case with default header values.

User not authorized or blocked
*/
type GetNumberForbidden struct {
	Payload *GetNumberForbiddenBody
}

func (o *GetNumberForbidden) Error() string {
	return fmt.Sprintf("[GET /number/][%d] getNumberForbidden  %+v", 403, o.Payload)
}

func (o *GetNumberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNumberForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNumberInternalServerError creates a GetNumberInternalServerError with default headers values
func NewGetNumberInternalServerError() *GetNumberInternalServerError {
	return &GetNumberInternalServerError{}
}

/*GetNumberInternalServerError handles this case with default header values.

Error
*/
type GetNumberInternalServerError struct {
	Payload *GetNumberInternalServerErrorBody
}

func (o *GetNumberInternalServerError) Error() string {
	return fmt.Sprintf("[GET /number/][%d] getNumberInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNumberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNumberInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetNumberForbiddenBody UnauthorizedResponse
swagger:model GetNumberForbiddenBody
*/
type GetNumberForbiddenBody struct {

	// error
	Error *GetNumberForbiddenBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this get number forbidden body
func (o *GetNumberForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNumberForbiddenBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNumberForbidden" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *GetNumberForbiddenBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNumberForbidden" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNumberForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNumberForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetNumberForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNumberForbiddenBodyError get number forbidden body error
swagger:model GetNumberForbiddenBodyError
*/
type GetNumberForbiddenBodyError struct {

	// Forbidden Message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get number forbidden body error
func (o *GetNumberForbiddenBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNumberForbiddenBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getNumberForbidden"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNumberForbiddenBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNumberForbiddenBodyError) UnmarshalBinary(b []byte) error {
	var res GetNumberForbiddenBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNumberInternalServerErrorBody ErrorResponse
swagger:model GetNumberInternalServerErrorBody
*/
type GetNumberInternalServerErrorBody struct {

	// error
	Error *GetNumberInternalServerErrorBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this get number internal server error body
func (o *GetNumberInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNumberInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNumberInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *GetNumberInternalServerErrorBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNumberInternalServerError" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNumberInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNumberInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetNumberInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNumberInternalServerErrorBodyError get number internal server error body error
swagger:model GetNumberInternalServerErrorBodyError
*/
type GetNumberInternalServerErrorBodyError struct {

	// Error Message
	// Required: true
	Message *string `json:"message"`

	// Parameter the error message is related to
	// `null` is the error is generic
	//
	Param *string `json:"param,omitempty"`
}

// Validate validates this get number internal server error body error
func (o *GetNumberInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNumberInternalServerErrorBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getNumberInternalServerError"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNumberInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNumberInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res GetNumberInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNumberOKBody AccountNumberListResponse
swagger:model GetNumberOKBody
*/
type GetNumberOKBody struct {

	// meta
	Meta *models.ArrayMetaResponse `json:"meta,omitempty"`

	// objects
	Objects []*models.AccountNumber `json:"objects"`
}

// Validate validates this get number o k body
func (o *GetNumberOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNumberOKBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNumberOK" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

func (o *GetNumberOKBody) validateObjects(formats strfmt.Registry) error {

	if swag.IsZero(o.Objects) { // not required
		return nil
	}

	for i := 0; i < len(o.Objects); i++ {
		if swag.IsZero(o.Objects[i]) { // not required
			continue
		}

		if o.Objects[i] != nil {
			if err := o.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNumberOK" + "." + "objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNumberOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNumberOKBody) UnmarshalBinary(b []byte) error {
	var res GetNumberOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
