package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Printf("cpu prof: %v\n", runtime.CPUProfile())
	fmt.Println("version:", runtime.Version())

	fmt.Println("pid:", os.Getpid())

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("pwd:", pwd)
}
