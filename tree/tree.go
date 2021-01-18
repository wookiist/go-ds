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

// Delete from tree // TEMP
func (n *Node) Delete(d int) {
	if !n.Search(d) {
		return
	}
	if n.left != nil {
		if n.left.Data == d {
			if n.left.left == nil && n.left.right == nil { // leaf node
				n.left = nil
			} else if n.left.left == nil && n.left.right != nil {
				n.left = n.left.right
			} else if n.left.left != nil && n.left.right == nil {
				n.left = n.left.left
			} else if n.left.left != nil && n.left.right != nil {
				// pass
			}
			return
		} else if n.left.Data > d {
			n.left.left.Delete(d)
		} else if n.left.Data < d {
			n.left.right.Delete(d)
		}
	} else if n.right != nil {
		if n.right.Data == d {
			if n.right.left == nil && n.right.right == nil { // leaf node
				n.right = nil
			} else if n.right.left == nil && n.right.right != nil {
				n.right = n.right.right
			} else if n.right.left != nil && n.right.right == nil {
				n.right = n.right.left
			} else if n.right.left != nil && n.right.right != nil {
				// pass
			}
		} else if n.right.Data > d {
			n.right.left.Delete(d)
		} else if n.right.Data < d {
			n.right.right.Delete(d)
		}
	}

	if n.Data == d {
		// Leaf Node
		if n.left == nil && n.right == nil {
			// parent의 left나 right를 삭제.
		}
	}
	if n.Data > d {
		// move left
	} else if n.Data < d {
		// move right
	}
}
