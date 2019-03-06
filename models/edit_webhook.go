// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EditWebhook edit webhook
// swagger:model EditWebhook
type EditWebhook struct {

	// Display name of the webhook
	Name string `json:"name,omitempty"`

	// HTTP method to use for fallback notification url
	SmsNotificationFallbackMethod *string `json:"sms_notification_fallback_method,omitempty"`

	// In case the service for `sms_notification_url` is down Karix will hit the
	// fallback url with the incoming message details
	//
	SmsNotificationFallbackURL *string `json:"sms_notification_fallback_url,omitempty"`

	// HTTP method to use for notifying API url in case of inbound message
	SmsNotificationMethod string `json:"sms_notification_method,omitempty"`

	// API url to notify in case of inbound message
	SmsNotificationURL string `json:"sms_notification_url,omitempty"`
}

// Validate validates this edit webhook
func (m *EditWebhook) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EditWebhook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditWebhook) UnmarshalBinary(b []byte) error {
	var res EditWebhook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
