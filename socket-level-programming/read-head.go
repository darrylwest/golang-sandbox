package main

// e.g.: go run read-head.go google.com:80

import (
    "net"
    "os"
    "fmt"
    "io/ioutil"
    "strings"
)

func parseArgs() string {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Use: %s host[:port]", os.Args[0])
        os.Exit(1)
    }

    host := os.Args[1]
    if strings.Contains(host, ":") {
        return host
    }

    return fmt.Sprintf("%s:80", host)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "\nFatal error: %s\n", err.Error())
        os.Exit(1)
    }
}

func main() {

    service := parseArgs()

    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    result, err := ioutil.ReadAll(conn)
    checkError(err)

    fmt.Printf("Read head response from %s\n\n", service)
    fmt.Println(string(result))
}

