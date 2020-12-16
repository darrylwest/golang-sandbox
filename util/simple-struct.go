package main

import (
	// "code.google.com/p/go-uuid/uuid"
	// "encoding/json"
	"fmt"
	// "strings"
	"reflect"
	"unsafe"
)

type UserStruct struct {
	id       string
	username string
	fullname string
	session  string
	status   string
}

func main() {
	var user UserStruct

	fmt.Println("user:", user)
	fmt.Println("size:", unsafe.Sizeof(user))

	v := reflect.TypeOf(user)
	fmt.Println("type:", v.Name())
}
