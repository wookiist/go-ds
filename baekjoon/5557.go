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
	D      [][]int
	A      []int
	N      int
	answer int
)

func main() {
	defer w.Flush()
	N = scanInt()
	A = scanInts()
	D = make([][]int, N+1)
	for i := range D {
		D[i] = make([]int, 21)
		for j := range D[i] {
			D[i][j] = -1
		}
	}
	fmt.Fprintln(w, solve(N-1, A[N]))
}

func solve(idx, sum int) int {
	if idx <= 0 || 0 > sum || sum > 20 {
		return 0
	}
	if idx == 1 && sum == A[idx] {
		D[idx][sum] = 1
		return D[idx][sum]
	}
	if D[idx][sum] != -1 {
		return D[idx][sum]
	}

	D[idx][sum] = solve(idx-1, sum-A[idx]) + solve(idx-1, sum+A[idx])
	return D[idx][sum]
}
func scanInts() []int {
	sc.Scan()
	s := "0 " + sc.Text()
	t := strings.Fields(s)
	res := make([]int, len(t))
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
