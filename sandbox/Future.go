package main

import (
    "io/ioutil"
    "fmt"
    "net/http"
)

// generic future wrapper to be used for any type of promise like implementation
func Future(fn func() (interface{}, error)) func() (interface{}, error) {
    var result interface{}
    var err error

    ch := make(chan struct{}, 1)
    go func() {
        defer close(ch)
        result, err = fn()
    }()

    return func() (interface{}, error) {
        <-ch
        return result, err
    }
}

func main() {
    // url := "http://raincity.s3-website-us-east-1.amazonaws.com/config/links.json"
    url := "https://data.bluelasso.com/VEStaging/status"
    future := Future(func() (interface{}, error) {
        req, qerr := http.NewRequest("GET", url, nil)
        if qerr != nil {
            return nil, qerr
        }

        req.Header.Set("x-api-key", "5a66072c-0f9d-11e5-826f-d77d588b4b69")
        client := &http.Client{}

        resp, err := client.Do(req)
        if err != nil {
            return nil, err
        }

        defer resp.Body.Close()
        return ioutil.ReadAll(resp.Body)
    })

    // do something here...

    b, err := future()
    if err != nil {
        fmt.Println(err)
        return
    }

    body, _ := b.([]byte)

    // should parse the payload and dump data...
    fmt.Printf("%s\n", body)
    fmt.Printf("body length: %d\n", len(body))
}
