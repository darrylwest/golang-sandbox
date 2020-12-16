package main

import (
	"fmt"
	"sync"
	"time"
)

type Stack struct {
	sync.RWMutex
	lifo    []string
	maxsize int
}

func main() {
	complete := make(chan bool)
	stack := Stack{}

	stack.maxsize = 40

	fmt.Printf("stack max size: %d\n", stack.maxsize)

	for i := 0; i < 100; i++ {
		// somehow wait until ready...
		go func(i int) {
			msg := fmt.Sprintf("element %d", i)
			stack.push(msg)
		}(i)
	}

	go func(max int) {

		for count := 0; count < max; count++ {
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("loop count: %d\n", count)
			if stack.empty() != true {
				v := stack.pop()
				fmt.Printf("poped %v\n", v)
			} else {
				stack.push("new thing...")
			}
		}

		fmt.Println("loop complete...")

		complete <- true
	}(50)

	// wait for complete event
	<-complete
}

func (s Stack) empty() bool {
	s.RLock()
	sz := len(s.lifo) == 0
	s.RUnlock()
	return sz
}

func (s *Stack) push(v string) {
	defer s.Unlock()
	s.Lock()

	size := len(s.lifo)
	fmt.Printf("push %v, size %d\n", v, size)

	if size >= s.maxsize {
		fmt.Printf("blocking, size %d\n", size)
		// wait for stack ready
		// <-s.ready
	}

	s.lifo = append(s.lifo, v)
}

func (s *Stack) pop() string {
	defer s.Unlock()
	s.Lock()
	top := len(s.lifo) - 1

	v := s.lifo[top]

	s.lifo = s.lifo[0:top]

	return v
}

func (s Stack) isfull() bool {
	defer s.RUnlock()
	s.RLock()
	return len(s.lifo) >= s.maxsize
}
