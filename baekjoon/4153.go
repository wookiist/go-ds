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
	for {
		res := getInts()
		if res[0] == 0 && res[1] == 0 && res[2] == 0 {
			break
		}
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		if res[2]*res[2] == res[1]*res[1]+res[0]*res[0] {
			fmt.Fprintln(w, "right")
		} else {
			fmt.Fprintln(w, "wrong")
		}
	}
}

func getInts() []int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	res := make([]int, 3)
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}
