// server is a minimal echo server with hits counter
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	countUp()
	fmt.Printf("url %q\n", r.URL)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Printf("counter now at %d\n", count)
}

func countUp() {
	mu.Lock()
	count++
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("shutdown...")
	fmt.Fprintf(w, "Shuting down afer count %d\n", count)
	os.Exit(0)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/shutdown", shutdown)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

