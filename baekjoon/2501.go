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
	var N, K int
	fmt.Fscanf(r, "%d %d\n", &N, &K)
	for i := 1; i <= N; i++ {
		if N%i == 0 {
			K--
		}
		if K == 0 {
			fmt.Fprintln(w, i)
			break
		}
	}
	if K != 0 {
		fmt.Fprintln(w, 0)
	}
}
