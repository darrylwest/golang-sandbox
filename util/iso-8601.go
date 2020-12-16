package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now().UTC()

    fmt.Println(now.Format(time.RFC3339Nano))

    fmt.Println(now.Format(time.RFC3339))
}
