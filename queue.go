package main

import "fmt"

// Item represents the data type of values stored in Queue
type Item int

// Queue represents a queue that holds a slice
type Queue struct {
	items []Item
}

// Enqueue adds a value at the end of the queue
func (queue *Queue) Enqueue(value Item) {
	queue.items = append(queue.items, value)
}

// Dequeue removes a value from the end of the queue and returns the removed value
func (queue *Queue) Dequeue() Item {
	removedValue := queue.items[0]
	queue.items = queue.items[1:]
	return removedValue
}

func (queue *Queue) Peek() Item {
	return queue.items[0]
}

func (queue *Queue) Contains(value Item) bool {
	for i := range queue.items {
		if value == queue.items[i] {
			return true
		}
	}
	return false
}

func (queue *Queue) Size() int {
	return len(queue.items)
}

func (queue *Queue) Clear() {
	queue.items = nil
}

func (queue *Queue) Empty() bool {
	return len(queue.items) == 0
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

	queue.Clear()
	fmt.Println(queue)
}
