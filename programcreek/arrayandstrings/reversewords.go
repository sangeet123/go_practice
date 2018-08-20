package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Split(s, " ")
	Reverse(words)
	r := strings.Join(words, " ")
	return r
}

func Reverse(words []string) {
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
}

func main() {
	fmt.Println(ReverseWords("testing reverse words using library strings"))
}
