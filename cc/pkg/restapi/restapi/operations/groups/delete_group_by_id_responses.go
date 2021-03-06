// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// DeleteGroupByIDOKCode is the HTTP code returned for type DeleteGroupByIDOK
const DeleteGroupByIDOKCode int = 200

/*DeleteGroupByIDOK Detailed Group and Group information.

swagger:response deleteGroupByIdOK
*/
type DeleteGroupByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Group `json:"body,omitempty"`
}

// NewDeleteGroupByIDOK creates DeleteGroupByIDOK with default headers values
func NewDeleteGroupByIDOK() *DeleteGroupByIDOK {

	return &DeleteGroupByIDOK{}
}

// WithPayload adds the payload to the delete group by Id o k response
func (o *DeleteGroupByIDOK) WithPayload(payload *models.Group) *DeleteGroupByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete group by Id o k response
func (o *DeleteGroupByIDOK) SetPayload(payload *models.Group) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteGroupByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteGroupByIDUnauthorizedCode is the HTTP code returned for type DeleteGroupByIDUnauthorized
const DeleteGroupByIDUnauthorizedCode int = 401

/*DeleteGroupByIDUnauthorized Unauthorized

swagger:response deleteGroupByIdUnauthorized
*/
type DeleteGroupByIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteGroupByIDUnauthorized creates DeleteGroupByIDUnauthorized with default headers values
func NewDeleteGroupByIDUnauthorized() *DeleteGroupByIDUnauthorized {

	return &DeleteGroupByIDUnauthorized{}
}

// WithPayload adds the payload to the delete group by Id unauthorized response
func (o *DeleteGroupByIDUnauthorized) WithPayload(payload *models.Error) *DeleteGroupByIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete group by Id unauthorized response
func (o *DeleteGroupByIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteGroupByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteGroupByIDNotFoundCode is the HTTP code returned for type DeleteGroupByIDNotFound
const DeleteGroupByIDNotFoundCode int = 404

/*DeleteGroupByIDNotFound Model with the given ID not found.

swagger:response deleteGroupByIdNotFound
*/
type DeleteGroupByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteGroupByIDNotFound creates DeleteGroupByIDNotFound with default headers values
func NewDeleteGroupByIDNotFound() *DeleteGroupByIDNotFound {

	return &DeleteGroupByIDNotFound{}
}

// WithPayload adds the payload to the delete group by Id not found response
func (o *DeleteGroupByIDNotFound) WithPayload(payload *models.Error) *DeleteGroupByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete group by Id not found response
func (o *DeleteGroupByIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteGroupByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
