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

// PatchNotebookRequest patch notebook request
// swagger:model PatchNotebookRequest
type PatchNotebookRequest struct {

	// User driver memory in BDAP Yarn Cluster
	DriverMemory string `json:"driverMemory,omitempty"`

	// User excutor core setting in BDAP Yarn Cluster
	ExecutorCores string `json:"executorCores,omitempty"`

	// User excutor memory setting in BDAP Yarn Cluster
	ExecutorMemory string `json:"executorMemory,omitempty"`

	// User excutors setting in BDAP Yarn Cluster
	Executors string `json:"executors,omitempty"`

	// Notebook name
	// Required: true
	Name *string `json:"name"`

	// Namesapce where notebook has been created
	// Required: true
	Namespace *string `json:"namespace"`

	// BDAP Yarn Queue Setting
	Queue string `json:"queue,omitempty"`
}

// Validate validates this patch notebook request
func (m *PatchNotebookRequest) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *PatchNotebookRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PatchNotebookRequest) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchNotebookRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchNotebookRequest) UnmarshalBinary(b []byte) error {
	var res PatchNotebookRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
