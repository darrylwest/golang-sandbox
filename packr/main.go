package main

import (
    "fmt"
    "github.com/gobuffalo/packr"
)

func main() {
    box := packr.NewBox("./public-html")

    html := box.String("index.html")

    fmt.Printf(">>\n%s\n", html)
}
