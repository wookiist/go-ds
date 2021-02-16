package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var N, mid, result int
	fmt.Fscanf(r, "%d\n", &N)
	P := getInts()
	sort.Slice(P, func(i, j int) bool {
		return P[i] < P[j]
	})
	for i := range P {
		mid += P[i]
		result += mid
	}
	fmt.Fprintln(w, result)
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
