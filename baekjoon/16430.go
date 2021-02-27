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
	var A, B int
	fmt.Fscanf(r, "%d %d\n", &A, &B)
	fmt.Fprintln(w, B-A, B)
}
