package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	fmt.Fprintf(w, "%d\n", n*(n+1)/2)
}
