package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "test.txt"

	stat, err := os.Stat(filename)

	if err != nil {
		panic(err)
	}

	file, err := os.Open(filename)
	data := make([]byte, stat.Size())
	count, err := file.Read(data)

	if err != nil {
		panic(err)
	}

	fmt.Printf("byte count: %d\n", count)
	fmt.Printf("%s\n", data[:count])
}
