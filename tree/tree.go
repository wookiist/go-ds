package tree

import "fmt"

// Node type
type Node struct {
	Data  int
	left  *Node
	right *Node
}

// Insert node into the Tree
// the data to add should not be already in the tree
func (n *Node) Insert(data int) {
	if n.Data > data {
		// move left
		if n.left == nil {
			n.left = &Node{Data: data}
		} else {
			n.left.Insert(data)
		}
	} else if n.Data < data {
		// move right
		if n.right == nil {
			n.right = &Node{Data: data}
		} else {
			n.right.Insert(data)
		}
	}
}

// Search a node
func (n *Node) Search(data int) bool {
	if n == nil {
		return false
	}
	if n.Data > data {
		// move left
		return n.left.Search(data)
	} else if n.Data < data {
		// move right
		return n.right.Search(data)
	}
	return true
}

// PreOrderTraverse func
func (n *Node) PreOrderTraverse() {
	if n == nil {
		return
	}
	fmt.Println(n.Data)
	n.left.PreOrderTraverse()
	n.right.PreOrderTraverse()
}

// InOrderTraverse func
func (n *Node) InOrderTraverse() {
	if n == nil {
		return
	}
	n.left.InOrderTraverse()
	fmt.Println(n.Data)
	n.right.InOrderTraverse()
}

// PostOrderTraverse func
func (n *Node) PostOrderTraverse() {
	if n == nil {
		return
	}
	n.left.PostOrderTraverse()
	n.right.PostOrderTraverse()
	fmt.Println(n.Data)
}
