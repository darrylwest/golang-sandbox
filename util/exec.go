package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if err := os.Chdir(os.Getenv("HOME")); err != nil {
		log.Fatal(err)
	}

    cmd := "ls /dev | fgrep ttys | fgrep -v 00 | wc -l"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Command finished without error")
	fmt.Printf("%s\n", out)
}
