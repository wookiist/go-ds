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
	givStrArr := getString()
	expStrArr := getString()
	lastExpStr := expStrArr[len(expStrArr)-1]
	stack := make([]string, 0)
	for i := range givStrArr {
		isExplode := true
		stack = append(stack, givStrArr[i])
		if givStrArr[i] == lastExpStr && len(stack) >= len(expStrArr) {
			for j := 0; j < len(expStrArr); j++ {
				if expStrArr[len(expStrArr)-j-1] != stack[len(stack)-j-1] {
					isExplode = false
					break
				}
			}
			if isExplode {
				for j := 0; j < len(expStrArr); j++ {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	if len(stack) == 0 {
		fmt.Fprintln(w, "FRULA")
	} else {
		for i := 0; i < len(stack); i++ {
			fmt.Fprintf(w, "%s", stack[i])
		}
	}
}

func getString() []string {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Split(s, "")
	return t
}
