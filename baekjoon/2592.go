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
	var x, avg, feq int
	feqMap := make(map[int]int)
	for i := 0; i < 10; i++ {
		fmt.Fscanf(r, "%d\n", &x)
		avg += x
		feqMap[x]++
	}
	for i := range feqMap {
		if feqMap[feq] < feqMap[i] {
			feq = i
		}
	}
	fmt.Fprintf(w, "%d\n%d", avg/10, feq)
}
