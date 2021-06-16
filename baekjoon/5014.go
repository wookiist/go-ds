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
	F, S, G, U, D int
	impossible    = "use the stairs"
	Dist          []int
	check         []bool
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	F, S, G, U, D = scanInt(), scanInt(), scanInt(), scanInt(), scanInt()
	Dist = make([]int, F+1)
	check = make([]bool, F+1)
	q := make([]int, 0)
	q = append(q, S)
	Dist[S] = 0
	check[S] = true
	for len(q) != 0 {
		nx := q[0]
		q = q[1:]
		// Up
		if nx+U <= F && !check[nx+U] {
			Dist[nx+U] = Dist[nx] + 1
			check[nx+U] = true
			q = append(q, nx+U)
		}
		// Down
		if 1 <= nx-D && !check[nx-D] {
			Dist[nx-D] = Dist[nx] + 1
			check[nx-D] = true
			q = append(q, nx-D)
		}
	}
	if check[G] {
		fmt.Fprintln(w, Dist[G])
	} else {
		fmt.Fprintln(w, impossible)
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
