package main

import "fmt"

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue adds a value at the end of the queue
func (queue *Queue) Enqueue(value int) {
	queue.items = append(queue.items, value)
}

// Dequeue removes a value from the end of the queue and returns the removed value
func (queue *Queue) Dequeue() int {
	removedValue := queue.items[0]
	queue.items = queue.items[1:]
	return removedValue
}


func main() {
	queue := Queue{}
	fmt.Println(queue)

	queue.Enqueue(12)
	queue.Enqueue(5)
	queue.Enqueue(43)
	fmt.Println(queue)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
}
