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
	var N, K int
	fmt.Fscanf(r, "%d %d\n", &N, &K)
	arr := getInts()
	sort.Ints(arr)
	fmt.Fprintln(w, arr[K-1])
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
