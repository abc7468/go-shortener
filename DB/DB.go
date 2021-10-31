package db

import (
	"github.com/boltdb/bolt"
)

const (
	dbName     = "shortener"
	dataBucket = "urls"
)

var db *bolt.DB

func InitDB() {
	if db == nil {
		dbPointer, err := bolt.Open("db.db", 0600, nil)
		db = dbPointer
		if err != nil {
			panic(err)
		}
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
			if err != nil {
				panic(err)
			}
			return err
		})
		if err != nil {
			panic(err)
		}
	}
}

func Close() {
	db.Close()

}

func SaveURL(shortener, original string) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(shortener), []byte(original))
		return err
	})
	if err != nil {
		panic(err)

	}
}
func LoadUrl(shortener string) []byte {
	var data []byte
	db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(shortener))
		return nil
	})
	return data
}
func LoadAllUrls() [][]string {
	var urls [][]string

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(dataBucket))

		b.ForEach(func(k, v []byte) error {
			val := make([]string, 2)
			val[0], val[1] = string(k), string(v)
			urls = append(urls, val)
			return nil
		})
		return nil
	})
	return urls
}
func DeleteUrl(shortener string) error {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Delete([]byte(shortener))
		return err
	})
	if err != nil {
		return err
	}
	return nil
}
