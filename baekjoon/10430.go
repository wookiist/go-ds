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
	var a, b, c int
	fmt.Fscanln(r, &a, &b, &c)
	fmt.Fprintf(w, "%d\n%d\n%d\n%d", (a+b)%c, (a%c+b%c)%c, (a*b)%c, (a%c)*(b%c)%c)
}
