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
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	arr := make([]int, N+1)
	maxSum := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscanf(r, "%d\n", &arr[i])
	}
	for i := range maxSum {
		maxSum[i] = make([]int, 3)
	}
	maxSum[1][1], maxSum[1][2] = arr[1], arr[1]
	for i := 1; i <= N; i++ {
		maxSum[i][1] = maxSum[i-1][2] + arr[i]
		if i-2 >= 0 {
			maxSum[i][2] = max(maxSum[i-2][1], maxSum[i-2][2]) + arr[i]
		}
	}
	fmt.Fprintln(w, max(maxSum[N][1], maxSum[N][2]))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
