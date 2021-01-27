package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	sum := func(n int) int {
		result := 0
		for i := 1; i <= n; i++ {
			result += i
		}
		return result
	}
	var T, result int
	defer writer.Flush()
	fmt.Fscanln(reader, &T)
	for i := 0; i < T; i++ {
		result = 0
		var tc string
		fmt.Fscanln(reader, &tc)
		splitWord := strings.Split(tc, "X")
		for _, word := range splitWord {
			result += sum(len(word))
		}
		fmt.Println(result)
	}
}
