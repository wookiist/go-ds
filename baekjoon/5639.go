package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

type node struct {
	data  int
	left  *node
	right *node
}

func (n *node) PostOrder() {
	if n == nil {
		return
	}
	n.left.PostOrder()
	n.right.PostOrder()
	fmt.Fprintln(w, n.data)
}

func (n *node) Insert(d int) {
	if n == nil {
		return
	}
	if n.data > d {
		if n.left == nil {
			n.left = &node{
				data: d,
			}
		} else {
			n.left.Insert(d)
		}
	} else if n.data < d {
		if n.right == nil {
			n.right = &node{
				data: d,
			}
		} else {
			n.right.Insert(d)
		}
	}
}

func main() {
	defer w.Flush()
	nodes := make([]int, 10000)
	i := 0
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		s = strings.TrimSuffix(s, "\n")
		nodes[i], _ = strconv.Atoi(s)
		i++
	}
	nodes = nodes[:i]
	rootNode := &node{
		data: nodes[0],
	}
	for i := 1; i < len(nodes); i++ {
		rootNode.Insert(nodes[i])
	}
	rootNode.PostOrder()
}
