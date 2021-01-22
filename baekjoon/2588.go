package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()
	var sa, sb string
	fmt.Fscanln(r, &sa)
	fmt.Fscanln(r, &sb)
	a, _ := strconv.Atoi(sa)
	bs := strings.Split(sb, "")
	for i := 2; i >= 0; i-- {
		b, _ := strconv.Atoi(bs[i])
		fmt.Fprintln(w, a*b)
	}
	b, _ := strconv.Atoi(sb)
	fmt.Fprintln(w, a*b)
}
