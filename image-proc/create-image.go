package main

import (
    "fmt"
    "image"
    "image/color"
    "image/jpeg"
    "os"
)

const (
    width = 1080
    height = 1080
)

var filename = "out.jpg"

func clear(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            img.Set(x, y, color.RGBA{255, 255, 255, 255})
        }
    }
}

func blur(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            img.Set(x, y, color.RGBA{0, 0, 255, 255})

            r := uint8(255 - y)
            b := uint8(y)

            a := uint8(255 - 128)
            img.Set(x, y, color.RGBA{r, 0, b, a})
        }
    }
}

func main() {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    clear(img)
    blur(img)

    file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Printf("error creating image file: %s\n", err)
        return
    }
    defer file.Close()

    jpeg.Encode(file, img, &jpeg.Options{ Quality:100 })
    fmt.Printf("wrote image to %s\n", filename)
}

