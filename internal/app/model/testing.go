package model

import "testing"

//TestUser ..
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

//TestDevice ...
func TestDevice(t *testing.T) *Device {
	return &Device{
		IP:          "10.61.129.144",
		Place:       "ЧНГКМ ВЗИС73км - Серврная1",
		Description: "Сервер 1С",
		MethodCheck: "ping",
	}
}

//TestDevice2 ...
func TestDevice2(t *testing.T) *Device {
	return &Device{
		IP:          "10.61.129.142",
		Place:       "ЧНГКМ ВЗИС73км - Серврная2",
		Description: "Сервер 2С",
		MethodCheck: "ping",
	}
}
