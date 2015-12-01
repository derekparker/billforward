package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

type GetSubscriptionByAccountIDReader struct {
	formats strfmt.Registry
}

func (o *GetSubscriptionByAccountIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result GetSubscriptionByAccountIDOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 500:
		var result GetSubscriptionByAccountIDInternalServerError
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getSubscriptionByAccountIdInternalServerError", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*
success
*/
type GetSubscriptionByAccountIDOK struct {
	Payload *models.SubscriptionQueryResultWrapper
}

func (o *GetSubscriptionByAccountIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionQueryResultWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*
error
*/
type GetSubscriptionByAccountIDInternalServerError struct {
	Payload *models.BFError
}

func (o *GetSubscriptionByAccountIDInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
