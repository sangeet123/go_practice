package main

import (
	"fmt"
)

func RotateRight(arr []int, by int) bool {

	if by < 0 {
		return false
	}

	if by > len(arr) {
		by = by % len(arr)
	}

	Reverse(arr)
	Reverse(arr[:by])
	Reverse(arr[by:])
	return true
}

func Reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	RotateRight(arr, 10)
	fmt.Println(arr)

	RotateRight(arr, 2)
	fmt.Println(arr)

	RotateRight(arr, 5)
	fmt.Println(arr)
}
