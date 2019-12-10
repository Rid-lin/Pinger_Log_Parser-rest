package teststore

import (
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store"
)

//Store ..
type Store struct {
	userRepository   *UserRepository
	deviceRepository *DeviceRepository
}

//New ..
func New() *Store {
	return &Store{}
}

//User ..
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}
	return s.userRepository
}

//Device ..
func (s *Store) Device() store.DeviceRepository {
	if s.deviceRepository != nil {
		return s.deviceRepository
	}

	s.deviceRepository = &DeviceRepository{
		store:   s,
		devices: make(map[int]*model.Device),
	}
	return s.deviceRepository
}
