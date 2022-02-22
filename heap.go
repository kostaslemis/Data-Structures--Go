package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (heap *MaxHeap) Insert(key int) {
	heap.array = append(heap.array, key)
	heap.maxHeapifyUp(len(heap.array)-1)
}

// Extract returns the largest key and removes it from the heap
func (heap *MaxHeap) Extract() int {
	if len(heap.array) > 0 {
		extracted := heap.array[0]

		lastIndex := len(heap.array)-1

		// take out the last index and put it in the root
		heap.array[0] = heap.array[lastIndex]
		heap.array = heap.array[:lastIndex]

		heap.maxHeapifyDown(0)
		return extracted
	}

	fmt.Println("Cannot extract because array is empty")
	return -1
}

// maxHeapifyUp will heapify from bottom to top
func (heap *MaxHeap) maxHeapifyUp(index int) {
	for heap.array[parent(index)] < heap.array[index] {
		heap.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (heap *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(heap.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex {	// when left child is the only child
			childToCompare = l
		} else if heap.array[l] > heap.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if heap.array[index] < heap.array[childToCompare] {
			heap.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Get the parent index
func  parent(i int) int {
	return (i-1)/2
}

// Get the left child index
func left(i int) int {
	return 2*i + 1
}

// Get the right child index
func right(i int) int {
	return  2*i + 2
}

// Swap keys in the array
func (heap *MaxHeap) swap(i1, i2 int) {
	heap.array[i1], heap.array[i2] = heap.array[i2], heap.array[i1]
}


func main() {
	heap := MaxHeap{}
	fmt.Println(heap)

	buildHeap := []int{10, 20 ,30, 40, 50, 7, 3, 1, 42, 17}
	for _,value := range buildHeap {
		heap.Insert(value)
		fmt.Println(heap)
	}

	for i := 0; i < 5; i++ {
		heap.Extract()
		fmt.Println(heap)
	}
}
