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

// DeviceSearchHitDTO device search hit DTO
//
// swagger:model DeviceSearchHitDTO
type DeviceSearchHitDTO struct {

	// client Ip
	ClientIP string `json:"clientIp,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// device Id
	DeviceID string `json:"deviceId,omitempty"`

	// last seen at
	// Format: date-time
	LastSeenAt strfmt.DateTime `json:"lastSeenAt,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// user agent
	UserAgent string `json:"userAgent,omitempty"`
}

// Validate validates this device search hit DTO
func (m *DeviceSearchHitDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSeenAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceSearchHitDTO) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceSearchHitDTO) validateLastSeenAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastSeenAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastSeenAt", "body", "date-time", m.LastSeenAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeviceSearchHitDTO) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device search hit DTO based on context it is used
func (m *DeviceSearchHitDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceSearchHitDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceSearchHitDTO) UnmarshalBinary(b []byte) error {
	var res DeviceSearchHitDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
