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
	P []int
)

func main() {
	defer w.Flush()
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	D = make([]int, N+1)
	P = getInts()
	fmt.Fprintln(w, solve(N))
}

func solve(n int) int {
	if n == 0 {
		return 0
	} else if D[n] != 0 {
		return D[n]
	}
	maxPrice := 0
	for i := 1; i <= n; i++ {
		temp := solve(n - i)
		if temp+P[i] > maxPrice {
			maxPrice = temp + P[i]
		}
	}
	D[n] = maxPrice
	return D[n]
}

func getInts() []int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	res := make([]int, len(t)+1)
	for i := 1; i <= len(t); i++ {
		res[i], _ = strconv.Atoi(t[i-1])
	}
	return res
}
