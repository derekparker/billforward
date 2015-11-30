package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*GetSubscriptionByAccountIDParams contains all the parameters to send to the API endpoint
for the get subscription by account ID operation typically these are written to a http.Request
*/
type GetSubscriptionByAccountIDParams struct {

	/*AccountID*/
	AccountID string
	/*IncludeRetired
	  Whether retired subscriptions should be returned.

	*/
	IncludeRetired bool
	/*Offset
	  The offset from the first subscription to return.

	*/
	Offset int32
	/*Order
	  Ihe direction of any ordering, either ASC or DESC.

	*/
	Order string
	/*OrderBy
	  Specify a field used to order the result set.

	*/
	OrderBy string
	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string
	/*Records
	  The maximum number of subscriptions to return.

	*/
	Records int32
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubscriptionByAccountIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param account-ID
	if err := r.SetPathParam("account-ID", o.AccountID); err != nil {
		return err
	}

	// query param include_retired
	if err := r.SetQueryParam("include_retired", swag.FormatBool(o.IncludeRetired)); err != nil {
		return err
	}

	// query param offset
	if err := r.SetQueryParam("offset", swag.FormatInt32(o.Offset)); err != nil {
		return err
	}

	// query param order
	if err := r.SetQueryParam("order", o.Order); err != nil {
		return err
	}

	// query param order_by
	if err := r.SetQueryParam("order_by", o.OrderBy); err != nil {
		return err
	}

	valuesOrganizations := o.Organizations

	// query array param organizations
	if err := r.SetQueryParam("organizations", swag.JoinByFormat(valuesOrganizations, "multi")...); err != nil {
		return err
	}

	// query param records
	if err := r.SetQueryParam("records", swag.FormatInt32(o.Records)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
