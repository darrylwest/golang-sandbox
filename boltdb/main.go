package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/darrylwest/go-unique/unique"
)

// Cache the cache structure
type Cache struct {
	db     *bolt.DB
	bucket []byte
}

// OpenDb will open the database for a single bucket, and create the bucket if necessary
func OpenDb(filename string, bucketName []byte) (*Cache, error) {

	db, err := bolt.Open(filename, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	// fmt.Printf("database '%s' opened...\n", db.Path());

	cache := Cache{db: db, bucket: bucketName}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			fmt.Printf("create bucket error: %s", err)
		}

		// fmt.Printf("bucket: %v\n", b);
		return nil
	})

	return &cache, err
}

// Put the value for the key
func (cache Cache) Put(key, value []byte) error {
	var err error
	cache.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(cache.bucket)
		err = b.Put(key, value)
		return err
	})

	return err
}

// Get the value for the key
func (cache Cache) Get(key []byte) ([]byte, error) {
	var value []byte
	var err error

	cache.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(cache.bucket)
		value = b.Get(key)
		return nil
	})

	return value, err
}

func main() {
	for {
		cache, err := OpenDb("production.db", []byte("Production"))
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		} else {

			key := []byte(unique.CreateULID())
			value := []byte(`{"status":"ok","ts":1494168301566,"version":"1.0"}`)
			if err = cache.Put(key, value); err != nil {
				log.Fatal(err)
			}

			val, _ := cache.Get(key)
			fmt.Printf("key: %s = %s\n", key, val)

			cache.db.Close()
			time.Sleep(5 * time.Second)
		}

	}
}
