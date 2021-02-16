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
	var N, money int
	fmt.Fscanf(r, "%d\n", &N)
	T := make([]int, N+1)
	P := make([]int, N+1)
	D := make([]int, N+2)
	for i := 1; i <= N; i++ {
		fmt.Fscanf(r, "%d %d\n", &T[i], &P[i])
	}
	for i := 1; i <= N; i++ {
		if i+T[i]-1 <= N {
			D[i+T[i]] = max(D[i]+P[i], D[i+T[i]])
			money = max(money, D[i+T[i]])
		}
		D[i+1] = max(D[i+1], D[i])
		money = max(money, D[i+1])
	}
	fmt.Fprintln(w, money)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
