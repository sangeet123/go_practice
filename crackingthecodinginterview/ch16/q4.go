package main

import (
	"fmt"
)

type TickTack [][]int

func checkRows(b [][]int) bool {
	for i := 0; i < len(b); i++ {
		if checkRow(b[i]) {
			return true
		}
	}
	return false
}

func checkRow(b []int) bool {
	for i := 1; i < len(b); i++ {
		if b[i] != b[i-1] {
			return false
		}
	}

	return true
}

func checkColumns(b [][]int) bool {
	for i := 0; i < len(b); i++ {
		if checkColumn(b, i) {
			return true
		}
	}

	return false
}

func checkColumn(b [][]int, c int) bool {
	for i := 1; i < len(b); i++ {
		if b[i][c] != b[i-1][c] {
			return false
		}
	}
	return true
}

func checkDiags(b [][]int) bool {
	return checkDiag(b) || checkOppDiag(b)
}

func checkDiag(b [][]int) bool {
	for i := 1; i < len(b); i++ {
		if b[i][i] != b[i-1][i-1] {
			return false
		}
	}

	return true
}

func checkOppDiag(b [][]int) bool {
	c := 1
	for i := len(b) - 2; i >= 0; i-- {
		if b[i+1][c-1] != b[i][c] {
			return false
		}
		c++
	}

	return true
}

func HasWon(b [][]int) bool {
	return checkDiags(b) || checkRows(b) || checkColumns(b)
}

func main() {
	// column win
	b := [][]int{
		{0, 0, 1},
		{1, 0, 1},
		{0, 1, 1},
	}

	fmt.Println(HasWon(b))

	// row win
	b = [][]int{
		{0, 1, 1},
		{0, 0, 0},
		{1, 1, 1},
	}

	fmt.Println(HasWon(b))

	// diag win
	b = [][]int{
		{1, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
	}

	fmt.Println(HasWon(b))

	// opp diag win
	b = [][]int{
		{0, 0, 1},
		{0, 1, 1},
		{1, 0, 1},
	}

	fmt.Println(HasWon(b))

	// draw
	b = [][]int{
		{1, 0, 1},
		{0, 0, 1},
		{1, 1, 0},
	}
	fmt.Println(HasWon(b))
}
