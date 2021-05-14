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
	D  []int
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N := scanInt()
	D = make([]int, 1000001)
	D[1], D[2] = 1, 2
	for i := 3; i <= N; i++ {
		D[i] = (D[i-1] + D[i-2]) % 15746
	}
	fmt.Fprintln(w, D[N]%15746)
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
