package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var result int
	N := getInt()
	for i := 0; i < N; i++ {
		s := getString()
		if checkGroupWord(s) {
			result++
		}
	}
	fmt.Fprintf(w, "%d\n", result)
}

func getString() string {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	return s
}

func getInt() int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	i, _ := strconv.Atoi(s)
	return i
}

func checkGroupWord(s string) bool {
	for i := 'a'; i <= 'z'; i++ {
		r, _ := regexp.Compile(string(i))
		f := r.FindAllStringIndex(s, -1)
		if l := len(f); l > 1 {
			lastIdx := strings.LastIndex(s, string(i))
			firstIdx := strings.Index(s, string(i))
			if lastIdx-firstIdx != l-1 {
				return false
			}
		}
	}
	return true
}

// 더 나은 솔루션은 다음과 같다. 4ms 944kb
/*
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var word string
	var groupCount int
	for i:=0; i<n; i++ {
		fmt.Fscanln(reader, &word)
		var letters = make(map[uint8]bool)
		var prev uint8
		var isGroup = true
		for j:=0; j<len(word); j++ {
			_, exist := letters[word[j]]
			if !exist {
				letters[word[j]] = true
			} else if prev != word[j] {
				isGroup = false
				break
			}
			prev = word[j]
		}
		if isGroup {
			groupCount++
		}
	}

	fmt.Println(groupCount)
}
*/
