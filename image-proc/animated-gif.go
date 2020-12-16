package main

import "image"
import "image/gif"
import "os"

func main() {
    files := []string{"g1.gif", "g2.gif","g3.gif", "g2.gif"}

    // load static image and construct outGif
    outGif := &gif.GIF{}
    for _, name := range files {
        f, _ := os.Open(name)
        inGif, _ := gif.Decode(f)
        f.Close()

        outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
        outGif.Delay = append(outGif.Delay, 10)
    }

    // save to out.gif
    f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    gif.EncodeAll(f, outGif)
}
