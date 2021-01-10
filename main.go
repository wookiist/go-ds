package main

import (
	"fmt"

	"github.com/wookiist/go-ds/linkedlist"
)

func main() {
	fmt.Println("Welcome to DS-Go")

	lList := linkedlist.LinkedList{}

	fmt.Println("len:", lList.Length())
	lList.Traverse()

	lList.Add(10)
	fmt.Println("len:", lList.Length())
	lList.Traverse()

	lList.Add(20)
	fmt.Println("len:", lList.Length())
	lList.Traverse()

	lList.Add(-10)
	fmt.Println("len:", lList.Length())
	lList.Traverse()

	lList.Add(30)
	fmt.Println("len:", lList.Length())
	lList.Traverse()
}
