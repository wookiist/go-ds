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
	var N, K, k, t, res int
	fmt.Fscanf(r, "%d %d\n", &N, &K)
	arr := make([]int, N)
	for i := range arr {
		fmt.Fscanf(r, "%d\n", &arr[i])
	}
	k = K
	t = K
	for i := len(arr) - 1; i >= 0; i-- {
		if K < arr[i] {
			continue
		}
		t = k / arr[i]
		k %= arr[i]
		res += t
		if k == 0 {
			break
		}
	}
	fmt.Fprintf(w, "%d\n", res)
}
