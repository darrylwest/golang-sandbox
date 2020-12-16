package main

import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

const (
    cost int = 12
)

func main() {
    fmt.Println("default cost:", bcrypt.DefaultCost, "my cost:", cost)
    fmt.Println("ya, here we go...")
}
