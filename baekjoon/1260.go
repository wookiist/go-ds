// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var (
// 	w = bufio.NewWriter(os.Stdout)
// 	r = bufio.NewReader(os.Stdin)
// )

// func main() {
// 	defer w.Flush()
// 	s, _ := r.ReadString('\n')
// 	t := strings.Fields(s)
// 	N, _ := strconv.Atoi(t[0])
// 	M, _ := strconv.Atoi(t[1])
// 	V, _ := strconv.Atoi(t[2])
// 	graphSlice := make([][]int, N+1)
// 	for i := range graphSlice {
// 		graphSlice[i] = make([]int, N+1)
// 	}
// 	for i := 0; i < M; i++ {
// 		s, _ = r.ReadString('\n')
// 		t = strings.Fields(s)
// 		a, _ := strconv.Atoi(t[0])
// 		b, _ := strconv.Atoi(t[1])
// 		graphSlice[a][b] = 1
// 		graphSlice[b][a] = 1
// 	}
// 	visited := make([]bool, N+1)
// 	dfs(V, &visited, &graphSlice)
// 	visited = make([]bool, N+1)
// 	fmt.Fprintln(w)
// 	bfs(V, &visited, &graphSlice)
// }

// func dfs(v int, visited *[]bool, graphSlice *[][]int) {
// 	(*visited)[v] = true
// 	fmt.Fprintf(w, "%d ", v)
// 	for i := 0; i < len((*graphSlice)[v]); i++ {
// 		if (*graphSlice)[v][i] == 1 && !(*visited)[i] {
// 			dfs(i, visited, graphSlice)
// 		}
// 	}
// }

// func bfs(v int, visited *[]bool, graphSlice *[][]int) {
// 	q := []int{v}
// 	(*visited)[v] = true
// 	for len(q) != 0 {
// 		n := q[0]
// 		q = q[1:]
// 		fmt.Fprintf(w, "%d ", n)
// 		for i := 0; i < len((*graphSlice)[n]); i++ {
// 			if (*graphSlice)[n][i] == 1 && !(*visited)[i] {
// 				q = append(q, i)
// 				(*visited)[i] = true
// 			}
// 		}
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w          = bufio.NewWriter(os.Stdout)
	r          = bufio.NewReader(os.Stdin)
	graphSlice [][]int
	visited    []bool
)

func main() {
	defer w.Flush()
	s, _ := r.ReadString('\n')
	t := strings.Fields(s)
	N, _ := strconv.Atoi(t[0])
	M, _ := strconv.Atoi(t[1])
	V, _ := strconv.Atoi(t[2])
	graphSlice = make([][]int, N+1)
	for i := range graphSlice {
		graphSlice[i] = make([]int, N+1)
	}
	for i := 0; i < M; i++ {
		s, _ = r.ReadString('\n')
		t = strings.Fields(s)
		a, _ := strconv.Atoi(t[0])
		b, _ := strconv.Atoi(t[1])
		graphSlice[a][b] = 1
		graphSlice[b][a] = 1
	}
	visited = make([]bool, N+1)
	dfs(V)
	visited = make([]bool, N+1)
	fmt.Fprintln(w)
	bfs(V)
}

func dfs(v int) {
	visited[v] = true
	fmt.Fprintf(w, "%d ", v)
	for i := 0; i < len(graphSlice[v]); i++ {
		if graphSlice[v][i] == 1 && !visited[i] {
			dfs(i)
		}
	}
}

func bfs(v int) {
	q := []int{v}
	visited[v] = true
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		fmt.Fprintf(w, "%d ", n)
		for i := 0; i < len(graphSlice[n]); i++ {
			if graphSlice[n][i] == 1 && !visited[i] {
				q = append(q, i)
				visited[i] = true
			}
		}
	}
}
