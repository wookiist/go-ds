package main

import (
	"bufio"
	"os"
	"strings"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	mbti := getString()
	for i := range mbti {
		switch mbti[i] {
		case "E":
			w.WriteByte(byte('I'))
		case "I":
			w.WriteByte(byte('E'))
		case "S":
			w.WriteByte(byte('N'))
		case "N":
			w.WriteByte(byte('S'))
		case "T":
			w.WriteByte(byte('F'))
		case "F":
			w.WriteByte(byte('T'))
		case "J":
			w.WriteByte(byte('P'))
		case "P":
			w.WriteByte(byte('J'))
		}
	}
}

func getString() []string {
	s, _ := r.ReadString('\n')
	t := strings.Split(s, "")
	return t
}
