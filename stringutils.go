package main

import (
	"fmt"
)

func HasPrefix(str, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}

func HasSuffix(str, prefix string) bool {
	return len(str) >= len(prefix) && str[len(str)-len(prefix):] == prefix
}

func SubString(str, sub string) bool {

	for i, _ := range str {
		if HasPrefix(str[i:], sub) {
			return true
		}
	}

	return false
}

func main() {

	fmt.Println("--------------HasPrefix----------------")
	fmt.Println(HasPrefix("abcdefgh", "abc"))
	fmt.Println(HasPrefix("abcdefgh", "abcdefgh"))
	fmt.Println(HasPrefix("abcdefgh", "a23"))
	fmt.Println("--------------HasPrefix Ends----------------")

	fmt.Println("--------------HasSuffix----------------")
	fmt.Println(HasSuffix("abcdefgh", "h"))
	fmt.Println(HasSuffix("abcdefgh", "fgh"))
	fmt.Println(HasSuffix("abcdefgh", "fg2"))
	fmt.Println("--------------HasSuffix Ends----------------")

	fmt.Println("--------------SubString----------------")
	fmt.Println(SubString("abcdefgh", "cde"))
	fmt.Println(SubString("abcdefgh", "fg"))
	fmt.Println(SubString("abcdefgh", "ab"))
	fmt.Println(SubString("abcdefgh", "gh"))
	fmt.Println(SubString("abcdefgh", "fg1"))
	fmt.Println("--------------SubString Ends----------------")

}
