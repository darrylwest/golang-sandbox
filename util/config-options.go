package main

// based on functional options...

import (
    "fmt"
)

type Server struct {
    timeout int
    port int
    name string
}

func NewServer(opts ...func(*Server)) *Server {
    // set the defaults
    svr := Server{
        port:8080,
        name:"no-name",
    }

    // now the overrides
    for _, fn := range opts {
        fn(&svr)
    }
    
    return &svr
}

func main() {
    s1 := NewServer()
    fmt.Println(s1)
    
    timeout := func(svr *Server) { svr.timeout = 100000 }
    name := func(svr *Server) { svr.name = "Flarber" }
    ttl := func(svr *Server) { fmt.Println("ttl...") }
    port := func(svr *Server) { svr.port = 9090 }
    
    s2 := NewServer(timeout, port, ttl, name)
    fmt.Println(s2)
    
}

