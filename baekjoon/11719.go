package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprint(w, s)
	}
}
