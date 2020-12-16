
package main

import (
    "bufio"
    "fmt"
    "image"
    // "image/color"
    _ "image/jpeg"
    "image/png"
    "os"

    "github.com/disintegration/gift"
)

func main() {
    // fin := "./mixer.jpg"
    fin := "./ebay-logo.png"
    ifile, err := os.Open(fin)
    if err != nil {
        fmt.Printf("error reading file: %s : %s\n", fin, err)
        os.Exit(1)
    }
    defer ifile.Close()

    src, _, err := image.Decode(ifile)

    g := gift.New(
        gift.Resize(250, 0, gift.LanczosResampling),
        gift.UnsharpMask(1, 1, 0),
    )

    dst := image.NewRGBA(g.Bounds(src.Bounds()))

    g.Draw(dst, src)

	// Save image to disk.
    fout := "ebay-logo-thumb.png"
	outFile, err := os.Create(fout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()

	b := bufio.NewWriter(outFile)
	err = png.Encode(b, dst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("output file: %s written.\n", fout)
}

