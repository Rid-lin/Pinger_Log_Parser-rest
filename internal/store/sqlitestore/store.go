package sqlitestore

import (
	"database/sql"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store"
	_ "github.com/mattn/go-sqlite3" //..
)

//Store ..
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
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
