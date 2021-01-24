package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	s := bufio.NewScanner(os.Stdin)
	defer w.Flush()
	s.Split(bufio.ScanWords)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	var numOnePlace, plusOnePlace int
	cnt := 1
	if n < 10 {
		numOnePlace = n
		plusOnePlace = n
	} else {
		numOnePlace = n % 10
		plusOnePlace = (n%10 + (n-n%10)/10) % 10
	}
	for true {
		nowNum := 10*numOnePlace + plusOnePlace
		if nowNum == n {
			break
		}
		numOnePlace = nowNum % 10
		plusOnePlace = (nowNum%10 + (nowNum-nowNum%10)/10) % 10
		cnt++
	}
	fmt.Fprintln(w, cnt)
}
