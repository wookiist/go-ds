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
	N              int
	D              [][]int
	r1, c1, r2, c2 int
	dx             = [6]int{-2, -2, 0, 0, 2, 2}
	dy             = [6]int{-1, 1, -2, 2, -1, 1}
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N = scanInt()
	r1, c1, r2, c2 = scanInt(), scanInt(), scanInt(), scanInt()
	D = make([][]int, N+1)
	for i := range D {
		D[i] = make([]int, N+1)
		for j := range D[i] {
			D[i][j] = -1
		}
	}
	q := make([]pair, 0)
	q = append(q, pair{r1, c1})
	D[r1][c1] = 0
	for len(q) != 0 {
		nx, ny := q[0].x, q[0].y
		q = q[1:]
		for k := range dx {
			nnx, nny := nx+dx[k], ny+dy[k]
			if 0 <= nnx && nnx < N && 0 <= nny && nny < N {
				if D[nnx][nny] == -1 {
					q = append(q, pair{nnx, nny})
					D[nnx][nny] = D[nx][ny] + 1
				}
			}
		}
		if D[r2][c2] != -1 {
			break
		}
	}
	fmt.Fprintln(w, D[r2][c2])
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
