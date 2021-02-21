package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var N, M int
	fmt.Fscanf(r, "%d %d\n", &N, &M)
	wordMap := make(map[string]int)
	wordSlice := make([]string, 0, N)
	for i := 0; i < N; i++ {
		word := getString()
		if len(word) < M {
			continue
		}
		if _, ok := wordMap[word]; ok {
			wordMap[word]++
			continue
		} else {
			wordMap[word] = 1
			wordSlice = append(wordSlice, word)
		}
	}
	sort.SliceStable(wordSlice, func(i, j int) bool {
		return wordSlice[i] < wordSlice[j]
	})
	sort.SliceStable(wordSlice, func(i, j int) bool {
		return len(wordSlice[i]) > len(wordSlice[j])
	})
	sort.SliceStable(wordSlice, func(i, j int) bool {
		return wordMap[wordSlice[i]] > wordMap[wordSlice[j]]
	})

	for i := range wordSlice {
		fmt.Fprintln(w, wordSlice[i])
	}
}

func getString() string {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	return s
}
