package main

// a multi-threaded time server

import (
    "net"
    "os"
    "fmt"
    "time"
)

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "\nFatal error: %s\n", err.Error())
        os.Exit(1)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()

    now := time.Now().UTC()
    json, err := now.MarshalJSON()
    if err != nil {
        fmt.Fprintf(os.Stderr, "\nTime error: %s\n", err.Error())
        return
    }

    fmt.Println("write the time: ", string(json))

    conn.Write(json)
}

func main() {

    service := ":1201"

    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    fmt.Printf("listening for time on %v\n", tcpAddr)
    fmt.Printf("use 'telnet localhost %v'\n", tcpAddr)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

        go handleClient(conn)
    }
}

