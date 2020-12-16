package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		fmt.Printf("create a go thread: %d\n", i)
		go func() {
			once.Do(onceBody)
			fmt.Printf("receive a go message: %d\n", i)
			ch <- true
		}()
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("send a go done message: %d\n", i)
		<-ch
	}
}
