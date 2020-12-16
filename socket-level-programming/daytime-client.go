package main

import (
    "bytes"
    "net"
    "io"
    "os"
    "fmt"
)

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "\nFatal error: %s\n", err.Error())
        os.Exit(1)
    }
}

func readDaytime(service string) ([]byte, error) {
    conn, err := net.Dial("tcp", service)
    checkError(err)

    defer conn.Close()

    result := bytes.NewBuffer(nil)
    var buf [512]byte

    for {
        n, err := conn.Read(buf[0:])
        result.Write(buf[0:n])
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
    }

    return result.Bytes(), nil
}

func parseArgs() string {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Use: %s host[:port]", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]

    return service
}

func main() {
    service := parseArgs()

    result, err := readDaytime(service)
    checkError(err)

    fmt.Printf("Time from %s is: %s\n", service, string(result))
}

