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
	var C, N int
	fmt.Fscanf(r, "%d\n", &C)
	for i := 0; i < C; i++ {
		str, _ := r.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")
		split := strings.Split(str, " ")
		N, _ = strconv.Atoi(split[0])
		avg := 0.0
		res := 0
		arr := make([]int, N)
		for j := 0; j < N; j++ {
			arr[j], _ = strconv.Atoi(split[j+1])
			avg += float64(arr[j])
		}
		avg = avg / float64(N)
		for k := 0; k < N; k++ {
			if arr[k] > int(avg) {
				res++
			}
		}
		fmt.Fprintf(w, "%.3f%%\n", float64(res)/float64(N)*100)
	}
}
