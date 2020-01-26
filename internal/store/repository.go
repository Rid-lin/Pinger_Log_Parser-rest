package store

import "github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"

//UserRepository ..
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	Find(int) (*model.User, error)
}

//DeviceRepository ...
type DeviceRepository interface {
	Create(*model.Device) error
	Delete(*model.Device) error
	Update(*model.Device, *model.Device) error
	// Check(*model.Device) error

	FindByIP(string) (*model.Device, error)
	FindIDByIP(string) (int, error)
	DeleteByIP(string) error
	UpdateByIP(string, *model.Device) error

	GetAllAsMap() (map[int](*model.Device), error)
	GetAllAsList() ([](*model.Device), error)
}
