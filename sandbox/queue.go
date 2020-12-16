package main

import (
    "fmt"
)

type QueueType interface{}
type Queue []QueueType

func (q *Queue) Enqueue(v QueueType) {
    (*q) = append(*q, v)
}

func (q *Queue) Dequeue() QueueType {
    slice := *q
    var value QueueType
    if len(slice) > 0 {
        value = slice[0]
        (*q) = slice[1:] // maybe more efficient to use copy...
    }

    return value
}

func (q Queue) Peek() QueueType {
    if len(q) > 0 {
        return q[0]
    }
    return nil
}

func (q Queue) Len() int {
    return len(q)
}

func main() {
    queue := Queue{}
    queue.Enqueue(42)
    fmt.Println(queue.Len(), queue.Peek(), queue)
    v := queue.Dequeue()
    fmt.Println(queue.Len(), v, queue)
    queue.Enqueue(14)
    fmt.Println(queue.Len(), queue.Peek(), queue)
    queue.Enqueue(28)
    fmt.Println(queue.Len(), queue.Peek(), queue)
    queue.Enqueue(60)
    queue.Enqueue(78)
    fmt.Println(queue.Len(), queue.Peek(), queue)
    queue.Dequeue()
    queue.Dequeue()
    fmt.Println(queue.Len(), queue.Peek(), queue)
}

