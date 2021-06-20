package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	N, M, R int
	X       [][]int
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	N, M, R = scanInt(), scanInt(), scanInt()
	X = make([][]int, N)
	for i := range X {
		X[i] = make([]int, M)
		for j := range X[i] {
			X[i][j] = scanInt()
		}
	}
	for ; R > 0; R-- {
		switch scanInt() {
		case 1:
			X = reverseUpDown(X)
		case 2:
			X = reverseLeftRight(X)
		case 3:
			X = rotateRight(X)
		case 4:
			X = rotateLeft(X)
		case 5:
			X = rotateQuadRight(X)
		case 6:
			X = rotateQuadLeft(X)
		}
	}
	printMatrix(X)
}

func reverseUpDown(D [][]int) [][]int {
	n, m := len(D), len(D[0])
	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
	}
	for i := range D {
		for j := range D[i] {
			A[i][j] = D[(n-1)-i][j]
		}
	}
	return A
}

func reverseLeftRight(D [][]int) [][]int {
	n, m := len(D), len(D[0])
	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
	}
	for i := range D {
		for j := range D[i] {
			A[i][j] = D[i][(m-1)-j]
		}
	}
	return A
}

func rotateRight(D [][]int) [][]int {
	n, m := len(D), len(D[0])
	A := make([][]int, m)
	for i := range A {
		A[i] = make([]int, n)
	}
	for i := range A {
		for j := range A[i] {
			A[i][j] = D[(n-1)-j][i]
		}
	}
	return A
}

func rotateLeft(D [][]int) [][]int {
	n, m := len(D), len(D[0])
	A := make([][]int, m)
	for i := range A {
		A[i] = make([]int, n)
	}
	for i := range A {
		for j := range A[i] {
			A[i][j] = D[j][(m-1)-i]
		}
	}
	return A
}

func rotateQuadRight(D [][]int) [][]int {
	n, m := len(D), len(D[0])
	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
	}
	for i := range D {
		for j := range D[i] {
			if i < n/2 && j < m/2 {
				A[i][m/2+j] = D[i][j]
			}
			if i < n/2 && m/2 <= j {
				A[n/2+i][j] = D[i][j]
			}
			if n/2 <= i && m/2 <= j {
				A[i][j-m/2] = D[i][j]
			}
			if n/2 <= i && j < m/2 {
				A[i-n/2][j] = D[i][j]
			}
		}
	}
	return A
}

func rotateQuadLeft(D [][]int) [][]int {
	n, m := len(D), len(D[0])
	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
	}
	for i := range D {
		for j := range D[i] {
			if i < n/2 && j < m/2 {
				A[n/2+i][j] = D[i][j]
			}
			if i < n/2 && m/2 <= j {
				A[i][j-m/2] = D[i][j]
			}
			if n/2 <= i && m/2 <= j {
				A[i-n/2][j] = D[i][j]
			}
			if n/2 <= i && j < m/2 {
				A[i][j+m/2] = D[i][j]
			}
		}
	}
	return A
}

func printMatrix(arr [][]int) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Fprintf(w, "%d ", arr[i][j])
		}
		fmt.Fprintln(w)
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
