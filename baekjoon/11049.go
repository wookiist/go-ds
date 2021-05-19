package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type matrix struct {
	r, c int
}

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	N int
	D [][]int
	A []matrix
	M = 125000002
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N = scanInt()
	A = make([]matrix, N+1)
	for i := 1; i <= N; i++ {
		A[i] = matrix{
			r: scanInt(),
			c: scanInt(),
		}
	}
	D = make([][]int, N+1)
	for i := range D {
		D[i] = make([]int, N+1)
		for j := range D[i] {
			D[i][j] = -1
		}
	}
	fmt.Fprintln(w, solve(1, N))
}

func solve(i, j int) int {
	if D[i][j] != -1 {
		return D[i][j]
	}
	if i == j {
		D[i][j] = 0
		return 0
	}
	if i+1 == j {
		D[i][j] = A[i].r * A[i].c * A[j].c
		return D[i][j]
	}
	for k := i; k < j; k++ {
		pre := solve(i, k)
		post := solve(k+1, j)
		if D[i][j] == -1 || D[i][j] > pre+post+(A[i].r*A[k].c*A[j].c) {
			D[i][j] = pre + post + (A[i].r * A[k].c * A[j].c)
		}
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
