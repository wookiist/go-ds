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

// Queue type
type Queue []int

func (q *Queue) push(x int) {
	*q = append(*q, x)
}

func (q *Queue) pop() int {
	if q.empty() == 1 {
		return -1
	}
	p := (*q)[0]
	(*q) = (*q)[1:]
	return p
}

func (q *Queue) size() int {
	return len(*q)
}

func (q *Queue) empty() int {
	if len(*q) == 0 {
		return 1
	}
	return 0
}

func (q *Queue) front() int {
	if q.empty() == 1 {
		return -1
	}
	return (*q)[0]
}

func (q *Queue) back() int {
	if q.empty() == 1 {
		return -1
	}
	return (*q)[q.size()-1]
}

func main() {
	defer w.Flush()
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	var q Queue
	for i := 0; i < N; i++ {
		s, _ := r.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		split := strings.Split(s, " ")

		switch split[0] {
		case "push":
			x, _ := strconv.Atoi(split[1])
			q.push(x)
		case "front":
			fmt.Fprintln(w, q.front())
		case "back":
			fmt.Fprintln(w, q.back())
		case "size":
			fmt.Fprintln(w, q.size())
		case "pop":
			fmt.Fprintln(w, q.pop())
		case "empty":
			fmt.Fprintln(w, q.empty())
		}
	}
}
