// 15552.go 
// 7024KB 360ms
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()
	var t, a, b int
	fmt.Fscanf(r, "%d\n", &t)
	for i := 0; i < t; i++ {
		str, _ := r.ReadString('\n')
		text := strings.TrimSuffix(str, "\n")
		split := strings.Split(text, " ")
		a, _ = strconv.Atoi(split[0])
		b, _ = strconv.Atoi(split[1])
		fmt.Fprintln(w, a+b)
	}
}


// 6744KB 1160ms
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()
	var t, a, b int
	fmt.Fscanf(r, "%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		fmt.Fprintln(w, a+b)
	}
}
