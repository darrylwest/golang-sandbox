package main

import (
    "fmt"
    "image"
    // "image/color"
    "image/jpeg"
    "os"
)

func init() {
    image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {
    filename := "./ebay.jpeg"
    ifile, err := os.Open(filename)
    if err != nil {
        fmt.Printf("error reading file: %s : %s\n", filename, err)
        os.Exit(1)
    }
    defer ifile.Close()

    img, name, err := image.Decode(ifile)

    co := img.At(10, 10)
    fmt.Printf("%s, color: %T %v\n", name, co, co)

    bounds := img.Bounds()

    fmt.Printf("bounds: %v\n", bounds)

    canvas := image.NewAlpha(bounds)
    op := canvas.Opaque()
    fmt.Printf("opaque: %v\n", op)
}
