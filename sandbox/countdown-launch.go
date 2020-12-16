// demonstrates thread/message multiplexing with select
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	n := 10
	fmt.Println("Commencing countdown from", n)

	tick := time.NewTicker(1 * time.Second)
	abort := make(chan struct{})

	go func() {
		fmt.Println("Press return to abort the countdown...")
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("key received...")
		abort <- struct{}{}
	}()

	for countdown := n; countdown > 0; countdown-- {
		select {
		case <-tick.C:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("Launch aborted!")

			return
		}
	}
	tick.Stop()
	launch()
}

func launch() {
	fmt.Println("BLAST OFF!")
}
