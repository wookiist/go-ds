package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w   = bufio.NewWriter(os.Stdout)
	r   = bufio.NewReader(os.Stdin)
	arr [][][]int
)

func main() {
	defer w.Flush()
	arr = make([][][]int, 51)
	for i := range arr {
		arr[i] = make([][]int, 51)
		for j := range arr[i] {
			arr[i][j] = make([]int, 51)
		}
	}
	arr[20][20][20] = 1048576
	var a, b, c int
	for {
		fmt.Fscanf(r, "%d %d %d\n", &a, &b, &c)
		if a == -1 && b == -1 && c == -1 {
			break
		}
		fmt.Fprintf(w, "w(%d, %d, %d) = %d\n", a, b, c, wFunc(a, b, c))
	}
}

func wFunc(a, b, c int) int {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	}
	if arr[a][b][c] != 0 {
		return arr[a][b][c]
	}
	for i := 0; i < 51; i++ {
		for j := 0; j < 51; j++ {
			for k := 0; k < 51; k++ {
				if i == 0 || j == 0 || k == 0 {
					arr[i][j][k] = 1
				} else {
					if i > 20 || j > 20 || k > 20 {
						arr[i][j][k] = arr[20][20][20]
					} else if i < j && j < k {
						arr[i][j][k] = arr[i][j][k-1] + arr[i][j-1][k-1] - arr[i][j-1][k]
					} else {
						arr[i][j][k] = arr[i-1][j][k] + arr[i-1][j-1][k] + arr[i-1][j][k-1] - arr[i-1][j-1][k-1]
					}
				}
				if i == a && j == b && k == c {
					return arr[i][j][k]
				}
			}
		}
	}
	return arr[a][b][c]
}
