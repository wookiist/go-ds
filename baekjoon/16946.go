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

/*
```go
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
	check  [][]bool
	answer [][]int
	G      [][]int
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
	check = make([][]bool, N+1)
	answer = make([][]int, N+1)
	G = make([][]int, N+1)
	for i := range check {
		if i != 0 {
			G[i] = scanGraph()
		}
		check[i] = make([]bool, M+1)
		answer[i] = make([]int, M+1)
	}
	for i := range G {
		for j := range G[i] {
			if G[i][j] == 1 {
				check = make([][]bool, N+1)
				for k := range check {
					check[k] = make([]bool, M+1)
				}
				q := make([]pair, 0)
				q = append(q, pair{i, j})
				check[i][j] = true
				answer[i][j] = bfs(i, j, &q)
			}
		}
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			fmt.Fprintf(w, "%d", answer[i][j])
		}
		fmt.Fprintln(w)
	}
}

func bfs(x, y int, q *[]pair) int {
	ans := 1
	for len(*q) != 0 {
		nx, ny := (*q)[0].x, (*q)[0].y
		(*q) = (*q)[1:]
		for k := range dx {
			nnx, nny := nx+dx[k], ny+dy[k]
			if 1 <= nnx && nnx <= N && 1 <= nny && nny <= M {
				if G[nnx][nny] == 0 && !check[nnx][nny] {
					(*q) = append((*q), pair{nnx, nny})
					check[nnx][nny] = true
					ans++
				}
			}
		}
	}
	return ans
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

func scanGraph() []int {
	sc.Scan()
	s := strings.Split("0"+sc.Text(), "")
	res := make([]int, len(s))
	for i := range s {
		res[i], _ = strconv.Atoi(s[i])
	}
	return res
}
```

# 두 번째 접근 코드

```go
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
	check  [][]bool
	answer [][]int
	D      [][]int
	G      [][]int
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
	check = make([][]bool, N+1)
	answer = make([][]int, N+1)
	G = make([][]int, N+1)
	D = make([][]int, N+1)
	for i := range check {
		if i != 0 {
			G[i] = scanGraph()
		}
		D[i] = make([]int, M+1)
		check[i] = make([]bool, M+1)
		answer[i] = make([]int, M+1)
	}
	// 먼저 벽이 아닌 공간을 탐색하며, 그룹별로 몇 개의 빈 방이 존재하는지 확인한다.

	for i := range G {
		for j := range G[i] {
			if G[i][j] == 0 && !check[i][j] {
				q := make([]pair, 0)
				q = append(q, pair{i, j})
				check[i][j] = true
				D[i][j] = 1
				D[i][j] = max(D[i][j], findArea(&q))
			}
		}
	}
	for i := range G {
		for j := range G[i] {
			if G[i][j] == 1 {
				answer[i][j] = 1
				for k := range dx {
					nx, ny := i+dx[k], j+dy[k]
					if 1 <= nx && nx <= N && 1 <= ny && ny <= M {
						answer[i][j] += D[nx][ny]
					}
				}
			}
		}
	}
	// fmt.Fprintln(w, D)
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			fmt.Fprintf(w, "%d", answer[i][j])
		}
		fmt.Fprintln(w)
	}
}

func findArea(q *[]pair) int {
	ans := 0
	area := make([]pair, 0)
	for len(*q) != 0 {
		nx, ny := (*q)[0].x, (*q)[0].y
		(*q) = (*q)[1:]
		for k := range dx {
			nnx, nny := nx+dx[k], ny+dy[k]
			if 1 <= nnx && nnx <= N && 1 <= nny && nny <= M {
				if G[nnx][nny] == 0 && !check[nnx][nny] {
					check[nnx][nny] = true
					(*q) = append((*q), pair{nnx, nny})
					area = append(area, pair{nnx, nny})
					D[nnx][nny] = D[nx][ny] + 1
					ans = max(ans, D[nnx][nny])
				}
			}
		}
	}
	for i := range area {
		D[area[i].x][area[i].y] = ans
	}
	return ans
}

func bfs(x, y int, q *[]pair) int {
	ans := 1
	for len(*q) != 0 {
		nx, ny := (*q)[0].x, (*q)[0].y
		(*q) = (*q)[1:]
		for k := range dx {
			nnx, nny := nx+dx[k], ny+dy[k]
			if 1 <= nnx && nnx <= N && 1 <= nny && nny <= M {
				if G[nnx][nny] == 0 && !check[nnx][nny] {
					(*q) = append((*q), pair{nnx, nny})
					check[nnx][nny] = true
					ans++
				}
			}
		}
	}
	return ans
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

func scanGraph() []int {
	sc.Scan()
	s := strings.Split("0"+sc.Text(), "")
	res := make([]int, len(s))
	for i := range s {
		res[i], _ = strconv.Atoi(s[i])
	}
	return res
}
```
*/
