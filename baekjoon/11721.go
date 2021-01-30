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
	for i := range s {
		if i != 0 && i%10 == 0 {
			fmt.Fprintln(w)
		}
		fmt.Fprintf(w, "%s", string(s[i]))
	}
}
