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

type pair struct {
	x, y int
}

var (
	N, M           int
	G              [][]int
	D              [][]int
	r1, c1, r2, c2 int
	dx             = [4]int{1, -1, 0, 0}
	dy             = [4]int{0, 0, 1, -1}
	virus          []pair
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N, M = scanInt(), scanInt()
	G = make([][]int, N+1)
	D = make([][]int, N+1)
	virus = make([]pair, 0)
	for i := range G {
		G[i] = make([]int, M+1)
		D[i] = make([]int, M+1)
		for j := range G[i] {
			D[i][j] = -1
			if !(i == 0 || j == 0) {
				G[i][j] = scanInt()
				if G[i][j] == 1 {
					D[i][j] = 1
				}
				if G[i][j] == 2 {
					virus = append(virus, pair{i, j})
				}
			}
		}
	}

	answer := 0
	for w1x := 1; w1x <= N; w1x++ {
		for w1y := 1; w1y <= M; w1y++ {
			if G[w1x][w1y] == 1 || G[w1x][w1y] == 2 {
				continue
			}
			G[w1x][w1y] = 1
			for w2x := 1; w2x <= N; w2x++ {
				for w2y := 1; w2y <= M; w2y++ {
					if w1x == w2x && w1y == w2y {
						continue
					}
					if G[w2x][w2y] == 1 || G[w2x][w2y] == 2 {
						continue
					}
					G[w2x][w2y] = 1
					for w3x := 1; w3x <= N; w3x++ {
						for w3y := 1; w3y <= M; w3y++ {
							if (w1x == w3x && w1y == w3y) || (w2x == w3x && w2y == w3y) {
								continue
							}
							if G[w3x][w3y] == 1 || G[w3x][w3y] == 2 {
								continue
							}
							G[w3x][w3y] = 1
							q := make([]pair, 0)
							for i := range virus {
								q = append(q, virus[i])
								D[virus[i].x][virus[i].y] = 2
							}
							for len(q) != 0 {
								nx, ny := q[0].x, q[0].y
								q = q[1:]
								for k := range dx {
									nnx, nny := nx+dx[k], ny+dy[k]
									if 1 <= nnx && nnx <= N && 1 <= nny && nny <= M {
										if D[nnx][nny] == -1 && G[nnx][nny] == 0 {
											q = append(q, pair{nnx, nny})
											D[nnx][nny] = 2
										}
									}
								}
							}
							answer = max(answer, checkArea())
							D = make([][]int, N+1)
							for i := range D {
								D[i] = make([]int, M+1)
								for j := range D[i] {
									D[i][j] = -1
								}
							}
							// finish up
							G[w3x][w3y] = 0
						}
					}
					G[w2x][w2y] = 0
				}
			}
			G[w1x][w1y] = 0
		}
	}
	fmt.Fprintln(w, answer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkArea() int {
	res := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if G[i][j] != 1 && D[i][j] == -1 {
				res++
			}
		}
	}
	return res
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
