package model

import (
	check "github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/checkmethods" //.
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"log"
	"os"
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

// CheckNLogStatus ...
func (d *Device) CheckNLogStatus(patchWorkLogs string) {
	// Проверка статуса одного устройства
	status := d.CheckStatus()
	d.LogStatus(status, patchWorkLogs)
}

// CheckStatus ...
func (d *Device) CheckStatus() string {
	switch d.MethodCheck {
	case "ping":
		check.Ping(d.IP)
	default:
		check.Ping(d.IP)
	}

	return ""
}

// LogStatus ...
func (d *Device) LogStatus(status, patchWorkLogs string) {
	f, err := os.OpenFile(patchWorkLogs, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.WriteString(status); err != nil {
		log.Fatal(err)
	}
}
