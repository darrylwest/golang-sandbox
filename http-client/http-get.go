package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

var host = "http://raincity.s3-website-us-east-1.amazonaws.com/config/links.json"

func search(id string, offset, limit int) ([]byte, error) {
    url := fmt.Sprintf("%s/%s?%d&%d", host, id, offset, limit)
    client := &http.Client{ 
        Timeout:time.Second * 10,
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := client.Do(req)
    defer resp.Body.Close()

    status, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return status, nil
}

func main() {
    status, err := search("58d581965c126cd8036a9206", 10, 40); if err != nil {
        panic(err)
    }

    fmt.Printf("%s", status)
}

