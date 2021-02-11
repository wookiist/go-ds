// https://www.acmicpc.net/problem/11727
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
	d []int
)

func main() {
	defer w.Flush()
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	d = make([]int, 1001)
	fmt.Fprintln(w, dp(n))
}

func dp(n int) int {
	d[1] = 1
	d[2] = 3
	if n == 1 {
		return 1
	} else if n == 2 {
		return 3
	}
	for i := 3; i <= 1001; i++ {
		d[i] = (d[i-1]%10007 + 2*d[i-2]%10007) % 10007
		if i == n {
			break
		}
	}
	return d[n]
}
