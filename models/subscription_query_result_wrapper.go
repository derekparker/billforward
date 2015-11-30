package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*SubscriptionQueryResultWrapper subscription query result wrapper

swagger:model SubscriptionQueryResultWrapper
*/
type SubscriptionQueryResultWrapper struct {

	/* {"description":"Paging parameter. 0-indexed. Describes your current location within a pageable list of query results.","verbs":["GET"]}
	 */
	CurrentOffset int32 `json:"currentOffset,omitempty"`

	/* {"default":10,"description":"Paging parameter. Describes how many records you requested.","verbs":["GET"]}
	 */
	RecordsRequested int32 `json:"recordsRequested,omitempty"`

	/* {"description":"Describes how many records were returned by your query.","verbs":["GET","POST"]}
	 */
	RecordsReturned int32 `json:"recordsReturned,omitempty"`

	/* {"description":"The results returned by your query.","verbs":["GET","PUT","POST"]}
	 */
	Results []*Subscription `json:"results,omitempty"`
}

// Validate validates this subscription query result wrapper
func (m *SubscriptionQueryResultWrapper) Validate(formats strfmt.Registry) error {
	return nil
}
