package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
	D []int
	N int
)

func main() {
	defer w.Flush()
	var T, n int
	D = make([]int, 10001)
	fmt.Fscanf(r, "%d\n", &T)
	D[0] = 1
	nums := [3]int{1, 2, 3}
	for j := 0; j < 3; j++ {
		for i := 1; i <= 10000; i++ {
			if i-nums[j] >= 0 {
				D[i] += D[i-nums[j]]
			}
		}
	}
	for i := 0; i < T; i++ {
		fmt.Fscanf(r, "%d\n", &n)
		fmt.Fprintln(w, D[n])
	}
}
