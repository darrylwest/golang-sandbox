
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    file, err := assets.Open("/index.html")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    bytes, err := ioutil.ReadAll(file)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", bytes)
}
