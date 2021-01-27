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
	var a, b, c, d int
	fmt.Fscanf(r, "%d\n", &a)
	fmt.Fscanf(r, "%d\n", &b)
	fmt.Fscanf(r, "%d\n", &c)
	fmt.Fscanf(r, "%d\n", &d)
	fmt.Fprintf(w, "%d\n%d", (a+b+c+d)/60, (a+b+c+d)%60)
}
