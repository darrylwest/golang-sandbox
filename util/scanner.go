package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader(`one
two
three
four`))

	var (
		text []byte
		n    int
	)

	for scanner.Scan() {
		n++
		text = append(text, fmt.Sprintf("%d. %s\n", n, scanner.Text())...)
	}

	// fmt.Println(text)
	os.Stdout.Write(text)
}
