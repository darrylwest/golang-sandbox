package main

// set the version like this: go run -ldflags '-X main.VersionString=0.0.1' version.go

import "fmt"

var VersionString = "unset"

func main() {
    if VersionString == "unset" {
        fmt.Println("USE: go run -ldflags '-X main.VersionString=0.0.1' version.go")
    } else {
        fmt.Println("Version:", VersionString)
    }
}
