// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewDcimInventoryItemsPartialUpdateParams creates a new DcimInventoryItemsPartialUpdateParams object
// with the default values initialized.
func NewDcimInventoryItemsPartialUpdateParams() *DcimInventoryItemsPartialUpdateParams {
	var ()
	return &DcimInventoryItemsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemsPartialUpdateParamsWithTimeout creates a new DcimInventoryItemsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInventoryItemsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimInventoryItemsPartialUpdateParams {
	var ()
	return &DcimInventoryItemsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimInventoryItemsPartialUpdateParamsWithContext creates a new DcimInventoryItemsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInventoryItemsPartialUpdateParamsWithContext(ctx context.Context) *DcimInventoryItemsPartialUpdateParams {
	var ()
	return &DcimInventoryItemsPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimInventoryItemsPartialUpdateParamsWithHTTPClient creates a new DcimInventoryItemsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInventoryItemsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimInventoryItemsPartialUpdateParams {
	var ()
	return &DcimInventoryItemsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimInventoryItemsPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim inventory items partial update operation typically these are written to a http.Request
*/
type DcimInventoryItemsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableInventoryItem
	/*ID
	  A unique integer value identifying this inventory item.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimInventoryItemsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) WithContext(ctx context.Context) *DcimInventoryItemsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimInventoryItemsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) WithData(data *models.WritableInventoryItem) *DcimInventoryItemsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) SetData(data *models.WritableInventoryItem) {
	o.Data = data
}

// WithID adds the id to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) WithID(id int64) *DcimInventoryItemsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim inventory items partial update params
func (o *DcimInventoryItemsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
