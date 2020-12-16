package main

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

	stack.Push("{")
	stack.Push("(")
	stack.Push("[")
	stack.Push("<")

	fmt.Println(stack)

	v := stack.Pop()
	fmt.Printf("poped %v expected <\n", v)

	v = stack.Pop()
	fmt.Printf("poped %v expected [\n", v)

	v = stack.Pop()
	fmt.Printf("poped %v expected (\n", v)

	v = stack.Pop()
	fmt.Printf("poped %v expected {\n", v)

    fmt.Println("stack empty?", stack.Empty())
}
