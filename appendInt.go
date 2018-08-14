package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {

	zLen := len(x) + 1

	z := []int(nil)

	if zLen <= cap(x) {
		z = x[:zLen]
	} else {
		zcap := zLen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zLen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}

func main() {
	x := []int(nil)

	for i := 0; i < 10; i++ {
		x = appendInt(x, i+1)
		fmt.Printf("cap of x = %d, len of x = %d\n", cap(x), len(x))
	}

}
