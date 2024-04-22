// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateQueryInQueryHistoryCommand CreateQueryInQueryHistoryCommand is the command for adding query history
//
// swagger:model CreateQueryInQueryHistoryCommand
type CreateQueryInQueryHistoryCommand struct {

	// UID of the data source for which are queries stored.
	// Example: PE1C5CBDA0504A6A3
	DatasourceUID string `json:"datasourceUid,omitempty"`

	// queries
	// Required: true
	Queries JSON `json:"queries"`
}

// Validate validates this create query in query history command
func (m *CreateQueryInQueryHistoryCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateQueryInQueryHistoryCommand) validateQueries(formats strfmt.Registry) error {

	if m.Queries == nil {
		return errors.Required("queries", "body", nil)
	}

	return nil
}

// ContextValidate validates this create query in query history command based on context it is used
func (m *CreateQueryInQueryHistoryCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateQueryInQueryHistoryCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateQueryInQueryHistoryCommand) UnmarshalBinary(b []byte) error {
	var res CreateQueryInQueryHistoryCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
