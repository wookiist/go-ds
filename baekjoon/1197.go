package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

type edge struct {
	u int
	v int
	w int
}

var (
	parent   map[int]int
	rank     map[int]int
	vertices []int
	edges    []edge
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	V, E := scanInt(), scanInt()
	vertices = make([]int, V+1)
	for i := 1; i <= V; i++ {
		vertices[i] = i
	}
	edges = make([]edge, E)
	for i := range edges {
		edges[i] = edge{
			u: scanInt(),
			v: scanInt(),
			w: scanInt(),
		}
	}
	parent = make(map[int]int)
	rank = make(map[int]int)
	mst := kruskal()
	ans := 0
	for i := range mst {
		ans += mst[i].w
	}
	fmt.Fprintln(w, ans)
}

func makeSet(node int) {
	parent[node] = node
	rank[node] = 0
}

func find(node int) int {
	// path compression 기법
	if parent[node] != node {
		parent[node] = find(parent[node])
	}
	return parent[node]
}

func union(nodeU, nodeV int) {
	rootU := find(nodeU)
	rootV := find(nodeV)

	// union-by-rank 기법
	if rank[rootU] > rank[rootV] {
		parent[rootV] = rootU
	} else {
		parent[rootU] = rootV
		if rank[rootU] == rank[rootV] {
			rank[rootV]++
		}
	}
}

func kruskal() []edge {
	mst := make([]edge, 0, 10001)

	// 초기화
	for i := 1; i < len(vertices); i++ {
		makeSet(vertices[i])
	}

	// 간선 가중치 기반으로 오름차순 정렬
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	// 사이클 없이 간선 연결
	for i := range edges {
		nodeU, nodeV := edges[i].u, edges[i].v
		if find(nodeU) != find(nodeV) {
			union(nodeU, nodeV)
			mst = append(mst, edges[i])
		}
	}
	return mst
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func scanInts(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = scanInt()
	}
	return res
}
