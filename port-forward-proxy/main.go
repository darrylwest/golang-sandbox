package main

// hangs on multiple calls; buggy because it doesn't close the connections when done (see simple.go)

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	target string
	port   int
)

func init() {
	flag.StringVar(&target, "target", "alameda.local:8181", "the target (<host>:<port>)")
	flag.IntVar(&port, "port", 7757, "the tunnelthing port")
}

func main() {
    fmt.Printf("pid %d\n", os.Getpid())
	flag.Parse()

	signals := make(chan os.Signal, 1)
	stop := make(chan bool)
	signal.Notify(signals, os.Interrupt)
	go func() {
		for _ = range signals {
			fmt.Println("\nReceived an interrupt, stopping...")
			stop <- true
		}
	}()

	incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("could not start server on %d: %v", port, err)
	}
	fmt.Printf("server running on %d\n", port)

	client, err := incoming.Accept()
	if err != nil {
		log.Fatal("could not accept client connection", err)
	}
	defer client.Close()
	fmt.Printf("client '%v' connected!\n", client.RemoteAddr())

	target, err := net.Dial("tcp", target)
	if err != nil {
		log.Fatal("could not connect to target", err)
	}
	defer target.Close()
	fmt.Printf("connection to server %v established!\n", target.RemoteAddr())

	go func() { io.Copy(target, client) }()
	go func() { io.Copy(client, target) }()

	<-stop
}
