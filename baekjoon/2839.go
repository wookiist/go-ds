package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	var threeBag, fiveBag int
	fiveBag = N / 5
	N %= 5
	for fiveBag >= 0 {
		if N%3 == 0 {
			threeBag = N / 3
			N %= 3
			break
		}
		fiveBag--
		N += 5
	}
	if N == 0 {
		fmt.Fprintf(w, "%d\n", fiveBag+threeBag)
	} else {
		fmt.Fprintf(w, "%d\n", -1)
	}
}
