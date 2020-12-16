
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var blob map[string]interface{}

    err := decoder.Decode(&blob)
    if err != nil {
        panic(err)
    }

    for k, v := range blob {
        fmt.Printf("%s = %s\n", k, v)
    }

    fmt.Fprintf(w, "%s\n", blob)
}

func main() {
    host := ":2020"
    http.HandleFunc("/", postHandler)
    fmt.Printf("listening for post on port: %s\n", host)
    fmt.Printf("test with: %s\n", "curl -XPOST -d @image-request.json http://localhost:2020/")
    http.ListenAndServe(host, nil)
}

