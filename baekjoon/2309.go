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
	var total int
	var f bool
	heights := make([]int, 9)
	for i := range heights {
		fmt.Fscanf(r, "%d\n", &heights[i])
		total += heights[i]
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})

	for i := 0; i < 9; i++ {
		for j := i; j < 9; j++ {
			if total-heights[i]-heights[j] == 100 {
				for k := 0; k < 9; k++ {
					if k != i && k != j {
						fmt.Fprintln(w, heights[k])
						f = true
					}
				}
			}
		}
		if f {
			break
		}
	}
}
