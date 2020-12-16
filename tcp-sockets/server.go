package main

// @see https://sahilm.com/tcp-servers-that-run-like-clockwork/ for an improved implementation with read/write buffers, scanner and shutdown...

import (
    "context"
	"fmt"
	"net"
	// "os"
    "time"
)

const bufsize = 1 << 14

func main() {
	// TODO : get from flags
	port := 24049
	host := fmt.Sprintf("0.0.0.0:%d", port)
	ss, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	defer ss.Close()
	fmt.Println("listening on: ", host)
	for {
		conn, err := ss.Accept()
		if err != nil {
			fmt.Println("error on accept: ", err.Error())
            continue
		}

		go handleRequest(conn)
	}
}

func readRequest(conn net.Conn) ([]byte, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 120 * time.Second)
    defer cancel()

    complete := make(chan bool, 1)
    var (
        request []byte
        err error
        ccount int
    )

    go func() {
        buf := make([]byte, bufsize)
        ccount, err = conn.Read(buf)
        if err == nil {
            request = buf[:ccount]
        }

        complete <- true
    }()

    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    case <-complete:
        return request, err
    }
} 

func handleRequest(conn net.Conn) {
    count := 0
	defer conn.Close()

	for {
        buf, err := readRequest(conn)
		if err != nil {
			fmt.Printf("connection lost: %s\n", err)
			return
		}

		fmt.Printf("%s\n", buf)

        // return the byte count
        fmt.Fprintf(conn, "%s:%d\n", "ok", len(buf))

        count++
	}
}

