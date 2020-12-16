package main

import (
    "net"
    "os"
    "fmt"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Use: %s dotted-ip-addr\n", os.Args[0])
        fmt.Printf("examples include %s\n", "127.0.0.1")
        os.Exit(1)
    }

    name := os.Args[1]
    addr := net.ParseIP(name)
    if addr == nil {
        fmt.Fprintf(os.Stderr, "invalid address: %s\n", name)
        os.Exit(1)
    }

    mask := addr.DefaultMask()
    network := addr.Mask(mask)
    ones, bits := mask.Size()

    fmt.Println("Addrss is ", addr.String())
    fmt.Println("Default mask length is ", bits)
    fmt.Println("leading ones count is ", ones)
    fmt.Println("Mask (hex) is ", mask.String())
    fmt.Println("Network is ", network.String())
}

