package main

import (
    "fmt"
    "os"

    "github.com/syndtr/goleveldb/leveldb"
)

func main() {
    db, err := leveldb.OpenFile(os.Getenv("HOME") + "/.spotcache/cachedb", nil)

    if err != nil {
        panic(err)
    }

    defer db.Close()

    value := []byte("My Test Value")
    key := []byte("MyTestKey")

    data, err := db.Get(key, nil)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("data: %s\n", data)

    err = db.Put(key, value, nil)
    if err != nil {
        panic(err)
    }

    data, err = db.Get(key, nil)
    fmt.Printf("%s = %s\n", key, value)
    has, _ := db.Has(key, nil)
    fmt.Printf("has key: '%s'? %v\n", key, has)

    key = []byte("second:key")
    db.Put(key, value, nil)

    data, err = db.Get(key, nil)
    fmt.Printf("%s = %s\n", key, value)

    err = db.Delete(key, nil)
    if err != nil {
        panic(err)
    }

    data, err = db.Get(key, nil)
    fmt.Printf("deleted data (should be empty byte array): %v\n", data)
    fmt.Printf("deleted err: %v\n", err)

    has, err = db.Has(key, nil)
    fmt.Printf("has key: '%s'? %v, err: %v\n", key, has, err)
}
