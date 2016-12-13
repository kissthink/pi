package db

import (
	"github.com/boltdb/bolt"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"encoding/binary"
	"errors"
)

type Device_t struct {
	ID		uint64		`json:"id"`
	Name		string		`json:"name" binding:"required,alphanum"`
	Status		string		`json:"status"`
	Password	string		`json:"password" binding:"required"`
	Description	string		`json:"description"`
	UserName	string		`json:"username"`
}

func Itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func (d *Device_t) Update() error {
	return session.Update(func (tx *bolt.Tx) error {
		b := tx.Bucket(device)

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		d.Password = string(hashedPassword)

		buf, err := json.Marshal(d)
		if err != nil {
			return err
		}

		return b.Put(Itob(d.ID), buf)
	})
}

func (d *Device_t) Create() error {
	return session.Update(func (tx *bolt.Tx) error {
		b := tx.Bucket(device)
		id, _ := b.NextSequence()
		d.ID = id

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		d.Password = string(hashedPassword)

		buf, err := json.Marshal(d)
		if err != nil {
			return err
		}

		return b.Put(Itob(id), buf)
	})
}

func (d *Device_t) Find() error {
	return session.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(device)
		device := b.Get(Itob(d.ID))
		if len(device) == 0 {
			return errors.New("Device not found")
		}

		return json.Unmarshal(device, d)
	})
}

func (d *Device_t) Delete() error {
	return session.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(device)
		return b.Delete(Itob(d.ID))
	})
}

func (d *Device_t) List() (*[]Device_t, error) {
	devices := make([]Device_t, 0)

	err := session.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(device)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			dv := Device_t{}
			err := json.Unmarshal(v, &dv)
			if err != nil {
				return err
			}
			if dv.UserName == d.UserName {
				devices = append(devices, dv)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &devices, nil
}