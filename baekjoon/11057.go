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
	fmt.Fscanf(r, "%d\n", &N)
	D = make([][]int, N+1)
	for i := range D {
		D[i] = make([]int, 10)
	}
	for i := range D[1] {
		D[1][i] = 1
	}
	for i := 2; i <= N; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= j; k++ {
				D[i][j] = (D[i][j] + D[i-1][k]) % 10_007
			}
		}
	}
	answer := 0
	for i := range D[N] {
		answer = (answer + D[N][i]) % 10_007
	}
	fmt.Fprintln(w, answer)
}
