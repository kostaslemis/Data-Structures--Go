package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

type linkedList struct {
	head *ListNode
	length int
}

func (list *linkedList) prepend(newNode *ListNode) {
	second := list.head
	list.head = newNode
	list.head.next = second
	list.length++
}

func (list linkedList) printListData() {
	node := list.head
	for  list.length !=0 {
		fmt.Printf("%d ", node.data)
		node = node.next
		list.length--
	}
	fmt.Printf("\n")
}

func (list *linkedList) deleteValue(value int) {
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
}


func main() {
	list := linkedList{}
	node1 := &ListNode{data: 31}
	node2 := &ListNode{data: 54}
	node3 := &ListNode{data: 11}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.printListData()

	list.deleteValue(11)
	list.printListData()
}
