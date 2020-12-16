package main

import "github.com/fogleman/gg"

var text = "https://ebay.co/2HDdQar"

func main() {
	const S = 500
	dc := gg.NewContext(S, S / 2)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
    face := "MarketSans-Light.ttf"
	if err := dc.LoadFontFace(face, 36); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(text, S/2, S/4, 0.5, 0.5)
	dc.SavePNG("out.png")
}
