package main

import "fmt"

// BSTNode represents the components of a binary search tree
type BSTNode struct {
	key int
	left *BSTNode
	right *BSTNode
}

// Insert will add a node to the tree, the key should not already be in the tree
func (node *BSTNode) Insert(key int) {
	if node.key < key {
		//move right
		if node.right == nil {
			node.right = &BSTNode{key: key}
		} else {
			node.right.Insert(key)
		}
	} else if node.key > key {
		// move left
		if node.left == nil {
			node.left = &BSTNode{key: key}
		} else {
			node.left.Insert(key)
		}
	}
}

// Search will take in a key value and return true if there is node with that value
func (node *BSTNode) Search(key int) bool {
	if node == nil {
		return false
	}

	if node.key < key {
		//move right
		return node.right.Search(key)
	} else if node.key > key {
		return node.left.Search(key)
	}
	return true
}


func main() {
	tree := &BSTNode{key: 4}
	tree.Insert(23)
	tree.Insert(1)
	tree.Insert(43)
	fmt.Println(tree)
}
