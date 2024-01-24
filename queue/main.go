package main

import "fmt"

//queue is a FIFO

// define a struct to represent a queue
type Queue struct {
	items []int
}

// create the enqueue method
func (queue *Queue) enqueue(item int) {
	queue.items = append(queue.items, item)
}

// create the dequeue method
func (queue *Queue) dequeue() int {
	res := queue.items[0]
	queue.items = queue.items[1:]
	return res
}

func main() {
	queue := Queue{items: []int{1, 2, 3, 4, 5, 6, 7}}
	fmt.Printf("Before enqueue: %v", queue.items)
	queue.enqueue(88)
	queue.enqueue(99)
	fmt.Printf("\nAfter enqueue: %v", queue.items)
	fmt.Printf("\nDequeue an item: %d", queue.dequeue())
	fmt.Printf("\nAfter dequeue: %v", queue.items) //expected result is 1

}
