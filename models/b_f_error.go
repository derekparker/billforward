package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
Standard BillForward error format.

swagger:model BFError
*/
type BFError struct {

	/* {"description":"Code describing the nature of the error. Currently unused; prefer `errorType`.","verbs":["GET","PUT","POST"]}
	 */
	ErrorCode int32 `json:"errorCode,omitempty"`

	/* {"description":"Human-readable description of the reason for the error.","verbs":["GET","PUT","POST"]}
	 */
	ErrorMessage string `json:"errorMessage,omitempty"`

	/* {"description":"List of erroneous parameters found in your input (if applicable).","verbs":["GET","PUT","POST"]}
	 */
	ErrorParameters []string `json:"errorParameters,omitempty"`

	/* {"description":"Enum categorizing the nature of the error.","verbs":["GET","PUT","POST"]}
	 */
	ErrorType string `json:"errorType,omitempty"`
}

// Validate validates this b f error
func (m *BFError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var bFErrorErrorTypeEnum []interface{}

func (m *BFError) validateErrorTypeEnum(path, location string, value string) error {
	if bFErrorErrorTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["BFError","ServerError","ValidationError","UnserializationException","Oauth","PermissionsError","PreconditionFailed","NotImplemented","InvocationError","NoSuchEntity","InconsistentState","StripeOperationFailure","BraintreeOperationFailure","BraintreeValidationError","SagePayOperationFailure","TokenizationAuthCaptureFailure","TokenizationPreAuthFailure","CouponException","CouponUniqueCodesRequestException","CouponUniqueCodesResponseException","RemoveCouponException","AddCouponCodeToSubscriptionRequestException","GatewayAuthenticationError","GatewayAuthorizationError","GatewayResourceNotFoundError","GatewayProtocolVersionError","GatewayInternalError","GatewayDownTemporarilyError","GatewayUnexpectedError","GatewayUnhandledError","GatewaySDKUnhandledError"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			bFErrorErrorTypeEnum = append(bFErrorErrorTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, bFErrorErrorTypeEnum)
}

func (m *BFError) validateErrorType(formats strfmt.Registry) error {

	if err := m.validateErrorTypeEnum("errorType", "body", m.ErrorType); err != nil {
		return err
	}

	return nil
}
