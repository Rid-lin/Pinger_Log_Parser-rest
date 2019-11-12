package model_test

import (
	"testing"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"
	"github.com/stretchr/testify/assert"
)

//TestUser_Validate ..
func TestDevice_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		d       func() *model.Device
		isValid bool
	}{
		{
			name: "valid",
			d: func() *model.Device {
				return model.TestDevice(t)
			},
			isValid: true,
		},
		{
			name: "empty IP",
			d: func() *model.Device {
				d := model.TestDevice(t)
				d.IP = ""

				return d
			},
			isValid: false,
		},
		{
			name: "IP invalid",
			d: func() *model.Device {
				d := model.TestDevice(t)
				d.IP = "invalid"

				return d
			},
			isValid: false,
		},
		{
			name: "empty place",
			d: func() *model.Device {
				d := model.TestDevice(t)
				d.Place = ""

				return d
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.d().Validate())
			} else {
				assert.Error(t, tc.d().Validate())
			}
		})
	}
}
