package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Context struct {
	url      string
	apikey   string
	mintime  time.Duration
	maxtime  time.Duration
	lasttime time.Duration
	avgtime  time.Duration
	count    int
}

func createContextList() []Context {
	min := time.Hour
	max := time.Nanosecond

	list := []Context{
		{"https://data.bluelasso.com/status", "", min, max, 0, 0, 0},
		{"https://data.bluelasso.com/VEAccess/status", "5a66072c-0f9d-11e5-826f-d77d588b4b69", min, max, 0, 0, 0},
		{"https://data.bluelasso.com/VEDevelop/status", "5a66072c-0f9d-11e5-826f-d77d588b4b69", min, max, 0, 0, 0},
		{"https://data.bluelasso.com/VEStaging/status", "5a66072c-0f9d-11e5-826f-d77d588b4b69", min, max, 0, 0, 0},
		{"https://data.bluelasso.com/VEProd01/status", "5a66072c-0f9d-11e5-826f-d77d588b4b69", min, max, 0, 0, 0},
	}

	return list
}

func readSite(context Context) (status int, err error) {
	url := context.url
	// fmt.Printf("get %s\n", url)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	if context.apikey != "" {
		req.Header.Add("X-API-Key", context.apikey)
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return resp.StatusCode, err
	}

	defer resp.Body.Close()

	fmt.Printf("status code: %d from URL: %s\n", resp.StatusCode, resp.Request.URL)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return resp.StatusCode, err
	}

	fmt.Printf("%s\n", body)

	return resp.StatusCode, nil
}

func main() {
	list := createContextList()

	for {
		for i := 0; i < len(list); i++ {
			// point to the context struct to enable updates
			var ctx *Context = &list[i]

			t0 := time.Now()
			_, err := readSite(*ctx)
			t := time.Now().Sub(t0)

			if ctx.count == 0 {
				ctx.lasttime = t
				ctx.avgtime = t
			} else {
				ctx.avgtime = (ctx.avgtime + t + ctx.lasttime) / 3
				ctx.lasttime = t
			}

			ctx.count++

			if err != nil {
				fmt.Println(err)
			}

			if t < ctx.mintime {
				ctx.mintime = t
			}

			if t > ctx.maxtime {
				ctx.maxtime = t
			}

			fmt.Printf("%s time: %v, min: %v max: %v, avg: %v\n\n", ctx.url, t, ctx.mintime, ctx.maxtime, ctx.avgtime)
		}

		time.Sleep(5 * time.Second)
	}
}
