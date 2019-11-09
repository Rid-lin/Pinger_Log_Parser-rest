package teststore

import (
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store"
)

//UserRepository ..
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

//Create ..
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users)
	r.users[u.Email] = u

	return nil
}

//FindByEmail ..
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if ok {
		return u, nil
	}
	return nil, store.ErrRecordNotFound
}
