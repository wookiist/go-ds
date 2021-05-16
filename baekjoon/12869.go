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
	H []int
	D [][][]int // D[i][j][k] = (i, j, k)의 체력이 남은 scv를 전멸시키는데 필요한 공격 횟수의 최솟값
	A = [4]int{0, 9, 3, 1}
	S = [7][4]int{
		{0, 0, 0, 0},
		{0, 1, 2, 3},
		{0, 1, 3, 2},
		{0, 2, 1, 3},
		{0, 2, 3, 1},
		{0, 3, 1, 2},
		{0, 3, 2, 1},
	}
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N = scanInt()
	H = make([]int, 4)
	for i := 1; i <= N; i++ {
		H[i] = scanInt()
	}
	D = make([][][]int, H[1]+1)
	for i := range D {
		D[i] = make([][]int, H[2]+1)
		for j := range D[i] {
			D[i][j] = make([]int, H[3]+1)
			for k := range D[i][j] {
				D[i][j][k] = -1
			}
		}
	}
	fmt.Fprintln(w, solve(H[1], H[2], H[3]))
}

func solve(scv1, scv2, scv3 int) int {
	if scv1 <= 0 && scv2 <= 0 && scv3 <= 0 {
		return 0
	}
	if scv1 <= 0 {
		scv1 = 0
	}
	if scv2 <= 0 {
		scv2 = 0
	}
	if scv3 <= 0 {
		scv3 = 0
	}
	if D[scv1][scv2][scv3] != -1 {
		return D[scv1][scv2][scv3]
	}
	temp := 999999
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 3; j++ {
			ret := solve(scv1-A[S[i][1]], scv2-A[S[i][2]], scv3-A[S[i][3]])
			if temp > ret {
				temp = ret
			}
		}
	}
	D[scv1][scv2][scv3] = temp + 1
	return D[scv1][scv2][scv3]
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
