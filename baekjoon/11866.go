package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var N, K int
	fmt.Fscanf(r, "%d %d\n", &N, &K)
	q := make([]int, N)
	for i := 0; i < N; i++ {
		q[i] = i + 1
	}
	fmt.Fprint(w, "<")
	for len(q) != 1 {
		for i := 0; i < K-1; i++ {
			n := q[0]
			q = q[1:]
			q = append(q, n)
		}
		fmt.Fprintf(w, "%d, ", q[0])
		q = q[1:]
	}
	fmt.Fprintf(w, "%d>", q[0])
}
