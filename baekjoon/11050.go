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
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	split := strings.Split(s, " ")
	n, _ := strconv.Atoi(split[0])
	k, _ := strconv.Atoi(split[1])
	binList := make([]int, 11)
	binList[0] = 1
	for i := 1; i < 11; i++ {
		binList[i] = binList[i-1] * i
	}
	fmt.Fprintf(w, "%d\n", binList[n]/(binList[n-k]*binList[k]))
}
