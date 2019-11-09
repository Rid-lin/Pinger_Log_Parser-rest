package sqlitestore

import (
	"database/sql"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store"
)

//UserRepository ..
type UserRepository struct {
	store *Store
}

//Create ..
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	additemSQL := "INSERT INTO users (email, encrypted_password) VALUES (?, ?)"

	stmt, err := r.store.db.Prepare(additemSQL)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(u.Email, u.EncryptedPassword)
	if err2 != nil {
		return err2
	}

	// return r.store.db.QueryRow(
	// 	"INSERT INTO users (email, encrypted_password) VALUES (?, ?)",
	// 	u.Email,
	// 	u.EncryptedPassword,
	// ).Scan(&u.ID); err != nil {
	// 	return nil, err
	// }
	return nil
}

//FindByEmail ..
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = ?",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
	}
	return u, nil
}
