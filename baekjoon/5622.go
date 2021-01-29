package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	table := []int{
		2, 2, 2, 3, 3, 3, 4, 4, 4,
		5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9}
	s := getString()
	var result int
	for _, v := range s {
		result += table[v-'A'] + 1
	}
	fmt.Fprintf(w, "%d\n", result)
}

func getString() string {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	return s
}
