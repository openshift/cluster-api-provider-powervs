// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVpnconnectionsPeersubnetsGetReader is a Reader for the PcloudVpnconnectionsPeersubnetsGet structure.
type PcloudVpnconnectionsPeersubnetsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsPeersubnetsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudVpnconnectionsPeersubnetsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudVpnconnectionsPeersubnetsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudVpnconnectionsPeersubnetsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPcloudVpnconnectionsPeersubnetsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudVpnconnectionsPeersubnetsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudVpnconnectionsPeersubnetsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudVpnconnectionsPeersubnetsGetOK creates a PcloudVpnconnectionsPeersubnetsGetOK with default headers values
func NewPcloudVpnconnectionsPeersubnetsGetOK() *PcloudVpnconnectionsPeersubnetsGetOK {
	return &PcloudVpnconnectionsPeersubnetsGetOK{}
}

/*PcloudVpnconnectionsPeersubnetsGetOK handles this case with default header values.

OK
*/
type PcloudVpnconnectionsPeersubnetsGetOK struct {
	Payload *models.PeerSubnets
}

func (o *PcloudVpnconnectionsPeersubnetsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeerSubnets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsGetBadRequest creates a PcloudVpnconnectionsPeersubnetsGetBadRequest with default headers values
func NewPcloudVpnconnectionsPeersubnetsGetBadRequest() *PcloudVpnconnectionsPeersubnetsGetBadRequest {
	return &PcloudVpnconnectionsPeersubnetsGetBadRequest{}
}

/*PcloudVpnconnectionsPeersubnetsGetBadRequest handles this case with default header values.

Bad Request
*/
type PcloudVpnconnectionsPeersubnetsGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsGetUnauthorized creates a PcloudVpnconnectionsPeersubnetsGetUnauthorized with default headers values
func NewPcloudVpnconnectionsPeersubnetsGetUnauthorized() *PcloudVpnconnectionsPeersubnetsGetUnauthorized {
	return &PcloudVpnconnectionsPeersubnetsGetUnauthorized{}
}

/*PcloudVpnconnectionsPeersubnetsGetUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudVpnconnectionsPeersubnetsGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsGetForbidden creates a PcloudVpnconnectionsPeersubnetsGetForbidden with default headers values
func NewPcloudVpnconnectionsPeersubnetsGetForbidden() *PcloudVpnconnectionsPeersubnetsGetForbidden {
	return &PcloudVpnconnectionsPeersubnetsGetForbidden{}
}

/*PcloudVpnconnectionsPeersubnetsGetForbidden handles this case with default header values.

Forbidden
*/
type PcloudVpnconnectionsPeersubnetsGetForbidden struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsGetNotFound creates a PcloudVpnconnectionsPeersubnetsGetNotFound with default headers values
func NewPcloudVpnconnectionsPeersubnetsGetNotFound() *PcloudVpnconnectionsPeersubnetsGetNotFound {
	return &PcloudVpnconnectionsPeersubnetsGetNotFound{}
}

/*PcloudVpnconnectionsPeersubnetsGetNotFound handles this case with default header values.

Not Found
*/
type PcloudVpnconnectionsPeersubnetsGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsGetInternalServerError creates a PcloudVpnconnectionsPeersubnetsGetInternalServerError with default headers values
func NewPcloudVpnconnectionsPeersubnetsGetInternalServerError() *PcloudVpnconnectionsPeersubnetsGetInternalServerError {
	return &PcloudVpnconnectionsPeersubnetsGetInternalServerError{}
}

/*PcloudVpnconnectionsPeersubnetsGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsPeersubnetsGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
