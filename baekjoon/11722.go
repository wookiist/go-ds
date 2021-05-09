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
	D = make([]int, N+1)
	A = getInts()
	for i := N; i >= 1; i-- {
		D[i] = 1
		for j := N; j > i; j-- {
			if A[i] > A[j] && D[i] < D[j]+1 {
				D[i] = D[j] + 1
			}
		}
	}
	fmt.Fprintln(w, max(D))
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
