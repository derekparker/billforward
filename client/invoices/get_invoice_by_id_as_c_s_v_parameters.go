package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
GetInvoiceByIDAsCSVParams contains all the parameters to send to the API endpoint
for the get invoice by ID as c s v operation typically these are written to a http.Request
*/
type GetInvoiceByIDAsCSVParams struct {
	/*
	  The ID of the invoice.
	*/
	ID string
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceByIDAsCSVParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param ID
	if err := r.SetPathParam("ID", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
