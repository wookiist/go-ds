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
	a       int
	b       int
	graph   [][]int
	visited [][]bool
	dx      = [8]int{1, -1, 0, 0, 1, -1, -1, 1}
	dy      = [8]int{0, 0, 1, -1, 1, 1, -1, -1}
)

type pair struct {
	x int
	y int
}

func main() {
	defer w.Flush()
	for {
		result := 0
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		if a == 0 && b == 0 {
			break
		}
		graph = make([][]int, b)
		visited = make([][]bool, b)
		for i := range graph {
			graph[i] = getInts()
			visited[i] = make([]bool, a)
		}
		for i := range graph {
			for j := range graph[i] {
				if graph[i][j] == 1 && !visited[i][j] {
					bfs(i, j)
					result++
				}
			}
		}
		fmt.Fprintln(w, result)
	}
}

func bfs(x, y int) {
	q := []pair{{x, y}}
	visited[x][y] = true

	for len(q) != 0 {
		nx, ny := q[0].x, q[0].y
		q = q[1:]
		for k := range dx {
			nnx, nny := nx+dx[k], ny+dy[k]
			if 0 <= nnx && nnx < b && 0 <= nny && nny < a {
				if graph[nnx][nny] == 1 && !visited[nnx][nny] {
					q = append(q, pair{nnx, nny})
					visited[nnx][nny] = true
				}
			}
		}
	}
}

func getInts() []int {
	res := make([]int, a)
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Split(s, " ")
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}
