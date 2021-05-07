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
	D []int
	A []int
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d\n", &N)
	D = make([]int, N+1)
	A = getInts()
	D[1] = 1
	for i := 2; i <= N; i++ {
		maxNum := 1
		for j := 1; j < i; j++ {
			if A[j] < A[i] {
				if maxNum < D[j]+1 {
					maxNum = D[j] + 1
				}
			}
		}
		D[i] = maxNum
	}
	fmt.Fprintln(w, max(D))
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

func max(arr []int) int {
	res := arr[0]
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
