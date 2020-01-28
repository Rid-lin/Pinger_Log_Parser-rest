package sqlitestore_test

import (
	"testing"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store/sqlitestore"
	"github.com/stretchr/testify/assert"
	// _ "github.com/mattn/go-sqlite3" // ..
)

func TestDeviceRepository_Create(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("devices")

	s := sqlitestore.New(db)
	u := model.TestDevice(t)
	assert.NoError(t, s.Device().Create(u))
	assert.NotNil(t, u)
}

func TestDeviceRepository_FindByIP(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("devices")

	s := sqlitestore.New(db)
	u1 := model.TestDevice(t)
	_, err := s.Device().FindByIP(u1.IP)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Device().Create(u1)
	u2, err := s.Device().FindByIP(u1.IP)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}

func TestDeviceRepository_Find(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("devices")

	s := sqlitestore.New(db)
	u1 := model.TestDevice(t)
	_, err := s.Device().Find(u1.ID)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.Device().Create(u1)
	u2, err := s.Device().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}

func TestDeviceRepository_DeleteByIP(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("devices")

	s := sqlitestore.New(db)
	d1 := model.TestDevice(t)
	err := s.Device().Create(d1)
	assert.NoError(t, err)
	d2, err2 := s.Device().FindByIP(d1.IP)
	assert.NoError(t, err2)
	err3 := s.Device().DeleteByIP(d2.IP)
	assert.NoError(t, err3)
	d4, err4 := s.Device().FindByIP(d2.IP)
	assert.Error(t, err4)
	assert.Nil(t, d4)

}

func TestDeviceRepository_Delete(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("devices")

	s := sqlitestore.New(db)
	d1 := model.TestDevice(t)
	err := s.Device().Create(d1)
	assert.NoError(t, err)
	d2, err2 := s.Device().FindByIP(d1.IP)
	assert.NoError(t, err2)
	err3 := s.Device().Delete(d2)
	assert.NoError(t, err3)
	d4, err4 := s.Device().FindByIP(d2.IP)
	assert.Error(t, err4)
	assert.Nil(t, d4)

}

func TestDeviceRepository_Update(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("devices")

	s := sqlitestore.New(db)

	d := model.TestDevice(t)
	s.Device().Create(d)

	err := s.Device().Create(d)
	assert.NoError(t, err)
	assert.NotNil(t, d)

	d2 := model.TestDevice2(t)
	s.Device().Create(d2)
	d2.ID = d.ID

	err2 := s.Device().Update(d, d2)
	assert.NoError(t, err2)

	d3, err3 := s.Device().FindByIP(d2.IP)

	assert.NoError(t, err3)
	assert.NotNil(t, d3)
}

func TestDeviceRepository_UpdateByIP(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("devices")

	s := sqlitestore.New(db)

	d := model.TestDevice(t)
	err := s.Device().Create(d)

	assert.NoError(t, err)
	assert.NotNil(t, d)

	d2 := model.TestDevice2(t)
	s.Device().Create(d2)

	err2 := s.Device().UpdateByIP(d.IP, d2)
	assert.NoError(t, err2)

	d3, err4 := s.Device().FindByIP(d2.IP)
	assert.NoError(t, err4)
	assert.NotNil(t, d3)

}

func TestDeviceRepository_GetAll(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("devices")

	s := sqlitestore.New(db)

	d1 := model.TestDevice(t)
	err1 := s.Device().Create(d1)
	assert.NoError(t, err1)

	d2 := model.TestDevice(t)
	err2 := s.Device().Create(d2)
	assert.NoError(t, err2)

	devices, err := s.Device().GetAllAsMap()
	assert.NoError(t, err)
	assert.NotNil(t, devices)

}

//GetAllAsList ..
func TestDeviceRepository_GetAllAsList(t *testing.T) {
	db, teardown := sqlitestore.TestDB(t, databaseURL)
	defer teardown("devices")

	s := sqlitestore.New(db)

	d1 := model.TestDevice(t)
	err1 := s.Device().Create(d1)
	assert.NoError(t, err1)

	d2 := model.TestDevice(t)
	err2 := s.Device().Create(d2)
	assert.NoError(t, err2)

	devices, err := s.Device().GetAllAsList()
	assert.NoError(t, err)
	assert.NotNil(t, devices)
}
