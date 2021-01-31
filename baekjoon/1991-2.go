package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

type node struct {
	data  string
	left  *node
	right *node
}

func (n *node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Fprintf(w, "%s", n.data)
	n.left.PreOrder()
	n.right.PreOrder()
}

func (n *node) InOrder() {
	if n == nil {
		return
	}
	n.left.InOrder()
	fmt.Fprintf(w, "%s", n.data)
	n.right.InOrder()
}

func (n *node) PostOrder() {
	if n == nil {
		return
	}
	n.left.PostOrder()
	n.right.PostOrder()
	fmt.Fprintf(w, "%s", n.data)
}

func main() {
	defer w.Flush()
	treeMap := make(map[string]*node)
	N := getInt()
	for i := 0; i < N; i++ {
		r, ln, rn := getString()
		createNode(r, &treeMap)
		if ln != "." {
			createNode(ln, &treeMap)
			treeMap[r].left = treeMap[ln]
		}
		if rn != "." {
			createNode(rn, &treeMap)
			treeMap[r].right = treeMap[rn]
		}
	}
	rootNode := treeMap["A"]
	(*rootNode).PreOrder()
	fmt.Fprintln(w)
	(*rootNode).InOrder()
	fmt.Fprintln(w)
	(*rootNode).PostOrder()
}

func createNode(r string, treeMap *map[string]*node) {
	if _, ok := (*treeMap)[r]; !ok {
		(*treeMap)[r] = &node{
			data: r,
		}
	}
}

func getInt() int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	n, _ := strconv.Atoi(s)
	return n
}

func getString() (string, string, string) {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Split(s, " ")
	return t[0], t[1], t[2]
}
