package main

import (
	"fmt"
	"time"
)

func main() {
	ready := make(chan bool)

	go block(ready, 1)
	go block(ready, 2)
	time.Sleep(time.Second)
	ready <- true
	time.Sleep(time.Second)
	ready <- true

	fmt.Println("ready for you now...")
	time.Sleep(time.Second)
}

func block(ok chan bool, id int) {
	fmt.Printf("%d waiting...\n", id)
	<-ok
	fmt.Printf("ok, %d I'm unblocked...\n", id)
}
