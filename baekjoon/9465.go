package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

var (
	A [][]int
	D [][]int
	N int
)

func main() {
	defer w.Flush()
	var T int
	fmt.Fscanf(r, "%d\n", &T)
	for ; T > 0; T-- {
		fmt.Fscanf(r, "%d\n", &N)
		A = make([][]int, 2)
		D = make([][]int, 2)
		for i := range A {
			A[i] = scanInts()
			D[i] = make([]int, N+1)
		}
		for j := 1; j <= N; j++ {
			for i := 0; i < 2; i++ {
				if j == 1 {
					D[i][j] = A[i][j]
					continue
				}
				D[i][j] = max(D[(i-1)*(-1)][j-1], D[(i-1)*(-1)][j-2]) + A[i][j]
			}
		}
		fmt.Fprintln(w, max(D[0][N], D[1][N]))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInts() []int {
	s, _ := r.ReadString('\n')
	s = "0 " + s
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	res := make([]int, len(t))
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}
