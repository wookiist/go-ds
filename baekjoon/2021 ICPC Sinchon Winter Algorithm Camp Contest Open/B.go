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
	if K == 0 {
		for i := 1; i <= N; i++ {
			fmt.Fprintf(w, "%d ", i)
		}
		return
	}
	if K == N*(N-1)/2 {
		for i := N; i >= 1; i-- {
			fmt.Fprintf(w, "%d ", i)
		}
		return
	}
	solve(N, K)
}

func solve(N, K int) {
	if K == 0 {
		for i := 1; i <= N; i++ {
			fmt.Fprintf(w, "%d ", i)
		}
		return
	}
	if K < N {
		f := K + 1
		fmt.Fprintf(w, "%d ", f)
		for i := 1; i <= N; i++ {
			if i == f {
				continue
			}
			fmt.Fprintf(w, "%d ", i)
		}
	} else if K >= N {
		fmt.Fprintf(w, "%d ", N)
		solve(N-1, K-N+1)
	}
}
