package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
	N int
	D [][]int
)

func main() {
	defer w.Flush()
	D = make([][]int, 91)
	for i := range D {
		D[i] = make([]int, 3)
	}
	fmt.Fscanf(r, "%d\n", &N)
	D[1][1] = 1
	D[2][0] = 1
	for i := 3; i <= N; i++ {
		for j := 0; j <= 1; j++ {
			if j == 0 {
				D[i][j] += D[i-1][1]
			}
			D[i][j] += D[i-1][0]
		}
	}
	fmt.Fprintln(w, sum(D[N]))
}

func sum(arr []int) int {
	res := 0
	for i := range arr {
		res += arr[i]
	}
	return res
}
