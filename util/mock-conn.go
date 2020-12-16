package main

import (
	"fmt"
	"io"
)

type MockConn struct {
	ServerReader *io.PipeReader
	ServerWriter *io.PipeWriter

	ClientReader *io.PipeReader
	ClientWriter *io.PipeWriter
}

func (c MockConn) Close() error {
	if err := c.ServerWriter.Close(); err != nil {
		return err
	}
	if err := c.ServerReader.Close(); err != nil {
		return err
	}

	return nil
}

func (c MockConn) Read(data []byte) (n int, err error) {
	return c.ServerReader.Read(data)
}

func (c MockConn) Write(data []byte) (n int, err error) {
	return c.ServerWriter.Write(data)
}

func NewMockConn() MockConn {
	serverRead, clientWrite := io.Pipe()
	clientRead, serverWrite := io.Pipe()

	conn := MockConn{}
	conn.ServerReader = serverRead
	conn.ServerWriter = serverWrite
	conn.ClientReader = clientRead
	conn.ClientWriter = clientWrite

	return conn
}

func write(w io.Writer, data []byte, c chan int) {
	n, err := w.Write(data)
	if err != nil {
		fmt.Println("error writing to pipe", err)
	}
	if n != len(data) {
		fmt.Printf("data length %d != written %d\n", len(data), n)
	}
	c <- 0
}

func main() {
	ch := make(chan int)
	conn := NewMockConn()
	var buf = make([]byte, 128)
	data := []byte("this is a server message to the client")

	fmt.Printf("write: %s\n", data)

	// write to the client...
	go write(conn.ServerWriter, data, ch)

	n, err := conn.ClientReader.Read(buf)
	if err != nil {
		fmt.Println("error reading message", err)
	}
	if n != len(data) {
		fmt.Printf("data length %d != read %d\n", len(data), n)
	}

	fmt.Printf("read: %s\n", buf[:n])
	fmt.Printf("%s\n", buf[:n])

	<-ch
	conn.Close()

}
