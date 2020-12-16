package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "./test.out"
	text := "this is my line of text...\n"

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}

}
