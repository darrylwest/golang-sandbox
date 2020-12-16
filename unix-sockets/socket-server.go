package main

import (
    "fmt"
	"net"
	"os"
    "time"
)

func getStatus() []byte {
    str := fmt.Sprintf(`{"status":"ok","ts":%d}`, time.Now().Unix())
    return []byte(str)
}

func echoServer(sock net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := sock.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println("rcvd:", string(data))

		_, err = sock.Write( getStatus() )
		if err != nil {
			panic("Write: " + err.Error())
		}
	}
}

func main() {
	sockfile := "echo.sock"

	os.Remove(sockfile)

	l, err := net.Listen("unix", sockfile)
	if err != nil {
		println("listen error", err.Error())
		println("it could be that you need to remove " + sockfile)
		return
	}

	defer os.Remove(sockfile)

	for {
		fd, err := l.Accept()
		if err != nil {
			println("accept error", err.Error())
			return
		}

		go echoServer(fd)
	}
}
