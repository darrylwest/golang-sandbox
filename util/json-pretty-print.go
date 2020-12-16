package main

// use: go build -o target json-pretty-print.go target.go

import (
    "encoding/json"
    "fmt"
)

func PrettyPrint(data []byte) {
    var hash map[string]interface{}

    json.Unmarshal(data, &hash)

    json, _ := json.MarshalIndent(hash, "", "  ")
    fmt.Printf("%s\n", json)
}

