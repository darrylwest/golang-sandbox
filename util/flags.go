package main

import (
    "fmt"
    "flag"
    "os"
)

func main() {
    // set the args for tests...
    // os.Args = append(os.Args, "--version")
    args := os.Args[1:]

    vflag := false
    port := 8080
    host := "localhost"

    flag.BoolVar(&vflag, "version", vflag, "show the version and exit")
    flag.IntVar(&port, "port", port, "set the port number")
    flag.StringVar(&host, "host", host, "set the host name")

    flag.Parse()

    if vflag {
        fmt.Println("Version 1.0")
        return
    }

    fmt.Printf("version: %v\n", vflag)
    fmt.Printf("port: %v\n", port)
    fmt.Printf("host: %v\n", host)

    fmt.Printf("args %v\n", args)
}
