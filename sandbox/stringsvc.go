package main

import (
    "fmt"
    "strings"
    "errors"
)

var (
    ErrEmpty = errors.New("hey, wtf?  I'm empty here!")
)

// define the interface to enable mocking...
type StringSvcType interface {
    Uppercase(string) (string, error)
    Count(string) int
}

// define the concrete struct
type StringSvc struct{}

// implement the interface
func (StringSvc) Uppercase(s string) (string, error) {
    if s == "" {
        return "", ErrEmpty
    }
    return strings.ToUpper(s), nil
}

func (StringSvc) Count(s string) int {
    return len(s)
}

type MockStringSvc struct{}

func (MockStringSvc) Uppercase(s string) (string, error) {
    return "fooled again", nil
}

func (MockStringSvc) Count(s string) int {
    return 127
}

func main() {
    svc := StringSvc{}

    uc, _ := svc.Uppercase("this is a test")
    fmt.Println(uc, "size", svc.Count(uc))

    _, err := svc.Uppercase("")
    fmt.Println(err)

    mock := MockStringSvc{}
    muc, _ := mock.Uppercase("this is a test")
    fmt.Println(muc, "size", mock.Count(muc))
}
