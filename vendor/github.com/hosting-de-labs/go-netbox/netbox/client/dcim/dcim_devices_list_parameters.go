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
)

// NewDcimDevicesListParams creates a new DcimDevicesListParams object
// with the default values initialized.
func NewDcimDevicesListParams() *DcimDevicesListParams {
	var ()
	return &DcimDevicesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesListParamsWithTimeout creates a new DcimDevicesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDevicesListParamsWithTimeout(timeout time.Duration) *DcimDevicesListParams {
	var ()
	return &DcimDevicesListParams{

		timeout: timeout,
	}
}

// NewDcimDevicesListParamsWithContext creates a new DcimDevicesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDevicesListParamsWithContext(ctx context.Context) *DcimDevicesListParams {
	var ()
	return &DcimDevicesListParams{

		Context: ctx,
	}
}

// NewDcimDevicesListParamsWithHTTPClient creates a new DcimDevicesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDevicesListParamsWithHTTPClient(client *http.Client) *DcimDevicesListParams {
	var ()
	return &DcimDevicesListParams{
		HTTPClient: client,
	}
}

/*DcimDevicesListParams contains all the parameters to send to the API endpoint
for the dcim devices list operation typically these are written to a http.Request
*/
type DcimDevicesListParams struct {

	/*AssetTag*/
	AssetTag *string
	/*ClusterID*/
	ClusterID *int64
	/*DeviceTypeID*/
	DeviceTypeID *int64
	/*HasPrimaryIP*/
	HasPrimaryIP *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*IsConsoleServer*/
	IsConsoleServer *string
	/*IsFullDepth*/
	IsFullDepth *string
	/*IsNetworkDevice*/
	IsNetworkDevice *string
	/*IsPdu*/
	IsPdu *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*MacAddress*/
	MacAddress *string
	/*Manufacturer*/
	Manufacturer *string
	/*ManufacturerID*/
	ManufacturerID *int64
	/*Model*/
	Model *string
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Platform*/
	Platform *string
	/*PlatformID*/
	PlatformID *int64
	/*Position*/
	Position *float64
	/*Q*/
	Q *string
	/*RackGroupID*/
	RackGroupID *int64
	/*RackID*/
	RackID *int64
	/*Region*/
	Region *string
	/*RegionID*/
	RegionID *int64
	/*Role*/
	Role *string
	/*RoleID*/
	RoleID *int64
	/*Serial*/
	Serial *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *int64
	/*Status*/
	Status *string
	/*Tag*/
	Tag *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *int64
	/*VirtualChassisID*/
	VirtualChassisID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim devices list params
func (o *DcimDevicesListParams) WithTimeout(timeout time.Duration) *DcimDevicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices list params
func (o *DcimDevicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices list params
func (o *DcimDevicesListParams) WithContext(ctx context.Context) *DcimDevicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices list params
func (o *DcimDevicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices list params
func (o *DcimDevicesListParams) WithHTTPClient(client *http.Client) *DcimDevicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices list params
func (o *DcimDevicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetTag adds the assetTag to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTag(assetTag *string) *DcimDevicesListParams {
	o.SetAssetTag(assetTag)
	return o
}

// SetAssetTag adds the assetTag to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTag(assetTag *string) {
	o.AssetTag = assetTag
}

// WithClusterID adds the clusterID to the dcim devices list params
func (o *DcimDevicesListParams) WithClusterID(clusterID *int64) *DcimDevicesListParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the dcim devices list params
func (o *DcimDevicesListParams) SetClusterID(clusterID *int64) {
	o.ClusterID = clusterID
}

// WithDeviceTypeID adds the deviceTypeID to the dcim devices list params
func (o *DcimDevicesListParams) WithDeviceTypeID(deviceTypeID *int64) *DcimDevicesListParams {
	o.SetDeviceTypeID(deviceTypeID)
	return o
}

// SetDeviceTypeID adds the deviceTypeId to the dcim devices list params
func (o *DcimDevicesListParams) SetDeviceTypeID(deviceTypeID *int64) {
	o.DeviceTypeID = deviceTypeID
}

// WithHasPrimaryIP adds the hasPrimaryIP to the dcim devices list params
func (o *DcimDevicesListParams) WithHasPrimaryIP(hasPrimaryIP *string) *DcimDevicesListParams {
	o.SetHasPrimaryIP(hasPrimaryIP)
	return o
}

// SetHasPrimaryIP adds the hasPrimaryIp to the dcim devices list params
func (o *DcimDevicesListParams) SetHasPrimaryIP(hasPrimaryIP *string) {
	o.HasPrimaryIP = hasPrimaryIP
}

// WithIDIn adds the iDIn to the dcim devices list params
func (o *DcimDevicesListParams) WithIDIn(iDIn *string) *DcimDevicesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the dcim devices list params
func (o *DcimDevicesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithIsConsoleServer adds the isConsoleServer to the dcim devices list params
func (o *DcimDevicesListParams) WithIsConsoleServer(isConsoleServer *string) *DcimDevicesListParams {
	o.SetIsConsoleServer(isConsoleServer)
	return o
}

// SetIsConsoleServer adds the isConsoleServer to the dcim devices list params
func (o *DcimDevicesListParams) SetIsConsoleServer(isConsoleServer *string) {
	o.IsConsoleServer = isConsoleServer
}

// WithIsFullDepth adds the isFullDepth to the dcim devices list params
func (o *DcimDevicesListParams) WithIsFullDepth(isFullDepth *string) *DcimDevicesListParams {
	o.SetIsFullDepth(isFullDepth)
	return o
}

// SetIsFullDepth adds the isFullDepth to the dcim devices list params
func (o *DcimDevicesListParams) SetIsFullDepth(isFullDepth *string) {
	o.IsFullDepth = isFullDepth
}

// WithIsNetworkDevice adds the isNetworkDevice to the dcim devices list params
func (o *DcimDevicesListParams) WithIsNetworkDevice(isNetworkDevice *string) *DcimDevicesListParams {
	o.SetIsNetworkDevice(isNetworkDevice)
	return o
}

// SetIsNetworkDevice adds the isNetworkDevice to the dcim devices list params
func (o *DcimDevicesListParams) SetIsNetworkDevice(isNetworkDevice *string) {
	o.IsNetworkDevice = isNetworkDevice
}

// WithIsPdu adds the isPdu to the dcim devices list params
func (o *DcimDevicesListParams) WithIsPdu(isPdu *string) *DcimDevicesListParams {
	o.SetIsPdu(isPdu)
	return o
}

// SetIsPdu adds the isPdu to the dcim devices list params
func (o *DcimDevicesListParams) SetIsPdu(isPdu *string) {
	o.IsPdu = isPdu
}

// WithLimit adds the limit to the dcim devices list params
func (o *DcimDevicesListParams) WithLimit(limit *int64) *DcimDevicesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim devices list params
func (o *DcimDevicesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMacAddress adds the macAddress to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddress(macAddress *string) *DcimDevicesListParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithManufacturer adds the manufacturer to the dcim devices list params
func (o *DcimDevicesListParams) WithManufacturer(manufacturer *string) *DcimDevicesListParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the dcim devices list params
func (o *DcimDevicesListParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithManufacturerID adds the manufacturerID to the dcim devices list params
func (o *DcimDevicesListParams) WithManufacturerID(manufacturerID *int64) *DcimDevicesListParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the dcim devices list params
func (o *DcimDevicesListParams) SetManufacturerID(manufacturerID *int64) {
	o.ManufacturerID = manufacturerID
}

// WithModel adds the model to the dcim devices list params
func (o *DcimDevicesListParams) WithModel(model *string) *DcimDevicesListParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the dcim devices list params
func (o *DcimDevicesListParams) SetModel(model *string) {
	o.Model = model
}

// WithName adds the name to the dcim devices list params
func (o *DcimDevicesListParams) WithName(name *string) *DcimDevicesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim devices list params
func (o *DcimDevicesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim devices list params
func (o *DcimDevicesListParams) WithOffset(offset *int64) *DcimDevicesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim devices list params
func (o *DcimDevicesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPlatform adds the platform to the dcim devices list params
func (o *DcimDevicesListParams) WithPlatform(platform *string) *DcimDevicesListParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the dcim devices list params
func (o *DcimDevicesListParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPlatformID adds the platformID to the dcim devices list params
func (o *DcimDevicesListParams) WithPlatformID(platformID *int64) *DcimDevicesListParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the dcim devices list params
func (o *DcimDevicesListParams) SetPlatformID(platformID *int64) {
	o.PlatformID = platformID
}

// WithPosition adds the position to the dcim devices list params
func (o *DcimDevicesListParams) WithPosition(position *float64) *DcimDevicesListParams {
	o.SetPosition(position)
	return o
}

// SetPosition adds the position to the dcim devices list params
func (o *DcimDevicesListParams) SetPosition(position *float64) {
	o.Position = position
}

// WithQ adds the q to the dcim devices list params
func (o *DcimDevicesListParams) WithQ(q *string) *DcimDevicesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim devices list params
func (o *DcimDevicesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRackGroupID adds the rackGroupID to the dcim devices list params
func (o *DcimDevicesListParams) WithRackGroupID(rackGroupID *int64) *DcimDevicesListParams {
	o.SetRackGroupID(rackGroupID)
	return o
}

// SetRackGroupID adds the rackGroupId to the dcim devices list params
func (o *DcimDevicesListParams) SetRackGroupID(rackGroupID *int64) {
	o.RackGroupID = rackGroupID
}

// WithRackID adds the rackID to the dcim devices list params
func (o *DcimDevicesListParams) WithRackID(rackID *int64) *DcimDevicesListParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the dcim devices list params
func (o *DcimDevicesListParams) SetRackID(rackID *int64) {
	o.RackID = rackID
}

// WithRegion adds the region to the dcim devices list params
func (o *DcimDevicesListParams) WithRegion(region *string) *DcimDevicesListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim devices list params
func (o *DcimDevicesListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionID adds the regionID to the dcim devices list params
func (o *DcimDevicesListParams) WithRegionID(regionID *int64) *DcimDevicesListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim devices list params
func (o *DcimDevicesListParams) SetRegionID(regionID *int64) {
	o.RegionID = regionID
}

// WithRole adds the role to the dcim devices list params
func (o *DcimDevicesListParams) WithRole(role *string) *DcimDevicesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the dcim devices list params
func (o *DcimDevicesListParams) SetRole(role *string) {
	o.Role = role
}

// WithRoleID adds the roleID to the dcim devices list params
func (o *DcimDevicesListParams) WithRoleID(roleID *int64) *DcimDevicesListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the dcim devices list params
func (o *DcimDevicesListParams) SetRoleID(roleID *int64) {
	o.RoleID = roleID
}

// WithSerial adds the serial to the dcim devices list params
func (o *DcimDevicesListParams) WithSerial(serial *string) *DcimDevicesListParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the dcim devices list params
func (o *DcimDevicesListParams) SetSerial(serial *string) {
	o.Serial = serial
}

// WithSite adds the site to the dcim devices list params
func (o *DcimDevicesListParams) WithSite(site *string) *DcimDevicesListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim devices list params
func (o *DcimDevicesListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the dcim devices list params
func (o *DcimDevicesListParams) WithSiteID(siteID *int64) *DcimDevicesListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim devices list params
func (o *DcimDevicesListParams) SetSiteID(siteID *int64) {
	o.SiteID = siteID
}

// WithStatus adds the status to the dcim devices list params
func (o *DcimDevicesListParams) WithStatus(status *string) *DcimDevicesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the dcim devices list params
func (o *DcimDevicesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTag adds the tag to the dcim devices list params
func (o *DcimDevicesListParams) WithTag(tag *string) *DcimDevicesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim devices list params
func (o *DcimDevicesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the dcim devices list params
func (o *DcimDevicesListParams) WithTenant(tenant *string) *DcimDevicesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim devices list params
func (o *DcimDevicesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the dcim devices list params
func (o *DcimDevicesListParams) WithTenantID(tenantID *int64) *DcimDevicesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim devices list params
func (o *DcimDevicesListParams) SetTenantID(tenantID *int64) {
	o.TenantID = tenantID
}

// WithVirtualChassisID adds the virtualChassisID to the dcim devices list params
func (o *DcimDevicesListParams) WithVirtualChassisID(virtualChassisID *int64) *DcimDevicesListParams {
	o.SetVirtualChassisID(virtualChassisID)
	return o
}

// SetVirtualChassisID adds the virtualChassisId to the dcim devices list params
func (o *DcimDevicesListParams) SetVirtualChassisID(virtualChassisID *int64) {
	o.VirtualChassisID = virtualChassisID
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssetTag != nil {

		// query param asset_tag
		var qrAssetTag string
		if o.AssetTag != nil {
			qrAssetTag = *o.AssetTag
		}
		qAssetTag := qrAssetTag
		if qAssetTag != "" {
			if err := r.SetQueryParam("asset_tag", qAssetTag); err != nil {
				return err
			}
		}

	}

	if o.ClusterID != nil {

		// query param cluster_id
		var qrClusterID int64
		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := swag.FormatInt64(qrClusterID)
		if qClusterID != "" {
			if err := r.SetQueryParam("cluster_id", qClusterID); err != nil {
				return err
			}
		}

	}

	if o.DeviceTypeID != nil {

		// query param device_type_id
		var qrDeviceTypeID int64
		if o.DeviceTypeID != nil {
			qrDeviceTypeID = *o.DeviceTypeID
		}
		qDeviceTypeID := swag.FormatInt64(qrDeviceTypeID)
		if qDeviceTypeID != "" {
			if err := r.SetQueryParam("device_type_id", qDeviceTypeID); err != nil {
				return err
			}
		}

	}

	if o.HasPrimaryIP != nil {

		// query param has_primary_ip
		var qrHasPrimaryIP string
		if o.HasPrimaryIP != nil {
			qrHasPrimaryIP = *o.HasPrimaryIP
		}
		qHasPrimaryIP := qrHasPrimaryIP
		if qHasPrimaryIP != "" {
			if err := r.SetQueryParam("has_primary_ip", qHasPrimaryIP); err != nil {
				return err
			}
		}

	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn string
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := qrIDIn
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
				return err
			}
		}

	}

	if o.IsConsoleServer != nil {

		// query param is_console_server
		var qrIsConsoleServer string
		if o.IsConsoleServer != nil {
			qrIsConsoleServer = *o.IsConsoleServer
		}
		qIsConsoleServer := qrIsConsoleServer
		if qIsConsoleServer != "" {
			if err := r.SetQueryParam("is_console_server", qIsConsoleServer); err != nil {
				return err
			}
		}

	}

	if o.IsFullDepth != nil {

		// query param is_full_depth
		var qrIsFullDepth string
		if o.IsFullDepth != nil {
			qrIsFullDepth = *o.IsFullDepth
		}
		qIsFullDepth := qrIsFullDepth
		if qIsFullDepth != "" {
			if err := r.SetQueryParam("is_full_depth", qIsFullDepth); err != nil {
				return err
			}
		}

	}

	if o.IsNetworkDevice != nil {

		// query param is_network_device
		var qrIsNetworkDevice string
		if o.IsNetworkDevice != nil {
			qrIsNetworkDevice = *o.IsNetworkDevice
		}
		qIsNetworkDevice := qrIsNetworkDevice
		if qIsNetworkDevice != "" {
			if err := r.SetQueryParam("is_network_device", qIsNetworkDevice); err != nil {
				return err
			}
		}

	}

	if o.IsPdu != nil {

		// query param is_pdu
		var qrIsPdu string
		if o.IsPdu != nil {
			qrIsPdu = *o.IsPdu
		}
		qIsPdu := qrIsPdu
		if qIsPdu != "" {
			if err := r.SetQueryParam("is_pdu", qIsPdu); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.MacAddress != nil {

		// query param mac_address
		var qrMacAddress string
		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {
			if err := r.SetQueryParam("mac_address", qMacAddress); err != nil {
				return err
			}
		}

	}

	if o.Manufacturer != nil {

		// query param manufacturer
		var qrManufacturer string
		if o.Manufacturer != nil {
			qrManufacturer = *o.Manufacturer
		}
		qManufacturer := qrManufacturer
		if qManufacturer != "" {
			if err := r.SetQueryParam("manufacturer", qManufacturer); err != nil {
				return err
			}
		}

	}

	if o.ManufacturerID != nil {

		// query param manufacturer_id
		var qrManufacturerID int64
		if o.ManufacturerID != nil {
			qrManufacturerID = *o.ManufacturerID
		}
		qManufacturerID := swag.FormatInt64(qrManufacturerID)
		if qManufacturerID != "" {
			if err := r.SetQueryParam("manufacturer_id", qManufacturerID); err != nil {
				return err
			}
		}

	}

	if o.Model != nil {

		// query param model
		var qrModel string
		if o.Model != nil {
			qrModel = *o.Model
		}
		qModel := qrModel
		if qModel != "" {
			if err := r.SetQueryParam("model", qModel); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Platform != nil {

		// query param platform
		var qrPlatform string
		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {
			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}

	}

	if o.PlatformID != nil {

		// query param platform_id
		var qrPlatformID int64
		if o.PlatformID != nil {
			qrPlatformID = *o.PlatformID
		}
		qPlatformID := swag.FormatInt64(qrPlatformID)
		if qPlatformID != "" {
			if err := r.SetQueryParam("platform_id", qPlatformID); err != nil {
				return err
			}
		}

	}

	if o.Position != nil {

		// query param position
		var qrPosition float64
		if o.Position != nil {
			qrPosition = *o.Position
		}
		qPosition := swag.FormatFloat64(qrPosition)
		if qPosition != "" {
			if err := r.SetQueryParam("position", qPosition); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.RackGroupID != nil {

		// query param rack_group_id
		var qrRackGroupID int64
		if o.RackGroupID != nil {
			qrRackGroupID = *o.RackGroupID
		}
		qRackGroupID := swag.FormatInt64(qrRackGroupID)
		if qRackGroupID != "" {
			if err := r.SetQueryParam("rack_group_id", qRackGroupID); err != nil {
				return err
			}
		}

	}

	if o.RackID != nil {

		// query param rack_id
		var qrRackID int64
		if o.RackID != nil {
			qrRackID = *o.RackID
		}
		qRackID := swag.FormatInt64(qrRackID)
		if qRackID != "" {
			if err := r.SetQueryParam("rack_id", qRackID); err != nil {
				return err
			}
		}

	}

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID int64
		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := swag.FormatInt64(qrRegionID)
		if qRegionID != "" {
			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
				return err
			}
		}

	}

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID int64
		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := swag.FormatInt64(qrRoleID)
		if qRoleID != "" {
			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
				return err
			}
		}

	}

	if o.Serial != nil {

		// query param serial
		var qrSerial string
		if o.Serial != nil {
			qrSerial = *o.Serial
		}
		qSerial := qrSerial
		if qSerial != "" {
			if err := r.SetQueryParam("serial", qSerial); err != nil {
				return err
			}
		}

	}

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID int64
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := swag.FormatInt64(qrSiteID)
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string
		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {
			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}

	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID int64
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := swag.FormatInt64(qrTenantID)
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}

	}

	if o.VirtualChassisID != nil {

		// query param virtual_chassis_id
		var qrVirtualChassisID int64
		if o.VirtualChassisID != nil {
			qrVirtualChassisID = *o.VirtualChassisID
		}
		qVirtualChassisID := swag.FormatInt64(qrVirtualChassisID)
		if qVirtualChassisID != "" {
			if err := r.SetQueryParam("virtual_chassis_id", qVirtualChassisID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}