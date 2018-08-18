package main

import (
	"fmt"
	"sort"
)

func SmallestDiff(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)

	diff := a[len(a)-1]
	r := []int{-1, -1}

	for i, j := 0, 0; i < len(a) && j < len(b); {
		d := a[i] - b[j]

		if diff > d && d > 0 {
			r[0] = a[i]
			r[1] = b[j]
			diff = d
		}

		if d < 0 {
			i++
		} else {
			j++
		}

	}

	return r
}

func main() {
	a := []int{1, 3, 15, 11, 2}
	b := []int{23, 127, 235, 19, 8}

	fmt.Println(SmallestDiff(a, b))
}
