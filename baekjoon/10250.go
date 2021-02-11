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
	var T, H, W, N, Y, X int
	fmt.Fscanf(r, "%d\n", &T)
	for i := 0; i < T; i++ {
		fmt.Fscanf(r, "%d %d %d\n", &H, &W, &N)
		if N%H == 0 {
			Y = H
			X = N / H
		} else {
			Y = N % H
			X = N/H + 1
		}
		if X < 10 {
			fmt.Fprintf(w, "%d0%d\n", Y, X)
		} else {
			fmt.Fprintf(w, "%d%d\n", Y, X)
		}
	}
}
