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
