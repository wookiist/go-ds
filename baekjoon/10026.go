package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w       = bufio.NewWriter(os.Stdout)
	r       = bufio.NewReader(os.Stdin)
	n       int
	graph   [][]byte
	visited [][]bool
	dx      = [4]int{1, -1, 0, 0}
	dy      = [4]int{0, 0, 1, -1}
)

type pair struct {
	x int
	y int
}

func main() {
	defer w.Flush()
	var normal, blind int
	fmt.Fscanf(r, "%d\n", &n)
	graph = make([][]byte, n)
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = getBytes()
		visited[i] = make([]bool, n)
	}
	// R = 82 / B = 66  / G = 77 / 적록 => R == G
	for i := range graph {
		for j := range graph[i] {
			if !visited[i][j] {
				bfs(i, j)
				if graph[i][j] == 'G' {
					graph[i][j] += 11
				}
				normal++
			}
		}
	}
	visited = make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
		for j := range graph[i] {
			if graph[i][j] == 'G' {
				graph[i][j] += 11
			}
		}
	}
	for i := range graph {
		for j := range graph[i] {
			if !visited[i][j] {
				bfs(i, j)
				blind++
			}
		}
	}
	fmt.Fprintln(w, normal, blind)
}

func getBytes() []byte {
	b, _ := r.ReadBytes('\n')
	return b[:len(b)-1]
}

func bfs(x, y int) {
	q := []pair{{x, y}}
	visited[x][y] = true

	for len(q) != 0 {
		nx, ny := q[0].x, q[0].y
		q = q[1:]
		for k := range dx {
			nnx, nny := nx+dx[k], ny+dy[k]
			if 0 <= nnx && nnx < n && 0 <= nny && nny < n {
				if graph[nnx][nny] == graph[nx][ny] && !visited[nnx][nny] {
					q = append(q, pair{nnx, nny})
					visited[nnx][nny] = true
				}
			}
		}
	}
}
