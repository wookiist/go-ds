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
	visited [][]bool
	D       [][]int
	dx      = [8]int{-2, -2, -1, -1, 1, 1, 2, 2}
	dy      = [8]int{1, -1, 2, -2, 2, -2, 1, -1}
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	T := scanInt()
	for ; T > 0; T-- {z
		L := scanInt()
		visited = make([][]bool, L)
		D = make([][]int, L)
		for i := range visited {
			visited[i] = make([]bool, L)
			D[i] = make([]int, L)
		}
		q := make([]pair, 0)
		fx, fy := scanInt(), scanInt()
		ex, ey := scanInt(), scanInt()
		q = append(q, pair{fx, fy})
		visited[fx][fy] = true
		for len(q) != 0 {
			nx, ny := q[0].x, q[0].y
			q = q[1:]
			for k := range dx {
				nnx, nny := nx+dx[k], ny+dy[k]
				if 0 <= nnx && nnx <= L-1 && 0 <= nny && nny <= L-1 {
					if !visited[nnx][nny] {
						visited[nnx][nny] = true
						D[nnx][nny] = D[nx][ny] + 1
						q = append(q, pair{nnx, nny})
					}
				}
			}
			if visited[ex][ey] {
				break
			}
		}
		fmt.Fprintln(w, D[ex][ey])
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
