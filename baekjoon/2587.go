package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var x, avg int
	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscanf(r, "%d\n", &x)
		avg += x
		arr[i] = x
	}
	sort.Ints(arr)
	fmt.Fprintf(w, "%d\n%d", avg/5, arr[2])
}
