package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	idx := strings.LastIndex(s, "/") + 1

	return s[idx:]
}

func main() {
	list := []string{"/usr/local/bin/go", "/etc/hosts", "/usr/sbin/haproxy"}

	for i := 0; i < len(list); i++ {
		fmt.Println(i+1, list[i], "basename ->", basename(list[i]))
	}
}
