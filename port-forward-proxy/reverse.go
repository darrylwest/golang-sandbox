package main

// works with local hosts

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    // "net/http/httptest"
    "net/http/httputil"
    "net/url"
)

func main() {

    // rpURL, err := url.Parse("http://192.168.33.11")
    rpURL, err := url.Parse("http://127.0.0.1:5980")
    if err != nil {
        panic(err)
    }

    fmt.Println(rpURL)

    proxy := httputil.NewSingleHostReverseProxy(rpURL)

    resp, err := http.Get(frontendProxy.URL)
    if err != nil {
        log.Fatal(err)
    }

    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    fmt.Printf("%s", b)
}
