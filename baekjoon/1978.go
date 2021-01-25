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

	checkArr := make([]bool, 1001)
	pArr := make([]bool, 1001)
	for i := 2; i < 1001; i++ {
		if !checkArr[i] {
			pArr[i] = true
			for j := i + i; j < 1001; j += i {
				checkArr[j] = true
			}
		}
	}
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	split := strings.Split(s, " ")
	res := 0
	for _, v := range split {
		idx, _ := strconv.Atoi(v)
		if pArr[idx] {
			res++
		}
	}
	fmt.Fprintf(w, "%d\n", res)
}
