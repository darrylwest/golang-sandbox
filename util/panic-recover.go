package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("recovered =>", recover())
	}()

	panic("first panic")
	fmt.Println("after the panic <won't print>")
}
