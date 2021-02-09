package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	fmt.Fprintln(w, "어느 한 컴퓨터공학과 학생이 유명한 교수님을 찾아가 물었다.")
	chatbot(N, N)
}

func chatbot(n, N int) {
	if n == 0 {
		s := strings.Repeat("____", N)
		fmt.Fprintln(w, s+"\"재귀함수가 뭔가요?\"")
		fmt.Fprintln(w, s+"\"재귀함수는 자기 자신을 호출하는 함수라네\"")
		fmt.Fprintln(w, s+"라고 답변하였지.")
		return
	}
	s := strings.Repeat("____", N-n)
	fmt.Fprintln(w, s+"\"재귀함수가 뭔가요?\"")
	fmt.Fprintln(w, s+"\"잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.")
	fmt.Fprintln(w, s+"마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.")
	fmt.Fprintln(w, s+"그의 답은 대부분 옳았다고 하네. 그런데 어느 날, 그 선인에게 한 선비가 찾아와서 물었어.\"")
	chatbot(n-1, N)
	fmt.Fprintln(w, s+"라고 답변하였지.")
}
