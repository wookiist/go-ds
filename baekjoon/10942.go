package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w          = bufio.NewWriter(os.Stdout)
	r          = bufio.NewReader(os.Stdin)
	N, M, S, E int
	D          [][]int
	A          []int
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d\n", &N)
	A = getInts()
	fmt.Fscanf(r, "%d\n", &M)
	D = make([][]int, N+1)
	for i := range D {
		D[i] = make([]int, N+1)
		for j := range D[i] {
			D[i][j] = -1
		}
	}
	for i := 1; i <= N; i++ {
		for j := i; j <= N; j++ {
			solve(i, j)
		}
	}
	// var a, b int -> memory 38460KB time 1052ms
	for ; M > 0; M-- {
		q := getInts() // -> memory 68768KB time 504ms
		// fmt.Fscanf(r, "%d %d\n", &a, &b)
		fmt.Fprintln(w, D[q[1]][q[2]])
		// fmt.Fprintln(w, D[a][b])
	}
}

func getInts() []int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	s = "0 " + s
	t := strings.Fields(s)
	res := make([]int, len(t))
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}

func solve(s, e int) int {
	if s-e > 0 {
		return 1
	}
	if s == e {
		D[s][e] = 1
	}
	if D[s][e] != -1 {
		return D[s][e]
	}
	res := (solve(s+1, e-1) == 1) && (A[s] == A[e])
	if res {
		D[s][e] = 1
	} else {
		D[s][e] = 0
	}
	return D[s][e]
}
