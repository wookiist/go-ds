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
	var money int
	fmt.Fscanf(r, "%d\n", &money)
	flow := getInts()
	jh := joonhyeon(money, flow)
	sm := seongmin(money, flow)
	if jh > sm {
		fmt.Fprintln(w, "BNP")
	} else if jh < sm {
		fmt.Fprintln(w, "TIMING")
	} else {
		fmt.Fprintln(w, "SAMESAME")
	}
}

func getInts() []int {
	s, _ := r.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Fields(s)
	res := make([]int, 14)
	for i := range t {
		res[i], _ = strconv.Atoi(t[i])
	}
	return res
}

func joonhyeon(money int, flow []int) int {
	res := 0
	for i := 0; i < 14; i++ {
		tmp := 0
		tmp += money / flow[i]
		res += tmp
		money -= flow[i] * tmp
	}
	return res*flow[13] + money
}

func seongmin(money int, flow []int) int {
	res := 0
	for i := 0; i < 14; i++ {
		if i > 2 {
			if flow[i-1] > flow[i-2] && flow[i-2] > flow[i-3] {
				// 3일 연속 상승장 <- 전량 매도
				money += res * flow[i]
				res = 0
			}
			if flow[i-1] < flow[i-2] && flow[i-2] < flow[i-3] {
				// 3일 연속 하락장 <- 전량 매수
				tmp := 0
				tmp += money / flow[i]
				res += tmp
				money -= flow[i] * tmp
			}
		}
	}
	return res*flow[13] + money
}
