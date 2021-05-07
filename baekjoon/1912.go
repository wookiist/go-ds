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
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d\n", &N)
	A = getInts()
	D = make([]int, N+1)

	for i := 1; i <= N; i++ {
		if D[i-1]+A[i] > A[i] {
			D[i] = D[i-1] + A[i]
		} else {
			D[i] = A[i]
		}
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
	res := arr[1]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
