package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>rain city server</h1><h6>a simple static http server powered by go!</h6>")
}

func main() {
	var host = ":5000"
	if len(os.Args) == 2 {
		host = os.Args[1]
	}

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	http.HandleFunc("/healthcheck", healthHandler)

	fmt.Println("listen on port: ", host)

	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bye...")
}
