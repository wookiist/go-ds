package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	defer w.Flush()
	A, B, C, D, E := scan(), scan(), scan(), scan(), scan()
	fmt.Fprintln(w, (A+B+C+D+E)/5)
}

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	if n < 40 {
		n = 40
	}
	return n
}
