package db

import "github.com/boltdb/bolt"

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

