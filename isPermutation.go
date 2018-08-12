package main

import "fmt"

func isPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	m := make(map[rune]int)

	for _, v := range str1 {
		m[v]++
	}

	for _, v := range str2 {
		_, ok := m[v]

		if !ok {
			return false
		}

		m[v]--

		if m[v] == 0 {
			delete(m, v)
		}
	}

	return true
}

func main() {
	fmt.Println(isPermutation("abcde", "eabcd"))
	fmt.Println(isPermutation("aabcde", "eeabcd"))
	fmt.Println(isPermutation("abcded", "deabcd"))
}
