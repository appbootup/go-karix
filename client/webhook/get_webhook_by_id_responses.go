// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/karixtech/go-karix/models"
)

// GetWebhookByIDReader is a Reader for the GetWebhookByID structure.
type GetWebhookByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebhookByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWebhookByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetWebhookByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetWebhookByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetWebhookByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetWebhookByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWebhookByIDOK creates a GetWebhookByIDOK with default headers values
func NewGetWebhookByIDOK() *GetWebhookByIDOK {
	return &GetWebhookByIDOK{}
}

/*GetWebhookByIDOK handles this case with default header values.

Webhook data
*/
type GetWebhookByIDOK struct {
	Payload *GetWebhookByIDOKBody
}

func (o *GetWebhookByIDOK) Error() string {
	return fmt.Sprintf("[GET /webhook/{uid}/][%d] getWebhookByIdOK  %+v", 200, o.Payload)
}

func (o *GetWebhookByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWebhookByIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookByIDUnauthorized creates a GetWebhookByIDUnauthorized with default headers values
func NewGetWebhookByIDUnauthorized() *GetWebhookByIDUnauthorized {
	return &GetWebhookByIDUnauthorized{}
}

/*GetWebhookByIDUnauthorized handles this case with default header values.

Authentication information is missing or invalid
*/
type GetWebhookByIDUnauthorized struct {
	WWWAuthenticate string
}

func (o *GetWebhookByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /webhook/{uid}/][%d] getWebhookByIdUnauthorized ", 401)
}

func (o *GetWebhookByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header WWW_Authenticate
	o.WWWAuthenticate = response.GetHeader("WWW_Authenticate")

	return nil
}

// NewGetWebhookByIDForbidden creates a GetWebhookByIDForbidden with default headers values
func NewGetWebhookByIDForbidden() *GetWebhookByIDForbidden {
	return &GetWebhookByIDForbidden{}
}

/*GetWebhookByIDForbidden handles this case with default header values.

User not authorized or blocked
*/
type GetWebhookByIDForbidden struct {
	Payload *GetWebhookByIDForbiddenBody
}

func (o *GetWebhookByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /webhook/{uid}/][%d] getWebhookByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetWebhookByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWebhookByIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookByIDNotFound creates a GetWebhookByIDNotFound with default headers values
func NewGetWebhookByIDNotFound() *GetWebhookByIDNotFound {
	return &GetWebhookByIDNotFound{}
}

/*GetWebhookByIDNotFound handles this case with default header values.

Resource not found
*/
type GetWebhookByIDNotFound struct {
	Payload *GetWebhookByIDNotFoundBody
}

func (o *GetWebhookByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /webhook/{uid}/][%d] getWebhookByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetWebhookByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWebhookByIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookByIDInternalServerError creates a GetWebhookByIDInternalServerError with default headers values
func NewGetWebhookByIDInternalServerError() *GetWebhookByIDInternalServerError {
	return &GetWebhookByIDInternalServerError{}
}

/*GetWebhookByIDInternalServerError handles this case with default header values.

Error
*/
type GetWebhookByIDInternalServerError struct {
	Payload *GetWebhookByIDInternalServerErrorBody
}

func (o *GetWebhookByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /webhook/{uid}/][%d] getWebhookByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebhookByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWebhookByIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWebhookByIDForbiddenBody UnauthorizedResponse
swagger:model GetWebhookByIDForbiddenBody
*/
type GetWebhookByIDForbiddenBody struct {

	// error
	Error *GetWebhookByIDForbiddenBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this get webhook by ID forbidden body
func (o *GetWebhookByIDForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *GetWebhookByIDForbiddenBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWebhookByIdForbidden" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *GetWebhookByIDForbiddenBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWebhookByIdForbidden" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWebhookByIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWebhookByIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetWebhookByIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWebhookByIDForbiddenBodyError get webhook by ID forbidden body error
swagger:model GetWebhookByIDForbiddenBodyError
*/
type GetWebhookByIDForbiddenBodyError struct {

	// Forbidden Message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get webhook by ID forbidden body error
func (o *GetWebhookByIDForbiddenBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWebhookByIDForbiddenBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getWebhookByIdForbidden"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWebhookByIDForbiddenBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWebhookByIDForbiddenBodyError) UnmarshalBinary(b []byte) error {
	var res GetWebhookByIDForbiddenBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWebhookByIDInternalServerErrorBody ErrorResponse
swagger:model GetWebhookByIDInternalServerErrorBody
*/
type GetWebhookByIDInternalServerErrorBody struct {

	// error
	Error *GetWebhookByIDInternalServerErrorBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this get webhook by ID internal server error body
func (o *GetWebhookByIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *GetWebhookByIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWebhookByIdInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *GetWebhookByIDInternalServerErrorBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWebhookByIdInternalServerError" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWebhookByIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWebhookByIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetWebhookByIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWebhookByIDInternalServerErrorBodyError get webhook by ID internal server error body error
swagger:model GetWebhookByIDInternalServerErrorBodyError
*/
type GetWebhookByIDInternalServerErrorBodyError struct {

	// Error Message
	// Required: true
	Message *string `json:"message"`

	// Parameter the error message is related to
	// `null` is the error is generic
	//
	Param *string `json:"param,omitempty"`
}

// Validate validates this get webhook by ID internal server error body error
func (o *GetWebhookByIDInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWebhookByIDInternalServerErrorBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getWebhookByIdInternalServerError"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWebhookByIDInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWebhookByIDInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res GetWebhookByIDInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWebhookByIDNotFoundBody NotFoundResponse
swagger:model GetWebhookByIDNotFoundBody
*/
type GetWebhookByIDNotFoundBody struct {

	// error
	Error *GetWebhookByIDNotFoundBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this get webhook by ID not found body
func (o *GetWebhookByIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetWebhookByIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWebhookByIdNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *GetWebhookByIDNotFoundBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWebhookByIdNotFound" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWebhookByIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWebhookByIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetWebhookByIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWebhookByIDNotFoundBodyError get webhook by ID not found body error
swagger:model GetWebhookByIDNotFoundBodyError
*/
type GetWebhookByIDNotFoundBodyError struct {

	// Error Message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get webhook by ID not found body error
func (o *GetWebhookByIDNotFoundBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWebhookByIDNotFoundBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getWebhookByIdNotFound"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWebhookByIDNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWebhookByIDNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res GetWebhookByIDNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWebhookByIDOKBody WebhookResponse
swagger:model GetWebhookByIDOKBody
*/
type GetWebhookByIDOKBody struct {

	// data
	Data *models.Webhook `json:"data,omitempty"`

	// meta
	Meta *models.ObjectMetaResponse `json:"meta,omitempty"`
}

// Validate validates this get webhook by ID o k body
func (o *GetWebhookByIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
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

func (o *GetWebhookByIDOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWebhookByIdOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

func (o *GetWebhookByIDOKBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWebhookByIdOK" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWebhookByIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWebhookByIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetWebhookByIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
