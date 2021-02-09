package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
	d []int
)

func main() {
	defer w.Flush()
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	d = make([]int, N+1)
	fmt.Fprintln(w, dp(N))
}

func dp(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	d[1] = 1
	d[2] = 2
	for i := 3; i <= n; i++ {
		d[i] = (d[i-1] + d[i-2]) % 10007
	}
	return d[n]
}
