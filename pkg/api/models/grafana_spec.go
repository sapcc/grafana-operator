// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GrafanaSpec grafana spec
// swagger:model GrafanaSpec
type GrafanaSpec struct {

	// admin password
	AdminPassword string `json:"AdminPassword,omitempty"`

	// admin user
	AdminUser string `json:"AdminUser,omitempty"`

	// anonymous
	Anonymous bool `json:"Anonymous,omitempty"`

	// basic auth
	BasicAuth bool `json:"BasicAuth,omitempty"`

	// disable login form
	DisableLoginForm bool `json:"DisableLoginForm,omitempty"`

	// disable signout menu
	DisableSignoutMenu bool `json:"DisableSignoutMenu,omitempty"`

	// enable auth proxy
	EnableAuthProxy bool `json:"EnableAuthProxy,omitempty"`

	// hostname
	Hostname string `json:"Hostname,omitempty"`

	// log level
	LogLevel string `json:"LogLevel,omitempty"`

	// secrets
	Secrets []string `json:"Secrets,omitempty"`
}

// Validate validates this grafana spec
func (m *GrafanaSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GrafanaSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GrafanaSpec) UnmarshalBinary(b []byte) error {
	var res GrafanaSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}