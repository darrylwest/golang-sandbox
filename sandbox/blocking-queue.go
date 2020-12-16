package main

import (
    "fmt"
)

// TODO fix this...

type QueueType interface{}
type BlockingQueue []QueueType

func (q *BlockingQueue) Enqueue(v QueueType) {
    (*q) = append(*q, v)
}

func (q *BlockingQueue) Dequeue(out chan<- QueueType) {
    for {
        slice := *q
        value := slice[0]
        (*q) = slice[1:]

        out <- value
    }
}

func (q *BlockingQueue) Peek() QueueType {
    slice := *q
    var value QueueType

    if len(slice) > 0 {
        value = slice[0]
    }

    return value
}

func (q *BlockingQueue) Len() int {
    return len(*q)
}

func main() {
    queue := BlockingQueue{}

    queue.Enqueue(42)
    queue.Enqueue(20)
    queue.Enqueue(24)

    fmt.Println(queue.Len(), queue.Peek(), queue)

    ch := make(chan QueueType)
    go queue.Dequeue(ch)

    v, _ := <- ch
    fmt.Println(queue.Len(), v, queue)

    v, _ = <- ch
    fmt.Println(queue.Len(), v, queue)

    /*
    queue.Enqueue(14)
    fmt.Println(queue.Len(), queue.Peek(), queue)
    queue.Enqueue(28)
    fmt.Println(queue.Len(), queue.Peek(), queue)
    queue.Enqueue(60)
    queue.Enqueue(78)
    fmt.Println(queue.Len(), queue.Peek(), queue)
    queue.Dequeue()
    fmt.Println(queue.Len(), queue.Peek(), queue)
    */
}

