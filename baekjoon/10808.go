package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	for i := 'a'; i <= 'z'; i++ {
		r, _ := regexp.Compile(string(i))
		fmt.Fprintf(w, "%d ", len(r.FindAllString(s, -1)))
	}
}
