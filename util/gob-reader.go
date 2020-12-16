package main

import (
    // "bytes"
    "encoding/gob"
    "fmt"
    "os"
    "time"
)

const (
    format = "ID=%d DateCreated=%v Name=%q Status=%q\n"
    iso8601 = "2006-01-02T15:04:05-0700"
)

type User struct {
    ID int
    DateCreated time.Time
    Name string
    Status string
}

/*
func (u *User) UnmarshalBinary(data []byte) error {
    b := bytes.NewBuffer(data)
    _, err :=fmt.Fscanf(b, &u.ID, &u.DateCreated, &u.Name, &u.Status)

    return err
}
*/

func main() {
    var users []User

    filename := "gob-user.data"
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    enc := gob.NewDecoder(file)
    if err = enc.Decode(&users); err != nil {
        panic(err)
    }

    fmt.Printf("read from file %s complete, total users %d...\n", filename, len(users))
    for _, user := range users {
        fmt.Printf("%v\n", user)
    }
}
