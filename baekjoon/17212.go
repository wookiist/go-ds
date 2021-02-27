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
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	d = make([]int, N+11)
	d[0] = 0
	d[1] = 1
	d[2] = 1
	d[3] = 2
	d[4] = 2
	d[5] = 1
	d[6] = 2
	d[7] = 1
	fmt.Fprintln(w, solve(N))
}

func solve(n int) int {
	if n == 1 || n == 2 || n == 5 || n == 7 {
		return 1
	}
	for i := 8; i <= n; i++ {
		if d[i-7] != 0 {
			d[i] = d[i-7] + 1
		}
		if d[i-5] != 0 {
			d[i] = min(d[i], d[i-5]+1)
		}
		if d[i-2] != 0 {
			d[i] = min(d[i], d[i-2]+1)
		}
		if d[i-1] != 0 {
			d[i] = min(d[i], d[i-1]+1)
		}
	}
	return d[n]
}

func min(a, b int) int {
	if a == 0 {
		return b
	}
	if a < b {
		return a
	}
	return b
}

/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dp    []int
	coins = []int{7, 5, 2, 1}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	dp = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = 100000
	}
	for i := 1; i < n+1; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	fmt.Fprintln(writer, dp[n])
}
*/
