package main

import (
    "gocv.io/x/gocv"
)

func main() {
    webcam, _ := gocf.VideoCaptureDevice(0)
    window := gocv.NewWindow("Howdy")
    img := gocv.NewMat()

    for {
        webcam.Read(&img)
        window.IMShow(img)
        window.WaitKey(1)
    }
}
