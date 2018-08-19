package main

import (
	"fmt"
)

func Flip(bit int32) int32 {
	return 1 ^ bit
}

func Sign(n int32) int32 {
	return Flip((n >> 31) & 0x1)
}

func MaxNaive(a, b int32) int32 {
	s := Sign(a - b)
	fs := Flip(s)
	return s*a + fs*b
}

func Max(a, b int32) int32 {
	sA := Sign(a)
	sB := Sign(b)
	sC := Sign(a - b)

	// if sign are different use sign of a (1,0) or (0,1)
	useSignOfA := sA ^ sB

	// if sign are same use sign of c (1,1) or (0,0)
	useSignOfC := Flip(sA ^ sB)

	k := sA*useSignOfA + sC*useSignOfC
	q := Flip(k)

	return k*a + q*b
}

func main() {
	// both positive case
	fmt.Println(Max(10, 2)) // expected 10
	fmt.Println(Max(2, 10)) // expected 10

	// one positive and one negative case
	fmt.Println(Max(10, -2)) // expected 10
	fmt.Println(Max(-2, 10)) // expected 10

	// both negative case
	fmt.Println(Max(-10, -15)) // expected -10
	fmt.Println(Max(-15, -10)) // expected -10
}
