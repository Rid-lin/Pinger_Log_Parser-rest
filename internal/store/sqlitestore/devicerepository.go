package sqlitestore

import (
	"database/sql"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store"
)

//DeviceRepository ..
type DeviceRepository struct {
	store *Store
}

//Create ..
func (r *DeviceRepository) Create(d *model.Device) error {
	if err := d.Validate(); err != nil {
		return err
	}

	if err := d.BeforeCreate(); err != nil {
		return err
	}

	additemSQL := "INSERT INTO devices (ip, place, description, methodcheck) VALUES (?, ?, ?, ?)"

	stmt, err := r.store.db.Prepare(additemSQL)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(d.IP, d.Place, d.Description, "ping")
	if err2 != nil {
		return err2
	}
	id64, err3 := result.LastInsertId()
	if err3 != nil {
		return err3
	}
	d.ID = int(id64)

	return nil
}

//FindByIP ..
func (r *DeviceRepository) FindByIP(ip string) (*model.Device, error) {
	d := &model.Device{}
	if err := r.store.db.QueryRow(
		"SELECT id, ip, place, description, methodcheck FROM devices WHERE ip = ?",
		ip,
	).Scan(
		&d.ID,
		&d.IP,
		&d.Place,
		&d.Description,
		&d.MethodCheck,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
	}
	return d, nil
}

//FindIDByIP ..
func (r *DeviceRepository) FindIDByIP(ip string) (int, error) {
	d := &model.Device{}
	if err := r.store.db.QueryRow(
		"SELECT id, methodcheck FROM devices WHERE ip = ?",
		ip,
	).Scan(
		&d.ID,
		&d.MethodCheck,
	); err != nil {
		if err == sql.ErrNoRows {
			return -1, store.ErrRecordNotFound
		}
	}
	return d.ID, nil
}

//DeleteByIP ..
func (r *DeviceRepository) DeleteByIP(ip string) error {
	id, err := r.FindIDByIP(ip)
	if err != nil {
		return err
	}

	delitemSQL := "DELETE FROM devices WHERE id = ?"

	stmt, err := r.store.db.Prepare(delitemSQL)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(id)
	if err2 != nil {
		return err2
	}

	return nil
}

//Update ..
func (r *DeviceRepository) Update(ip string, dNew *model.Device) error {
	id, err := r.FindIDByIP(ip)
	if err != nil {
		return err
	}

	// обновляем строку с id=1
	_, err2 := r.store.db.Exec(
		"UPDATE devices SET (ip = $2, place = $3, description = $4, methodcheck = $5) WHERE id = $1",
		id, dNew.IP, dNew.Place, dNew.Description, dNew.MethodCheck)
	if err2 != nil {
		return err2
	}

	return nil
}

//GetAll ..
func (r *DeviceRepository) GetAll(sortBy string) (map[interface{}](*model.Device), error) {
	//...
	return nil, nil
}
