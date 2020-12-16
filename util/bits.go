package main

import (
	"fmt"
)

func main() {
	var x uint8
	var mask uint8

	x = 0xFF
	mask = 0x80

	fmt.Printf("raw: %0X push mask: %0X push? %t value: %d\n", x, x&mask, (x&mask == mask), x & ^mask)
}
