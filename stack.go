package main

import "fmt"

// Stack represents a stack that holds a slice
type Stack struct {
	items []int
}

// Push adds a value on top of the stack
func (stack *Stack) Push(value int) {
	stack.items = append(stack.items, value)
}

//Pop removes a value from the top of the stack and returns the removed value
func (stack *Stack) Pop() int {
	l := len(stack.items)-1
	removedValue := stack.items[l]
	stack.items = stack.items[:l]
	return removedValue
}


func main() {
	stack := Stack{}
	fmt.Println(stack)

	stack.Push(10)
	stack.Push(12)
	stack.Push(13)
	fmt.Println(stack)

	fmt.Println(stack.Pop())
	fmt.Println(stack)
}
