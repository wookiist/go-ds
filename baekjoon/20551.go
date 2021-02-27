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
	var N, M int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	Ns := make([]int, N)
	Ms := make([]int, M)
	for i := range Ns {
		fmt.Fscanf(r, "%d\n", &Ns[i])
	}
	for i := range Ms {
		fmt.Fscanf(r, "%d\n", &Ms[i])
	}
	sort.Slice(Ns, func(i, j int) bool {
		return Ns[i] < Ns[j]
	})
	for i := range Ms {
		fmt.Fprintln(w, binarySearch(Ms[i], Ns))
	}
}

func binarySearch(target int, arr []int) int {
	s, e, m := 0, len(arr)-1, 0
	for s <= e {
		m = (s + e) / 2
		if arr[m] < target {
			s = m + 1
		} else {
			e = m - 1
		}
	}
	if s == len(arr) {
		return -1
	}
	if arr[s] != target {
		return -1
	}
	return s
}
