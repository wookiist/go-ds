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

type pair struct {
	x, y int
}

var (
	A, B, C, M, S int
	D             [][]bool
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	A, B, C = scanInt(), scanInt(), scanInt()
	S = A + B + C
	if S%3 != 0 {
		fmt.Fprintln(w, 0)
		return
	}
	M = S / 3
	D = make([][]bool, S+1)
	for i := range D {
		D[i] = make([]bool, S+1)
	}
	q := make([]pair, 0)
	D[A][B] = true
	q = append(q, pair{A, B})
	for len(q) != 0 {
		nx, ny := q[0].x, q[0].y
		q = q[1:]
		// nx, ny
		if nx < ny {
			// nx+nx, ny-nx, nz
			if !D[nx+nx][ny-nx] {
				D[nx+nx][ny-nx] = true
				q = append(q, pair{nx + nx, ny - nx})
			}
		} else {
			// nx-ny, ny+ny, nz
			if !D[nx-ny][ny+ny] {
				D[nx-ny][ny+ny] = true
				q = append(q, pair{nx - ny, ny + ny})
			}
		}
		// nx, nz
		if nx < S-(nx+ny) {
			// nx+nx, ny, nz-nx
			if !D[nx+nx][ny] {
				D[nx+nx][ny] = true
				q = append(q, pair{nx + nx, ny})
			}
		} else {
			// nx-nz, ny, nz+nz
			if !D[nx-S+(nx+ny)][ny] {
				D[nx-S+(nx+ny)][ny] = true
				q = append(q, pair{nx - S + (nx + ny), ny})
			}
		}
		// ny, nz
		if ny < S-(nx+ny) {
			// nx, ny+ny, nz-ny
			if !D[nx][ny+ny] {
				D[nx][ny+ny] = true
				q = append(q, pair{nx, ny + ny})
			}
		} else {
			// nx, ny-nz, nz+nz
			if !D[nx][ny-S+(nx+ny)] {
				D[nx][ny-S+(nx+ny)] = true
				q = append(q, pair{nx, ny - S + (nx + ny)})
			}
		}
		if D[M][M] {
			break
		}
	}
	if D[M][M] {
		fmt.Fprintln(w, 1)
	} else {
		fmt.Fprintln(w, 0)
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
