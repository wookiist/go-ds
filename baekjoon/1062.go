package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	w       = bufio.NewWriter(os.Stdout)
	r       = bufio.NewReader(os.Stdin)
	result  int
	visited []bool
	words   []string
)

func main() {
	defer w.Flush()
	var N, K int
	fmt.Fscanf(r, "%d %d\n", &N, &K)
	words = make([]string, N)
	for i := range words {
		words[i] = getString()
	}
	if K < 5 {
		fmt.Fprintln(w, 0)
		return
	} else if K == 26 {
		fmt.Fprintln(w, N)
		return
	} else {
		visited = make([]bool, 26)
		visited['a'-'a'] = true
		visited['n'-'a'] = true
		visited['t'-'a'] = true
		visited['i'-'a'] = true
		visited['c'-'a'] = true
		solve(K-5, 'a')
		fmt.Fprintln(w, result)
	}
}

func solve(remain int, idx rune) {
	if remain == 0 {
		// 계산은 여기에서
		temp := 0
		for i := range words {
			flag := true
			for _, j := range words[i] {
				if !visited[j-'a'] {
					flag = false
					break
				}
			}
			if flag {
				temp++
			}
		}
		result = max(result, temp)
		return
	}
	// 아직 remain인 경우
	for i := idx; i <= 'z'; i++ {
		if !visited[i-'a'] {
			visited[i-'a'] = true
			solve(remain-1, i)
			visited[i-'a'] = false
		}
	}
}

func getString() string {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	return s[4 : len(s)-4]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
