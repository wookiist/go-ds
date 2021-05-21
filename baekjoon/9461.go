package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	D [101]int
	N int
)

func main() {
	defer w.Flush()
	D[1], D[2], D[3] = 1, 1, 1
	for i := 4; i <= 100; i++ {
		D[i] = D[i-2] + D[i-3]
	}
	N = scanInt()
	for ; N > 0; N-- {
		fmt.Fprintln(w, D[scanInt()])
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
