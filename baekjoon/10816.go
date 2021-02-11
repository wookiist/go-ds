package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w    = bufio.NewWriter(os.Stdout)
	r    = bufio.NewReader(os.Stdin)
	nMap map[string]int
)

func main() {
	defer w.Flush()
	var N, M int
	nMap = make(map[string]int)
	fmt.Fscanf(r, "%d\n", &N)
	getArrMap(N)
	fmt.Fscanf(r, "%d\n", &M)
	mCards := getArr(M)
	for i := range mCards {
		fmt.Fprintf(w, "%d ", nMap[mCards[i]])
	}
}

func getArr(n int) []string {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	return t
}

func getArrMap(n int) {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	for i := range t {
		nMap[t[i]]++
	}
}

func getInts(n int) []int {
	arr := make([]int, n)
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	for i := range t {
		arr[i], _ = strconv.Atoi(t[i])
	}
	return arr
}
