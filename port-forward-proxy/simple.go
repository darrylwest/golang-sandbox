package main

/**
 * listen on the tunnel port and proxy to/from the target; currently it only copies in/out
 * but ideally it would read in-evaluate-then write out.
 *
 * seems to work well; multiple calls to simple endpoints; the strategy of open/close for each reqeust/response works well
 */


import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	target string
	port   int
)

func init() {
	flag.StringVar(&target, "target", "127.0.0.1:8080", "the target (<host>:<port>)")
	// flag.StringVar(&target, "target", "ebay.local:8000", "the target (<host>:<port>)")

	flag.IntVar(&port, "port", 3030, "the tunnelthing port")
}

func main() {
	fmt.Printf("pid %d\n", os.Getpid())
	flag.Parse()

	incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("could not start server on %d: %v", port, err)
	}
	fmt.Printf("server running on %d\n", port)

    for {
        client, err := incoming.Accept()
        if err != nil {
            log.Fatal("could not accept client connection", err)
        }

        fmt.Printf("client '%v' connected!\n", client.RemoteAddr())

        target, err := net.Dial("tcp", target)
        if err != nil {
            log.Fatal("could not connect to target", err)
        }

        fmt.Printf("connection to server %v established!\n", target.RemoteAddr())

        go func() { 
            defer client.Close()
            io.Copy(target, client) 
            fmt.Println("read complete...")
        }()

        go func() {
            defer target.Close()
            io.Copy(client, target)
            fmt.Println("copy complete...")
        }()
    }
}
