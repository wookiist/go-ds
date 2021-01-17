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

// Stack type
type Stack []int

// Pop from stack
func (s *Stack) Pop() int {
	if len(*s) != 0 {
		data := (*s)[len(*s)-1]
		(*s) = (*s)[:len(*s)-1]
		return data
	}
	return -1
}

// Push to stack
func (s *Stack) Push(k int) {
	(*s) = append((*s), k)
}

// Empty stack
func (s *Stack) Empty() int {
	if len(*s) == 0 {
		return 1
	}
	return 0
}

// Size stack
func (s *Stack) Size() int {
	return len(*s)
}

// Top stack
func (s *Stack) Top() int {
	if s.Empty() == 1 {
		return -1
	}
	return (*s)[s.Size()-1]
}

func main() {
	defer w.Flush()
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	var stack Stack
	for i := 0; i < n; i++ {
		text, _ := r.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		split := strings.Split(text, " ")
		switch split[0] {
		case "push":
			item, _ := strconv.Atoi(split[1])
			stack.Push(item)
		case "top":
			fmt.Fprintf(w, "%d\n", stack.Top())
		case "pop":
			fmt.Fprintf(w, "%d\n", stack.Pop())
		case "empty":
			fmt.Fprintf(w, "%d\n", stack.Empty())
		case "size":
			fmt.Fprintf(w, "%d\n", stack.Size())
		}
	}
}
