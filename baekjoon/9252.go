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
	A, B   string
	D      [][]int
	result []string
)

func main() {
	defer w.Flush()
	A, B = " "+scan(), " "+scan()
	n, m := len(A)-1, len(B)-1
	D = make([][]int, n+1)
	result = make([]string, 0)
	for i := range D {
		D[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if A[i] == B[j] {
				D[i][j] = D[i-1][j-1] + 1
			} else {
				D[i][j] = max(D[i-1][j], D[i][j-1])
			}
		}
	}
	fmt.Fprintln(w, D[n][m])
	solve(n, m)
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Fprintf(w, "%s", result[i])
	}
	fmt.Fprintln(w)
}

func solve(i, j int) {
	if D[i][j] == 0 {
		return
	}
	if D[i][j] == D[i-1][j] {
		solve(i-1, j)
	} else if D[i][j] == D[i][j-1] {
		solve(i, j-1)
	} else {
		result = append(result, string(A[i]))
		solve(i-1, j-1)
	}
}

func scan() string {
	sc.Scan()
	return sc.Text()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
