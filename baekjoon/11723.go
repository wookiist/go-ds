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

func main() {
	defer w.Flush()
	var M int
	var S uint32
	fmt.Fscanf(r, "%d\n", &M)
	for i := 0; i < M; i++ {
		s, n := getString()
		switch s {
		case "add":
			S = S | (1 << n)
		case "remove":
			S = S & ^(1 << n)
		case "toggle":
			S = S ^ (1 << n)
		case "check":
			if S&(1<<n) == 0 {
				fmt.Fprintf(w, "%d\n", 0)
			} else {
				fmt.Fprintf(w, "%d\n", 1)
			}
		case "all":
			S = (1 << 21) - 1
		case "empty":
			S = 0
		}
	}
}

func getString() (string, int) {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Split(s, " ")
	if len(t) == 1 {
		return t[0], -1
	}
	n, _ := strconv.Atoi(t[1])
	return t[0], n
}
