package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	w     = bufio.NewWriter(os.Stdout)
	r     = bufio.NewReader(os.Stdin)
	vowel = []string{"a", "e", "i", "o", "u"}
)

func main() {
	defer w.Flush()
	var L, C int
	fmt.Fscanf(r, "%d %d\n", &L, &C)
	arr := getStr()
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	solve(0, L, "", arr)
}

func getStr() []string {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	return t
}

func check(s string) bool {
	var v int
	for i := range vowel {
		if strings.Contains(s, vowel[i]) {
			v++
		}
	}
	if v < 1 || len(s)-v < 2 {
		return false
	}
	return true
}

func solve(idx, goal int, str string, arr []string) {
	if len(str) == goal {
		if check(str) {
			fmt.Fprintln(w, str)
		}
		return
	}
	if idx >= len(arr) {
		return
	}
	solve(idx+1, goal, str+arr[idx], arr)
	solve(idx+1, goal, str, arr)
}
