package store

import "github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"

//UserRepository ..
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	Find(int) (*model.User, error)
}
