package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func worker(id int, data <-chan string) {
	fmt.Printf("worker %d starting...\n", id)
	defer func() {
		fmt.Printf("worker %d complete...\n", id)
		waitGroup.Done()
	}()

	for {
		value, ok := <-data
		if !ok {
			fmt.Printf("channel closed for worker %d\n", id)
			break
		}
		time.Sleep(time.Millisecond * 10)
		fmt.Printf("worker: %d) %s\n", id, value)
	}
}

func main() {
	fmt.Println("app startup...")

	data := make(chan string)

	// this starts to break down for long runs (1e6) with more than 1,000 go routines
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go worker((i + 1), data)
	}

	t0 := time.Now().UnixNano()
	for i := 0; i < 1e6; i++ {
		data <- fmt.Sprintf("test # %d", i)
	}

	close(data)
	waitGroup.Wait()
	fmt.Printf("completed in %f ms\n", float64(time.Now().UnixNano()-t0)/1e6)
}
