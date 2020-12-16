package main

// TODO turn this into a package/struct
// TODO ability to balance
// TODO ability to remove a node

import (
	"fmt"
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

func Traverse(t *Node) []int {
	return traverse([]int{}, t)
}

func traverse(values []int, t *Node) []int {
	if t != nil {
		values = traverse(values, t.left)
		values = append(values, t.value)
		values = traverse(values, t.right)
	}

	return values
}

func Reverse(t *Node) []int {
	return reverse([]int{}, t)
}

func reverse(values []int, t *Node) []int {
	if t != nil {
		values = reverse(values, t.right)
		values = append(values, t.value)
		values = reverse(values, t.left)
	}

	return values
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

func main() {
	list := []int{5, 23, 4, 66, 1, 9}
	root := Create(list)

	fmt.Println(list)
	add(root, 60)

	sorted := Traverse(root)
	fmt.Println(sorted)

	reversed := Reverse(root)
	fmt.Println(reversed)
}
