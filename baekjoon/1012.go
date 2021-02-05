package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w       = bufio.NewWriter(os.Stdout)
	r       = bufio.NewReader(os.Stdin)
	graph   [][]int
	visited [][]bool
	dx      = []int{1, -1, 0, 0}
	dy      = []int{0, 0, 1, -1}
)

type axis struct {
	X int
	Y int
}

func main() {
	defer w.Flush()
	var T int
	fmt.Fscanf(r, "%d\n", &T)
	for i := 0; i < T; i++ {
		fmt.Fprintln(w, solve())
	}
}

func solve() int {
	var M, N, K, X, Y, result int
	fmt.Fscanf(r, "%d %d %d\n", &M, &N, &K)
	graph = make([][]int, N)
	visited = make([][]bool, N)
	for i := range graph {
		graph[i] = make([]int, M)
		visited[i] = make([]bool, M)
	}
	for i := 0; i < K; i++ {
		fmt.Fscanf(r, "%d %d\n", &X, &Y)
		graph[Y][X] = 1
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if graph[i][j] == 1 && !visited[i][j] {
				result++
				bfs(j, i, M, N)
			}
		}
	}
	return result
}

func bfs(x, y, M, N int) {
	q := []axis{{x, y}}
	visited[y][x] = true
	for len(q) != 0 {
		nx, ny := q[0].X, q[0].Y
		q = q[1:]
		for i := 0; i < len(dx); i++ {
			nnx, nny := nx+dx[i], ny+dy[i]
			if 0 <= nnx && nnx < M && 0 <= nny && nny < N {
				if graph[nny][nnx] == 1 && !visited[nny][nnx] {
					q = append(q, axis{nnx, nny})
					visited[nny][nnx] = true
				}
			}
		}
	}
}
