# Tree
Create 2021년 01월 16일
Update 2021년 01월 16일

---

## Tree
트리는 노드와 브랜치를 이용해서 사이클을 이루지 않도록 구성한 데이터 구조이다.
주로 탐색 알고리즘의 구현을 위해 사용된다.

## Binary Tree (BT) & Binary Search Tree (BST)
* BT - 이진 트리 - 노드의 최대 브랜치가 2인 트리
* BST - 이진 탐색 트리 - BT에 추가적인 조건이 존재하는 트리
    - 왼쪽 자식 노드는 부모 노드보다 작은 값이고, 오른족 자식 노드는 부모 노드보다 큰 값을 가지고 있다.

## 장점과 주요 용도
- 주요 용도는 데이터 검색(탐색)
- 장점은 탐색의 속도를 개선할 수 있다는 점.
- O(logN)  / Unbalanced Tree라면 O(n)

## binarysearchtree code
```go
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

```

## binarysearchtree main code
```go
package main

import "fmt"

func main() {
	tree := &Node{key: 100}
	tree.Insert(5)
	fmt.Println(tree)
	fmt.Println(tree.Search(5))
}
```