package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	N     int
	G     [][]string
	color [][]int
	ans   int
	dx    = [6]int{-1, 0, 1, 1, 0, -1}
	dy    = [6]int{0, -1, -1, 0, 1, 1}
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N = scanInt()
	G = make([][]string, N+1)
	G[0] = make([]string, N+1)
	color = make([][]int, N+1)
	for i := range G {
		if i != 0 {
			G[i] = scanGraph()
		}
		color[i] = make([]int, N+1)
		for j := range color[i] {
			color[i][j] = -1
		}
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if color[i][j] == -1 && G[i][j] == "X" {
				dfs(i, j, 0)
			}
		}
	}

	fmt.Fprintln(w, ans)
}

func dfs(x, y, c int) {
	color[x][y] = c
	ans = max(ans, 1)
	for k := range dx {
		nx, ny := x+dx[k], y+dy[k]
		if 1 <= nx && nx <= N && 1 <= ny && ny <= N {
			if G[nx][ny] == "X" {
				if color[nx][ny] == -1 {
					dfs(nx, ny, 1-c)
				}
				ans = max(ans, 2)
				if color[nx][ny] == c {
					ans = max(ans, 3)
				}
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func scanGraph() []string {
	sc.Scan()
	s := strings.Split(" "+sc.Text(), "")
	return s
}
