package main

import (
	tm "github.com/buger/goterm"
	"time"
)

func main() {
	tm.Clear()

	for {
		tm.MoveCursor(1, 1)
		tm.Println("now: ", time.Now().Format(time.RFC1123))
		tm.Flush()
		time.Sleep(time.Second)
	}
}
