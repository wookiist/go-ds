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
	var S string
	var stack []int
	for S != "고무오리 디버깅 끝" {
		S, _ = r.ReadString('\n')
		S = strings.TrimSuffix(S, "\n")
		switch S {
		case "문제":
			stack = append(stack, 1)
		case "고무오리":
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, 1)
				stack = append(stack, 1)
			}
		}
	}
	if len(stack) == 0 {
		fmt.Fprintln(w, "고무오리야 사랑해")
	} else {
		fmt.Fprintln(w, "힝구")
	}
}
