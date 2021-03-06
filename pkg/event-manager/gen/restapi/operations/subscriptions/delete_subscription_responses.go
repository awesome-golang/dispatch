///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/event-manager/gen/models"
)

// DeleteSubscriptionOKCode is the HTTP code returned for type DeleteSubscriptionOK
const DeleteSubscriptionOKCode int = 200

/*DeleteSubscriptionOK successful operation

swagger:response deleteSubscriptionOK
*/
type DeleteSubscriptionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Subscription `json:"body,omitempty"`
}

// NewDeleteSubscriptionOK creates DeleteSubscriptionOK with default headers values
func NewDeleteSubscriptionOK() *DeleteSubscriptionOK {
	return &DeleteSubscriptionOK{}
}

// WithPayload adds the payload to the delete subscription o k response
func (o *DeleteSubscriptionOK) WithPayload(payload *models.Subscription) *DeleteSubscriptionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete subscription o k response
func (o *DeleteSubscriptionOK) SetPayload(payload *models.Subscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSubscriptionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSubscriptionBadRequestCode is the HTTP code returned for type DeleteSubscriptionBadRequest
const DeleteSubscriptionBadRequestCode int = 400

/*DeleteSubscriptionBadRequest Invalid ID supplied

swagger:response deleteSubscriptionBadRequest
*/
type DeleteSubscriptionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSubscriptionBadRequest creates DeleteSubscriptionBadRequest with default headers values
func NewDeleteSubscriptionBadRequest() *DeleteSubscriptionBadRequest {
	return &DeleteSubscriptionBadRequest{}
}

// WithPayload adds the payload to the delete subscription bad request response
func (o *DeleteSubscriptionBadRequest) WithPayload(payload *models.Error) *DeleteSubscriptionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete subscription bad request response
func (o *DeleteSubscriptionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSubscriptionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSubscriptionNotFoundCode is the HTTP code returned for type DeleteSubscriptionNotFound
const DeleteSubscriptionNotFoundCode int = 404

/*DeleteSubscriptionNotFound Subscription not found

swagger:response deleteSubscriptionNotFound
*/
type DeleteSubscriptionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSubscriptionNotFound creates DeleteSubscriptionNotFound with default headers values
func NewDeleteSubscriptionNotFound() *DeleteSubscriptionNotFound {
	return &DeleteSubscriptionNotFound{}
}

// WithPayload adds the payload to the delete subscription not found response
func (o *DeleteSubscriptionNotFound) WithPayload(payload *models.Error) *DeleteSubscriptionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete subscription not found response
func (o *DeleteSubscriptionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSubscriptionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSubscriptionInternalServerErrorCode is the HTTP code returned for type DeleteSubscriptionInternalServerError
const DeleteSubscriptionInternalServerErrorCode int = 500

/*DeleteSubscriptionInternalServerError Internal server error

swagger:response deleteSubscriptionInternalServerError
*/
type DeleteSubscriptionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSubscriptionInternalServerError creates DeleteSubscriptionInternalServerError with default headers values
func NewDeleteSubscriptionInternalServerError() *DeleteSubscriptionInternalServerError {
	return &DeleteSubscriptionInternalServerError{}
}

// WithPayload adds the payload to the delete subscription internal server error response
func (o *DeleteSubscriptionInternalServerError) WithPayload(payload *models.Error) *DeleteSubscriptionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete subscription internal server error response
func (o *DeleteSubscriptionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSubscriptionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteSubscriptionDefault Generic error response

swagger:response deleteSubscriptionDefault
*/
type DeleteSubscriptionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSubscriptionDefault creates DeleteSubscriptionDefault with default headers values
func NewDeleteSubscriptionDefault(code int) *DeleteSubscriptionDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSubscriptionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete subscription default response
func (o *DeleteSubscriptionDefault) WithStatusCode(code int) *DeleteSubscriptionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete subscription default response
func (o *DeleteSubscriptionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete subscription default response
func (o *DeleteSubscriptionDefault) WithPayload(payload *models.Error) *DeleteSubscriptionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete subscription default response
func (o *DeleteSubscriptionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSubscriptionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
