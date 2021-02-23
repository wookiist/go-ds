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
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	arr := makeSliceUnique(getInts())
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := range arr {
		fmt.Fprintf(w, "%d ", arr[i])
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

func makeSliceUnique(arr []int) []int {
	res := make([]int, 0, len(arr))
	d := make(map[int]struct{})
	for i := range arr {
		if _, ok := d[arr[i]]; ok {
			continue
		}
		d[arr[i]] = struct{}{}
		res = append(res, arr[i])
	}
	return res
}
