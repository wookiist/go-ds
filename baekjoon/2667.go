package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	w       = bufio.NewWriter(os.Stdout)
	r       = bufio.NewReader(os.Stdin)
	graph   [][]int
	visited [][]bool
	dx      = [4]int{1, -1, 0, 0}
	dy      = [4]int{0, 0, 1, -1}
)

type axis struct {
	X int
	Y int
}

func main() {
	defer w.Flush()
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	graph = make([][]int, N)
	visited = make([][]bool, N)
	for i := range graph {
		graph[i] = getStrings(N)
		visited[i] = make([]bool, N)
	}
	results := make([]int, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if !visited[i][j] && graph[i][j] == 1 {
				result := bfs(i, j)
				results = append(results, result)
			}
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})
	fmt.Fprintln(w, len(results))
	for _, v := range results {
		fmt.Fprintln(w, v)
	}
}

func bfs(x, y int) int {
	result := 1
	q := []axis{{X: x, Y: y}}
	visited[x][y] = true
	for len(q) != 0 {
		nx, ny := q[0].X, q[0].Y
		q = q[1:]
		for k := range dx {
			nnx, nny := nx+dx[k], ny+dy[k]
			if 0 <= nnx && nnx < len(graph) && 0 <= nny && nny < len(graph) {
				if graph[nnx][nny] == 1 && !visited[nnx][nny] {
					q = append(q, axis{X: nnx, Y: nny})
					visited[nnx][nny] = true
					result++
				}
			}
		}
	}
	return result
}

func getStrings(n int) []int {
	arr := make([]int, n)
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Split(s, "")
	for i, v := range t {
		j, _ := strconv.Atoi(v)
		arr[i] = j
	}
	return arr
}
