
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    if content, err := ioutil.ReadFile("test-file.txt"); err != nil {
        panic(err)
    } else {
        fmt.Printf("%s\n", content)
        fmt.Printf("Content type: %T\n", content)
    }
}

