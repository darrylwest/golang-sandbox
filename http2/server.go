package main

import (
    "log"
    "net/http"
    "os"

    "golang.org/x/net/http2"
)

func main() {
    cwd, err := os.Getwd()
    if err != nil {
        log.Fatal( err )
    }

    srv := &http.Server {
        Addr: ":8000",
        Handler: http.FileServer(http.Dir( cwd )),
    }

    http2.ConfigureServer( srv, &http2.Server{} )
    log.Fatal(srv.ListenAndServeTLS("raincity.com.crt", "raincity.com.key"))
}
