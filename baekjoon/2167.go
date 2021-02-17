package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

type pair struct {
	x int
	y int
}

func main() {
	defer w.Flush()
	var N, M, K int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	arr := make([][]int, N)
	for i := range arr {
		arr[i] = getInts()
	}
	d := make([][]int, N+1)
	d[0] = make([]int, M+1)
	for i := 1; i <= N; i++ {
		d[i] = make([]int, M+1)
		for j := 1; j <= M; j++ {
			d[i][j] = d[i-1][j] + d[i][j-1] - d[i-1][j-1] + arr[i-1][j-1]
		}
	}
	fmt.Fscanf(r, "%d\n", &K)
	for i := 0; i < K; i++ {
		temp := getInts()
		x1, y1, x2, y2 := temp[0], temp[1], temp[2], temp[3]
		sum := d[x2][y2] - d[x1-1][y2] - d[x2][y1-1] + d[x1-1][y1-1]
		fmt.Fprintln(w, sum)
	}
}

func getInts() []int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Split(s, " ")
	res := make([]int, len(t))
	for i := range res {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}
