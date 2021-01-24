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
	a := make([]Axis, N)
	for i := 0; i < N; i++ {
		var ta Axis
		s, _ := r.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		split := strings.Split(s, " ")
		ta.X, _ = strconv.Atoi(split[0])
		ta.Y, _ = strconv.Atoi(split[1])
		a[i] = ta
	}
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].Y < a[j].Y
	})
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].X < a[j].X
	})
	for _, v := range a {
		fmt.Fprintln(w, v.X, v.Y)
	}

}
