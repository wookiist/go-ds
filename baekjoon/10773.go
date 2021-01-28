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
	var K, res, tmp int
	fmt.Fscanf(r, "%d\n", &K)
	stack := make([]int, 0)
	for i := 0; i < K; i++ {
		fmt.Fscanf(r, "%d\n", &tmp)
		if tmp != 0 {
			stack = append(stack, tmp)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	for _, v := range stack {
		res += v
	}
	fmt.Fprintf(w, "%d\n", res)
}
