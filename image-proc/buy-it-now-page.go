package main

import (
    "log"
    "github.com/fogleman/gg"
    "github.com/disintegration/imaging"
)

func drawThumb(dc *gg.Context) {
    sz := 380

    thumb, err := gg.LoadPNG("dress-composite.png")
    if err != nil {
        log.Fatalf("failed to open image: %v", err)
    }

    scaled := imaging.Resize(thumb, sz, sz, imaging.Lanczos)

    dc.DrawImage(scaled, 80, 80)
}

func drawBackground(dc *gg.Context) {
    // change to get blurred background and paint at 5% alpha...
    dc.SetRGBA(0.5, 0.5, 0.5, 0.1)
    x := float64(1080 / 2)
    dc.DrawRectangle(x, 0.0, x, 1080)
    dc.Fill()
}

func drawLogo(dc *gg.Context) {
    logo, err := gg.LoadPNG("ebay-logo.png")
    if err != nil {
        log.Fatalf("failed to load logo: %v", err)
    }

    dc.DrawImage(logo, 776, 80)
}

func createContext() *gg.Context {
	const sz = 1080
	dc := gg.NewContext(sz, sz)
    dc.SetRGB(1, 1, 1)
    dc.Clear()

    return dc
}

func drawBuyItNow(dc *gg.Context) {
    face := "MarketSans-Bold.ttf"
	if err := dc.LoadFontFace(face, 48); err != nil {
		panic(err)
	}

    dc.SetRGB(0, 0, 0)
	dc.DrawString("Buy It Now", 80, 564) 
}

func drawLink(dc *gg.Context, link string) {
    face := "MarketSans-Light.ttf"
	if err := dc.LoadFontFace(face, 48); err != nil {
		panic(err)
	}

    dc.SetRGB(0, 0, 0)
	dc.DrawString(link, 80, 628) 
}

func drawPrice(dc *gg.Context, price string) {
    face := "MarketSans-Bold.ttf"
	if err := dc.LoadFontFace(face, 72); err != nil {
		panic(err)
	}

    dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(price, 1000, 1000, 1.0, 0.0) 
}

func main() {
    dc := createContext()
    drawBackground(dc)
    drawThumb(dc)
    drawLogo(dc)

    drawBuyItNow(dc)
    drawLink(dc, "ebay.to/2HdPTV")
    drawPrice(dc, "US $480")

	dc.SavePNG("buy-now.png")
}
