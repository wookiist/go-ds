package main

import "fmt"

// Node type
type Node struct {
	key   int
	left  *Node
	right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if k < n.key {
		// move left
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	} else if k > n.key {
		// move right
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.Insert(k)
		}
	}
}

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.key < k {
		// move right
		return n.right.Search(k)
	} else if n.key > k {
		// move left
		return n.left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{key: 100}
	tree.Insert(5)
	fmt.Println(tree)
	fmt.Println(tree.Search(5))
}
