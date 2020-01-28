package teststore

import (
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

	d.ID = len(r.devices) + 1
	r.devices[d.ID] = d

	return nil
}

//FindByIP ..
func (r *DeviceRepository) FindByIP(ip string) (*model.Device, error) {
	for _, d := range r.devices {
		if d.IP == ip {
			return d, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

//Find ..
func (r *DeviceRepository) Find(id int) (*model.Device, error) {
	for _, d := range r.devices {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

//DeleteByIP ..
func (r *DeviceRepository) DeleteByIP(ip string) error {
	for _, d := range r.devices {
		if d.IP == ip {
			delete(r.devices, d.ID)
			return nil
		}
	}

	return store.ErrRecordNotFound
}

//Delete ..
func (r *DeviceRepository) Delete(d *model.Device) error {
	d2, err := r.Find(d.ID)
	if err != nil {
		d2, err = r.FindByIP(d.IP)
		if err != nil {
			return store.ErrRecordNotFound
		}
		delete(r.devices, d2.ID)
		return nil
	}
	delete(r.devices, d2.ID)
	return nil
}

//Update ..
func (r *DeviceRepository) Update(dOld, dNew *model.Device) error {
	if err := dNew.Validate(); err != nil {
		return err
	} else if dNew.ID == dOld.ID {
		return errors.New("Invalid input data")
	}

	r.devices[dOld.ID] = dNew

	return nil
}

//UpdateByIP ..
func (r *DeviceRepository) UpdateByIP(ip string, dNew *model.Device) error {
	d, err := r.FindByIP(ip)
	if err != nil {
		return err
	}

	r.devices[d.ID] = dNew

	return nil
}

//GetAllAsMap ..
func (r *DeviceRepository) GetAllAsMap() (map[int](*model.Device), error) {
	devices := make(map[int](*model.Device))

	for id, row := range r.devices {
		devices[id] = row
	}

	return devices, nil
}

//GetAllAsList ..
func (r *DeviceRepository) GetAllAsList() ([](*model.Device), error) {
	var devicesList [](*model.Device)

	for _, device := range r.devices {
		devicesList = append(devicesList, device)
	}

	return devicesList, nil
}
