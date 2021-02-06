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
	solve(0, N, M)
}

func solve(idx, n, m int) {
	if idx == m {
		for i := range seq {
			w.WriteByte(byte(seq[i]) + '0')
			w.WriteByte(' ')
		}
		w.WriteByte('\n')
		return
	}
	for i := 1; i <= n; i++ {
		seq[idx] = i
		solve(idx+1, n, m)
	}
}
