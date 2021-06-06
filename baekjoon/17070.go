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
	N int
	G [][]int
	D [][][]int
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N = scanInt()
	G = make([][]int, N+1)
	D = make([][][]int, N+1)
	for i := range G {
		G[i] = make([]int, N+1)
		D[i] = make([][]int, N+1)
		for j := range D[i] {
			D[i][j] = make([]int, 3)
		}
		if i != 0 {
			for j := range G[i] {
				if j != 0 {
					G[i][j] = scanInt()
				}
			}
		}
	}

	if G[N][N] == 1 {
		fmt.Fprintln(w, 0)
		return
	}

	// D[i][j] = 끝자락이 (i, j)에 놓이도록 이동하는 방법의 수
	// k == 0 -> 가로 , k == 1 -> 세로, k == 2 -> 대각선
	D[1][2][0] = 1
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			// 가로로 i, j에 놓여있다는 것이고. 가로로 있다는 소리는
			// 1) 가로->가로 2) 대각선->가로
			if G[i][j-1] != 1 {
				D[i][j][0] += D[i][j-1][0] + D[i][j-1][2]
			}
			// 세로로 i, j에 놓여있다는 것이고, 세로로 있다는 소리는
			// 1) 세로->세로 2) 대각선->세로
			if G[i-1][j] != 1 {
				D[i][j][1] += D[i-1][j][1] + D[i-1][j][2]
			}
			// 대각선으로 i,j에 놓여있다는 것이고, 대각선으로 있다는 소리는
			// 1) 가로->대각선 2) 세로->대각선 3) 대각선->대각선
			if G[i-1][j-1] != 1 && G[i][j-1] != 1 && G[i-1][j] != 1 {
				D[i][j][2] += (D[i-1][j-1][0] + D[i-1][j-1][1] + D[i-1][j-1][2])
			}
		}
	}
	fmt.Fprintln(w, sum(D[N][N]))
}

func sum(arr []int) int {
	res := 0
	for i := range arr {
		res += arr[i]
	}
	return res
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
