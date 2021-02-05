// distance 1차원 배열로 풀기 - 메모리 사용량만 조금 줄고, 시간은 4ms로 같음.
// 2021-02-05-Fri
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
	var n, a, b, m int
	fmt.Fscanf(r, "%d\n", &n)
	fmt.Fscanf(r, "%d %d\n", &a, &b)
	fmt.Fscanf(r, "%d\n", &m)
	graph = make([][]int, n+1)
	visited = make([]bool, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscanf(r, "%d %d\n", &x, &y)
		graph[x][y] = 1
		graph[y][x] = 1
	}
	fmt.Fprintf(w, "%d\n", bfs(a, b, n))
}

func bfs(a, b, n int) int {
	distance := make([]int, n+1)
	q := []int{a}
	visited[a] = true
	for len(q) != 0 {
		p := q[0]
		q = q[1:]
		for i := 0; i <= n; i++ {
			if graph[p][i] == 1 && !visited[i] {
				q = append(q, i)
				visited[i] = true
				distance[i] = distance[p] + 1
			}
		}
	}
	if visited[b] {
		return distance[b]
	}
	return -1
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var (
// 	w        = bufio.NewWriter(os.Stdout)
// 	r        = bufio.NewReader(os.Stdin)
// 	graph    [][]int
// 	distance [][]int
// 	visited  []bool
// )

// func main() {
// 	defer w.Flush()
// 	var n, m int
// 	fmt.Fscanf(r, "%d\n", &n)
// 	p1, p2 := getInts()
// 	fmt.Fscanf(r, "%d\n", &m)
// 	graph = make([][]int, n+1)
// 	distance = make([][]int, n+1)
// 	for i := range graph {
// 		graph[i] = make([]int, n+1)
// 		distance[i] = make([]int, n+1)
// 	}
// 	for i := 0; i < m; i++ {
// 		x, y := getInts()
// 		graph[x][y] = 1
// 		graph[y][x] = 1
// 	}
// 	visited = make([]bool, n+1)
// 	fmt.Fprintf(w, "%d\n", bfs(p1, p2))
// }

// func getInts() (int, int) {
// 	s, _ := r.ReadString('\n')
// 	t := strings.Fields(s)
// 	res1, _ := strconv.Atoi(t[0])
// 	res2, _ := strconv.Atoi(t[1])
// 	return res1, res2
// }

// func bfs(p1, p2 int) int {
// 	q := []int{p1}
// 	visited[p1] = true
// 	for len(q) != 0 {
// 		n := q[0]
// 		q = q[1:]
// 		for i := 0; i < len(graph[n]); i++ {
// 			if graph[n][i] == 1 && !visited[i] {
// 				q = append(q, i)
// 				visited[i] = true
// 				distance[p1][i] = distance[p1][n] + 1
// 			}
// 		}
// 	}
// 	if distance[p1][p2] == 0 {
// 		return -1
// 	}
// 	return distance[p1][p2]
// }
