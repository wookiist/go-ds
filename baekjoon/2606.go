package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w       = bufio.NewWriter(os.Stdout)
	r       = bufio.NewReader(os.Stdin)
	visited []bool
	graph   [][]int
)

func main() {
	defer w.Flush()
	var N, E int
	fmt.Fscanf(r, "%d\n", &N)
	fmt.Fscanf(r, "%d\n", &E)
	visited = make([]bool, N+1)
	graph = make([][]int, N+1)
	for i := range graph {
		graph[i] = make([]int, N+1)
	}
	for i := 0; i < E; i++ {
		a, b := getInts()
		graph[a][b] = 1
		graph[b][a] = 1
	}
	fmt.Fprintf(w, "%d", bfs(1))
}

func bfs(n int) (res int) {
	q := []int{n}
	visited[n] = true
	for len(q) != 0 {
		n = q[0]
		q = q[1:]
		for i := 0; i < len(graph[n]); i++ {
			if graph[n][i] == 1 && !visited[i] {
				visited[i] = true
				q = append(q, i)
				res++
			}
		}
	}
	return res
}

func getInts() (a int, b int) {
	s, _ := r.ReadString('\n')
	t := strings.Fields(s)
	a, _ = strconv.Atoi(t[0])
	b, _ = strconv.Atoi(t[1])
	return a, b
}
