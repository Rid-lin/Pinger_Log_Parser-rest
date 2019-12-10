package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

//Device ...
type Device struct {
	ID          int    `json:"id"`
	IP          string `json:"ip"`
	Place       string `json:"place"`
	Description string `json:"description"`
	MethodCheck string `json:"methodcheck"`
}

//Validate ..
func (d *Device) Validate() error {
	return validation.ValidateStruct(
		d,
		validation.Field(&d.IP, validation.Required, is.IPv4),
		validation.Field(&d.Place, validation.Required),
	)
}

//BeforeCreate ..
func (d *Device) BeforeCreate() error {
	// if len(d.Password) > 0 {
	// 	enc, err := encryptString(d.Password)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	d.EncryptedPassword = enc
	// }
	return nil
}
