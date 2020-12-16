//
//
//
package main


import (
    crand "crypto/rand"
    "fmt"
    "strconv"
    "time"
)

const radix = 36

// RandomBytes generates a byte buffer of the specified size and populates it with crypo-strength random bytes
func RandomBytes(size int) ([]byte, error) {
	buf := make([]byte, size)
	_, err := crand.Read(buf)

	return buf, err
}

// CreateTXID generates a 16 character time-stamp / base 36 id
func CreateTXID() string {
	id := strconv.FormatInt(time.Now().UnixNano(), radix)
	buf, _ := RandomBytes(2)
	str := fmt.Sprintf("%s%x", id, buf)

	return str
}

func main() {
    fmt.Println(CreateTXID())
}

