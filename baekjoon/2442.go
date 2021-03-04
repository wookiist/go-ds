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
	for i := 1; i <= N; i++ {
		for j := N - i; j > 0; j-- {
			fmt.Fprint(w, " ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Fprint(w, "*")
		}
		fmt.Fprintln(w)
	}
}
