package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/go-zoo/bone"
)

func main() {
    host := ":1900"
    mux := bone.New()

    mux.NotFoundFunc(NotFoundHandler)

    // static stuff first
    mux.GetFunc("/", dfltHandler)

    fs := http.FileServer(http.Dir("public-html"))
    mux.Get("/static/", http.StripPrefix("/static/", fs))

    mux.PostFunc("/user/:id", UserHandler)
    mux.PutFunc("/user/:id", UserHandler)

    fmt.Println("listening on host:", host)
    http.ListenAndServe(host, mux)
}

func dfltHandler(w http.ResponseWriter, req *http.Request) {
    file, _ := ioutil.ReadFile("public-html/index.html")
    w.Write(file)
}

func UserHandler(w http.ResponseWriter, req *http.Request) {
    val := bone.GetValue(req, "id")

    w.Write([]byte(val))
    w.Write([]byte("\n\n"))
}

func NotFoundHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("nothing here to see...\nmove along now...\n\n"))
}
