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
		dbPointer, err := bolt.Open("myDb.db", 0600, nil)
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

func SaveURL(original, shortener string) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(shortener), []byte(original))
		return err
	})
	if err != nil {
		panic(err)
	}
}
