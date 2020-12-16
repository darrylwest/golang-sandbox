package main

/**
 * BufferCopyProxy implments an L4 proxy with request/response type connections to a remote target.  The prirmary purpose
 * is to capture requests and responses and write to data files to enable re-play or analysis.  Since the proxy is L4, a
 * a full proxy my be required to control L7 traffic to/from an actual remote host.
 *
 * listen on the local client and proxy to/from the target. use of buffer enables the use of man-in-the-middle evaluation
 * and stream modification.  buffers are enabled on the client and target read/writes
 *
 * Note: this works for a simple request and response but fails for multiple requests with delayed responses.  this is
 *       fixed in MockDataProxy.
 *
 * TODO:
 *  Should refactor to read entire request and response prior to relaying; or, consider using pipes.
 *  Should ping the destination and, if rejected, return a mock response
 *  Figure out how to plug-in a request process hook
 *
 * @author darryl.west <darwest@ebay.com>
 * @created 2018-02-26 13:47:16
 */

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/darrylwest/cassava-logger/logger"
)

const VERSION = "2018.02.27-b"
const logo = `
______________________________________________________________________
 ___       __  __            ___                 ___                  
| _ )_  _ / _|/ _|___ _ _   / __|___ _ __ _  _  | _ \_ _ _____ ___  _ 
| _ \ || |  _|  _/ -_) '_| | (__/ _ \ '_ \ || | |  _/ '_/ _ \ \ / || |
|___/\_,_|_| |_| \___|_|    \___\___/ .__/\_, | |_| |_| \___/_\_\\_, |
                                    |_|   |__/                   |__/ 
______________________________________________________________________
`

var (
	log     *logger.Logger
	target  string
	port    int
	datadir string
	txcount int
	txid    int64
	bufsize int
)

func init() {
	handler, _ := logger.NewStreamHandler(os.Stdout)
	log = logger.NewLogger(handler)

	flag.StringVar(&target, "target", "alameda.local:8181", "the target (<host>:<port>)")
	flag.IntVar(&port, "port", 3400, "the proxy's listener port")
	flag.StringVar(&datadir, "data", "data", "the folder to store request/response data")
	flag.IntVar(&bufsize, "bufsize", 96, "the response buffer size in 1K blocks")
}

// this is the response from the target
func targetRead(dst io.Writer, src io.Reader) (written int64, err error) {
	size := bufsize * 1024
	for {
		buf := make([]byte, size)
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = fmt.Errorf("short write from client: %d != %d", nr, nw)
				break
			}
		}

		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}

		filename := fmt.Sprintf("%s/%d-response.log", datadir, txid)
		if err = writeFile(filename, buf[0:nr]); err != nil {
			log.Error("response file write error: %s", err)
		}
	}

	return written, err
}

// this is the request read from client and written to the target
func targetWrite(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
	filename := fmt.Sprintf("%s/%d-request.log", datadir, txid)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)

				if er = writeFile(filename, buf[0:nr]); er != nil {
					log.Error("request file write error: %s", er)
				}
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = fmt.Errorf("short write from client: %d != %d", nr, nw)
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}

	return written, err
}

func writeFile(filename string, buf []byte) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		return err
	}

	_, err = f.WriteString("\n")
	return err
}

func main() {
	fmt.Println(logo)
	fmt.Printf("Buffer Copy Proxy, Version %s\n", VERSION)

	log.Info("pid %d\n", os.Getpid())
	flag.Parse()

	incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Error("could not start server on %d: %v", port, err)
		panic(err)
	}

	log.Info("proxy listening on %d, proxy to %s\n", port, target)

	for {
		client, err := incoming.Accept()
		if err != nil {
			log.Error("could not accept client connection", err)
			panic(err)
		}

		log.Info("client '%v' connected\n", client.RemoteAddr())

		target, err := net.Dial("tcp", target)
		if err != nil {
			log.Error("could not connect to target, %s, please try again later...", err)
			continue
		}

		log.Info("connection to server %v established\n", target.RemoteAddr())

		go func() {
			defer client.Close()
			txcount++
			txid = time.Now().UnixNano() + int64(txcount%1000)

			size := 32 * 1024
			buf := make([]byte, size)
			written, err := targetWrite(target, client, buf)

			if err != nil {
				log.Error("client error : %s\n", err)
			}

			log.Info("total response written: %d", written)
		}()

		go func() {
			defer target.Close()

			written, err := targetRead(client, target)

			if err != nil {
				log.Error("client error : %s\n", err)
			}

			log.Info("target read %d bytes and written back to client...\n", written)
		}()
	}
}
