package main

import (
	"context"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer func() {
        cancel()
        println(time.Now().Unix())
    }()

	println("wait for 5 seconds?")

    println(time.Now().Unix())
	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
		println("the timeout is set to 1 second, so the 5 second sleep fails (as it should)...")
	} else {
        println("no timeout...")
    }

}
