package main

import (
	"fmt"
	"log"
	"net/http"
    "time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// middleware provides a convenient mechanism for filtering HTTP requests
// entering the application. It returns a new handler which performs various
// operations and finishes with calling the next HTTP handler.
type middleware func(http.HandlerFunc) http.HandlerFunc

// chainMiddleware provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func chainMiddleware(mw ...middleware) middleware {
    count := 0
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
            count++
            start := time.Now()
            defer func() { log.Println(r.URL.Path, time.Since(start)) }()
            fmt.Printf("statement: %d\n", count)
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}

func withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        if r.RequestURI != "/login" {
            log.Printf("authenticate connection from %s", r.RemoteAddr)
            // do the authentication for all but login...
            // simulate a delay to validate authentication
            time.Sleep(time.Millisecond * 100)
        }

		next.ServeHTTP(w, r)
	}
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func withTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s", r.RequestURI)
        if r.RequestURI == "/hello" {
            log.Printf("terminating...")
            fmt.Fprintf(w, "terminated!")
            return
        }

		next.ServeHTTP(w, r)
	}
}

func home() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("reached home")
        fmt.Fprintf(w, "welcome")
    }
}

func login() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("reached login")
        time.Sleep(time.Millisecond * 1500)
        fmt.Fprintf(w, "login")
    }
}

func tester() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("reached hello")
        fmt.Fprintf(w, "test test test")
    }
}

func hello() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("reached hello")
        fmt.Fprintf(w, "howdy")
    }
}

func main() {
	mw := chainMiddleware(withLogging, withAuth, withTracing)
	http.Handle("/home", mw(home()))
	http.Handle("/hello", mw(hello()))
    http.Handle("/login", mw(login()))
    port := ":2200"
    log.Printf("listen on port: %s for: %s", port, "/home, /hello, /login")
	log.Fatal(http.ListenAndServe(port, nil))
}
