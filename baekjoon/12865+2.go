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
	W    []int
	V    []int
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
			D[i][k] = max(D[i][k], D[i-1][k]) // 담지 않았을 때
			if k-W[i] >= 0 {
				D[i][k] = max(D[i][k], D[i-1][k-W[i]]+V[i])
			}
		}
	}
	fmt.Fprintln(w, D[N][K])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
