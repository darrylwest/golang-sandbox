package main

import (
	"fmt"
	"reflect"
	"time"
)

type DocumentIdentifierType interface {
	GetId() string
	GetDateCreated() time.Time
	GetLastUpdated() time.Time
	GetVersion() int64
}

type DocumentIdentifier struct {
	id          string
	dateCreated time.Time
	lastUpdated time.Time
	version     int64
}

func (doi DocumentIdentifier) GetId() string {
	return doi.id
}

type UserType interface {
	GetDoi() DocumentIdentifier
}

type User struct {
	doi      DocumentIdentifier
	username string
	session  string
	status   string
}

// returns a pointer to user
func createUser() *User {
	user := new(User)

	user.doi.id = "123412341234234"
	user.doi.dateCreated = time.Now().UTC()
	user.doi.lastUpdated = time.Now().UTC()
	user.doi.version = 1234

	user.username = "dpw@rcs.com"
	user.session = "43243243432432432"
	user.status = "active"

	return user
}

func main() {
	user := *createUser()

	tp := reflect.TypeOf(user)

	fmt.Printf("user type: %s\n", tp)

	field := tp.Field(0)
	fmt.Printf("field : %v\n", field)
	field = tp.Field(1)
	fmt.Printf("field : %v\n", field)
	field = tp.Field(2)
	fmt.Printf("field : %v\n", field)
	field = tp.Field(3)
	fmt.Printf("field : %v\n", field)
	field = tp.Field(4)
	fmt.Printf("field : %v\n", field)
}
