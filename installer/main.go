package main

import (
    "fmt"

    "./install"
)

func main() {
    data, err := install.Asset("dist/test-file.txt")
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", data)
}
