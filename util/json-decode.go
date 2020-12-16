
package main

import (
    "fmt"
    "time"
    "github.com/francoispqt/gojay"
)

type User struct {
    ID int
    Name string
    Email string
    Created time.Time
}

func (u *User) NKeys() int {
    return 4
}

func (u *User) UnmarshalObject(dec *gojay.Decoder, key string) error {

    switch key {
    case "id":
        return dec.AddInt(&u.ID)
    case "name":
        return dec.AddString(&u.Name)
    case "email":
        return dec.AddString(&u.Email)
    case "created":
        var dt string
        err := dec.AddString(&dt)
        if err != nil {
            return err
        }
        u.Created, err = time.Parse(time.RFC3339, dt)
        return err
    }

    return nil
}

func (u *User) MarshalObject(enc *gojay.Encoder) {
    enc.AddIntKey("id", u.ID)
    enc.AddStringKey("name", u.Name)
    enc.AddStringKey("email", u.Email)
    enc.AddStringKey("created", u.Created.Format(time.RFC3339))
}

func (u *User) IsNil() bool {
    return u == nil
}

func (u User) String() string {
    return fmt.Sprintf("id:%d,name:%s,email:%s,created:%s", u.ID, u.Name, u.Email, u.Created.Format(time.RFC3339))
}

func main() {
    u := &User{}
    d := []byte(`{"id":1,"name":"dpw","email":"dpw@ebay.com","created":"2018-04-27T10:07:03+00:00"}`)

    err := gojay.UnmarshalObject(d, u)
    if err != nil {
        fmt.Printf("ERROR! %s", err)
        return
    }

    fmt.Println(u)

    b, err := gojay.MarshalObject(u)
    if err != nil {
        fmt.Printf("Marshal error: %s\n", err)
        return 
    }

    fmt.Printf("%s\n", b)
}
