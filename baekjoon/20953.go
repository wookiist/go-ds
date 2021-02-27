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
	var T, a, b int
	fmt.Fscanf(r, "%d\n", &T)
	for ; T > 0; T-- {
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		fmt.Fprintln(w, (a+b)*(a+b)*(a+b-1)/2)
	}
}
