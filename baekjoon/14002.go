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
	N int
	A []int
	D []int
	P []int
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d\n", &N)
	A = getInts()
	D = make([]int, N+1)
	P = make([]int, N+1)
	for i := 1; i <= N; i++ {
		D[i] = 1
		for j := 1; j < i; j++ {
			if A[j] < A[i] {
				if D[i] < D[j]+1 {
					D[i] = D[j] + 1
					P[i] = j
				}
			}
		}
	}
	idx := maxElem(D)
	fmt.Fprintln(w, D[idx])
	R := make([]int, D[idx])
	pIdx := idx
	for i := 0; i < D[idx]; i++ {
		R[i] = A[pIdx]
		pIdx = P[pIdx]
	}
	for i := D[idx] - 1; i >= 0; i-- {
		fmt.Fprintf(w, "%d ", R[i])
	}
	fmt.Fprintln(w)
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

func maxElem(arr []int) int {
	res := arr[0]
	idx := 0
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
			idx = i
		}
	}
	return idx
}
