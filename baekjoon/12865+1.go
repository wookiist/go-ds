package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	N, K int
	W, V []int
	D    [][]int
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N, K = scanInt(), scanInt()
	W = make([]int, N+1)
	V = make([]int, N+1)
	D = make([][]int, N+1)
	D[0] = make([]int, K+1)
	for i := 1; i <= N; i++ {
		W[i], V[i] = scanInt(), scanInt()
		D[i] = make([]int, K+1)
		for j := range D[i] {
			D[i][j] = -1
		}
	}
	for i := 1; i <= N; i++ {
		for k := 0; k <= K; k++ {
			if k-W[i] >= 0 {
				D[i][k] = max(D[i][k], V[i]+D[i-1][k-W[i]])
			}
			D[i][k] = max(D[i][k], D[i-1][k])
		}
	}
	res := D[N][0]
	for i := range D[N] {
		if res < D[N][i] {
			res = D[N][i]
		}
	}
	fmt.Fprintln(w, res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
