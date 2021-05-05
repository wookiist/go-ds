package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w   = bufio.NewWriter(os.Stdout)
	r   = bufio.NewReader(os.Stdin)
	D   [][]int
	N   int
	Div = 1000000009
)

func main() {
	defer w.Flush()
	var T int
	fmt.Fscanf(r, "%d\n", &T)
	D = make([][]int, 100001)
	for i := range D {
		D[i] = make([]int, 4)
	}
	D[1][1] = 1
	D[2][2] = 1
	D[3][1] = 1
	D[3][2] = 1
	D[3][3] = 1
	for i := 4; i <= 100000; i++ {
		D[i][1] = (D[i-1][2] + D[i-1][3]) % Div
		D[i][2] = (D[i-2][1] + D[i-2][3]) % Div
		D[i][3] = (D[i-3][1] + D[i-3][2]) % Div
	}
	for ; T > 0; T-- {
		fmt.Fscanf(r, "%d\n", &N)
		fmt.Fprintln(w, sum(D[N]))
	}
}

func sum(arr []int) int {
	res := 0
	for i := range arr {
		res += (arr[i]) % Div
	}
	return res % Div
}
