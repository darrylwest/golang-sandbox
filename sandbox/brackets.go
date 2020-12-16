package main

import (
	"fmt"
    "strings"
)


const (
    yes = "YES"
    no = "NO"
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

func testBracketBalance(s string) string {
    stack := Stack{}
    chars := strings.Split(s, "")

    for _,char := range chars {
        switch(char) {
        case "}":
            if stack.Pop() != "{" {
                return no
            }
        case ")":
            if stack.Pop() != "(" {
                return no
            }
        case "]":
            if stack.Pop() != "[" {
                return no
            }
        default: // an opener
            stack.Push(char)
        }
    }

    if stack.Empty() {
        return yes
    } else {
        return no
    }
}

func main() {
    s := "{[()]}"
    fmt.Println(testBracketBalance(s))
}
