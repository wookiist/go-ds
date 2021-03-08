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
	for {
		s, _ := r.ReadString('.')
		s = strings.TrimSuffix(s, "\n")
		if len(s) == 2 && s[0] == 10 && s[1] == '.' {
			break
		}
		stack := make([]byte, 0)
		for i := range s {
			switch s[i] {
			case 40: // (
				stack = append(stack, 40)
			case 91: // [
				stack = append(stack, 91)
			case 41: // )
				if len(stack) == 0 {
					stack = append(stack, 41)
					break
				}
				if stack[len(stack)-1] == 40 {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, 93)
					break
				}
			case 93: // ]
				if len(stack) == 0 {
					stack = append(stack, 93)
					break
				}
				if stack[len(stack)-1] == 91 {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, 93)
					break
				}
			}
		}
		if len(stack) == 0 {
			fmt.Fprintln(w, "yes")
		} else {
			fmt.Fprintln(w, "no")
		}
	}
}
