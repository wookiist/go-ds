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

// Deque type
type Deque []int

// PopFront from Deque
func (d *Deque) PopFront() int {
	if len(*d) != 0 {
		data := (*d)[0]
		(*d) = (*d)[1:]
		return data
	}
	return -1
}

// PopBack from Deque
func (d *Deque) PopBack() int {
	if len(*d) != 0 {
		data := (*d)[len(*d)-1]
		(*d) = (*d)[:len(*d)-1]
		return data
	}
	return -1
}

// PushBack to Deque
func (d *Deque) PushBack(k int) {
	(*d) = append((*d), k)
}

// PushFront to Deque
func (d *Deque) PushFront(k int) {
	n := &Deque{k}
	*d = append(*n, (*d)...)
}

// Empty Deque
func (d *Deque) Empty() int {
	if len(*d) == 0 {
		return 1
	}
	return 0
}

// Size Deque
func (d *Deque) Size() int {
	return len(*d)
}

// Front Deque
func (d *Deque) Front() int {
	if d.Empty() == 1 {
		return -1
	}
	return (*d)[0]
}

// Back Deque
func (d *Deque) Back() int {
	if d.Empty() == 1 {
		return -1
	}
	return (*d)[d.Size()-1]
}

func main() {
	defer w.Flush()
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	var deque Deque
	for i := 0; i < n; i++ {
		text, _ := r.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		split := strings.Split(text, " ")
		switch split[0] {
		case "push_front":
			item, _ := strconv.Atoi(split[1])
			deque.PushFront(item)
		case "push_back":
			item, _ := strconv.Atoi(split[1])
			deque.PushBack(item)
		case "front":
			fmt.Fprintf(w, "%d\n", deque.Front())
		case "pop_front":
			fmt.Fprintf(w, "%d\n", deque.PopFront())
		case "pop_back":
			fmt.Fprintf(w, "%d\n", deque.PopBack())
		case "empty":
			fmt.Fprintf(w, "%d\n", deque.Empty())
		case "size":
			fmt.Fprintf(w, "%d\n", deque.Size())
		case "back":
			fmt.Fprintf(w, "%d\n", deque.Back())
		}
	}
}
