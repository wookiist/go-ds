package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	numMap := make(map[int]string)
	strMap := make(map[string]int)
	var N, M int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	for i := 1; i <= N; i++ {
		s, _ := r.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		numMap[i] = s
		strMap[s] = i
	}
	for i := 0; i < M; i++ {
		s, _ := r.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		if no, err := strconv.Atoi(s); err != nil {
			fmt.Fprintln(w, strMap[s])
		} else {
			fmt.Fprintln(w, numMap[no])
		}
	}
}
