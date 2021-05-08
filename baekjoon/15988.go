package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w   = bufio.NewWriter(os.Stdout)
	r   = bufio.NewReader(os.Stdin)
	D   []int
	Div = 1_000_000_009
)

func main() {
	defer w.Flush()
	D = make([]int, 1_000_001)
	D[0] = 1
	D[1] = 1
	D[2] = 2
	D[3] = 4

	for i := 4; i <= 1_000_000; i++ {
		D[i] = (D[i-1] + D[i-2] + D[i-3]) % Div
	}

	var T, N int
	fmt.Fscanf(r, "%d\n", &T)
	for ; T > 0; T-- {
		fmt.Fscanf(r, "%d\n", &N)
		fmt.Fprintln(w, D[N]%Div)
	}
}
