package main

import "fmt"

// ArraySize is the size of the hash table array
const ArraySize = 7

// HashTable structure
type HashTable struct {
	array [ArraySize]*Bucket
}

// Bucket is a linked list in each slot of the hash table array
type Bucket struct {
	head *BucketNode
}

// BucketNode structure
type BucketNode struct {
	key string
	next *BucketNode
}

// insert will take in a key and create a node with the key and insert the node in the bucket
func (bucket *Bucket) insert(key string) {
	if !bucket.search(key) {
		newNode := &BucketNode{key: key}
		newNode.next = bucket.head
		bucket.head = newNode
	} else {
		fmt.Println("Key already exists")
	}
}

// search will take in a key and return true if the bucket has that key
func (bucket *Bucket) search(key string) bool {
	node := bucket.head
	for node != nil {
		if node.key == key {
			return true
		}
		node = node.next
	}
	return false
}

//delete will take a key and delete the node from the bucket
func (bucket *Bucket) delete(key string) {
	if bucket.head.key == key {
		bucket.head = bucket.head.next
		return
	}

	previousNode := bucket.head
	for previousNode.next != nil {
		if previousNode.next.key == key {
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}


// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

// Insert will take in a key and add it to the hash table array
func (hashTable *HashTable) Insert(key string) {
	index := hash(key)
	hashTable.array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (hashTable *HashTable) Search(key string) bool {
	index := hash(key)
	return hashTable.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (hashTable *HashTable) Delete(key string) {
	index := hash(key)
	hashTable.array[index].delete(key)
}

// hash function
func hash(key string) int {
	sum := 0
	for _,char := range key {
		sum += int(char)
	}
	return sum % ArraySize
}


func main() {
	hashTable := Init()

	fruits := []string{"apple", "pear", "orange", "banana", "strawberry", "cherry"}
	for _,fruit := range fruits {
		hashTable.Insert(fruit)
	}

	hashTable.Delete("cherry")

	result := hashTable.Search("cherry")
	fmt.Println(result)
}
