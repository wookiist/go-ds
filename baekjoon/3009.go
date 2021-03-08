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
	var ansX, ansY int
	xSlice := make([]int, 3)
	ySlice := make([]int, 3)
	fmt.Fscanln(r, &xSlice[0], &ySlice[0])
	fmt.Fscanln(r, &xSlice[1], &ySlice[1])
	fmt.Fscanln(r, &xSlice[2], &ySlice[2])
	for i := 0; i < 3; i++ {
		if count(xSlice[i], xSlice) == 1 {
			ansX = xSlice[i]
		}
		if count(ySlice[i], ySlice) == 1 {
			ansY = ySlice[i]
		}
	}
	fmt.Fprintln(w, ansX, ansY)
}

func count(n int, arr []int) int {
	res := 0
	for i := range arr {
		if arr[i] == n {
			res++
		}
	}
	return res
}
