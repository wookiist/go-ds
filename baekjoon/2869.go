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
	var A, B, V int64
	fmt.Fscanf(r, "%d %d %d\n", &A, &B, &V)
	fmt.Fprintln(w, (V-B-1)/(A-B)+1)
}
