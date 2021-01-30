package main

import (
	"bufio"
	"fmt"
	"os"
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
	var res int
	for _, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			res++
		}
	}
	fmt.Fprintf(w, "%d", res)
}
