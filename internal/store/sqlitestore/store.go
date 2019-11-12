package sqlitestore

import (
	"database/sql"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store"
	_ "github.com/mattn/go-sqlite3" //..
)

//Store ..
type Store struct {
	db               *sql.DB
	userRepository   *UserRepository
	deviceRepository *DeviceRepository
}

//New ..
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//User ..
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

//Device ..
func (s *Store) Device() store.DeviceRepository {
	if s.deviceRepository != nil {
		return s.deviceRepository
	}

	s.deviceRepository = &DeviceRepository{
		store: s,
	}
	return s.deviceRepository
}
