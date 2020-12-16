package main

import (
    "image"
    "log"
    "github.com/disintegration/imaging"
)

func resize(src image.Image) image.Image {
    width := 1080
    height := 1080

    size := src.Bounds().Size()
    log.Printf("size: %v\n", size)

    // if the image is toll, set width to zero
    // if the image is wide, set the height to zero
    if size.X < size.Y {
        width = 0
    } else {
        height = 0
    }

    return imaging.Resize(src, width, height, imaging.Lanczos)
}

func main() {
    src, err := imaging.Open("dress.png")
    if err != nil {
        log.Fatalf("failed to open image: %v", err)
    }


    background := imaging.Resize(src, 1080, 1080, imaging.Lanczos)
    background = imaging.Blur(background, 50.0)

    // based on the src size, determine how to resize
    scaled := resize(src)

    // now do a centered composite
    dst := imaging.OverlayCenter(background, scaled, 1.0)
    err = imaging.Save(dst, "composite.png")
    if err != nil {
        log.Fatalf("failed to saveimage: %v", err)
    }
}
