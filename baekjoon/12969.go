package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

var (
	D    [31][31][31][436]bool
	ans  []rune
	N, K int
)

func main() {
	defer w.Flush()
	fmt.Fscanf(r, "%d %d\n", &N, &K)
	ans = make([]rune, N+2)
	if solve(0, 0, 0, 0) {
		fmt.Fprintln(w, string(ans))
	} else {
		fmt.Fprintln(w, -1)
	}
}

func solve(idx, a, b, p int) bool {
	// 종료 조건
	if idx == N {
		if p == K {
			return true
		}
		return false
	}

	// 중복 탐색 여부
	if D[idx][a][b][p] {
		return false // 이미 방문했는데 아직까지도 안 끝났다는 건, 해당 노드가 정답이 아니라는 뜻
	}

	// 탐색 체크
	D[idx][a][b][p] = true // 현재 노드에 처음 방문했으므로 방문했음을 체크

	// 탐색
	// A를 넣어본다
	ans[idx] = 'A'
	if solve(idx+1, a+1, b, p) {
		return true
	}

	// B를 넣어본다
	ans[idx] = 'B'
	if solve(idx+1, a, b+1, p+a) {
		return true
	}

	// C를 넣어본다
	ans[idx] = 'C'
	if solve(idx+1, a, b, p+a+b) {
		return true
	}

	// 모두 실패시 실패 리턴
	return false
}
