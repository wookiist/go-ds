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
	D [1004][1004]int
)

type pair struct {
	x, c int
}

func main() {
	defer w.Flush()
	S := scanInt()
	for i := range D {
		for j := range D[i] {
			D[i][j] = -1
		}
	}
	D[1][0] = 0
	q := make([]pair, 0)
	q = append(q, pair{1, 0})
	for len(q) != 0 {
		nx, nc := q[0].x, q[0].c
		q = q[1:]
		// clipboard copy
		if D[nx][nx] == -1 {
			D[nx][nx] = D[nx][nc] + 1
			q = append(q, pair{nx, nx})
		}
		// clipboard paste
		if nx+nc < 1004 && D[nx+nc][nc] == -1 {
			D[nx+nc][nc] = D[nx][nc] + 1
			if nx+nc == S {
				fmt.Fprintln(w, D[nx+nc][nc])
				break
			}
			q = append(q, pair{nx + nc, nc})
		}
		// delete
		if 0 <= nx-1 && D[nx-1][nc] == -1 {
			D[nx-1][nc] = D[nx][nc] + 1
			if nx-1 == S {
				fmt.Fprintln(w, D[nx-1][nc])
				break
			}
			q = append(q, pair{nx - 1, nc})
		}
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
