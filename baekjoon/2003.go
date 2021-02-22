package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var N, M int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	arr := make([]int, N)
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	for i := range t {
		arr[i], _ = strconv.Atoi(t[i])
	}
	startPointer, endPointer := 0, 0
	partialSum := 0
	result := 0
	for {
		if partialSum < M {
			partialSum += arr[endPointer]
			endPointer++
		} else if partialSum >= M {
			partialSum -= arr[startPointer]
			startPointer++
		}
		if partialSum == M {
			result++
		}
		if (partialSum < M && endPointer == N) || (partialSum > M && startPointer == endPointer) {
			break
		}
	}
	fmt.Fprintln(w, result)
}
