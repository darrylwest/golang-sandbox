package main

// TODO turn this into a package/struct
// TODO ability to balance
// TODO ability to remove a node

import (
	"fmt"
)

var (
    collection []int
)

type Node struct {
	value       int
	left, right *Node
}

func Create(values []int) *Node {
	var root *Node
	for _, v := range values {
		root = add(root, v)
	}

	return root
}

func traverse(t *Node, fn func(int)) {
	if t != nil {
		traverse(t.left, fn)
        fn(t.value)
		traverse(t.right, fn)
	}
}

func reverse(t *Node, fn func(int)) {
	if t != nil {
		reverse(t.right, fn)
		fn(t.value)
		reverse(t.left, fn)
	}
}

func add(t *Node, value int) *Node {
	if t == nil {
		t = new(Node)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func collect(value int) {
    if collection == nil {
        collection = make([]int, 0)
    }

    collection = append(collection, value)
}

func main() {
	list := []int{5, 23, 4, 66, 1, 9}
	root := Create(list)

	fmt.Println(list)
	add(root, 60)

	traverse(root, collect)
	fmt.Println(collection)

    collection = make([]int, 0)
	reverse(root, collect)
	fmt.Println(collection)
}
