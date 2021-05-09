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
	r  = bufio.NewReader(os.Stdin)
	A  []int
	D  []int
	Dr []int
	N  int
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d\n", &N)
	A = getInts()
	D = make([]int, N+1)
	Dr = make([]int, N+1)

	// 가장 긴 증가하는 부분 수열
	for i := 1; i <= N; i++ {
		D[i] = 1
		for j := 1; j < i; j++ {
			if A[i] > A[j] && D[i] < D[j]+1 {
				D[i] = D[j] + 1
			}
		}
	}

	// 가장 긴 감소하는 부분 수열
	for i := N; i >= 1; i-- {
		Dr[i] = 1
		for j := N; j > i; j-- {
			if A[i] > A[j] && Dr[i] < Dr[j]+1 {
				Dr[i] = Dr[j] + 1
			}
		}
	}
	for i := range D {
		D[i] += Dr[i] - 1
	}
	fmt.Fprintln(w, max(D))
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

func max(arr []int) int {
	res := arr[0]
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
