package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	N int
	G []string
	D []int
)

func main() {
	defer w.Flush()
	N, G = scanInt(), scanString()
	D = make([]int, N+1)
	for i := range D {
		D[i] = -1
	}
	D[1] = 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			if ((G[i] == "O" && G[j] == "B") || (G[i] == "J" && G[j] == "O") || (G[i] == "B" && G[j] == "J")) && D[j] != -1 {
				if D[i] == -1 {
					D[i] = D[j] + (i-j)*(i-j)
				} else if D[i] > D[j]+(i-j)*(i-j) {
					D[i] = D[j] + (i-j)*(i-j)
				}
			}
		}
	}
	fmt.Fprintln(w, D[N])
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func scanString() []string {
	sc.Scan()
	s := strings.Split(" "+sc.Text(), "")
	return s
}
