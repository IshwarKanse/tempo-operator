// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// SignatureAlgorithm signature algorithm
//
// swagger:model SignatureAlgorithm
type SignatureAlgorithm int64

// Validate validates this signature algorithm
func (m SignatureAlgorithm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this signature algorithm based on context it is used
func (m SignatureAlgorithm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
