package main

import (
	"crypto/tls"
	"flag"
    "fmt"
	"log"

	"github.com/kahlys/proxy"
)

var (
	// addresses
	localAddr  = flag.String("lhost", ":3030", "proxy local address")
	targetAddr = flag.String("rhost", ":4000", "proxy remote address")

	// tls configuration for proxy as a server (listen)
	localTLS  = flag.Bool("ltls", false, "tls/ssl between client and proxy, you must set 'lcert' and 'lkey'")
	localCert = flag.String("lcert", "", "certificate file for proxy server side")
	localKey  = flag.String("lkey", "", "key x509 file for proxy server side")

	// tls configuration for proxy as a client (connection to target)
	targetTLS  = flag.Bool("rtls", false, "tls/ssl between proxy and target, you must set 'rcert' and 'rkey'")
	targetCert = flag.String("rcert", "", "certificate file for proxy client side")
	targetKey  = flag.String("rkey", "", "key x509 file for proxy client side")

    bufSize = flag.Int("bufsize", 128000, "set the buffer size")
)

func main() {
    fmt.Println(getLogo())
	flag.Parse()

	p := proxy.Server{
		Addr:   *localAddr,
		Target: *targetAddr,
        BufferSize: *bufSize,
	}


	if *targetTLS {
		cert, err := tls.LoadX509KeyPair(*targetCert, *targetKey)
		if err != nil {
			log.Fatalf("configuration tls for target connection: %v", err)
		}
		p.TLSConfigTarget = &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	}

	log.Println("Proxying from " + p.Addr + " to " + p.Target)
    log.Println(fmt.Sprintf("socket buffer size: %d", p.BufferSize))

	if *localTLS {
		p.ListenAndServeTLS(*localCert, *localKey)
	} else {
		p.ListenAndServe()
	}
}

func getLogo() string {
    logo := `
         ___                  ___           
        | _ \_ _ _____ ___  _/ __|_ __ _  _ 
        |  _/ '_/ _ \ \ / || \__ \ '_ \ || |
        |_| |_| \___/_\_\\_, |___/ .__/\_, |
                         |__/    |_|   |__/ 

                  Version 19.02.12
`

    return logo
}

