package main

/**
 * available in multiple languages including a node module here: https://github.com/alizain/ulid
 */

import (
	"fmt"
	"github.com/oklog/ulid"
	"io"
	"math/rand"
	"time"
)

func createULID(entropy io.Reader) (ulid.ULID, error) {
	var ts uint64 = uint64(time.Now().UnixNano() / 1000000)
	value, err := ulid.New(ts, entropy)
	return value, err
}

var entropy io.Reader = rand.New(rand.NewSource(time.Now().UnixNano()))

func testit() {
	hash := make(map[string]int)

	for i := 0; i < 100; i++ {
		v, e := createULID(entropy)

		if e != nil {
			panic(e)
		}

		hash[v.String()]++

		fmt.Println(v)
	}

	fmt.Println(hash)
}

func t2() {
	t0 := time.Now().UnixNano() / 1000000
	fmt.Printf("%v\n", t0)
	v, _ := createULID(entropy)

	t1 := v.Time()
	fmt.Printf("%d %v %v %d\n", t0, v, t1, ulid.Now())
}

func main() {
	// t2();
	// testit()
	v, _ := createULID(entropy)
	fmt.Println(v)
	fmt.Println(v.String())
}
