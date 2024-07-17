// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PrometheusRemoteWriteTargetJSON prometheus remote write target JSON
//
// swagger:model PrometheusRemoteWriteTargetJSON
type PrometheusRemoteWriteTargetJSON struct {

	// data source uid
	DataSourceUID string `json:"data_source_uid,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// remote write path
	RemoteWritePath string `json:"remote_write_path,omitempty"`
}

// Validate validates this prometheus remote write target JSON
func (m *PrometheusRemoteWriteTargetJSON) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this prometheus remote write target JSON based on context it is used
func (m *PrometheusRemoteWriteTargetJSON) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusRemoteWriteTargetJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusRemoteWriteTargetJSON) UnmarshalBinary(b []byte) error {
	var res PrometheusRemoteWriteTargetJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
