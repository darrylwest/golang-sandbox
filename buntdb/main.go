package main

import (
    "fmt"
    "time"

    "github.com/tidwall/buntdb"
)

func  main() {
    db, err := buntdb.Open("bunt.db")
    if err != nil {
        fmt.Printf("error opening db: %s\n", err)
        return
    }

    defer db.Close()

    mykey := fmt.Sprintf("mykey:%d", time.Now().Unix())

    err = db.Update(func(tx *buntdb.Tx) error {
        _, _, err := tx.Set(mykey, "my value", nil)
        return err
    })

    err = db.View(func(tx *buntdb.Tx) error {
        val, err := tx.Get(mykey)
        if err != nil {
            return err
        }
        fmt.Printf("value for '%s' is '%s'\n", mykey, val)
        return err
    })

    err = db.View(func(tx *buntdb.Tx) error {
        err := tx.Ascend("", func(key, value string) bool {
            fmt.Printf("key: %s, value: %s\n", key, value)
            return true
        })
        return err
    })

    time.Sleep(time.Second)
}
