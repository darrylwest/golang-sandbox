package main

import (
    "fmt"
    "container/list"
)

/*
    this implementation is not thread safe; it probably isn't the best solution
    because the list object is really meant for quick inserts within the list
    not just at the top.
 */

type Stack struct {
    values *list.List
}

func (s *Stack) Push(v interface{}) {
    s.values.PushFront(v)
}

func (s *Stack) Pop() interface{} {
    e := s.values.Front()
    return s.values.Remove(e)
}

func (s Stack) AsList() []interface{} {
    lst := make([]interface{}, s.values.Len())

    i := s.values.Len() - 1
    for e := s.values.Front(); e != nil; e = e.Next() {
        lst[i] = e.Value
        i--
    }

    return lst
}

func (s Stack) Len() int {
    return s.values.Len()
}

func main() {
    stack := Stack{values:list.New()}

    // works like a stack
    stack.Push("{")
    stack.Push("(")
    stack.Push("[")
    stack.Push("<")

    fmt.Println("pushed...")
    lst := stack.AsList()
    for _,v := range lst {
        fmt.Printf("%v", v)
    }
    fmt.Println("\nstack size", stack.Len())

    fmt.Println("pops...")
    v := stack.Pop()
    fmt.Printf("%v should be <, stack size: %d\n", v, stack.Len())
    v = stack.Pop()
    fmt.Printf("%v should be [, stack size: %d\n", v, stack.Len())
    v = stack.Pop()
    fmt.Printf("%v should be (, stack size: %d\n", v, stack.Len())
    v = stack.Pop()
    fmt.Printf("%v should be {, stack size: %d\n", v, stack.Len())

    fmt.Printf("stack size should now be zero(%d)\n", stack.Len())
}
