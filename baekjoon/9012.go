package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

// Stack type
type Stack []string

// Pop from stack
func (s *Stack) Pop() string {
	if len(*s) != 0 {
		data := (*s)[len(*s)-1]
		(*s) = (*s)[:len(*s)-1]
		return data
	}
	return ""
}

// Push to stack
func (s *Stack) Push(str string) {
	(*s) = append((*s), str)
}

// Clean stack
func (s *Stack) Clean() {
	(*s) = (*s)[:0]
}

func main() {
	defer w.Flush()
	var t, l int
	fmt.Fscanf(r, "%d\n", &t)
	var stack Stack
	for i := 0; i < t; i++ {
		text, _ := r.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		split := strings.Split(text, "")
		for idx, item := range split {
			l = idx - 1
			if item == ")" {
				if ret := stack.Pop(); ret == "" {
					// 비어있는데 )가 들어오면 에러
					break
				}
			} else if item == "(" {
				stack.Push("(")
			}
			l = idx
		}
		if len(stack) == 0 && l == len(split)-1 {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}
		stack.Clean()
	}
}
