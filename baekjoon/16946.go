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
	N, M   int
	answer [][]int
	D      [][]int
	G      [][]int
	group  map[int]int
	dx     = [4]int{1, -1, 0, 0}
	dy     = [4]int{0, 0, 1, -1}
)

type pair struct {
	x, y int
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N, M = scanInt(), scanInt()
	answer = make([][]int, N+1)
	G = make([][]int, N+1)
	D = make([][]int, N+1)
	group = make(map[int]int)
	for i := range G {
		if i != 0 {
			G[i] = scanGraph()
		}
		D[i] = make([]int, M+1)
		answer[i] = make([]int, M+1)
	}

	id := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if G[i][j] == 0 && D[i][j] == 0 {
				id++
				group[id] = 0
				dfs(i, j, id)
			}
		}
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if G[i][j] == 1 {
				answer[i][j] = 1
				checkedGroup := make([]int, 0, 4)
				for k := range dx {
					nx, ny := i+dx[k], j+dy[k]
					if 1 <= nx && nx <= N && 1 <= ny && ny <= M {
						groupID := D[nx][ny]
						if isCheckedGroup(groupID, checkedGroup) {
							continue
						}
						checkedGroup = append(checkedGroup, groupID)
						answer[i][j] += group[groupID]
					}
				}
			}
		}
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			fmt.Fprintf(w, "%d", answer[i][j]%10)
		}
		fmt.Fprintln(w)
	}
}

func isCheckedGroup(target int, group []int) bool {
	for i := range group {
		if target == group[i] {
			return true
		}
	}
	return false
}

func dfs(x, y, id int) {
	D[x][y] = id
	group[id]++
	for k := range dx {
		nx, ny := x+dx[k], y+dy[k]
		if 1 <= nx && nx <= N && 1 <= ny && ny <= M {
			if G[nx][ny] == 0 && D[nx][ny] == 0 {
				dfs(nx, ny, id)
			}
		}
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func scanGraph() []int {
	sc.Scan()
	s := strings.Split("0"+sc.Text(), "")
	res := make([]int, len(s))
	for i := range s {
		res[i], _ = strconv.Atoi(s[i])
	}
	return res
}
