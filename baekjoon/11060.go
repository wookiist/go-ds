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
	A = getInts()
	D = make([]int, N+1)
	for i := range D {
		D[i] = 10001
	}
	if N == 1 {
		fmt.Fprintln(w, 0)
		return
	}
	res := solve(1)
	if res == 10001 {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, res)
	}
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

func solve(n int) int {
	if n >= N {
		return 0
	}
	if D[n] != 10001 {
		return D[n]
	}
	for i := 1; i <= A[n]; i++ {
		temp := solve(n + i)
		if D[n] > temp+1 {
			D[n] = temp + 1
		}
	}
	return D[n]
}
