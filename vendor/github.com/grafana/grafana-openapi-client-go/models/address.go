// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Address address
//
// swagger:model Address
type Address struct {

	// address1
	Address1 string `json:"address1,omitempty"`

	// address2
	Address2 string `json:"address2,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// zip code
	ZipCode string `json:"zipCode,omitempty"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this address based on context it is used
func (m *Address) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Address) UnmarshalBinary(b []byte) error {
	var res Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
