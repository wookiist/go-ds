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
	K, W, H int
	G       [][]int
	D       [][][]int
	hx      = [8]int{-1, -2, -2, -1, 1, 2, 2, 1}
	hy      = [8]int{-2, -1, 1, 2, 2, 1, -1, -2}
	dx      = [4]int{1, -1, 0, 0}
	dy      = [4]int{0, 0, 1, -1}
)

type pair struct {
	x, y, k int
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	K, W, H = scanInt(), scanInt(), scanInt()
	G = make([][]int, H+1)
	D = make([][][]int, H+1)
	for i := range D {
		if i != 0 {
			G[i] = scanGraph()
		}
		D[i] = make([][]int, W+1)
		for j := range D[i] {
			D[i][j] = make([]int, K+1)
			for k := range D[i][j] {
				D[i][j][k] = -1
			}
		}
	}
	q := make([]pair, 0)
	D[1][1][K] = 0
	q = append(q, pair{1, 1, K})
	for len(q) != 0 {
		nx, ny, nk := q[0].x, q[0].y, q[0].k
		q = q[1:]
		if nk > 0 {
			for k := range hx {
				nnx, nny := nx+hx[k], ny+hy[k]
				if 1 <= nnx && nnx <= H && 1 <= nny && nny <= W {
					if G[nnx][nny] == 0 && D[nnx][nny][nk-1] == -1 {
						q = append(q, pair{nnx, nny, nk - 1})
						D[nnx][nny][nk-1] = D[nx][ny][nk] + 1
					}
				}
			}
		}
		for k := range dx {
			nnx, nny := nx+dx[k], ny+dy[k]
			if 1 <= nnx && nnx <= H && 1 <= nny && nny <= W {
				if G[nnx][nny] == 0 && D[nnx][nny][nk] == -1 {
					q = append(q, pair{nnx, nny, nk})
					D[nnx][nny][nk] = D[nx][ny][nk] + 1
				}
			}
		}
	}
	fmt.Fprintln(w, min(D[H][W]))
}

func min(arr []int) int {
	cnt := 0
	res := 100_000_000
	for i := range arr {
		if arr[i] == -1 {
			cnt++
		} else if res > arr[i] {
			res = arr[i]
		}
	}
	if cnt == K+1 {
		return -1
	}
	return res
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func scanGraph() []int {
	res := make([]int, W+1)
	for i := range res {
		if i != 0 {
			res[i] = scanInt()
		}
	}
	return res
}
