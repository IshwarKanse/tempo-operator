// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteAlertNotificationChannelByUIDOKBody delete alert notification channel by Uid Ok body
//
// swagger:model deleteAlertNotificationChannelByUidOkBody
type DeleteAlertNotificationChannelByUIDOKBody struct {

	// ID Identifier of the deleted notification channel.
	// Example: 65
	// Required: true
	ID *int64 `json:"id"`

	// Message Message of the deleted notificatiton channel.
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this delete alert notification channel by Uid Ok body
func (m *DeleteAlertNotificationChannelByUIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteAlertNotificationChannelByUIDOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DeleteAlertNotificationChannelByUIDOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete alert notification channel by Uid Ok body based on context it is used
func (m *DeleteAlertNotificationChannelByUIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteAlertNotificationChannelByUIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteAlertNotificationChannelByUIDOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteAlertNotificationChannelByUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
