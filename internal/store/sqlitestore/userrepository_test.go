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
	u1 := model.TestUser(t)
	_, err := s.User().FindByEmail(u1.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlitestore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}
