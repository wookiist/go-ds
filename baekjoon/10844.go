package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
	d [][]int
)

func main() {
	defer w.Flush()
	d = make([][]int, 101)
	for i := range d {
		d[i] = make([]int, 10)
	}
	for i := range d[1] {
		d[1][i] = 1
	}
	d[1][0] = 0
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	fmt.Fprintln(w, solve(N))
}

func solve(n int) int {
	for i := 2; i <= n; i++ {
		for j := 0; j <= 9; j++ {
			if j == 0 {
				d[i][j] = d[i-1][1] % 1000000000
			} else if j == 9 {
				d[i][j] = d[i-1][8] % 1000000000
			} else {
				d[i][j] = (d[i-1][j-1] + d[i-1][j+1]) % 1000000000
			}
		}
	}
	return sum(d[n])
}

func sum(arr []int) int {
	res := 0
	for i := range arr {
		res += arr[i]
	}
	return res % 1000000000
}
