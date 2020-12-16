package main

import (
    . "fmt"
    . "net/http"
)

const SECURE_ADDRESS = ":1025"

func main() {
    message := "i am secure..."
    HandleFunc("/hello", func(w ResponseWriter, r *Request) {
        w.Header().Set("Content-Type", "text/plain")
        Fprintf(w, message)
    })

    ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
    // ListenAndServeTLS(SECURE_ADDRESS, "localhost.crt", "localhost.key", nil)
}

