package main

import (
    // "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
    // "io"
    "io/ioutil"
    "strings"
    "time"
)

var counter = createCounter(0)
var serviceName = "saas"
var port = ":3030"

func handler(w http.ResponseWriter, r *http.Request) {
    dumpRequest(r)
	fmt.Fprintf(w, "%s\n", "fobar")
}

func createCounter(start int) func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func dumpRequest(req *http.Request) {
    t := time.Now().UnixNano() / 1000

    fn := fmt.Sprintf("%s-%d-%d.req", serviceName, counter(), t)
    f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0644)
    if  err != nil {
        fmt.Printf("error opening file: %s", err)
        return
    }
    defer f.Close()

    fmt.Fprintf(f, "t0: %d\n", t)
    fmt.Fprintf(f, "%s %s %s %s\n", req.Method, req.Proto, req.Host, req.URL)
    fmt.Fprintf(f, "length: %d\n", req.ContentLength)

    for _, cookie := range req.Cookies() {
        fmt.Fprintf(f, "cookie: %s %s\n", cookie.Name, cookie.Value)
    }

    for k, v := range req.Header {
        fmt.Fprintf(f, "header: %s = %s\n", k, strings.Join(v, ""))
    }

    content, err := ioutil.ReadAll(req.Body)
    if err != nil {
        return
    }

    fmt.Fprintf(f, "content: %s", content)
}

/* replace with gjson reader
func parseRequestBody(reader *io.Reader) (map[string]interface{}, error) {
    var data map[string]interface{}

    hash, err := parseRequestBody(req.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    json, err := json.Marshal(hash)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(json))
    err := json.NewDecoder(reader).Decode(&data)

    return data, err
}
*/

func main() {
    if (len(os.Args) == 2) {
        port = os.Args[1]
    }

	fmt.Println("listen on port: ", port)

    http.HandleFunc("/", handler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bye...")
}
