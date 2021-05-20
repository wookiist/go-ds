package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	defer w.Flush()
	for {
		s := scan()
		if s == "END" {
			break
		}
		for i := len(s) - 1; i >= 0; i-- {
			fmt.Fprintf(w, "%s", string(s[i]))
		}
		fmt.Fprintln(w)
	}
}

func scan() string {
	sc.Scan()
	return sc.Text()
}
