package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d %d", msg, i, time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	bob := boring("bob:")
	sam := boring("sam:")

	for i := 1; i <= 5; i++ {
		fmt.Println(<-bob)
		fmt.Println(<-sam)
	}

	fmt.Println("I'm outta here...")
}
