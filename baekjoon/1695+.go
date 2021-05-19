package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	N int
	D [][]int
	A []int
)

func main() {
	defer w.Flush()
	N = scanInt()
	A = scanInts()
	D = make([][]int, N+1)
	for i := range D {
		D[i] = make([]int, N+1)
	}
	fmt.Fprintln(w, solve(1, N))
}

func solve(i, j int) int {
	if i > j {
		return 0
	}
	if i == j {
		D[i][j] = 0
		return D[i][j]
	}
	if D[i][j] != 0 {
		return D[i][j]
	}
	if A[i] == A[j] {
		D[i][j] = solve(i+1, j-1)
	} else {
		D[i][j] = min(solve(i+1, j), solve(i, j-1)) + 1
	}
	return D[i][j]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func scanInts() []int {
	sc.Scan()
	t := strings.Fields("0 " + sc.Text())
	res := make([]int, len(t))
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}
