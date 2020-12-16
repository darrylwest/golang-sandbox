package main

import (
    "log"
    "fmt"
    "time"
    "github.com/dgraph-io/badger"
    "github.com/darrylwest/go-unique/unique"
)

func handle(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func insert(db *badger.DB) ([]byte, error) {
    key := []byte(unique.CreateXUID())
    err := db.Update(func(txn *badger.Txn) error {
        val := []byte(fmt.Sprintf("this is my value %s", unique.CreateGUID()))
        err := txn.Set(key, val)

        return err
    })

    return key, err
}

func main() {
    opts := badger.DefaultOptions
    // opts.Dir = "/tmp/badger"
    // opts.ValueDir = "/tmp/badger"
    opts.Dir = "./bdb"
    opts.ValueDir = "./bdb"
    db, err := badger.Open(opts)
    handle(err)

    defer db.Close()

    keys := make([][]byte, 0)
    t0 := time.Now()
    for i := 0; i < 1000; i++ {
        if key, err := insert(db); err == nil {
            keys = append(keys, key)
        } else {
            handle(err)
        }
    }
    t1 := time.Now()

    count := 0
    err = db.View(func(txn *badger.Txn) error {
        opts := badger.DefaultIteratorOptions
        opts.PrefetchSize = 100
        it := txn.NewIterator(opts)
        defer it.Close()
        for it.Rewind(); it.Valid(); it.Next() {
            item := it.Item()
            // k := item.Key()
            err := item.Value(func(v []byte) error {
                // log.Printf("key=%s, val=%s\n", k, v)
                return nil
            })

            if err != nil {
                return err
            }

            count++
        }


        return nil
    })
    t2 := time.Now()

    fmt.Printf("inserted %d keys in %s\n", len(keys), t1.Sub(t0).String())
    fmt.Printf("fetched %d items in %s\n", count, t2.Sub(t1).String())
}

