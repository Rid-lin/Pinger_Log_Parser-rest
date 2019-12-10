package teststore

import (
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

//FindIDByIP ..
func (r *DeviceRepository) FindIDByIP(ip string) (int, error) {
	for _, d := range r.devices {
		if d.IP == ip {
			return d.ID, nil
		}
	}
	return -1, store.ErrRecordNotFound
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

//Update ..
func (r *DeviceRepository) Update(ip string, dNew *model.Device) error {
	id, err := r.FindIDByIP(ip)
	if err != nil {
		return err
	}

	r.devices[id] = dNew

	return nil
}

//GetAll ..
func (r *DeviceRepository) GetAll(sortBy string) (map[interface{}](*model.Device), error) {
	devices := make(map[interface{}](*model.Device))

	for id, row := range r.devices {
		devices[id] = row
	}

	return devices, nil
}
