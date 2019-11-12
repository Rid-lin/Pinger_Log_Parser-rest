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
	FindByIP(string) (*model.Device, error)
	FindIDByIP(string) (int, error)
	DeleteByIP(string) error
	Update(string, *model.Device) error
	GetAll(string) (map[interface{}](*model.Device), error)
}
