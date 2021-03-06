// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_service_d_h_c_p

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
)

// NewPcloudDhcpPostParams creates a new PcloudDhcpPostParams object
// with the default values initialized.
func NewPcloudDhcpPostParams() *PcloudDhcpPostParams {
	var ()
	return &PcloudDhcpPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudDhcpPostParamsWithTimeout creates a new PcloudDhcpPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudDhcpPostParamsWithTimeout(timeout time.Duration) *PcloudDhcpPostParams {
	var ()
	return &PcloudDhcpPostParams{

		timeout: timeout,
	}
}

// NewPcloudDhcpPostParamsWithContext creates a new PcloudDhcpPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudDhcpPostParamsWithContext(ctx context.Context) *PcloudDhcpPostParams {
	var ()
	return &PcloudDhcpPostParams{

		Context: ctx,
	}
}

// NewPcloudDhcpPostParamsWithHTTPClient creates a new PcloudDhcpPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudDhcpPostParamsWithHTTPClient(client *http.Client) *PcloudDhcpPostParams {
	var ()
	return &PcloudDhcpPostParams{
		HTTPClient: client,
	}
}

/*PcloudDhcpPostParams contains all the parameters to send to the API endpoint
for the pcloud dhcp post operation typically these are written to a http.Request
*/
type PcloudDhcpPostParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud dhcp post params
func (o *PcloudDhcpPostParams) WithTimeout(timeout time.Duration) *PcloudDhcpPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud dhcp post params
func (o *PcloudDhcpPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud dhcp post params
func (o *PcloudDhcpPostParams) WithContext(ctx context.Context) *PcloudDhcpPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud dhcp post params
func (o *PcloudDhcpPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud dhcp post params
func (o *PcloudDhcpPostParams) WithHTTPClient(client *http.Client) *PcloudDhcpPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud dhcp post params
func (o *PcloudDhcpPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud dhcp post params
func (o *PcloudDhcpPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudDhcpPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud dhcp post params
func (o *PcloudDhcpPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudDhcpPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
