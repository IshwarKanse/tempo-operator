// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Assignments assignments
//
// swagger:model Assignments
type Assignments struct {

	// built in roles
	BuiltInRoles bool `json:"builtInRoles,omitempty"`

	// service accounts
	ServiceAccounts bool `json:"serviceAccounts,omitempty"`

	// teams
	Teams bool `json:"teams,omitempty"`

	// users
	Users bool `json:"users,omitempty"`
}

// Validate validates this assignments
func (m *Assignments) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this assignments based on context it is used
func (m *Assignments) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Assignments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Assignments) UnmarshalBinary(b []byte) error {
	var res Assignments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
