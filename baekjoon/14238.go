package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	w  = bufio.NewWriter(os.Stdout)
	sc = bufio.NewScanner(os.Stdin)
)

var (
	A, B, C, N int
	D          [51][51][51][3][3]int
)

func main() {
	defer w.Flush()
	for i := range D {
		for j := range D[i] {
			for k := range D[i][j] {
				for l := range D[i][j][k] {
					for m := range D[i][j][k][l] {
						D[i][j][k][l][m] = -1
					}
				}
			}
		}
	}
	countString()
	if solve(0, 0, 0, 0, 0) == 1 {
		fmt.Fprintln(w, makeAnswer(0, 0, 0, 0, 0))
	} else {
		fmt.Fprintln(w, -1)
	}
}

func makeAnswer(a, b, c, p1, p2 int) string {
	if a+b+c == N {
		return ""
	}
	if a+1 <= A && solve(a+1, b, c, 0, p1) == 1 {
		return "A" + makeAnswer(a+1, b, c, 0, p1)
	}
	if b+1 <= B && p1 != 1 {
		if solve(a, b+1, c, 1, p1) == 1 {
			return "B" + makeAnswer(a, b+1, c, 1, p1)
		}
	}
	if c+1 <= C && p1 != 2 && p2 != 2 {
		if solve(a, b, c+1, 2, p1) == 1 {
			return "C" + makeAnswer(a, b, c+1, 2, p1)
		}
	}
	return ""
}

func solve(a, b, c, p1, p2 int) int {
	temp := &D[a][b][c][p1][p2]
	if a+b+c == N {
		return 1
	}

	if *temp != -1 {
		return *temp
	}

	if a+1 <= A {
		if solve(a+1, b, c, 0, p1) == 1 {
			*temp = 1
			return *temp
		}
	}

	if b+1 <= B && p1 != 1 {
		if solve(a, b+1, c, 1, p1) == 1 {
			*temp = 1
			return *temp
		}
	}

	if c+1 <= C && p1 != 2 && p2 != 2 {
		if solve(a, b, c+1, 2, p1) == 1 {
			*temp = 1
			return *temp
		}
	}
	*temp = 0
	return *temp
}

func countString() {
	sc.Scan()
	s := strings.Split(sc.Text(), "")
	N = len(s)
	for i := range s {
		switch s[i] {
		case "A":
			A++
		case "B":
			B++
		case "C":
			C++
		}
	}
}
