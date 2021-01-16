# InsertionSort with Go

InsertionSort 삽입 정렬은 마치 책장 넘기듯이 수행한다.
책장을 넘기다가. 내가 들어갈 차례가 되면 그 사이에 끼워 넣는 방식이다.

## insertionsort code
```go
package insertionsort

// InsertionSort sorts an integer slice with insertionsort-way
func InsertionSort(arr *[]int) {
	var tempVal int
	n := len(*arr)
	for i := 1; i < n; i++ {
		tempVal = (*arr)[i]
		for j := i - 1; j >= 0; j-- {
			if (*arr)[j] > tempVal {
				(*arr)[j+1], (*arr)[j] = (*arr)[j], (*arr)[j+1]
			} else {
				(*arr)[j+1] = tempVal
				break
			}
		}
	}
}
```

## insertionsort main code
```go
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wookiist/go-ds/sort/insertionsort"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	arr := []int{23, 4, 5, 1, -1, 3, 100, 55, 66, 45, 46, 47, 44}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	insertionsort.InsertionSort(&arr)
	fmt.Fprintln(w, arr)
	insertionsort.InsertionSort(&arr2)
	fmt.Fprintln(w, arr2)
}
```