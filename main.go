package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wookiist/go-ds/sort/bubblesort"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	arr := []int{23, 4, 5, 1, -1, 3, 100, 55, 66, 45, 46, 47, 44}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	bubblesort.BubbleSort(&arr)
	fmt.Fprintln(w, arr)
	bubblesort.BubbleSort(&arr2)
	fmt.Fprintln(w, arr2)
}
