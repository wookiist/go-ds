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
	S := getString()
	for i := range S {
		n := adder(rune(S[i]))
		fmt.Fprintln(w, strings.Repeat(string(S[i]), n))
	}
}

func getString() string {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	return s
}

func adder(a rune) int {
	x := int(a)
	res := 0
	for {
		res += x % 10
		x /= 10
		if x == 0 {
			break
		}
	}
	return res
}
