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
	fmt.Fprintln(w, getInts())
}

func getInts() int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	res := make([]int, 3)
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	sort.Ints(res)
	return res[1]
}
