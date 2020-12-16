// fetchall fetchs URLs in parallel
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var urls = []string{}

// auto-magically called...
func init() {
	list := []string{"http://raincitysoftware.com", "http://gopl.io"}

	urls = append(list, "http://yahoo.com")
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)

	fmt.Printf("fetch data from a list of %d urls...\n", len(urls))

	for _, url := range urls {
		fmt.Printf("fetch %s in the background...\n", url)
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch) // receive from ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

