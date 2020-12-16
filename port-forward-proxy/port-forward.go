//
// PortForward : https://zupzup.org/go-prot-forwarding
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-19 15:47:39
//

// broken, should not consider...

package main

import (
    "fmt"
    "io"
    "net"
    "os"
    "os/signal"
)

var (
    target string
    port int
)

func handleClient(client net.Conn) {
    defer func() { 
        client.Close()
        fmt.Printf("client socket closed...\n")
    }()

    target, err := net.Dial("tcp", target)
    if err != nil {
        panic(err)
    }
    defer target.Close()
    fmt.Printf("connect to target server %v\n", target.RemoteAddr())

    go func() { io.Copy(target, client) }()
    go func() { io.Copy(client, target) }()
}

func ListenAndForward() {
    incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        panic(err)
    }

    fmt.Printf("server running on %d\n", port)

    // this needs to be in a forever loop
    for {
        client, err := incoming.Accept()
        if err != nil {
            panic(err)
        }

        fmt.Printf("client '%v' connected\n", client.RemoteAddr())
        handleClient(client)
    }
}

func main() {
    // forward code...
    // target = "192.168.33.11:8181"
    target = "alameda.local:8181"
    port = 1338

    pid := os.Getpid()
    fmt.Printf("pid %d\n", pid)

    signals := make(chan os.Signal, 1)
    stop := make(chan bool)
    signal.Notify(signals, os.Interrupt)

    go func() {
        for sig := range signals {
            fmt.Printf("\nReceived %v...\n", sig)
            stop <- true
        }
    }()

    fmt.Printf("forward %d to %s\n", port, target)
    ListenAndForward()
    fmt.Printf("navigate to http://127.0.0.1:%d\n", port)
    fmt.Printf("stop with kill -1 %d\n", pid)

    // wait for the stopper
    <-stop

    fmt.Println("stopped...")
}
