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
	T := getInt()
	for i := 0; i < T; i++ {
		fmt.Fprintln(w, getPlus())
	}
}

func getInt() int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	r, _ := strconv.Atoi(s)
	return r
}

func getPlus() int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	sArr := strings.Split(s, ",")
	r1, _ := strconv.Atoi(sArr[0])
	r2, _ := strconv.Atoi(sArr[1])
	return r1 + r2
}

/*
package main

import (
	"bufio"
	_ "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var N int
	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}
	for i := 0; i < N; i++ {

		if sc.Scan() {
			s := strings.Split(sc.Text(), ",")
			A, _ := strconv.Atoi(s[0])
			B, _ := strconv.Atoi(s[1])
			wr.WriteString(strconv.Itoa(A + B))
			wr.WriteString("\n")
		}

	}

}
*/
