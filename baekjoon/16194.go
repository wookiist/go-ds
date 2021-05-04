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
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			D[i] = min(D[i], D[i-j]+P[j])
		}
	}
	fmt.Fprintln(w, D[N])
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

func min(a, b int) int {
	if a < b && a != 0 {
		return a
	} else if a > b && b == 0 {
		return a
	}
	return b
}
