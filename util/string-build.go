package main

import (
    "fmt"
    "strings"
)

func main() {
    var b strings.Builder
    for i := 3; i > 0; i-- {
        fmt.Fprintf(&b, "%d...", i)
    }

    b.WriteString("ignition...")
    b.WriteString("blast off!")

    fmt.Println(b.String())
}
