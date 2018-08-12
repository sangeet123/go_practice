package main

import "fmt"

func isUnique(str string) bool {
	m := make(map[rune]bool)

	for _, v := range str {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = true
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcdefgh"))
	fmt.Println(isUnique(""))
	fmt.Println(isUnique("aaaa"))
	fmt.Println(isUnique("abcdea"))
}
