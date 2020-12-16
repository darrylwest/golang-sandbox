package main

import (
    "encoding/xml"
    "io/ioutil"
    "fmt"
    // "strings"
    // "strconv"
)

func main() {
    // we just want to pick off the attributes of this node...
    // <testsuite name="TestSuite" time="331.85" tests="7" errors="0" skipped="0" failures="2">
    bytes, err := ioutil.ReadFile("test-data.xml")
    if err != nil {
        panic(err)
    }

    type Results struct {
        Time string `xml:"time,attr"`
        Tests int `xml:"tests,attr"`
        Errors int `xml:"errors,attr"`
        Skipped int `xml:"skipped,attr"`
        Failures int `xml:"failures,attr"`
    }

    type TestResults struct {
        TestSuite xml.Name `xml:"testsuite"`
        Results
    }

    var tr TestResults

    err = xml.Unmarshal(bytes, &tr)

    if err != nil {
        panic(err)
    }

    fmt.Println(tr)
    fmt.Printf("tests: %d, failures: %d, time: %s\n", tr.Tests, tr.Failures, tr.Time)
}
