package main

import (
    "io"
    "net/http"
    "github.com/bmizerany/pat"
    "fmt"
    "log"
)

func HomeServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "<!doctype html><html><h2>hello server</h2></html>\n")
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "<!doctype html><html><h2>hello, " + req.URL.Query().Get(":name") + "!</h2></html>\n")
}

func main() {
    port := ":9001"

    m := pat.New()
    m.Get("/", http.HandlerFunc( HomeServer ))
    m.Get("/hello/:name", http.HandlerFunc( HelloServer ))

    http.Handle("/", m)

    fmt.Println("starting server at port: " + port)

    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

