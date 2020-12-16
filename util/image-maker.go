package main

import (
    "os"
    "github.com/campoy/tools/imgcat"
)

func main() {
    enc, err := imgcat.NewEncoder(os.Stdout, imgcat.Width(imgcat.Pixels(600)), imgcat.Inline(true), imgcat.Name("sasha.jpg"))
    if err != nil {
        panic(err)
    }

    f, err := os.Open("sasha.jpg")
    if err != nil {
        panic(err)
    }
    defer func () { _ = f.Close() }()

    if err := enc.Encode(f); err != nil {
        panic(err)
    }
}
