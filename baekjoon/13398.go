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
	D [][2]int
	A []int
	N int
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d\n", &N)
	D = make([][2]int, N+1)
	for i := range D {
		D[i][0] = -100_000_000
		D[i][1] = -100_000_000
	}
	A = getInts()
	D[1][1] = A[1]
	for i := 2; i <= N; i++ {
		D[i][0] = max(D[i-1][1], D[i-1][0]+A[i])
		D[i][1] = max(D[i-1][1]+A[i], A[i])
	}
	fmt.Fprintln(w, maxElem(D))
}

func getInts() []int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	s = "0 " + s
	t := strings.Fields(s)
	res := make([]int, len(t))
	for i := 1; i < len(t); i++ {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxElem(arr [][2]int) int {
	res := arr[1][1]
	for i := 1; i <= N; i++ {
		for j := range arr[i] {
			if res < arr[i][j] {
				res = arr[i][j]
			}
		}
	}
	return res
}
