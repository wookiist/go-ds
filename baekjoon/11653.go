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
	for i := 2; i <= N; i++ {
		for N%i == 0 {
			N /= i
			fmt.Fprintln(w, i)
		}
	}
}
