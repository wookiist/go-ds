package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var N, M int
	fmt.Fscanf(r, "%d\n", &N)
	sk := makeMap(N)
	fmt.Fscanf(r, "%d\n", &M)
	res := getString()
	for i := range res {
		if _, ok := (*sk)[res[i]]; ok {
			fmt.Fprintf(w, "%d ", 1)
		} else {
			fmt.Fprintf(w, "%d ", 0)
		}
	}
}

func makeMap(N int) *map[string]struct{} {
	sk := make(map[string]struct{}, N)
	t := getString()
	for i := range t {
		sk[t[i]] = struct{}{}
	}
	return &sk
}

func getString() []string {
	s, _ := r.ReadString('\n')
	t := strings.Fields(s)
	return t
}
