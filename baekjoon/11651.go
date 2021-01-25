package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

// Axis struct
type Axis struct {
	X int
	Y int
}

func main() {
	defer w.Flush()
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	arr := make([]Axis, N)
	for i := 0; i < N; i++ {
		s, _ := r.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		split := strings.Split(s, " ")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		a := Axis{
			X: x,
			Y: y,
		}
		arr[i] = a
	}

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].X < arr[j].X
	})

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].Y < arr[j].Y
	})

	for _, v := range arr {
		fmt.Fprintln(w, v.X, v.Y)
	}
}
