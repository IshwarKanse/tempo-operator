// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetHomeDashboardResponseAllOf1 get home dashboard response all of1
//
// swagger:model getHomeDashboardResponseAllOf1
type GetHomeDashboardResponseAllOf1 struct {

	// redirect Uri
	RedirectURI string `json:"redirectUri,omitempty"`
}

// Validate validates this get home dashboard response all of1
func (m *GetHomeDashboardResponseAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get home dashboard response all of1 based on context it is used
func (m *GetHomeDashboardResponseAllOf1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetHomeDashboardResponseAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetHomeDashboardResponseAllOf1) UnmarshalBinary(b []byte) error {
	var res GetHomeDashboardResponseAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}