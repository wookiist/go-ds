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

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N = scanInt()
	D[1], D[2], D[3] = 1, 2, 3
	for i := 4; i <= N; i++ {
		temp := 0
		for j := 1; j <= i-2; j++ {
			temp = max(temp, (j+1)*D[i-(j+2)])
		}
		D[i] = max(D[i-1]+1, temp)
	}
	fmt.Fprintln(w, D[N])
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func max(arr ...int) int {
	res := arr[0]
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
