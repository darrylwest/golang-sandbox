
package main

import (
    "fmt"
    "github.com/google/tcpproxy"
)

func main() {

    p := tcpproxy.Proxy{}

    port := ":8080"

    p.AddRoute(port, tcpproxy.To("127.0.0.1:9601"))
    p.AddRoute(port, tcpproxy.To("127.0.0.1:9602"))
    p.AddRoute(port, tcpproxy.To("127.0.0.1:9603"))

    fmt.Printf("listening on port %s\n", port)

    err := p.Run()
    fmt.Println(err)
}
