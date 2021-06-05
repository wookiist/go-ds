package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	N      int
	check  []int
	check2 []bool
	D      []int
	G      [][]int
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N = scanInt()
	check = make([]int, N+1)
	check2 = make([]bool, N+1)
	D = make([]int, N+1)
	G = make([][]int, N+1)
	for i := 0; i < N; i++ {
		u, v := scanInt(), scanInt()
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	// cycle부터 찾는다.
	solve(1, -1)

	// cycle에서 시작하여 연결된 노드를 찾는다.
	for i := 1; i <= N; i++ {
		if check[i] == 2 && !check2[i] {
			q := make([]int, 0)
			q = append(q, i)
			check2[i] = true
			for len(q) != 0 {
				cur := q[0]
				q = q[1:]
				for j := range G[cur] {
					next := G[cur][j]
					if check[next] != 2 && !check2[next] {
						check2[next] = true
						D[next] = D[cur] + 1
						q = append(q, next)
					}
				}
			}
		}
	}

	for i := 1; i <= N; i++ {
		fmt.Fprintf(w, "%d ", D[i])
	}
}

func solve(node, prev int) int {
	if check[node] == 1 {
		return node
	}
	check[node] = 1
	for i := range G[node] {
		next := G[node][i]
		if next == prev {
			continue
		}
		res := solve(next, node)
		if res == -2 {
			return -2
		}
		if res >= 0 {
			check[node] = 2
			if node == res {
				return -2
			}
			return res
		}
	}
	return -1
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
