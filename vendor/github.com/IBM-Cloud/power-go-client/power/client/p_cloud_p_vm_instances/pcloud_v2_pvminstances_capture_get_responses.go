// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV2PvminstancesCaptureGetReader is a Reader for the PcloudV2PvminstancesCaptureGet structure.
type PcloudV2PvminstancesCaptureGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2PvminstancesCaptureGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudV2PvminstancesCaptureGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPcloudV2PvminstancesCaptureGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudV2PvminstancesCaptureGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudV2PvminstancesCaptureGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudV2PvminstancesCaptureGetOK creates a PcloudV2PvminstancesCaptureGetOK with default headers values
func NewPcloudV2PvminstancesCaptureGetOK() *PcloudV2PvminstancesCaptureGetOK {
	return &PcloudV2PvminstancesCaptureGetOK{}
}

/*PcloudV2PvminstancesCaptureGetOK handles this case with default header values.

OK
*/
type PcloudV2PvminstancesCaptureGetOK struct {
	Payload *models.Job
}

func (o *PcloudV2PvminstancesCaptureGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCaptureGetOK  %+v", 200, o.Payload)
}

func (o *PcloudV2PvminstancesCaptureGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesCaptureGetUnauthorized creates a PcloudV2PvminstancesCaptureGetUnauthorized with default headers values
func NewPcloudV2PvminstancesCaptureGetUnauthorized() *PcloudV2PvminstancesCaptureGetUnauthorized {
	return &PcloudV2PvminstancesCaptureGetUnauthorized{}
}

/*PcloudV2PvminstancesCaptureGetUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudV2PvminstancesCaptureGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudV2PvminstancesCaptureGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCaptureGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2PvminstancesCaptureGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesCaptureGetNotFound creates a PcloudV2PvminstancesCaptureGetNotFound with default headers values
func NewPcloudV2PvminstancesCaptureGetNotFound() *PcloudV2PvminstancesCaptureGetNotFound {
	return &PcloudV2PvminstancesCaptureGetNotFound{}
}

/*PcloudV2PvminstancesCaptureGetNotFound handles this case with default header values.

Not Found
*/
type PcloudV2PvminstancesCaptureGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudV2PvminstancesCaptureGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCaptureGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2PvminstancesCaptureGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesCaptureGetInternalServerError creates a PcloudV2PvminstancesCaptureGetInternalServerError with default headers values
func NewPcloudV2PvminstancesCaptureGetInternalServerError() *PcloudV2PvminstancesCaptureGetInternalServerError {
	return &PcloudV2PvminstancesCaptureGetInternalServerError{}
}

/*PcloudV2PvminstancesCaptureGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudV2PvminstancesCaptureGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudV2PvminstancesCaptureGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCaptureGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2PvminstancesCaptureGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
