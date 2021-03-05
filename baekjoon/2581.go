package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w       = bufio.NewWriter(os.Stdout)
	r       = bufio.NewReader(os.Stdin)
	isPrime []bool
)

func main() {
	defer w.Flush()
	var N, M, res, min int
	fmt.Fscanf(r, "%d\n", &N)
	fmt.Fscanf(r, "%d\n", &M)
	isPrime = make([]bool, M+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	primeNum(M)
	for i := N; i <= M; i++ {
		if isPrime[i] {
			if min == 0 {
				min = i
			}
			res += i
		}
	}
	if res == 0 {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, res)
		fmt.Fprintln(w, min)
	}
}

func primeNum(m int) {
	for i := 2; i <= m; i++ {
		if isPrime[i] {
			for j := i + i; j <= m; j += i {
				isPrime[j] = false
			}
		}
	}
}
