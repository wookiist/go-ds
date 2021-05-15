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
	n, k int
	C    []int
	D    []int
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	n, k = scanInt(), scanInt()
	C = make([]int, n+1)
	for i := 1; i <= n; i++ {
		C[i] = scanInt()
	}
	D = make([]int, k+1)
	for i := range D {
		D[i] = 10001
	}
	for i := 1; i <= n; i++ {
		if C[i] <= k {
			D[C[i]] = 1
		}
	}
	for i := 1; i <= k; i++ {
		for j := 1; j <= n; j++ {
			if i-C[j] < 0 {
				continue
			}
			if D[i] > D[i-C[j]]+1 {
				D[i] = D[i-C[j]] + 1
			}
		}
	}
	if D[k] == 10001 {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, D[k])
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
