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
	D    [][]int
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
	D = make([][]int, k+1)
	for i := range D {
		D[i] = make([]int, n+1)
		for j := range D[i] {
			D[i][j] = 10001
		}
	}
	for i := 1; i <= n; i++ {
		if C[i] <= k {
			D[C[i]][i] = 1
		}
	}
	for i := 1; i <= k; i++ {
		for j := 1; j <= n; j++ {
			if i-C[j] < 0 {
				continue
			}
			for l := 1; l <= n; l++ {
				if D[i][j] > D[i-C[j]][l]+1 {
					D[i][j] = D[i-C[j]][l] + 1
				}
			}
		}
	}
	result := min(D[k][1:])
	if result == 10001 {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, result)
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func min(arr []int) int {
	res := arr[0]
	for i := range arr {
		if res > arr[i] && arr[i] != -1 {
			res = arr[i]
		}
	}
	return res
}
