package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	A []int
	P []pair
	D []int
	N int
)

type pair struct {
	A, B int
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N = scan()
	P = make([]pair, N)
	A = make([]int, N+1)
	D = make([]int, N+1)
	for i := range P {
		P[i] = pair{
			A: scan(),
			B: scan(),
		}
	}
	sort.Slice(P, func(i, j int) bool {
		return P[i].A < P[j].A
	})
	for i := range P {
		A[i+1] = P[i].B
	}
	for i := 1; i <= N; i++ {
		D[i] = 1
		for j := 1; j < i; j++ {
			if A[i] > A[j] {
				D[i] = max(D[i], D[j]+1)
			}
		}
	}
	fmt.Fprintln(w, N-max(D...))
}

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
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
