
package main

import (
    . "fmt"
    . "net/http"
    "os"
    "os/signal"
    "syscall"
    . "sync"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

var servers WaitGroup

func init() {
    go SignalHandler(make(chan os.Signal, 1))
}

func main() {
    message := "I am secure.  kind of..."
    HandleFunc("/hello", func(w ResponseWriter, r *Request) {
        w.Header().Set("Content-Type", "text/plain")
        Fprintf(w, message)
    })

    Launch("http", func() error {
        return ListenAndServe(ADDRESS, nil)
    })
        
    Launch("https", func() error {
        return ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
    })

    servers.Wait()
}

func Launch(name string, f func() error) {
    servers.Add(1)
    go func() {
        defer servers.Done()
        if err := f(); err != nil {
            Println(name, "->", err)
            syscall.Kill(syscall.Getpid(), syscall.SIGABRT)
        }
    }()
}

func shutdown() {
    Println("clean shutdown")
}

func SignalHandler(c chan os.Signal) {
    signal.Notify(c, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)
    for s := <-c; ; s = <- c {
        switch s {
        case syscall.SIGABRT:
            Println("SIGKILL received")
            os.Exit(1)
        case os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT:
            shutdown()
            os.Exit(0)
        }
    }
}

