package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer fmt.Printf("defered message was created at %v...\n", time.Now())

	fmt.Printf("wait message created %v...", time.Now().Unix())

	sigchan := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigchan
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	go func() {
		for n := 0; n < 5; n++ {
			time.Sleep(time.Duration(1 * time.Second))
			fmt.Printf("wait %v...\n", time.Now().Unix())
		}

		fmt.Println("timeout...")
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting...")
}
