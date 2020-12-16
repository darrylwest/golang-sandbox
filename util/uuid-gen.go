package main

import (
	"fmt"
	"github.com/pborman/uuid"
)

func main() {
	uid := uuid.New()
	fmt.Println(uid)
}
