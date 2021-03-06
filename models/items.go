package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Items map of items

swagger:model items
*/
type Items map[string]Item

// Validate validates this items
func (m Items) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", Items(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
