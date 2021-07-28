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

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N := scanInt()
	A := make(map[int]int)
	for i := 0; i < N; i++ {
		n := scanInt()
		if _, ok := A[n]; !ok {
			A[n] = 1
		} else {
			A[n]++
		}
	}
	target := scanInt()
	if result, ok := A[target]; !ok {
		fmt.Fprintln(w, 0)
	} else {
		fmt.Fprintln(w, result)
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
