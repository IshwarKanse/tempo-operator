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

// Matcher Matcher matcher
//
// swagger:model matcher
type Matcher struct {

	// is equal
	IsEqual bool `json:"isEqual,omitempty"`

	// is regex
	// Required: true
	IsRegex *bool `json:"isRegex"`

	// name
	// Required: true
	Name *string `json:"name"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this matcher
func (m *Matcher) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsRegex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Matcher) validateIsRegex(formats strfmt.Registry) error {

	if err := validate.Required("isRegex", "body", m.IsRegex); err != nil {
		return err
	}

	return nil
}

func (m *Matcher) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Matcher) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this matcher based on context it is used
func (m *Matcher) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Matcher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Matcher) UnmarshalBinary(b []byte) error {
	var res Matcher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
