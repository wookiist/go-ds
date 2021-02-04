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

func main() {
	defer w.Flush()
	var N, M, i, j, result, sumValue int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	arr := getArray(N)
	sumArr := make([]int, N+1)
	for k := 1; k <= N; k++ {
		sumValue += arr[k-1]
		sumArr[k] = sumValue
	}
	for k := 0; k < M; k++ {
		fmt.Fscanf(r, "%d %d\n", &i, &j)
		result = sumArr[j] - sumArr[i-1]
		fmt.Fprintln(w, result)
	}
}

func getArray(n int) []int {
	arr := make([]int, n)
	s, _ := r.ReadString('\n')
	t := strings.Fields(s)
	for i := range arr {
		arr[i], _ = strconv.Atoi(t[i])
	}
	return arr
}
