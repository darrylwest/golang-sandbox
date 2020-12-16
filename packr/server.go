package main

import (
    "fmt"
    "net/http"

    "github.com/gobuffalo/packr"
)

func main() {
    host := ":3000"
    box := packr.NewBox("./public-html")

    http.Handle("/", http.FileServer(box))
    fmt.Printf("start servie at host: %s\n", host)

    http.ListenAndServe(host, nil)
}
