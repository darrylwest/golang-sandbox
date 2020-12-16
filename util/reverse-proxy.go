package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"
    "fmt"
)

// director modifies the request
func director(r *http.Request) {
    fmt.Printf("request %v\n", r)

    r.Host = "localhost:3001"
    r.URL.Host = r.Host
    r.URL.Scheme = "http"
}

// responder modifies the response
func responder(w *http.Response) error {
    fmt.Printf("response %v\n", w)

    return nil
}

func main() {
    http.HandleFunc("/vp1/status", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "{\"status\":\"ok\"}")
    })

    targetHost := "http://localhost:3001"
    url, _ := url.Parse(targetHost)
    proxy := httputil.NewSingleHostReverseProxy(url)
    proxy.Director = director
    proxy.ModifyResponse = responder

    http.Handle("/", proxy)

    port := ":3100"
    fmt.Printf("listening on %s\n", port)
    http.ListenAndServe(port, nil)
}

