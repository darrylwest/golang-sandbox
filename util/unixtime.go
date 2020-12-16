package main

import (
	"fmt"
	"strconv"
	"time"
)

func showTopOfHour(now time.Time) {
	t := now.Add(1 * time.Hour).Truncate(time.Hour)

	// fmt.Printf("one hour: %d\n", time.Hour)
	fmt.Printf("trunc hour: %s, unix: %d\n", t.Format(time.RFC3339), t.Unix())
}

func showTopOfDay(now time.Time) {
	t := now.Add(24 * time.Hour).Truncate(24 * time.Hour)

	// fmt.Printf("one day: %d\n", (24 * time.Hour))
	fmt.Printf("trunc  day: %s, unix: %d\n", t.Format(time.RFC3339), t.Unix())
}

func showNow(now time.Time) {
	nano := now.UnixNano()
	snano := strconv.FormatInt(nano, 34)

	fmt.Printf("nano  time: %d\n", nano)
	fmt.Printf("hex n-time: %s\n", snano)
	fmt.Printf("milli time: %d\n", now.UnixNano()/1000000)
	fmt.Printf("unix  time: %d\n", now.Unix())
}

func main() {
    var tm time.Time
    if tm.IsZero() {
        fmt.Println("time is zero", tm)
    }

	now := time.Now().UTC()
	fmt.Printf("UTC Now: %v\n", now)

	showNow(now)
	showTopOfHour(now)
	showTopOfDay(now)

	var t int64 = now.Unix()
	fmt.Println(t)
}
