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
	N, M, K int
	D       [][][]int
	G       [][]int
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
	N, M, K = scanInt(), scanInt(), scanInt()
	G = make([][]int, N+1)
	D = make([][][]int, N+1)
	for i := range G {
		if i != 0 {
			G[i] = scanGraph()
		}
		D[i] = make([][]int, M+1)
		for j := range D[i] {
			D[i][j] = make([]int, K+1)
			for k := range D[i][j] {
				D[i][j][k] = -1
			}
		}
	}
	q := make([]pair, 0)
	D[1][1][K] = 1
	q = append(q, pair{1, 1, K})
	for len(q) != 0 {
		nx, ny, nk := q[0].x, q[0].y, q[0].k
		q = q[1:]
		for k := range dx {
			nnx, nny := nx+dx[k], ny+dy[k]
			if 1 <= nnx && nnx <= N && 1 <= nny && nny <= M {
				if G[nnx][nny] == 1 {
					if nk > 0 && D[nnx][nny][nk-1] == -1 {
						D[nnx][nny][nk-1] = D[nx][ny][nk] + 1
						q = append(q, pair{nnx, nny, nk - 1})
					}
				} else if D[nnx][nny][nk] == -1 {
					D[nnx][nny][nk] = D[nx][ny][nk] + 1
					q = append(q, pair{nnx, nny, nk})
				}
			}
		}
	}
	if !isPossible(D[N][M]) {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, min(D[N][M]))
	}
}

func min(arr []int) int {
	res := 100_000_000
	for i := range arr {
		if arr[i] != -1 && res > arr[i] {
			res = arr[i]
		}
	}
	return res
}

func isPossible(arr []int) bool {
	for i := range arr {
		if arr[i] != -1 {
			return true
		}
	}
	return false
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func scanGraph() []int {
	sc.Scan()
	s := strings.Split("0"+sc.Text(), "")
	res := make([]int, len(s))
	for i := range res {
		res[i], _ = strconv.Atoi(s[i])
	}
	return res
}
