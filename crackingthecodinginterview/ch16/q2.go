package main

import (
	"fmt"
	"strings"
)

var wordCount map[string]int

func findWordFreq(book, word string) int {
	count := 0

	if wordCount != nil {
		count = wordCount[word]
		return count
	}

	wordCount = make(map[string]int)

	for _, w := range strings.Split(book, " ") {
		wordCount[w]++
	}

	count = wordCount[word]
	return count
}

func main() {
	book := "this is a very small portion of the book so that we can test book easily."
	fmt.Println(findWordFreq(book, "book"))
	fmt.Println(findWordFreq(book, "portion"))
	fmt.Println(findWordFreq(book, "none"))
}
