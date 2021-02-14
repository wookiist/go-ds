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
	visited []bool
)

func main() {
	defer w.Flush()
	var N, M, u, v, result int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	graph = make([][]int, N+1)
	visited = make([]bool, N+1)
	for i := range graph {
		graph[i] = make([]int, N+1)
	}
	for i := 0; i < M; i++ {
		fmt.Fscanf(r, "%d %d\n", &u, &v)
		graph[u][v] = 1
		graph[v][u] = 1
	}
	for i := 1; i <= N; i++ {
		if !visited[i] {
			bfs(i)
			result++
		}
	}
	fmt.Fprintln(w, result)
}

func bfs(x int) {
	q := []int{x}
	visited[x] = true
	for len(q) != 0 {
		nx := q[0]
		q = q[1:]
		for i := range graph[nx] {
			if graph[nx][i] == 1 && !visited[i] {
				q = append(q, i)
				visited[i] = true
			}
		}
	}
}
