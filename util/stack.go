package main

// TODO: refactor to use channels for concurrent use

import (
	"fmt"
)

type StackType interface{}
type Stack []StackType

func (s Stack) Empty() bool {
	return len(s) == 0
}

func (s *Stack) Push(v StackType) {
	(*s) = append(*s, v)
}

func (s *Stack) Pop() StackType {
	idx := len(*s) - 1
	v := (*s)[idx]

	(*s) = (*s)[0:idx]

	return v
}

func main() {
	stack := Stack{}
	fmt.Println(stack)

	stack.Push(5)
	stack.Push("any")
	stack.Push(6)
	stack.Push(7)

	fmt.Println(stack)

	v := stack.Pop()
	fmt.Println("poped", v)
	fmt.Println(stack)
}
