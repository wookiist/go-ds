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
	D [][]int
)

func main() {
	defer w.Flush()
	A, B := " "+scan(), " "+scan()
	lenA, lenB := len(A), len(B)
	D := make([][]int, lenA)
	for i := range D {
		D[i] = make([]int, lenB)
	}
	answer := 0
	for i := 1; i < lenA; i++ {
		for j := 1; j < lenB; j++ {
			if A[i] == B[j] {
				D[i][j] = D[i-1][j-1] + 1
			} else {
				D[i][j] = 0
			}
			if answer < D[i][j] {
				answer = D[i][j]
			}
		}
	}
	fmt.Fprintln(w, answer)
}

func scan() string {
	sc.Scan()
	return sc.Text()
}
