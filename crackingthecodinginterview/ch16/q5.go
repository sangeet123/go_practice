package main

import (
	"fmt"
)

func NumTrailZeroes(n int) int {
	// factorial of negative no not defined
	if n < 0 {
		return -1
	}

	z := 0

	for fm := 5; fm <= n; fm *= 5 {
		z += (n / fm)
	}

	return z
}

func main() {
	fmt.Printf("-5:%d\n", NumTrailZeroes(-5))
	fmt.Printf("1:%d\n", NumTrailZeroes(1))
	fmt.Printf("5:%d\n", NumTrailZeroes(5))
	fmt.Printf("7:%d\n", NumTrailZeroes(7))
	fmt.Printf("10:%d\n", NumTrailZeroes(10))
	fmt.Printf("12:%d\n", NumTrailZeroes(12))
	fmt.Printf("15:%d\n", NumTrailZeroes(15))
	fmt.Printf("20:%d\n", NumTrailZeroes(20))
	fmt.Printf("25:%d\n", NumTrailZeroes(25))
	fmt.Printf("30:%d\n", NumTrailZeroes(30))
}
