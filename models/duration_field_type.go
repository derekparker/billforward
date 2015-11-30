package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*DurationFieldType duration field type

swagger:model DurationFieldType
*/
type DurationFieldType struct {

	/* Name name
	 */
	Name string `json:"name,omitempty"`
}

// Validate validates this duration field type
func (m *DurationFieldType) Validate(formats strfmt.Registry) error {
	return nil
}
