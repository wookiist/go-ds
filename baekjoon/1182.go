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
	var N, S int
	fmt.Fscanf(r, "%d %d\n", &N, &S)
	arr := getInts()
	res := solve(0, 0, S, arr)
	if S == 0 {
		res--
	}
	fmt.Fprintln(w, res)
}

func getInts() []int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	res := make([]int, len(t))
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}

func solve(idx, sum, goal int, arr []int) int {
	if idx == len(arr) {
		if sum == goal {
			return 1
		}
		return 0
	}
	return solve(idx+1, sum+arr[idx], goal, arr) + solve(idx+1, sum, goal, arr)
}
