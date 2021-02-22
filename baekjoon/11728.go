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

func main() {
	defer w.Flush()
	var N, M int
	var aPointer, bPointer int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	A := getInts()
	B := getInts()
	for {
		if aPointer >= N && bPointer >= M {
			break
		} else if aPointer >= N {
			fmt.Fprintf(w, "%d ", B[bPointer])
			bPointer++
		} else if bPointer >= M {
			fmt.Fprintf(w, "%d ", A[aPointer])
			aPointer++
		} else {
			if A[aPointer] < B[bPointer] {
				fmt.Fprintf(w, "%d ", A[aPointer])
				aPointer++
			} else if A[aPointer] >= B[bPointer] {
				fmt.Fprintf(w, "%d ", B[bPointer])
				bPointer++
			}
		}
	}
}

func getInts() []int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	res := make([]int, len(t))
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}
