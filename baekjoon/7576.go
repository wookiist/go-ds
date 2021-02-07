package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w        = bufio.NewWriter(os.Stdout)
	r        = bufio.NewReader(os.Stdin)
	graph    [][]int
	distance [][]int
	n        int
	m        int
	dx       = []int{1, -1, 0, 0}
	dy       = []int{0, 0, 1, -1}
)

type pair struct {
	X int
	Y int
}

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d %d\n", &m, &n)
	graph = make([][]int, n)
	distance = make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, m)
		distance[i] = make([]int, m)
	}
	makeGraph()
	q := make([]pair, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] == 1 {
				q = append(q, pair{i, j})
				distance[i][j] = 0
			}
		}
	}
	bfs(q)
	fmt.Fprintln(w, checkGraph())
}

func makeGraph() {
	for i := range graph {
		s, _ := r.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		t := strings.Split(s, " ")
		for j := range t {
			graph[i][j], _ = strconv.Atoi(t[j])
			distance[i][j] = -1
		}
	}
}

func checkGraph() (max int) {
	for i := range graph {
		for j := range graph[i] {
			if max < distance[i][j] {
				max = distance[i][j]
			}
			if graph[i][j] == 0 && distance[i][j] == -1 {
				return -1
			}
		}
	}
	return max
}

func bfs(q []pair) {
	for len(q) != 0 {
		nx, ny := q[0].X, q[0].Y
		q = q[1:]
		for k := range dx {
			nnx, nny := nx+dx[k], ny+dy[k]
			if 0 <= nnx && nnx < n && 0 <= nny && nny < m {
				if graph[nnx][nny] == 0 && distance[nnx][nny] == -1 {
					q = append(q, pair{nnx, nny})
					distance[nnx][nny] = distance[nx][ny] + 1
				}
			}
		}
	}
}
