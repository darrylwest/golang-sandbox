package main

import (
    "io"
    "os"
    "strings"
)

func main() {
    r1 := strings.NewReader("data from reader 1")
    dots := strings.NewReader("...")
    r2 := strings.NewReader("data from reader 2")

    reader := io.MultiReader(r1, dots, r2)
    io.Copy(os.Stdout, reader)

}
