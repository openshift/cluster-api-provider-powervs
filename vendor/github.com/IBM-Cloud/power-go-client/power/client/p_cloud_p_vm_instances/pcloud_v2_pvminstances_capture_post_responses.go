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

// PcloudV2PvminstancesCapturePostReader is a Reader for the PcloudV2PvminstancesCapturePost structure.
type PcloudV2PvminstancesCapturePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2PvminstancesCapturePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewPcloudV2PvminstancesCapturePostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudV2PvminstancesCapturePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudV2PvminstancesCapturePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudV2PvminstancesCapturePostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPcloudV2PvminstancesCapturePostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudV2PvminstancesCapturePostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudV2PvminstancesCapturePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudV2PvminstancesCapturePostAccepted creates a PcloudV2PvminstancesCapturePostAccepted with default headers values
func NewPcloudV2PvminstancesCapturePostAccepted() *PcloudV2PvminstancesCapturePostAccepted {
	return &PcloudV2PvminstancesCapturePostAccepted{}
}

/*PcloudV2PvminstancesCapturePostAccepted handles this case with default header values.

Accepted, pvm-instance capture successfully added to the jobs queue
*/
type PcloudV2PvminstancesCapturePostAccepted struct {
	Payload *models.JobReference
}

func (o *PcloudV2PvminstancesCapturePostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCapturePostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudV2PvminstancesCapturePostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesCapturePostBadRequest creates a PcloudV2PvminstancesCapturePostBadRequest with default headers values
func NewPcloudV2PvminstancesCapturePostBadRequest() *PcloudV2PvminstancesCapturePostBadRequest {
	return &PcloudV2PvminstancesCapturePostBadRequest{}
}

/*PcloudV2PvminstancesCapturePostBadRequest handles this case with default header values.

Bad Request
*/
type PcloudV2PvminstancesCapturePostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudV2PvminstancesCapturePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCapturePostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2PvminstancesCapturePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesCapturePostUnauthorized creates a PcloudV2PvminstancesCapturePostUnauthorized with default headers values
func NewPcloudV2PvminstancesCapturePostUnauthorized() *PcloudV2PvminstancesCapturePostUnauthorized {
	return &PcloudV2PvminstancesCapturePostUnauthorized{}
}

/*PcloudV2PvminstancesCapturePostUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudV2PvminstancesCapturePostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudV2PvminstancesCapturePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCapturePostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2PvminstancesCapturePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesCapturePostNotFound creates a PcloudV2PvminstancesCapturePostNotFound with default headers values
func NewPcloudV2PvminstancesCapturePostNotFound() *PcloudV2PvminstancesCapturePostNotFound {
	return &PcloudV2PvminstancesCapturePostNotFound{}
}

/*PcloudV2PvminstancesCapturePostNotFound handles this case with default header values.

pvm instance id not found
*/
type PcloudV2PvminstancesCapturePostNotFound struct {
	Payload *models.Error
}

func (o *PcloudV2PvminstancesCapturePostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCapturePostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2PvminstancesCapturePostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesCapturePostConflict creates a PcloudV2PvminstancesCapturePostConflict with default headers values
func NewPcloudV2PvminstancesCapturePostConflict() *PcloudV2PvminstancesCapturePostConflict {
	return &PcloudV2PvminstancesCapturePostConflict{}
}

/*PcloudV2PvminstancesCapturePostConflict handles this case with default header values.

Conflict, a conflict has prevented adding the pvm-instance capture job
*/
type PcloudV2PvminstancesCapturePostConflict struct {
	Payload *models.Error
}

func (o *PcloudV2PvminstancesCapturePostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCapturePostConflict  %+v", 409, o.Payload)
}

func (o *PcloudV2PvminstancesCapturePostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesCapturePostUnprocessableEntity creates a PcloudV2PvminstancesCapturePostUnprocessableEntity with default headers values
func NewPcloudV2PvminstancesCapturePostUnprocessableEntity() *PcloudV2PvminstancesCapturePostUnprocessableEntity {
	return &PcloudV2PvminstancesCapturePostUnprocessableEntity{}
}

/*PcloudV2PvminstancesCapturePostUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudV2PvminstancesCapturePostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudV2PvminstancesCapturePostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCapturePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudV2PvminstancesCapturePostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2PvminstancesCapturePostInternalServerError creates a PcloudV2PvminstancesCapturePostInternalServerError with default headers values
func NewPcloudV2PvminstancesCapturePostInternalServerError() *PcloudV2PvminstancesCapturePostInternalServerError {
	return &PcloudV2PvminstancesCapturePostInternalServerError{}
}

/*PcloudV2PvminstancesCapturePostInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudV2PvminstancesCapturePostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudV2PvminstancesCapturePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture][%d] pcloudV2PvminstancesCapturePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2PvminstancesCapturePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
