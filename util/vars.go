//usr/bin/env go run $0 $@ ; exit

package main

import (
	"fmt"
	"os"
)

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
    xyz string
)

func main() {
	fmt.Println(home)
	fmt.Println(user)
	fmt.Println(gopath)

    // go run -ldflags "-X main.xyz=my-test" vars.go
    fmt.Printf("xyz = %s\n", xyz) 
}
