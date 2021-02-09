package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var N int
	fmt.Fscanf(r, "%d\n", &N)

	for i := 1; i <= N; i++ {
		n := i
		sum := 0
		for n != 0 {
			sum += n % 10
			n /= 10
		}
		if sum+i == N {
			fmt.Fprintln(w, i)
			return
		}
	}
	fmt.Fprintln(w, 0)
}
