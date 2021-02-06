package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w   = bufio.NewWriter(os.Stdout)
	r   = bufio.NewReader(os.Stdin)
	seq []int
)

func main() {
	defer w.Flush()
	var N, M int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	seq = make([]int, M)
	solve(0, 1, N, M)
}

func solve(idx, start, n, m int) {
	if idx == m {
		for i := range seq {
			w.WriteByte(byte(seq[i]) + '0')
			w.WriteByte(' ')
		}
		w.WriteByte('\n')
		return
	}
	for i := start; i <= n; i++ {
		seq[idx] = i
		solve(idx+1, i, n, m)
	}
}
