package main

import "fmt"

func main() {
	var K, N, M int
	fmt.Scanf("%d %d %d\n", &K, &N, &M)
	if K*N > M {
		fmt.Println(K*N - M)
	} else {
		fmt.Println(0)
	}
}
