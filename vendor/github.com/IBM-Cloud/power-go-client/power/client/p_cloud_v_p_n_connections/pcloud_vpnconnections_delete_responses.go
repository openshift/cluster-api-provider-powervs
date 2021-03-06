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

// PcloudVpnconnectionsDeleteReader is a Reader for the PcloudVpnconnectionsDelete structure.
type PcloudVpnconnectionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewPcloudVpnconnectionsDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudVpnconnectionsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudVpnconnectionsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPcloudVpnconnectionsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudVpnconnectionsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudVpnconnectionsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudVpnconnectionsDeleteAccepted creates a PcloudVpnconnectionsDeleteAccepted with default headers values
func NewPcloudVpnconnectionsDeleteAccepted() *PcloudVpnconnectionsDeleteAccepted {
	return &PcloudVpnconnectionsDeleteAccepted{}
}

/*PcloudVpnconnectionsDeleteAccepted handles this case with default header values.

Accepted
*/
type PcloudVpnconnectionsDeleteAccepted struct {
	Payload *models.JobReference
}

func (o *PcloudVpnconnectionsDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteAccepted  %+v", 202, o.Payload)
}

func (o *PcloudVpnconnectionsDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsDeleteBadRequest creates a PcloudVpnconnectionsDeleteBadRequest with default headers values
func NewPcloudVpnconnectionsDeleteBadRequest() *PcloudVpnconnectionsDeleteBadRequest {
	return &PcloudVpnconnectionsDeleteBadRequest{}
}

/*PcloudVpnconnectionsDeleteBadRequest handles this case with default header values.

Bad Request
*/
type PcloudVpnconnectionsDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsDeleteUnauthorized creates a PcloudVpnconnectionsDeleteUnauthorized with default headers values
func NewPcloudVpnconnectionsDeleteUnauthorized() *PcloudVpnconnectionsDeleteUnauthorized {
	return &PcloudVpnconnectionsDeleteUnauthorized{}
}

/*PcloudVpnconnectionsDeleteUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudVpnconnectionsDeleteUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsDeleteForbidden creates a PcloudVpnconnectionsDeleteForbidden with default headers values
func NewPcloudVpnconnectionsDeleteForbidden() *PcloudVpnconnectionsDeleteForbidden {
	return &PcloudVpnconnectionsDeleteForbidden{}
}

/*PcloudVpnconnectionsDeleteForbidden handles this case with default header values.

Forbidden
*/
type PcloudVpnconnectionsDeleteForbidden struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsDeleteNotFound creates a PcloudVpnconnectionsDeleteNotFound with default headers values
func NewPcloudVpnconnectionsDeleteNotFound() *PcloudVpnconnectionsDeleteNotFound {
	return &PcloudVpnconnectionsDeleteNotFound{}
}

/*PcloudVpnconnectionsDeleteNotFound handles this case with default header values.

Not Found
*/
type PcloudVpnconnectionsDeleteNotFound struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVpnconnectionsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsDeleteInternalServerError creates a PcloudVpnconnectionsDeleteInternalServerError with default headers values
func NewPcloudVpnconnectionsDeleteInternalServerError() *PcloudVpnconnectionsDeleteInternalServerError {
	return &PcloudVpnconnectionsDeleteInternalServerError{}
}

/*PcloudVpnconnectionsDeleteInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}][%d] pcloudVpnconnectionsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
