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
	// Логирование результата
	d.LogStatus(status, patchWorkLogs)
}

// CheckStatus ...
func (d *Device) CheckStatus() string {
	// В зависимости от указанного метода проверки вызываем нужный
	switch d.MethodCheck {
	case "ping": // пока что только пинг
		return check.PingIP(d.IP)
	default:
		return check.PingIP(d.IP)
	}

	// return ""
}

// LogStatus добавляет полученную строку "status" в файл находящийся по пути "patchWorkLogs"
func (d *Device) LogStatus(status, patchWorkLogs string) {
	// f, err := os.OpenFile(patchWorkLogs, os.O_APPEND|os.O_WRONLY, 0600)
	f, err := os.OpenFile(patchWorkLogs, os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.WriteString(status); err != nil {
		log.Fatal(err)
	}
}
