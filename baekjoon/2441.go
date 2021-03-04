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
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			fmt.Fprint(w, " ")
		}
		for j := N - i; j > 0; j-- {
			fmt.Fprint(w, "*")
		}
		fmt.Fprintln(w)
	}
}
