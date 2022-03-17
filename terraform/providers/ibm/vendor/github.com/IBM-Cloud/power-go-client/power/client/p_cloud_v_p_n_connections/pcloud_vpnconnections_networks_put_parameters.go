// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudVpnconnectionsNetworksPutParams creates a new PcloudVpnconnectionsNetworksPutParams object
// with the default values initialized.
func NewPcloudVpnconnectionsNetworksPutParams() *PcloudVpnconnectionsNetworksPutParams {
	var ()
	return &PcloudVpnconnectionsNetworksPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudVpnconnectionsNetworksPutParamsWithTimeout creates a new PcloudVpnconnectionsNetworksPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudVpnconnectionsNetworksPutParamsWithTimeout(timeout time.Duration) *PcloudVpnconnectionsNetworksPutParams {
	var ()
	return &PcloudVpnconnectionsNetworksPutParams{

		timeout: timeout,
	}
}

// NewPcloudVpnconnectionsNetworksPutParamsWithContext creates a new PcloudVpnconnectionsNetworksPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudVpnconnectionsNetworksPutParamsWithContext(ctx context.Context) *PcloudVpnconnectionsNetworksPutParams {
	var ()
	return &PcloudVpnconnectionsNetworksPutParams{

		Context: ctx,
	}
}

// NewPcloudVpnconnectionsNetworksPutParamsWithHTTPClient creates a new PcloudVpnconnectionsNetworksPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudVpnconnectionsNetworksPutParamsWithHTTPClient(client *http.Client) *PcloudVpnconnectionsNetworksPutParams {
	var ()
	return &PcloudVpnconnectionsNetworksPutParams{
		HTTPClient: client,
	}
}

/*PcloudVpnconnectionsNetworksPutParams contains all the parameters to send to the API endpoint
for the pcloud vpnconnections networks put operation typically these are written to a http.Request
*/
type PcloudVpnconnectionsNetworksPutParams struct {

	/*Body
	  network to attach

	*/
	Body *models.NetworkID
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*VpnConnectionID
	  ID of a VPN connection

	*/
	VpnConnectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) WithTimeout(timeout time.Duration) *PcloudVpnconnectionsNetworksPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) WithContext(ctx context.Context) *PcloudVpnconnectionsNetworksPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) WithHTTPClient(client *http.Client) *PcloudVpnconnectionsNetworksPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) WithBody(body *models.NetworkID) *PcloudVpnconnectionsNetworksPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) SetBody(body *models.NetworkID) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) WithCloudInstanceID(cloudInstanceID string) *PcloudVpnconnectionsNetworksPutParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithVpnConnectionID adds the vpnConnectionID to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) WithVpnConnectionID(vpnConnectionID string) *PcloudVpnconnectionsNetworksPutParams {
	o.SetVpnConnectionID(vpnConnectionID)
	return o
}

// SetVpnConnectionID adds the vpnConnectionId to the pcloud vpnconnections networks put params
func (o *PcloudVpnconnectionsNetworksPutParams) SetVpnConnectionID(vpnConnectionID string) {
	o.VpnConnectionID = vpnConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudVpnconnectionsNetworksPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param vpn_connection_id
	if err := r.SetPathParam("vpn_connection_id", o.VpnConnectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}