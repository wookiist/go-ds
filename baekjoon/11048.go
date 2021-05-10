package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w    = bufio.NewWriter(os.Stdout)
	r    = bufio.NewReader(os.Stdin)
	N, M int
	A    [][]int
	D    [][]int
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	D = make([][]int, N+1)
	A = make([][]int, N+1)
	for i := range D {
		D[i] = make([]int, M+1)
		if i == 0 {
			continue
		}
		A[i] = getInts()
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			D[i][j] = max(D[i-1][j], D[i][j-1], D[i-1][j-1]) + A[i][j]
		}
	}

	fmt.Fprintln(w, D[N][M])
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

func max(arr ...int) int {
	res := arr[0]
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
