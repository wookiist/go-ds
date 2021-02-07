package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var X, result int
	fmt.Fscanf(r, "%d\n", &X)
	result = solve(X)
	fmt.Fprintf(w, "%d", result)
}

func solve(x int) int {
	memo := make([]int, x+1)
	memo[1] = 0
	for i := 2; i <= x; i++ {
		tmp := 10000000
		if i%3 == 0 {
			if tmp > memo[i/3]+1 {
				tmp = memo[i/3] + 1
			}
		}
		if i%2 == 0 {
			if tmp > memo[i/2]+1 {
				tmp = memo[i/2] + 1
			}
		}
		if tmp > memo[i-1]+1 {
			tmp = memo[i-1] + 1
		}
		memo[i] = tmp
	}
	return memo[x]
}
