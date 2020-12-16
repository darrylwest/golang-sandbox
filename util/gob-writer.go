package main

import (
    // "bytes"
    "encoding/gob"
    "fmt"
    "math/rand"
    "os"
    "time"
)

type User struct {
    ID int
    DateCreated time.Time
    Name string
    Status string
}

const (
    format = "ID=%d DateCreated=%v Name=%q Status=%q\n"
    iso8601 = "2006-01-02T15:04:05-0700"
)

func getDate() time.Time {
    days := rand.Intn(365) * -1
    return time.Now().AddDate(0, 0, days)
}

/*
func (u User) MarshalBinary() ([]byte, error) {
    var b bytes.Buffer
    _, err := fmt.Fprintf(&b, format, u.ID, u.DateCreated.Format(iso8601), u.Name, u.Status)
    return b.Bytes(), err
}
*/

func main() {
    users := []User{
        User{1, getDate(), "dpw", "active"},
        User{2, getDate(), "john", "inactive"},
        User{3, getDate(), "sally", "active"},
        User{4, getDate(), "jane", "active"},
    }

    fmt.Printf("%v\n", users)

    filename := "gob-user.data"
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    enc := gob.NewEncoder(file)
    if err = enc.Encode(users); err != nil {
        panic(err)
    }

    fmt.Printf("write to file %s complete...\n", filename)
}
