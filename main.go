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
	bubblesort.BubbleSort(&arr)
	fmt.Fprintln(w, arr)
}
