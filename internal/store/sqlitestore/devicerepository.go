package sqlitestore

import (
	"database/sql"
	"errors"

	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/app/model"
	"github.com/Rid-lin/Pinger_Log_Parser-rest/internal/store"
)

//DeviceRepository ..
type DeviceRepository struct {
	store   *Store
	devices map[int]*model.Device
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

//Find ..
func (r *DeviceRepository) Find(id int) (*model.Device, error) {
	d := &model.Device{}
	if err := r.store.db.QueryRow(
		"SELECT id, ip, place, description, methodcheck FROM devices WHERE id = ?",
		id,
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

//DeleteByIP ..
func (r *DeviceRepository) DeleteByIP(ip string) error {
	d, err := r.FindByIP(ip)
	if err != nil {
		return err
	}

	delitemSQL := "DELETE FROM devices WHERE id = ?"

	stmt, err := r.store.db.Prepare(delitemSQL)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(d.ID)
	if err2 != nil {
		return err2
	}

	return nil
}

//Delete ..
func (r *DeviceRepository) Delete(d *model.Device) error {
	if d == nil || d.IP == "0" || d.IP == "" || d.ID == 0 {
		return errors.New("Invalid input data")
	}

	delitemSQL := "DELETE FROM devices WHERE id = ?"

	stmt, err := r.store.db.Prepare(delitemSQL)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(d.ID)
	if err2 != nil {
		return err2
	}

	return nil
}

//Update ..
func (r *DeviceRepository) Update(dOld, dNew *model.Device) error {
	if err := dNew.Validate(); err != nil {
		return err
	} else if dNew.ID != dOld.ID {
		return errors.New("Invalid input data")
	}

	// обновляем строку с id=
	_, err := r.store.db.Exec(
		"UPDATE devices SET ip = ?, place = ?, description = ?, methodcheck = ? WHERE id = ?",
		dNew.IP, dNew.Place, dNew.Description, dNew.MethodCheck, dNew.ID)
	if err != nil {
		return err
	}
	return nil
}

//UpdateByIP ...
func (r *DeviceRepository) UpdateByIP(ip string, dNew *model.Device) error {
	d, err := r.FindByIP(ip)
	if err != nil {
		return err
	}

	// обновляем строку с id=1
	_, err2 := r.store.db.Exec(
		"UPDATE devices SET ip = ?, place = ?, description = ?, methodcheck = ? WHERE id = ?",
		dNew.IP, dNew.Place, dNew.Description, dNew.MethodCheck, d.ID)
	if err2 != nil {
		return err2
	}

	return nil
}

//GetAllAsMap ..
func (r *DeviceRepository) GetAllAsMap() (map[int](*model.Device), error) {
	d := &model.Device{}
	// query
	rows, err := r.store.db.Query("SELECT * FROM devices")
	if err != nil {
		return nil, err
	}

	devices := make(map[int](*model.Device))

	for rows.Next() {
		err = rows.Scan(&d.ID, &d.IP, &d.Place, &d.Description, &d.MethodCheck)
		if err != nil {
			return nil, err
		}
		devices[d.ID] = d
	}

	rows.Close() //good habit to close
	return devices, nil
}

//GetAllAsList ..
func (r *DeviceRepository) GetAllAsList() ([](*model.Device), error) {
	rows, err := r.store.db.Query("SELECT * FROM devices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devicesList [](*model.Device)

	for rows.Next() {
		d := &model.Device{}
		err := rows.Scan(&d.ID, &d.IP, &d.Place, &d.Description, &d.MethodCheck)
		if err != nil {
			return nil, err
		}
		devicesList = append(devicesList, d)
	}

	return devicesList, nil
}
