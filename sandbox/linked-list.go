package main

import (
    "fmt"
)

/*
    Simple linked list implementation.  Much better to use a standard slice...
*/

type LinkNode struct {
    Value interface{}
    Next *LinkNode
}

type LinkedList struct {
    Root *LinkNode
    Last *LinkNode
    Size int
}

func (l *LinkedList) Add(v interface{}) *LinkNode {
    // or to create a pointer to the node use this: node := new(LinkNode)

    node := LinkNode{ Value:v }

    if l.Root == nil {
        l.Root = &node
    } else {
        l.Last.Next = &node
    }

    l.Last = &node
    l.Size++

    return &node
}

func (l *LinkedList) Remove(el *LinkNode) *LinkNode {
    //      : iterate to fine this element; preserving the previous Next
    //      : find the element's next (if it exists)
    //      : assign curr Next to prev Next
    //      : decrement size
    //      : clear the pointer
    return el.Next
}

func main() {
    list := LinkedList{}

    fmt.Println( list )
    list.Add("my root item")
    fmt.Printf("item: %v size: %d address %v\n", list.Root.Value, list.Size, &list.Root )

    n2 := list.Add("my next item")
    n3 := list.Add("my third")
    fmt.Printf("list size: %d last values %v %v %v %v\n", list.Size, n2.Value, &n2, n3.Value, &n3)

    fmt.Printf("iterate through list starting at root...\n")
    count := 0
    for el := list.Root; el != nil; el = el.Next {
        count++
        fmt.Printf("%d element = %v address %v\n", count, el.Value, &el.Next)
    }
}
