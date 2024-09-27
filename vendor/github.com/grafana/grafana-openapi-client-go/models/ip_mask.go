// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IPMask An IPMask is a bitmask that can be used to manipulate
// IP addresses for IP addressing and routing.
//
// See type IPNet and func ParseCIDR for details.
//
// swagger:model IPMask
type IPMask []uint8

// Validate validates this IP mask
func (m IPMask) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP mask based on context it is used
func (m IPMask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}