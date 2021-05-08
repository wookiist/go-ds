package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w   = bufio.NewWriter(os.Stdout)
	r   = bufio.NewReader(os.Stdin)
	N   int
	D   []int
	MAX = 999999999
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d\n", &N)
	D = make([]int, N+1)
	for i := 1; i*i <= N; i++ {
		D[i*i] = 1
	}
	for i := 1; i <= N; i++ {
		if D[i] == 1 {
			continue
		}
		for j := 1; j*j <= i; j++ {
			if D[i] == 0 || D[i] > D[j*j]+D[i-j*j] {
				D[i] = D[j*j] + D[i-j*j]
			}
		}
	}
	fmt.Fprintln(w, D[N])
}
