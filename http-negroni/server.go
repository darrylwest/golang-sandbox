package main

import (
    "github.com/codegangsta/negroni"
    "github.com/phyber/negroni-gzip/gzip"
    "gopkg.in/tylerb/graceful.v1"
    "github.com/darrylwest/cassava-logger/logger"
    "net/http"
    "log"
    "fmt"
    "os"
    "os/exec"
    "flag"
    "encoding/json"
    "strings"
    "time"
)

type Context struct {
    port int
    static string
    lg *logger.Logger
}

func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "shutdown requested...")

    // check for post and token
    if (r.Method == "POST") {
        log.Printf("shutdown in a graceful way...")

        cmd := exec.Command("kill", "-2", fmt.Sprintf( "%d", os.Getpid() ))
        cmd.Run()
    } else {
        log.Printf("shudown denied, method %s", r.Method)
        fmt.Fprintf(w, "shutdown request denied...")
    }
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
    var m = map[string]interface{}{
        "status":"ok",
        "ts":time.Now().UnixNano() / 1000000,
        "version":"1.0",
        "webStatus":map[string]interface{}{
            "version":"2015-10-21",
            "pid":os.Getpid(),
            "host":r.Host,
            "path":r.URL.Path,
            "agent":r.UserAgent(),
            "ips":ParseRequestIP( r ),
        },
    }

    json, err := json.Marshal( m )

    if err != nil {
        fmt.Fprintf(w, "json error")
    } else {
        headers := w.Header()
        headers.Set("Content-Type", "application/json")
        log.Printf( "headers: %v", headers)

        w.Write( json )
    }
}

func ParseRequestIP(req *http.Request) string {
    ips := []string{ req.RemoteAddr }

    forwarded := req.Header.Get("X-Forwarded-For")
    if forwarded != ""  {
        ips = append(ips, forwarded)
    }

    return strings.Join( ips, ", " )
}

func startServer(context Context) {
    mux := http.NewServeMux()
    mux.HandleFunc("/status", StatusHandler)

    server := negroni.New()
    server.Use( negroni.NewRecovery() )
    server.Use( logger.NewMiddlewareLogger( context.lg ) )

    server.Use( gzip.Gzip( gzip.DefaultCompression ))
    server.Use( negroni.NewStatic( http.Dir("staging") ))

    server.UseHandler( mux )

    log.Printf("starting server at port: %d", context.port)
    graceful.Run( fmt.Sprintf( ":%v", context.port ), 0, server )
}

func startShutdownServer(context Context) {
    mux := http.NewServeMux()
    mux.HandleFunc("/", StatusHandler)
    mux.HandleFunc("/shutdown", ShutdownHandler)
    
    server := negroni.New()
    server.Use( negroni.NewRecovery() )
    server.Use( negroni.NewLogger() )

    server.UseHandler( mux )

    log.Printf("running, shutown at port: %d", context.port)
    graceful.Run( fmt.Sprintf( ":%v", context.port ), 0, server )
}

func main() {
    handler,_ := logger.NewRotatingDayHandler( "negroni" )
    log := logger.NewLogger( handler )

    var port = 5001
    var servers = 2
    var static = "staging"

    // baseport := flag.Int("baseport", 5001, "set the server's base port number (e.g., 5001)...")
    // serverCount := flag.Int("servers", 2, "set the number of servers")
    // shutdownPort := flag.Int("shutdownPort", 5009, "set the service shutdown port")
    // daemon := flag.Bool("daemon", false, "run in the background...")
    // quiet := flag.Bool("quiet", false, "error level logging...")

    flag.Parse()

    for idx := 0; idx < servers; idx++ {
        p := port + idx
        log.Info("start service on port %d", p)

        ctx := Context{ p, static, log }
        go startServer( ctx )
    }

    log.Info("start shutdown service on port 5009")
    startShutdownServer( Context{ 5009, "", log } )
}

