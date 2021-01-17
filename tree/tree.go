package tree

// Node type
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Insert node into the Tree
// the data to add should not be already in the tree
func (n *Node) Insert(data int) {
	if n.data > data {
		// move left
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.Insert(data)
		}
	} else if n.data < data {
		// move right
		if n.right == nil {
			n.right = &Node{data: data}
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
	if n.data > data {
		// move left
		return n.left.Search(data)
	} else if n.data < data {
		// move right
		return n.right.Search(data)
	}
	return true
}
