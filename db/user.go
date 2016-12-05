package db

import (
	"github.com/boltdb/bolt"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)

type user_t struct {
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
}

func createAdmin() error {
	err := session.Update(func(tx *bolt.Tx) error {
		key := []byte("admin")
		b := tx.Bucket(user)
		adm := b.Get(key)
		if len(adm) == 0 {
			hashedPassword, err := bcrypt.GenerateFromPassword(key, bcrypt.DefaultCost)
			if err != nil {
				return err
			}

			adm, err = json.Marshal(user_t{
				Name:		"admin",
				Email:		"admin@localhost",
				Password:	string(hashedPassword),
			})

			err = b.Put(key, adm)
			return err
		}

		return nil
	})

	return err
}