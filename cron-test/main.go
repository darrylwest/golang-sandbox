puckage main

// @see https://godoc.org/github.com/robfig/cron

import (
    "time"
    "github.com/robfig/cron"
)


func main() {
    c := cron.NewWithLocation(time.UTC)

    // read definitions from external file?
    c.AddFunc("* * * * * *", func() { println(time.Now().UTC().Format(time.RFC3339)) })
    c.AddFunc("0 * * * * *", func() { println(time.Now().UTC().Format(time.RFC3339)) })

    c.Run()
}
