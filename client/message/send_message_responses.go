// Code generated by go-swagger; DO NOT EDIT.

package message

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

// SendMessageReader is a Reader for the SendMessage structure.
type SendMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewSendMessageAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSendMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewSendMessagePaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSendMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendMessageAccepted creates a SendMessageAccepted with default headers values
func NewSendMessageAccepted() *SendMessageAccepted {
	return &SendMessageAccepted{}
}

/*SendMessageAccepted handles this case with default header values.

A list of message records
*/
type SendMessageAccepted struct {
	Payload *SendMessageAcceptedBody
}

func (o *SendMessageAccepted) Error() string {
	return fmt.Sprintf("[POST /message/][%d] sendMessageAccepted  %+v", 202, o.Payload)
}

func (o *SendMessageAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SendMessageAcceptedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendMessageBadRequest creates a SendMessageBadRequest with default headers values
func NewSendMessageBadRequest() *SendMessageBadRequest {
	return &SendMessageBadRequest{}
}

/*SendMessageBadRequest handles this case with default header values.

Error
*/
type SendMessageBadRequest struct {
	Payload *SendMessageBadRequestBody
}

func (o *SendMessageBadRequest) Error() string {
	return fmt.Sprintf("[POST /message/][%d] sendMessageBadRequest  %+v", 400, o.Payload)
}

func (o *SendMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SendMessageBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendMessageUnauthorized creates a SendMessageUnauthorized with default headers values
func NewSendMessageUnauthorized() *SendMessageUnauthorized {
	return &SendMessageUnauthorized{}
}

/*SendMessageUnauthorized handles this case with default header values.

Authentication information is missing or invalid
*/
type SendMessageUnauthorized struct {
	WWWAuthenticate string
}

func (o *SendMessageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /message/][%d] sendMessageUnauthorized ", 401)
}

func (o *SendMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header WWW_Authenticate
	o.WWWAuthenticate = response.GetHeader("WWW_Authenticate")

	return nil
}

// NewSendMessagePaymentRequired creates a SendMessagePaymentRequired with default headers values
func NewSendMessagePaymentRequired() *SendMessagePaymentRequired {
	return &SendMessagePaymentRequired{}
}

/*SendMessagePaymentRequired handles this case with default header values.

Insufficient Balance Error
*/
type SendMessagePaymentRequired struct {
	Payload *SendMessagePaymentRequiredBody
}

func (o *SendMessagePaymentRequired) Error() string {
	return fmt.Sprintf("[POST /message/][%d] sendMessagePaymentRequired  %+v", 402, o.Payload)
}

func (o *SendMessagePaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SendMessagePaymentRequiredBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendMessageForbidden creates a SendMessageForbidden with default headers values
func NewSendMessageForbidden() *SendMessageForbidden {
	return &SendMessageForbidden{}
}

/*SendMessageForbidden handles this case with default header values.

User not authorized or blocked
*/
type SendMessageForbidden struct {
	Payload *SendMessageForbiddenBody
}

func (o *SendMessageForbidden) Error() string {
	return fmt.Sprintf("[POST /message/][%d] sendMessageForbidden  %+v", 403, o.Payload)
}

func (o *SendMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SendMessageForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendMessageInternalServerError creates a SendMessageInternalServerError with default headers values
func NewSendMessageInternalServerError() *SendMessageInternalServerError {
	return &SendMessageInternalServerError{}
}

/*SendMessageInternalServerError handles this case with default header values.

Error
*/
type SendMessageInternalServerError struct {
	Payload *SendMessageInternalServerErrorBody
}

func (o *SendMessageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /message/][%d] sendMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *SendMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SendMessageInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SendMessageAcceptedBody MessageCreatedResponse
swagger:model SendMessageAcceptedBody
*/
type SendMessageAcceptedBody struct {

	// meta
	Meta *models.MetaResponseWithBalance `json:"meta,omitempty"`

	// objects
	Objects []*models.Message `json:"objects"`
}

// Validate validates this send message accepted body
func (o *SendMessageAcceptedBody) Validate(formats strfmt.Registry) error {
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

func (o *SendMessageAcceptedBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sendMessageAccepted" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

func (o *SendMessageAcceptedBody) validateObjects(formats strfmt.Registry) error {

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
					return ve.ValidateName("sendMessageAccepted" + "." + "objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendMessageAcceptedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendMessageAcceptedBody) UnmarshalBinary(b []byte) error {
	var res SendMessageAcceptedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SendMessageBadRequestBody ErrorResponse
swagger:model SendMessageBadRequestBody
*/
type SendMessageBadRequestBody struct {

	// error
	Error *SendMessageBadRequestBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this send message bad request body
func (o *SendMessageBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *SendMessageBadRequestBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sendMessageBadRequest" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *SendMessageBadRequestBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sendMessageBadRequest" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendMessageBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendMessageBadRequestBody) UnmarshalBinary(b []byte) error {
	var res SendMessageBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SendMessageBadRequestBodyError send message bad request body error
swagger:model SendMessageBadRequestBodyError
*/
type SendMessageBadRequestBodyError struct {

	// Error Message
	// Required: true
	Message *string `json:"message"`

	// Parameter the error message is related to
	// `null` is the error is generic
	//
	Param *string `json:"param,omitempty"`
}

// Validate validates this send message bad request body error
func (o *SendMessageBadRequestBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SendMessageBadRequestBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("sendMessageBadRequest"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendMessageBadRequestBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendMessageBadRequestBodyError) UnmarshalBinary(b []byte) error {
	var res SendMessageBadRequestBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SendMessageForbiddenBody UnauthorizedResponse
swagger:model SendMessageForbiddenBody
*/
type SendMessageForbiddenBody struct {

	// error
	Error *SendMessageForbiddenBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this send message forbidden body
func (o *SendMessageForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *SendMessageForbiddenBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sendMessageForbidden" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *SendMessageForbiddenBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sendMessageForbidden" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendMessageForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendMessageForbiddenBody) UnmarshalBinary(b []byte) error {
	var res SendMessageForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SendMessageForbiddenBodyError send message forbidden body error
swagger:model SendMessageForbiddenBodyError
*/
type SendMessageForbiddenBodyError struct {

	// Forbidden Message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this send message forbidden body error
func (o *SendMessageForbiddenBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SendMessageForbiddenBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("sendMessageForbidden"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendMessageForbiddenBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendMessageForbiddenBodyError) UnmarshalBinary(b []byte) error {
	var res SendMessageForbiddenBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SendMessageInternalServerErrorBody ErrorResponse
swagger:model SendMessageInternalServerErrorBody
*/
type SendMessageInternalServerErrorBody struct {

	// error
	Error *SendMessageInternalServerErrorBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this send message internal server error body
func (o *SendMessageInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

func (o *SendMessageInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sendMessageInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *SendMessageInternalServerErrorBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sendMessageInternalServerError" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendMessageInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendMessageInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SendMessageInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SendMessageInternalServerErrorBodyError send message internal server error body error
swagger:model SendMessageInternalServerErrorBodyError
*/
type SendMessageInternalServerErrorBodyError struct {

	// Error Message
	// Required: true
	Message *string `json:"message"`

	// Parameter the error message is related to
	// `null` is the error is generic
	//
	Param *string `json:"param,omitempty"`
}

// Validate validates this send message internal server error body error
func (o *SendMessageInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SendMessageInternalServerErrorBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("sendMessageInternalServerError"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendMessageInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendMessageInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res SendMessageInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SendMessagePaymentRequiredBody InsufficientBalanceResponse
swagger:model SendMessagePaymentRequiredBody
*/
type SendMessagePaymentRequiredBody struct {

	// error
	Error *SendMessagePaymentRequiredBodyError `json:"error,omitempty"`

	// meta
	Meta *models.MetaResponse `json:"meta,omitempty"`
}

// Validate validates this send message payment required body
func (o *SendMessagePaymentRequiredBody) Validate(formats strfmt.Registry) error {
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

func (o *SendMessagePaymentRequiredBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sendMessagePaymentRequired" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *SendMessagePaymentRequiredBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sendMessagePaymentRequired" + "." + "meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendMessagePaymentRequiredBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendMessagePaymentRequiredBody) UnmarshalBinary(b []byte) error {
	var res SendMessagePaymentRequiredBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SendMessagePaymentRequiredBodyError send message payment required body error
swagger:model SendMessagePaymentRequiredBodyError
*/
type SendMessagePaymentRequiredBodyError struct {

	// Insufficient Balance Message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this send message payment required body error
func (o *SendMessagePaymentRequiredBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SendMessagePaymentRequiredBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("sendMessagePaymentRequired"+"."+"error"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendMessagePaymentRequiredBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendMessagePaymentRequiredBodyError) UnmarshalBinary(b []byte) error {
	var res SendMessagePaymentRequiredBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
