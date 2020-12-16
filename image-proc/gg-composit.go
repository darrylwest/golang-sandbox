package main

import (
    "fmt"

    "github.com/fogleman/gg"
    "github.com/disintegration/imaging"
)

func main() {
    fin := "ebay-logo.png"
    title := "DJI Mavic Pro Folding Drone, GPS"
    subtitle := "Brand New, Free Shipping"
    price := "$889 USD"
    link := "https://ebay.to/2rT6o0T"
    fout := "out.png"

    data, err := gg.LoadPNG(fin)
    if err  != nil {
        fmt.Printf("error reading image data: %s : %s\n", fin, err)
        return
    }

    size := data.Bounds().Size()
	width, height := 1600, 1600
	dc := gg.NewContext(width, height)

    if err = dc.LoadFontFace("MarketSans-Light.ttf", 96); err != nil {
        fmt.Printf("error loading font: %s\n", err)
        return
    }

    // draw a white background
	dc.SetRGBA(1.0, 1.0, 1.0, 1.0)

    // draw a white background...
    dc.DrawRectangle(0, 0, float64(width), float64(height))
    dc.Fill()

    dc.DrawImage(data, int(width - size.X - 40), int(height - size.Y - 80))

    // draw a black fill
	dc.SetRGBA(0.0, 0.0, 0.0, 1.0)
    dc.DrawString(title, 60.0, 200.0)
    dc.DrawString(subtitle, 60.0, 400.0)
    dc.DrawString(price, 60.0, 600.0)

    dc.DrawString(link, 60.0, 800.0)

    // instagram preferred size : 1080 X 1080
    image := imaging.Resize(dc.Image(), 1080, 0, imaging.Lanczos)

	gg.SavePNG(fout, image)

    fmt.Println("saved file to", fout)
}
