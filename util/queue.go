package main

import (
	"fmt"
)

type QueueType interface{}
type Queue []QueueType

func (s Queue) Empty() bool {
	return len(s) == 0
}

func (s *Queue) Enqueue(v QueueType) {
	(*s) = append(*s, v)
}

func (s *Queue) Dequeue() QueueType {
	idx := len(*s) - 1
	v := (*s)[idx]

	(*s) = (*s)[0:idx]

	return v
}

func main() {
	queue := Queue{}
	fmt.Println(queue)

	queue.Enqueue(5)
	queue.Enqueue("any")
	queue.Enqueue(6)
	queue.Enqueue(7)

	fmt.Println(queue)

	v := queue.Dequeue()
	fmt.Println("dequeued", v)
	fmt.Println(queue)
}
