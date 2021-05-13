package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	w = bufio.NewWriter(os.Stdout)
	// r  = bufio.NewReader(os.Stdin)
	sc = bufio.NewScanner(os.Stdin)
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	// var T, N, M int
	// fmt.Fscanf(r, "%d\n", &T)
	T := scan()
	for ; T > 0; T-- {
		// fmt.Fscanf(r, "%d %d\n", &N, &M)
		N := scan()
		M := scan()
		U := 2*M - N
		fmt.Fprintln(w, U, M-U)
	}
}

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
