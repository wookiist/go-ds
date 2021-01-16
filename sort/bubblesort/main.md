# BubbleSort with Go

bubblesort는 현재 값과 바로 다음 값을 계속 비교해 나가며 현재 값이 더 큰 경우 다음 값과 자리 바꿈을 한다.

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
	bubblesort.BubbleSort(&arr)
	fmt.Fprintln(w, arr)
}
```