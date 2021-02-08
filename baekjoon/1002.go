package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var x1, y1, r1, x2, y2, r2 float64
	var T int
	fmt.Fscanf(r, "%d\n", &T)
	for i := 0; i < T; i++ {
		fmt.Fscanf(r, "%f %f %f %f %f %f\n", &x1, &y1, &r1, &x2, &y2, &r2)
		if x1 == x2 && y1 == y2 && r1 == r2 {
			fmt.Fprintln(w, -1)
		} else if x1 == x2 && y1 == y2 && r1 != r2 {
			fmt.Fprintln(w, 0)
		} else {
			d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
			if math.Abs(r1-r2) < d && d < r1+r2 {
				fmt.Fprintln(w, 2)
			} else if math.Abs(r1-r2) == d || r1+r2 == d {
				fmt.Fprintln(w, 1)
			} else {
				fmt.Fprintln(w, 0)
			}
		}
	}
}
