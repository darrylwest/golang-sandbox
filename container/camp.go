package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

// a better impl: https://medium.com/@teddyking/namespaces-in-go-basics-e3f0fc1ff69a
// docker run <container> cmd args
// this works for linux only
// go run camp.go
func main() {
    switch os.Args[1] {
    case "run":
        run()
    default:
        panic("what?")
    }
}

func run() {
    fmt.Printf("running %v\n", os.Args[2:])
    
    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    cmd.Env = []string{"PS1=-[ns-process-] # "}

    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS,
    }

    must(cmd.Run())
}

func must(err error) {
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        panic(err)
    }
}
