
package main

import (
	"fmt"
	"strings"
)

func main() {
    str := "aspects,attributes,title"

    ss := strings.Split(str, ",")

    fmt.Println(str, ss)
}
