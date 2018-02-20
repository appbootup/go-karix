// Code generated by go-swagger; DO NOT EDIT.

package number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

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
	Payload *models.GetNumberOKBody
}

func (o *GetNumberOK) Error() string {
	return fmt.Sprintf("[GET /number/][%d] getNumberOK  %+v", 200, o.Payload)
}

func (o *GetNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetNumberOKBody)

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
	Payload *models.GetNumberForbiddenBody
}

func (o *GetNumberForbidden) Error() string {
	return fmt.Sprintf("[GET /number/][%d] getNumberForbidden  %+v", 403, o.Payload)
}

func (o *GetNumberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetNumberForbiddenBody)

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
	Payload *models.GetNumberInternalServerErrorBody
}

func (o *GetNumberInternalServerError) Error() string {
	return fmt.Sprintf("[GET /number/][%d] getNumberInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNumberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetNumberInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
