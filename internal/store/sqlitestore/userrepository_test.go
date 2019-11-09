package sqlitestore_test

import (
	"testing"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store/sqlitestore"
	"github.com/stretchr/testify/assert"
	// _ "github.com/mattn/go-sqlite3" // ..
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlitestore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlitestore.New(db)
	email := "user@examle.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
