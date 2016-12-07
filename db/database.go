package db

import (
	"github.com/boltdb/bolt"
	"github.com/satori/go.uuid"
)

var session *bolt.DB

func OpenDatabase(path string) error {
	var err error
	session, err = bolt.Open(path, 0600, nil)
	if err != nil {
		return err
	}

	err = session.Update(func(tx *bolt.Tx) error {
		for _, buck := range buckets {
			_, err := tx.CreateBucketIfNotExists(buck)
			if err != nil {
				return err
			}
		}

		return err
	})
	if err != nil {
		return err
	}

	return err
}

func GetSecret() (secret *string, rerr error) {
	err := session.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(general)
		key := []byte("secret")
		s := b.Get(key)
		if len(s) != 0 {
			sstr := string(s)
			secret = &sstr
			return nil
		}

		suuid := uuid.NewV4()
		err := b.Put(key, suuid.Bytes())
		if err != nil {
			return err
		}

		sstr := suuid.String()
		secret = &sstr
		return nil
	})

	if err != nil {
		rerr = err
		return
	}

	return
}
