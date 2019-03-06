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

// DeleteWebhookByIDReader is a Reader for the DeleteWebhookByID structure.
type DeleteWebhookByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWebhookByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteWebhookByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteWebhookByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteWebhookByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteWebhookByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteWebhookByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteWebhookByIDNoContent creates a DeleteWebhookByIDNoContent with default headers values
func NewDeleteWebhookByIDNoContent() *DeleteWebhookByIDNoContent {
	return &DeleteWebhookByIDNoContent{}
}

/*DeleteWebhookByIDNoContent handles this case with default header values.

Deleted
*/
type DeleteWebhookByIDNoContent struct {
}

func (o *DeleteWebhookByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /webhook/{uid}/][%d] deleteWebhookByIdNoContent ", 204)
}

func (o *DeleteWebhookByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWebhookByIDUnauthorized creates a DeleteWebhookByIDUnauthorized with default headers values
func NewDeleteWebhookByIDUnauthorized() *DeleteWebhookByIDUnauthorized {
	return &DeleteWebhookByIDUnauthorized{}
}

/*DeleteWebhookByIDUnauthorized handles this case with default header values.

Authentication information is missing or invalid
*/
type DeleteWebhookByIDUnauthorized struct {
	WWWAuthenticate string
}

func (o *DeleteWebhookByIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /webhook/{uid}/][%d] deleteWebhookByIdUnauthorized ", 401)
}

func (o *DeleteWebhookByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header WWW_Authenticate
	o.WWWAuthenticate = response.GetHeader("WWW_Authenticate")

	return nil
}

// NewDeleteWebhookByIDForbidden creates a DeleteWebhookByIDForbidden with default headers values
func NewDeleteWebhookByIDForbidden() *DeleteWebhookByIDForbidden {
	return &DeleteWebhookByIDForbidden{}
}

/*DeleteWebhookByIDForbidden handles this case with default header values.

User not authorized or blocked
*/
type DeleteWebhookByIDForbidden struct {
	Payload *DeleteWebhookByIDForbiddenBody
}

func (o *DeleteWebhookByIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /webhook/{uid}/][%d] deleteWebhookByIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWebhookByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteWebhookByIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebhookByIDNotFound creates a DeleteWebhookByIDNotFound with default headers values
func NewDeleteWebhookByIDNotFound() *DeleteWebhookByIDNotFound {
	return &DeleteWebhookByIDNotFound{}
}

/*DeleteWebhookByIDNotFound handles this case with default header values.

Resource not found
*/
type DeleteWebhookByIDNotFound struct {
	Payload *DeleteWebhookByIDNotFoundBody
}

func (o *DeleteWebhookByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /webhook/{uid}/][%d] deleteWebhookByIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWebhookByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteWebhookByIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebhookByIDInternalServerError creates a DeleteWebhookByIDInternalServerError with default headers values
func NewDeleteWebhookByIDInternalServerError() *DeleteWebhookByIDInternalServerError {
	return &DeleteWebhookByIDInternalServerError{}
}

/*DeleteWebhookByIDInternalServerError handles this case with default header values.

Error
*/
type DeleteWebhookByIDInternalServerError struct {
	Payload *DeleteWebhookByIDInternalServerErrorBody
}

func (o *DeleteWebhookByIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /webhook/{uid}/][%d] deleteWebhookByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWebhookByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteWebhookByIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteWebhookByIDForbiddenBody UnauthorizedResponse
swagger:model DeleteWebhookByIDForbiddenBody
*/
type DeleteWebhookByIDForbiddenBody struct {

	// error
	Error *DeleteWebhookByIDForbiddenBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this delete webhook by ID forbidden body
func (o *DeleteWebhookByIDForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteWebhookByIDForbiddenBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteWebhookByIdForbidden" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *DeleteWebhookByIDForbiddenBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteWebhookByIdForbidden" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteWebhookByIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteWebhookByIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DeleteWebhookByIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteWebhookByIDForbiddenBodyError delete webhook by ID forbidden body error
swagger:model DeleteWebhookByIDForbiddenBodyError
*/
type DeleteWebhookByIDForbiddenBodyError struct {

	// Forbidden Message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete webhook by ID forbidden body error
func (o *DeleteWebhookByIDForbiddenBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteWebhookByIDForbiddenBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteWebhookByIdForbidden"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteWebhookByIDForbiddenBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteWebhookByIDForbiddenBodyError) UnmarshalBinary(b []byte) error {
	var res DeleteWebhookByIDForbiddenBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteWebhookByIDInternalServerErrorBody ErrorResponse
swagger:model DeleteWebhookByIDInternalServerErrorBody
*/
type DeleteWebhookByIDInternalServerErrorBody struct {

	// error
	Error *DeleteWebhookByIDInternalServerErrorBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this delete webhook by ID internal server error body
func (o *DeleteWebhookByIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteWebhookByIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteWebhookByIdInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *DeleteWebhookByIDInternalServerErrorBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteWebhookByIdInternalServerError" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteWebhookByIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteWebhookByIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeleteWebhookByIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteWebhookByIDInternalServerErrorBodyError delete webhook by ID internal server error body error
swagger:model DeleteWebhookByIDInternalServerErrorBodyError
*/
type DeleteWebhookByIDInternalServerErrorBodyError struct {

	// Error Message
	// Required: true
	Message *string `json:"message"`

	// Parameter the error message is related to
	// `null` is the error is generic
	//
	Param *string `json:"param,omitempty"`
}

// Validate validates this delete webhook by ID internal server error body error
func (o *DeleteWebhookByIDInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteWebhookByIDInternalServerErrorBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteWebhookByIdInternalServerError"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteWebhookByIDInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteWebhookByIDInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res DeleteWebhookByIDInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteWebhookByIDNotFoundBody NotFoundResponse
swagger:model DeleteWebhookByIDNotFoundBody
*/
type DeleteWebhookByIDNotFoundBody struct {

	// error
	Error *DeleteWebhookByIDNotFoundBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this delete webhook by ID not found body
func (o *DeleteWebhookByIDNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteWebhookByIDNotFoundBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteWebhookByIdNotFound" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *DeleteWebhookByIDNotFoundBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteWebhookByIdNotFound" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteWebhookByIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteWebhookByIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteWebhookByIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteWebhookByIDNotFoundBodyError delete webhook by ID not found body error
swagger:model DeleteWebhookByIDNotFoundBodyError
*/
type DeleteWebhookByIDNotFoundBodyError struct {

	// Error Message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete webhook by ID not found body error
func (o *DeleteWebhookByIDNotFoundBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteWebhookByIDNotFoundBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteWebhookByIdNotFound"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteWebhookByIDNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteWebhookByIDNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res DeleteWebhookByIDNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
