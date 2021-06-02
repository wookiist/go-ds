package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	P [100001]int
	D [100001]int
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N, K := scanInt(), scanInt()
	for i := range P {
		P[i] = -2
	}
	q := make([]int, 0)
	q = append(q, N)
	P[N] = -1
	D[N] = 0
	for len(q) != 0 {
		nx := q[0]
		q = q[1:]
		if 0 <= nx-1 {
			if P[nx-1] == -2 {
				q = append(q, nx-1)
				D[nx-1] = D[nx] + 1
				P[nx-1] = nx
			}
		}
		if nx+1 <= 100000 {
			if P[nx+1] == -2 {
				q = append(q, nx+1)
				D[nx+1] = D[nx] + 1
				P[nx+1] = nx
			}
		}
		if 2*nx <= 100000 {
			if P[2*nx] == -2 {
				q = append(q, 2*nx)
				D[2*nx] = D[nx] + 1
				P[2*nx] = nx
			}
		}
		if P[K] != -2 {
			break
		}
	}
	fmt.Fprintln(w, D[K])
	printRoute(K)
}

func printRoute(S int) {
	if P[S] == -1 {
		fmt.Fprintf(w, "%d ", S)
		return
	}
	printRoute(P[S])
	fmt.Fprintf(w, "%d ", S)
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
