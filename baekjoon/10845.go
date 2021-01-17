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

// Pop from Queue
func (q *Queue) Pop() int {
	if len(*q) != 0 {
		data := (*q)[0]
		(*q) = (*q)[1:]
		return data
	}
	return -1
}

// Push to Queue
func (q *Queue) Push(k int) {
	(*q) = append((*q), k)
}

// Empty Queue
func (q *Queue) Empty() int {
	if len(*q) == 0 {
		return 1
	}
	return 0
}

// Size Queue
func (q *Queue) Size() int {
	return len(*q)
}

// Front Queue
func (q *Queue) Front() int {
	if q.Empty() == 1 {
		return -1
	}
	return (*q)[0]
}

// Back Queue
func (q *Queue) Back() int {
	if q.Empty() == 1 {
		return -1
	}
	return (*q)[q.Size()-1]
}

func main() {
	defer w.Flush()
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	var queue Queue
	for i := 0; i < n; i++ {
		text, _ := r.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		split := strings.Split(text, " ")
		switch split[0] {
		case "push":
			item, _ := strconv.Atoi(split[1])
			queue.Push(item)
		case "front":
			fmt.Fprintf(w, "%d\n", queue.Front())
		case "pop":
			fmt.Fprintf(w, "%d\n", queue.Pop())
		case "empty":
			fmt.Fprintf(w, "%d\n", queue.Empty())
		case "size":
			fmt.Fprintf(w, "%d\n", queue.Size())
		case "back":
			fmt.Fprintf(w, "%d\n", queue.Back())
		}
	}
}
