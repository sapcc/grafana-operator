// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/integr8ly/grafana-operator/v3/pkg/api/models"
)

// ListGrafanasOKCode is the HTTP code returned for type ListGrafanasOK
const ListGrafanasOKCode int = 200

/*ListGrafanasOK OK

swagger:response listGrafanasOK
*/
type ListGrafanasOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Grafana `json:"body,omitempty"`
}

// NewListGrafanasOK creates ListGrafanasOK with default headers values
func NewListGrafanasOK() *ListGrafanasOK {

	return &ListGrafanasOK{}
}

// WithPayload adds the payload to the list grafanas o k response
func (o *ListGrafanasOK) WithPayload(payload []*models.Grafana) *ListGrafanasOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list grafanas o k response
func (o *ListGrafanasOK) SetPayload(payload []*models.Grafana) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListGrafanasOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Grafana, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*ListGrafanasDefault Error

swagger:response listGrafanasDefault
*/
type ListGrafanasDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListGrafanasDefault creates ListGrafanasDefault with default headers values
func NewListGrafanasDefault(code int) *ListGrafanasDefault {
	if code <= 0 {
		code = 500
	}

	return &ListGrafanasDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list grafanas default response
func (o *ListGrafanasDefault) WithStatusCode(code int) *ListGrafanasDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list grafanas default response
func (o *ListGrafanasDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list grafanas default response
func (o *ListGrafanasDefault) WithPayload(payload *models.Error) *ListGrafanasDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list grafanas default response
func (o *ListGrafanasDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListGrafanasDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
