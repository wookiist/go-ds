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
	D []int
	A []int
	N int
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d\n", &N)
	D = make([]int, N+1)
	A = getInts()
	for i := 1; i <= N; i++ {
		temp := A[i]
		for j := 1; j < i; j++ {
			if A[i] > A[j] {
				temp = max(temp, D[j]+A[i])
			}
		}
		D[i] = temp
	}
	fmt.Fprintln(w, maxElem(D))
}

func maxElem(arr []int) int {
	res := arr[0]
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
