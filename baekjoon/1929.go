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

	checkArr := make([]bool, 1000001)
	pArr := make([]bool, 1000001)
	for i := 2; i < 1000001; i++ {
		if !checkArr[i] {
			pArr[i] = true
			for j := i + i; j < 1000001; j += i {
				checkArr[j] = true
			}
		}
	}
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	split := strings.Split(s, " ")
	m, _ := strconv.Atoi(split[0])
	n, _ := strconv.Atoi(split[1])
	for i := m; i <= n; i++ {
		if pArr[i] {
			fmt.Fprintln(w, i)
		}
	}
}
