package db

import (
	"github.com/boltdb/bolt"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type User_t struct {
	Name		string		`json:"name" binding:"required"`
	Email		string		`json:"email" binding:"required,email"`
	Password	string		`json:"-" binding:"required"`
}

func (u *User_t) Init(name string, email string, password string) {
	u.Name = name
	u.Email = email
	u.Password = password
}


func (u *User_t) Create() error {
	err := session.Update(func(tx *bolt.Tx) error {
		key := []byte(u.Name)
		b := tx.Bucket(user)
		exists := b.Get(key)
		if len(exists) != 0 {
			return errors.New("User already exists")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)

		usr, err := json.Marshal(u)
		if err != nil {
			return err
		}

		return b.Put(key, usr)
	})

	return err
}

func (u *User_t) Find() error {
	return session.View(func(tx *bolt.Tx) error {
		key := []byte(u.Name)
		b := tx.Bucket(user)
		usr := b.Get(key)
		if len(usr) == 0 {
			return errors.New("User not found")
		}

		return json.Unmarshal(usr, u)
	})
}

func (u *User_t) Delete() error {
	err := session.Update(func(tx *bolt.Tx) error {
		key := []byte(u.Name)
		b := tx.Bucket(user)
		return b.Delete(key)
	})

	return err
}

func CreateAdmin() error {
	adm := User_t{Name: "admin"}
	err := adm.Find()
	if err != nil {
		adm.Email = "admin@admin.com"
		adm.Password = "123456"
		return adm.Create()
	}

	return nil
}