package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("version: ", runtime.Version())

	m := new(runtime.MemStats)

	runtime.ReadMemStats(m)
	fmt.Printf("alloc  %d\n", m.TotalAlloc)
	fmt.Printf("system %d\n", m.Sys)

	fmt.Printf("cpus %d\n", runtime.NumCPU())
}
