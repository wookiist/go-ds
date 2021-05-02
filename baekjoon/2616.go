package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w          = bufio.NewWriter(os.Stdout)
	r          = bufio.NewReader(os.Stdin)
	sumArr     []int
	guestsArr  []int
	D          [][]int
	maxTrain   int
	numOfTrain int
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d\n", &numOfTrain)
	guestsArr = getInts()
	sumArr = make([]int, numOfTrain+1)
	for i := 1; i <= numOfTrain; i++ {
		sumArr[i] = sumArr[i-1] + guestsArr[i]
	}
	fmt.Fscanf(r, "%d\n", &maxTrain)
	D = make([][]int, 3+1)
	for i := range D {
		D[i] = make([]int, numOfTrain+1)
	}
	solve()
	fmt.Fprintln(w, D[3][numOfTrain])
}

func solve() {
	for i := 1; i <= 3; i++ {
		for j := i * maxTrain; j <= numOfTrain; j++ {
			D[i][j] = max(D[i][j-1], D[i-1][j-maxTrain]+(sumArr[j]-sumArr[j-maxTrain]))
		}
	}
}

func getInts() []int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	res := make([]int, len(t))
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	res = append([]int{0}, res...)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
