
package main

// from https://github.com/thedevsaddam/govalidator

import (
    "encoding/json"
    "fmt"
    "net/http"

    validator "github.com/thedevsaddam/govalidator"
)

const stduse = `USE: curl "http://localhost:9000?web=&phone=&zip=&dob=&agree="`

func handler(w http.ResponseWriter, r *http.Request) {
    rules := validator.MapData{
        "username": []string{"required", "between:3,12"},
        "email":    []string{"required", "min:4", "max:30", "email"},
        "web":      []string{"url"},
        "phone":    []string{"digits:10"},
        "agree":    []string{"bool"},
        // "dob":      []string{"date"},
    }

    messages := validator.MapData{
        "username": []string{"required:this is required...", "between:x and y"},
        "phone":    []string{"digits:needs to smell like a phone number"},
    }

    opts := validator.Options{
        Request:         r,
        Rules:           rules,
        Messages:        messages,
        RequiredDefault: true,
    }

    v := validator.New(opts)
    e := v.Validate()
    err := map[string]interface{}{"validationError": e}
    w.Header().Set("Content-type", "application/json")
    json.NewEncoder(w).Encode(err)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("listening on port: 9000")
    fmt.Println(stduse)
    http.ListenAndServe(":9000", nil)
}

