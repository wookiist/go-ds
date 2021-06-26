package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	A, B string
	D    [][]int
)

func main() {
	defer w.Flush()
	A = scanText()
	B = scanText()
	D = make([][]int, len(A))
	for i := range D {
		D[i] = make([]int, len(B))
	}
	for i := 1; i < len(A); i++ {
		for j := 1; j < len(B); j++ {
			if A[i] == B[j] {
				D[i][j] = D[i-1][j-1] + 1
			} else {
				D[i][j] = max(D[i-1][j], D[i][j-1])
			}
		}
	}
	res := D[len(A)-1][1]
	for i := range B {
		if res < D[len(A)-1][i] {
			res = D[len(A)-1][i]
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

func scanText() string {
	sc.Scan()
	return " " + sc.Text()
}
