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
	C  []int
	D  []int
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N, K := scanInt(), scanInt()
	C = make([]int, N+1)
	D = make([]int, K+1)
	for i := 1; i <= N; i++ {
		C[i] = scanInt()
	}
	D[0] = 1
	for i := 1; i <= N; i++ {
		for j := 1; j <= K; j++ {
			if j-C[i] >= 0 {
				D[j] += D[j-C[i]]
			}
		}
	}
	fmt.Fprintln(w, D[K])
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
