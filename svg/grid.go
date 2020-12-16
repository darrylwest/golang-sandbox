
// see https://www.w3.org/TR/SVG11/styling.html for styling tips

// 96px = 1.0 inch
// 8.0 x 10.0 margins...

package main

import (
    "flag"
    "fmt"
    "github.com/ajstarks/svgo"
    "os"
    "strings"
)

type Parameters struct {
    FontSize int
    ShowGrid bool
    Address  string
}

var params Parameters

func init() {
    flag.IntVar(&params.FontSize, "fs", 16, "font size (px)")
    flag.BoolVar(&params.ShowGrid, "grid", false, "show or hide the grid")
    flag.StringVar(&params.Address, "address", "First Last|Street Address|City State|Zip Code", "pipe delimited address lines (4 max)")

    flag.Parse()
}

func main() {

    width := 8 * 96
    height := 10 * 96
    canvas := svg.New(os.Stdout)
    canvas.Start(width, height)

    canvas.Rect(0, 0, width, height, canvas.RGB(255, 255, 255))

    w, h := width / 3, 96

    // lines := []string{ "Darryl West", "24 The Plaza Street", "Berkeley, CA 94705" }
    lines := strings.Split(params.Address, "|")
    yinc := 20
    tformat := fmt.Sprintf("text-anchor:left;font-family:Arial,Helvetica;font-size:%dpx;fill:black", params.FontSize)

    for row := 0; row < 10; row++ {
        y := row * 96
        for col := 0; col < 3; col++ {
            x := col * w
            if params.ShowGrid {
                canvas.Rect(x, y, w - 2, h - 2, "stroke:black;fill:none;opacity:0.1") 
            }

            x += 10
            for i, v := range lines {
                canvas.Text(x, y + yinc * (i + 1), v, tformat)
            }
        }
    }
    
    canvas.End()
}

