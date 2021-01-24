package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanf(r, "%d\n", &arr[i])
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 0; i < N; i++ {
		fmt.Fprintln(w, arr[i])
	}
}
