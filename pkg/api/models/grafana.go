// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Grafana grafana
// swagger:model Grafana
type Grafana struct {

	// config
	Config GrafanaConfig `json:"config,omitempty"`

	// name of Grafana instance
	// Max Length: 20
	// Pattern: ^[a-z]([-a-z0-9]*[a-z0-9])?$
	Name *string `json:"name,omitempty"`

	// name of Grafana instance
	// Max Length: 20
	// Pattern: ^[a-z]([-a-z0-9]*[a-z0-9])?$
	Namespace *string `json:"namespace,omitempty"`
}

// Validate validates this grafana
func (m *Grafana) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Grafana) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if err := m.Config.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("config")
		}
		return err
	}

	return nil
}

func (m *Grafana) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 20); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[a-z]([-a-z0-9]*[a-z0-9])?$`); err != nil {
		return err
	}

	return nil
}

func (m *Grafana) validateNamespace(formats strfmt.Registry) error {

	if swag.IsZero(m.Namespace) { // not required
		return nil
	}

	if err := validate.MaxLength("namespace", "body", string(*m.Namespace), 20); err != nil {
		return err
	}

	if err := validate.Pattern("namespace", "body", string(*m.Namespace), `^[a-z]([-a-z0-9]*[a-z0-9])?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Grafana) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Grafana) UnmarshalBinary(b []byte) error {
	var res Grafana
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
