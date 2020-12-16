package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
    "strings"
)

func main() {
    params := []string{ "run", "-d", "-p", "5001:9623", "--name", "gohttp-5001", "gohttp-scratch"}
	if err := os.Chdir(os.Getenv("HOME")); err != nil {
		log.Fatal(err)
	}

    cmd := exec.Command("docker")
    for _, arg := range params {
        cmd.Args = append(cmd.Args, arg)
    }

    fmt.Printf("%v\n", cmd)

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Command finished without error")
    id := strings.Split(string(out), "\n")[0]

	fmt.Printf("%s\n", id)
}
