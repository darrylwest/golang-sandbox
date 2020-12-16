package main

import (
    "fmt"
)

type Stacker interface {
    Push(interface{})
    Pop() interface{}
    IsEmpty() bool
    Len() int
}

type Stack struct {
    value interface{}
    next *Stack
}

func NewStack() *Stack  {
    return &Stack{}
}

var top *Stack

func (s Stack) Push(value interface{}) {
    item := Stack{ value: value, next: top }
    top = &item
}

func (s Stack) Pop() interface{} {
    if top == nil {
        return nil
    }

    item := top

    top = item.next
    return item.value
}

func (s Stack) IsEmpty() bool {
    return top == nil
}

func (s Stack) Peek() interface{} {
    if top == nil {
        return nil
    }
    return top.value
}

func (s Stack) Len() int {
    if top == nil {
        return 0
    }

    count := 1
    p := top.next

    for p != nil {
        count++
        p = p.next
    }

    return count
}

func main() {
    stack := NewStack()

    stack.Push("first value")
    fmt.Printf("stack size: %d\n", stack.Len())

    stack.Push("second value")
    fmt.Printf("stack size: %d\n", stack.Len())

    stack.Push("third value")
    fmt.Printf("stack size: %d\n", stack.Len())

    fmt.Printf("peek at top: %s\n", stack.Peek())

    for item := stack.Pop() ; item != nil; item = stack.Pop() {
        fmt.Printf("poped: %s, size: %d\n", item, stack.Len())
    }
}