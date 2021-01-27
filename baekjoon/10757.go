package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	sum := big.NewInt(0)
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	split := strings.Split(s, " ")
	x, y := &big.Int{}, &big.Int{}
	x, _ = x.SetString(split[0], 10)
	y, _ = y.SetString(split[1], 10)
	sum = sum.Add(sum, x)
	sum = sum.Add(sum, y)
	fmt.Fprintln(w, sum)
}
