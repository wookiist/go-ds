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
	var N, B, C, result int
	fmt.Fscanf(r, "%d\n", &N)
	A := make([]int, N)
	s, _ := r.ReadString('\n')
	t := strings.Fields(s)
	for i, v := range t {
		A[i], _ = strconv.Atoi(v)
	}
	fmt.Fscanf(r, "%d %d\n", &B, &C)
	for _, v := range A {
		if v-B < 0 {
			continue
		}
		if (v-B)%C == 0 {
			result += (v - B) / C
		} else {
			result += (v-B)/C + 1
		}
	}
	fmt.Fprintln(w, result+N)
}
