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

func hanSoo(n int) bool {
	if n >= 100 {
		if (n/100 - n%100/10) == (n%100/10 - n%10) {
			return true
		}
		return false
	}
	return true
}

func main() {
	defer w.Flush()
	var N, result int
	fmt.Fscanf(r, "%d\n", &N)
	for i := 1; i <= N; i++ {
		if hanSoo(i) {
			result++
		}
	}
	fmt.Fprintln(w, result)
}
