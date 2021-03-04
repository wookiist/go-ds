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
	var T int
	fmt.Fscanf(r, "%d\n", &T)
	for T > 0 {
		arr := getInts()
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
		fmt.Fprintln(w, arr[2])
		T--
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
	return res
}
