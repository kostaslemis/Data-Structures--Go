package main

import "fmt"

// Item represents the data type of values stored in Queue
type Item int

type Queue struct {
	items []Item
}

// Enqueue adds value at end of Queue
func (queue *Queue) Enqueue(value Item) {
	queue.items = append(queue.items, value)
}

// Dequeue removes value from start of Queue and returns value
func (queue *Queue) Dequeue() Item {
	removedValue := queue.items[0]
	queue.items = queue.items[1:]
	return removedValue
}

// Peek returns value from start of Queue
func (queue *Queue) Peek() Item {
	return queue.items[0]
}

// Contains takes in value and returns true if value is stored in Queue
func (queue *Queue) Contains(value Item) bool {
	for i := range queue.items {
		if value == queue.items[i] {
			return true
		}
	}
	return false
}

// Size returns size of Queue
func (queue *Queue) Size() int {
	return len(queue.items)
}

// Clear removes every value stored in Queue
func (queue *Queue) Clear() {
	queue.items = nil
}

// Empty returns true if no value is stored in Queue
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
