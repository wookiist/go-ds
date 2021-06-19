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
	S, T  int64
	check map[int64]struct{}
	D     map[int64]string
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	S, T = scanInt(), scanInt()
	if S == T {
		fmt.Fprintln(w, 0)
		return
	}
	check = make(map[int64]struct{})
	D = make(map[int64]string)
	q := make([]int64, 0)
	q = append(q, S)
	check[S] = struct{}{}
	D[S] = ""
	for len(q) != 0 {
		nx := q[0]
		if nx == T {
			break
		}
		q = q[1:]
		if nx*nx <= 1_000_000_000 {
			if _, ok := check[nx*nx]; !ok {
				q = append(q, nx*nx)
				check[nx*nx] = struct{}{}
				D[nx*nx] = D[nx] + "*"
			}
		}
		if nx+nx <= 1_000_000_000 {
			if _, ok := check[nx+nx]; !ok {
				q = append(q, nx+nx)
				check[nx+nx] = struct{}{}
				D[nx+nx] = D[nx] + "+"
			}
		}
		if _, ok := check[0]; !ok {
			q = append(q, 0)
			check[0] = struct{}{}
			D[0] = D[nx] + "-"
		}
		if S != 0 {
			if _, ok := check[1]; !ok {
				q = append(q, 1)
				check[1] = struct{}{}
				D[1] = D[nx] + "/"
			}
		}
	}
	if _, ok := D[T]; !ok {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, D[T])
	}
}

func scanInt() int64 {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return int64(n)
}
