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
	var N int
	for {
		fmt.Fscanf(r, "%d\n", &N)
		if N == -1 {
			break
		}
		if arr, ok := isPerfect(N); ok {
			fmt.Fprintf(w, "%d %s", N, "=")
			for i := range arr {
				fmt.Fprintf(w, " %d ", arr[i])
				if i != len(arr)-1 {
					fmt.Fprintf(w, "%s", "+")
				}
			}
			fmt.Fprintln(w)
		} else {
			fmt.Fprintln(w, N, "is NOT perfect.")
		}

	}
}

func isPerfect(n int) ([]int, bool) {
	arr := make([]int, 0, n)
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
			arr = append(arr, i)
		}
	}
	if sum == n {
		return arr, true
	}
	return nil, false
}
