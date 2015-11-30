package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Seconds seconds

swagger:model Seconds
*/
type Seconds struct {

	/* FieldType field type
	 */
	FieldType *DurationFieldType `json:"fieldType,omitempty"`

	/* PeriodType period type
	 */
	PeriodType *PeriodType `json:"periodType,omitempty"`

	/* Seconds seconds
	 */
	Seconds int32 `json:"seconds,omitempty"`
}

// Validate validates this seconds
func (m *Seconds) Validate(formats strfmt.Registry) error {
	return nil
}
