package main

import "fmt"

// AlphabetSize is the number of different characters in the trie
const AlphabetSize = 26

// TrieNode represents each node in the trie
type TrieNode struct {
	children [AlphabetSize]*TrieNode
	isEnd bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *TrieNode
}

// InitTrie creates a new Trie
func InitTrie() *Trie {
	result := &Trie{root: &TrieNode{}}
	return result
}

// Insert will take in a word and add it to the trie
func (trie *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := trie.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil{
			currentNode.children[charIndex] = &TrieNode{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and return true if the word is included is the trie
func (trie *Trie) Search(word string) bool {
	wordLength := len(word)
	currentNode := trie.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil{
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}


func main() {
	trie := InitTrie()
	trie.Insert("pear")
	fmt.Println(trie.Search("pear"))
}
