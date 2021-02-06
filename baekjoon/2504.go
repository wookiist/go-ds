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

func main() {
	defer w.Flush()
	result := 0
	tmpResult := 1
	s := getString()
	if !checkString(s) {
		fmt.Fprintln(w, 0)
		return
	}
	for i := range s {
		switch s[i] {
		case "(":
			tmpResult *= 2
		case "[":
			tmpResult *= 3
		case ")":
			if s[i-1] == "(" {
				result += tmpResult
			}
			tmpResult /= 2
		case "]":
			if s[i-1] == "[" {
				result += tmpResult
			}
			tmpResult /= 3
		}
	}
	fmt.Fprintln(w, result)
}

func getString() []string {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Split(s, "")
	return t
}

func checkString(s []string) bool {
	stack := make([]string, 0, len(s))
	for i := range s {
		switch s[i] {
		case "(":
			stack = append(stack, "(")
		case "[":
			stack = append(stack, "[")
		case ")":
			if len(stack) != 0 {
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, ")")
			}
		case "]":
			if len(stack) != 0 {
				if stack[len(stack)-1] == "[" {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, "]")
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
