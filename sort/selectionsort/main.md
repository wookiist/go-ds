# SelectionSort with Go

Selection sort는 0번부터 n-1번까지 차례로 증가하며, 해당 차례의 value와 해당 index부터 n-1번째까지의 값 중에 가장 작은 값의 위치를 교환하는 방식이다.

## selectionsort code
```go
package selectionsort

// SelectionSort sorts an integer array with selectionsort-way
func SelectionSort(arr *[]int) {
	var min int
	n := len(*arr)
	for i := 0; i < n; i++ {
		min = i
		for j := i; j < n; j++ {
			if (*arr)[min] > (*arr)[j] {
				min = j
			}
		}
		(*arr)[i], (*arr)[min] = (*arr)[min], (*arr)[i]
	}
}
```

## selectionsort main code
```go
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wookiist/go-ds/sort/selectionsort"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	arr := []int{23, 4, 5, 1, -1, 3, 100, 55, 66, 45, 46, 47, 44}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	selectionsort.SelectionSort(&arr)
	fmt.Fprintln(w, arr)
	selectionsort.SelectionSort(&arr2)
	fmt.Fprintln(w, arr2)
}
```