package main

import (
	"io"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("rcvd:", string(buf[0:n]))
	}
}

func main() {
	c, err := net.Dial("unix", "echo.sock")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	go reader(c)
	for {
		s := `{"get":"status"}`

		_, err := c.Write([]byte(s))
		if err != nil {
			println(err.Error())
			break
		}

		time.Sleep(10e9)
	}
}
