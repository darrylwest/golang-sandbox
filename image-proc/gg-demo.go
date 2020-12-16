package main

import (
    "fmt"
    "github.com/fogleman/gg"
)

func draw(name string, r, g, b float64) {
    fout := name
	const sz = 1080
	dc := gg.NewContext(sz, sz)
	dc.SetRGBA(r, g, b, 0.25)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), sz/2, sz/2)
		dc.DrawEllipse(sz/2, sz/2, sz*7/16, sz/8)
		dc.Fill()
		dc.Pop()
	}
	dc.SavePNG(fout)

    fmt.Println("saved file to", fout)
}

func main() {
    draw("image-1.png", 0, 0, 0)
    draw("image-2.png", 100, 0, 0)
    draw("image-3.png", 0, 100, 0)
    draw("image-4.png", 0, 0, 100)
}
