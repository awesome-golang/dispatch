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
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSubscriptionParams creates a new DeleteSubscriptionParams object
// with the default values initialized.
func NewDeleteSubscriptionParams() *DeleteSubscriptionParams {
	var ()
	return &DeleteSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubscriptionParamsWithTimeout creates a new DeleteSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSubscriptionParamsWithTimeout(timeout time.Duration) *DeleteSubscriptionParams {
	var ()
	return &DeleteSubscriptionParams{

		timeout: timeout,
	}
}

// NewDeleteSubscriptionParamsWithContext creates a new DeleteSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSubscriptionParamsWithContext(ctx context.Context) *DeleteSubscriptionParams {
	var ()
	return &DeleteSubscriptionParams{

		Context: ctx,
	}
}

// NewDeleteSubscriptionParamsWithHTTPClient creates a new DeleteSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSubscriptionParamsWithHTTPClient(client *http.Client) *DeleteSubscriptionParams {
	var ()
	return &DeleteSubscriptionParams{
		HTTPClient: client,
	}
}

/*DeleteSubscriptionParams contains all the parameters to send to the API endpoint
for the delete subscription operation typically these are written to a http.Request
*/
type DeleteSubscriptionParams struct {

	/*SubscriptionName
	  Name of the subscription to work on

	*/
	SubscriptionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete subscription params
func (o *DeleteSubscriptionParams) WithTimeout(timeout time.Duration) *DeleteSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete subscription params
func (o *DeleteSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete subscription params
func (o *DeleteSubscriptionParams) WithContext(ctx context.Context) *DeleteSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete subscription params
func (o *DeleteSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete subscription params
func (o *DeleteSubscriptionParams) WithHTTPClient(client *http.Client) *DeleteSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete subscription params
func (o *DeleteSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubscriptionName adds the subscriptionName to the delete subscription params
func (o *DeleteSubscriptionParams) WithSubscriptionName(subscriptionName string) *DeleteSubscriptionParams {
	o.SetSubscriptionName(subscriptionName)
	return o
}

// SetSubscriptionName adds the subscriptionName to the delete subscription params
func (o *DeleteSubscriptionParams) SetSubscriptionName(subscriptionName string) {
	o.SubscriptionName = subscriptionName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param subscriptionName
	if err := r.SetPathParam("subscriptionName", o.SubscriptionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
