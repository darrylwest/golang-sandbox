package main

// BROKEN...

import (
    "fmt"
    "github.com/tfriedel6/canvas"
    "image"
    "image/jpeg"
)

func init() {
    image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {
    filename := "./ebay.jpeg"
    img, err := canvas.LoadImage(filename)
    if err != nil {
        fmt.Printf("image %s load error: %s\n", filename, err)
        return
    }

    fmt.Printf("width: %d, height: %d\n", img.Width(), img.Height())
}
