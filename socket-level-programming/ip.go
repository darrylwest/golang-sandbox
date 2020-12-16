package main

import (
    "net"
    "os"
    "fmt"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Use: %s ip-addr\n", os.Args[0])
        fmt.Printf("examples include %s or %s\n", "127.0.0.1", "0:0:0:0:0:0:0:1")
        os.Exit(1)
    }
    name := os.Args[1]
    addr := net.ParseIP(name)
    if addr == nil {
        fmt.Fprintf(os.Stderr, "invalid address: %s\n", name)
        os.Exit(1)
    }

    fmt.Println(addr.String())
}

