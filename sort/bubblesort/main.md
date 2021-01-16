# BubbleSort with Go

bubblesort는 현재 값과 바로 다음 값을 계속 비교해 나가며 현재 값이 더 큰 경우 다음 값과 자리 바꿈을 한다.

반복문이 두 개이기 때문에 O(n^2)
최악의 경우에는 n * (n-1) / 2
완전 정렬이 되어 있는 상태면 최선은 O(n)

추가 코드
- swapped bool 변수를 추가했다. swap되는 값이 없으면 더 이상 sorting을 진행할 이유가 없이 때문에 시간을 아끼고자 함.


## bubblesort code
```go
package bubblesort

// BubbleSort sorts an integer array with bubblesort-way
func BubbleSort(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
```

## bubblesort main code
```go
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
```