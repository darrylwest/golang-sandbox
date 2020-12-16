package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>a very simple http server</h1><h6>powered by go!</h6>")
}

func main() {
    fmt.Println("pid", os.Getpid())
	host := ":8181"

	fmt.Println("host: ", host)

	var h Hello
	err := http.ListenAndServe(host, h)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ready for listeners...")
	}

	fmt.Println("...")
}
