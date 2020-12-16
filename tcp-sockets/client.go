package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// id := time.Now().UnixNano()

	port := 24049
	host := fmt.Sprintf(":%d", port)
	fmt.Println("dailing: ", host)

	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("error connecting to host ", host)
		os.Exit(1)
	}

	defer conn.Close()
	count := 1

    buf := make([]byte, 2048)

	for {
		count++

		text := fmt.Sprintf("put %d\n", time.Now().Unix())
		_, err := fmt.Fprintf(conn, text)
		if err != nil {
			fmt.Println("lost connection...")
			return
		}

		fmt.Printf("sent: %s", text)

        n, err := conn.Read(buf)
        if err != nil {
			fmt.Println("lost connection...")
			return
        }

        fmt.Printf("recd: %s", buf[:n]);

		time.Sleep(time.Millisecond * 10)

        if count > 2000 {
            fmt.Println("end the conversation and exit...")
            break
        }
	}
}
