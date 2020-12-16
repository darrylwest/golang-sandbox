// print a decimal number in any base/radix to emulate toString(n)
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2015
	fmt.Println(x, strconv.FormatInt(int64(x), 2))

	x = 5964570133443222
	fmt.Println(x, strconv.FormatInt(int64(x), 36))
}
