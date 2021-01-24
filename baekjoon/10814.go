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

// Member struct
type Member struct {
	Name     string
	Age      int
	Register int
}

func main() {
	defer w.Flush()
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	m := []Member{}
	for i := 0; i < N; i++ {
		var tm Member
		s, _ := r.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		split := strings.Split(s, " ")
		tm.Age, _ = strconv.Atoi(split[0])
		tm.Name = split[1]
		tm.Register = i
		m = append(m, tm)
	}

	sort.SliceStable(m, func(i, j int) bool {
		return m[i].Register < m[j].Register
	})

	sort.SliceStable(m, func(i, j int) bool {
		return m[i].Age < m[j].Age
	})

	for _, v := range m {
		fmt.Fprintln(w, v.Age, v.Name)
	}
}
