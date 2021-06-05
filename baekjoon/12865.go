package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	w    = bufio.NewWriter(os.Stdout)
	r    = bufio.NewReader(os.Stdin)
	s    = bufio.NewScanner(os.Stdin)
	N, K int
	D    [][]int
	W    []int
	V    []int
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	N = scanInt()
	K = scanInt()
	D = make([][]int, N+1)
	for i := range D {
		D[i] = make([]int, K+1)
	}
	W = make([]int, N+1)
	V = make([]int, N+1)
	for i := 1; i <= N; i++ {
		W[i] = scanInt()
		V[i] = scanInt()
	}
	for i := 1; i <= N; i++ {
		for j := 0; j <= K; j++ {
			if W[i] > j {
				D[i][j] = D[i-1][j]
			} else {
				D[i][j] = max(D[i-1][j-W[i]]+V[i], D[i-1][j])
			}
		}
	}
	fmt.Fprintln(w, max(D[N]...))
}

func scanInt() int {
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	return n
}

func max(arr ...int) int {
	res := arr[0]
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
======================