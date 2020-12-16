package main

import (
    "fmt"
    "log"
    "time"

    "github.com/boltdb/bolt"
    // "github.com/darrylwest/go-unique/unique"
)

type Cache struct {
    db *bolt.DB
    bucket []byte
}

func OpenDb(filename string, bucketName []byte) (*Cache, error) {

    fmt.Printf("open database '%s'...\n", filename);

    opts := &bolt.Options{}
    opts.ReadOnly = true
    opts.Timeout = 1 * time.Second
    // fmt.Printf("bolt options: %v\n", opts)

    db, err := bolt.Open(filename, 0600, opts)
    if err != nil {
        return nil, err
    }

    fmt.Printf("database '%s' opened...\n", db.Path());

    cache := Cache{ db:db, bucket:bucketName }

    return &cache, err
}

func (cache Cache) Query() ([]byte, error) {
    var value []byte
    var err error

    cache.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket(cache.bucket)

        b.ForEach(func(k, v []byte) error {
            fmt.Printf("%s = %s\n", k, v)
            return nil
        })
        
        return nil
    })

    return value, err
}

func main() {
    for {
        // file := "/data/hub-backup.db"
        file := "production.db"
        cache, err := OpenDb(file, []byte("Production"))
        if err != nil {
            log.Fatal(err)
            time.Sleep(1 * time.Second)
        } else {
            cache.Query()
            cache.db.Close()
            time.Sleep(5 * time.Second)
        }
    }
}
