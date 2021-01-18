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

## 순회
### 전위 순회 : PreOrderTraverse
전위 순회는 DLR (출력 후, 왼쪽으로 이동, 오른쪽으로 이동)
### 중위 순회 : InOrderTraverse
중위 순회는 LDR (왼쪽으로 이동 후, 출력, 오른쪽으로 이동)
### 후위 순회 : PostOrderTraverse
후위 순회는 LRD (왼쪽으로 이동 후, 오른쪽으로 이동, 출력)

## 트리의 노드 삭제
매우 복잡하므로 경우를 나누어 생각하는 것이 좋다.

## binarysearchtree code
```go
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


```

## binarysearchtree main code
```go
package main

import (
	"fmt"

	"github.com/wookiist/go-ds/tree"
)

func main() {
	treeTest := &tree.Node{Data: 100}
	treeTest.Insert(5)
	treeTest.Insert(48)
	treeTest.Insert(1)
	treeTest.Insert(-1)
	treeTest.Insert(99)
	treeTest.Insert(57)
	treeTest.Insert(66)
	treeTest.Insert(88)
	treeTest.Insert(24)
	treeTest.Insert(35)
	treeTest.Insert(27)
	treeTest.Insert(28)
	treeTest.Insert(95)
	treeTest.Insert(17)
	treeTest.Insert(12)
	treeTest.Insert(72)
	fmt.Println(treeTest)
	treeTest.PreOrderTraverse()
	fmt.Println()
	treeTest.InOrderTraverse()
	fmt.Println()
	treeTest.PostOrderTraverse()
}

```