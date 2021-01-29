package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	alpha, _ := r.ReadString('\n')
	alpha = strings.TrimSuffix(alpha, "\n")
	var result int
	croArray := [8]string{"dz=", "c=", "s=", "z=", "c-", "d-", "lj", "nj"}
	for _, v := range croArray {
		r, _ := regexp.Compile(v)
		if l := len(r.FindAllString(alpha, -1)); l > 0 {
			result += l
			alpha = r.ReplaceAllString(alpha, " ")
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		r, _ := regexp.Compile(string(i))
		if l := len(r.FindAllString(alpha, -1)); l > 0 {
			result += l
			alpha = r.ReplaceAllString(alpha, " ")
		}
	}
	fmt.Fprintln(w, result)
}
