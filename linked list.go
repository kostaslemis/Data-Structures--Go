package main

import (
	"fmt"
)

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func (list *LinkedList) prepend(newNode *ListNode) {
	second := list.head
	list.head = newNode
	list.head.next = second
	if list.length == 0 {
		list.tail = newNode
	}
	list.length++
}

func (list *LinkedList) append(newNode *ListNode) {
	list.tail.next = newNode
	list.tail = newNode
	list.length++
}

func (list *LinkedList) pop() (value int) {
	if list.length == 0 {
		return 0
	} else if list.length == 1 {
		value = list.head.data
		list.head = nil
		list.tail = nil
	}

	node := list.head
	for node.next != list.tail {
		node = node.next
	}

	value = list.tail.data
	list.tail = node
	return
}

func (list LinkedList) printListData() {
	node := list.head
	for list.length != 0 {
		fmt.Printf("%d ", node.data)
		node = node.next
		list.length--
	}
	fmt.Printf("\n")
}

func (list *LinkedList) deleteValue(value int) {
	if list.length == 0 {
		return
	}

	if list.head.data == value {
		list.head = list.head.next
		list.length--
		return
	}

	node := list.head
	for node.next.data != value {
		if node.next.next == nil {
			return
		}
		node = node.next
	}
	node.next = node.next.next
	list.length--
	// Check tail
}

func (list *LinkedList) getLength() int {
	return list.length
}

func (list *LinkedList) empty() bool {
	return list.length == 0
}

func main() {
	list := LinkedList{}
	node1 := &ListNode{data: 31}
	node2 := &ListNode{data: 54}
	node3 := &ListNode{data: 11}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.printListData()

	list.deleteValue(11)
	list.printListData()

	fmt.Println(list.getLength())
	fmt.Println(list.empty())

	node4 := &ListNode{data: 5}
	list.append(node4)
	list.printListData()
	fmt.Println(list.getLength())

	node5 := &ListNode{data: 100}
	list.append(node5)
	list.printListData()
	fmt.Println(list.getLength())
}

