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
	N, M  int
	check [][]bool
	G     [][]string
	dx    = [4]int{1, -1, 0, 0}
	dy    = [4]int{0, 0, 1, -1}
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N, M = scanInt(), scanInt()
	G = make([][]string, N+1)
	G[0] = make([]string, M+1)
	check = make([][]bool, N+1)
	for i := range G {
		if i != 0 {
			G[i] = scanGraph()
		}
		check[i] = make([]bool, M+1)
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if !check[i][j] {
				if solve(i, j, -1, -1, G[i][j]) {
					fmt.Fprintln(w, "Yes")
					return
				}
			}
		}
	}
	fmt.Fprintln(w, "No")
}

func solve(nx, ny, px, py int, color string) bool {
	if check[nx][ny] {
		return true
	}
	check[nx][ny] = true
	for k := range dx {
		nnx, nny := nx+dx[k], ny+dy[k]
		if 1 <= nnx && nnx <= N && 1 <= nny && nny <= M {
			if !(nnx == px && nny == py) && G[nnx][nny] == color {
				if solve(nnx, nny, nx, ny, color) {
					return true
				}
			}
		}
	}
	return false
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
