package main

import "fmt"

// Item represents the data type of values stored in Stack
//type Item int

// Stack represents a stack that holds a slice
type Stack struct {
	items []Item
}

// Push adds a value on top of the Stack
func (stack *Stack) Push(value Item) {
	stack.items = append(stack.items, value)
}

// Pop removes a value from the top of the Stack and returns the removed value
func (stack *Stack) Pop() Item {
	l := len(stack.items) - 1
	removedValue := stack.items[l]
	stack.items = stack.items[:l]
	return removedValue
}

// Peek returns the value from the top of the Stack
func (stack *Stack) Peek() Item {
	return stack.items[len(stack.items)-1]
}

// Contains takes in a value and returns true if that value is stored in Stack
func (stack *Stack) Contains(value Item) bool {
	for i := range stack.items {
		if value == stack.items[i] {
			return true
		}
	}
	return false
}

// Size returns the size of the Stack
func (stack *Stack) Size() int {
	return len(stack.items)
}

// Clear removes every value stored in Stack
func (stack *Stack) Clear() {
	stack.items = nil
}

// Empty returns true if no value is stored in Stack
func (stack *Stack) Empty() bool {
	return len(stack.items) == 0
}

func main() {
	stack := Stack{}
	fmt.Println(stack)

	stack.Push(10)
	stack.Push(12)
	stack.Push(13)
	fmt.Println(stack)
	fmt.Println(stack.Peek())

	fmt.Println(stack.Pop())
	fmt.Println(stack)
	fmt.Println(stack.Peek())

	stack.Clear()
	fmt.Println(stack)
	fmt.Println(stack.Size())
}
